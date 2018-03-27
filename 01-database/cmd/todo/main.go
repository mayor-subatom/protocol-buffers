package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"protocol-buffers/01-database/todo"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"io"
	"encoding/gob"
	"bytes"
)

const dbFilePath = "mydb.pb"


func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing command: either list or add")
		os.Exit(1)
	}

	var err error

	switch  cmd := flag.Arg(0); cmd {
	case "list":
		err = listTasks()

	case "add":
		listOfTasks := flag.Args()[1:]
		tasks := strings.Join(listOfTasks, " ")
		err = addTask(tasks)

	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
func addTask(text string) error {
	task := &todo.Task{
		Text: text,
		Done:false,
	}
	fmt.Println(proto.MarshalTextString(task))
	serialized, err := proto.Marshal(task)
	if err != nil {
		return fmt.Errorf("could not encoded task: %v",err)
	}

	dbFile, err := os.OpenFile(dbFilePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666 )
	if err != nil {
		return fmt.Errorf("could not encoded db file %s: %v",dbFilePath, err)
	}

	if err:=gob.NewEncoder(dbFile).Encode((int64(len(serialized))));err != nil {
		return fmt.Errorf("could not encode length of message: %v", err)
	}

	_,err = dbFile.Write(serialized)
	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	if err:=dbFile.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v",dbFilePath, err)
	}

	return nil
}
func listTasks() error {
	fileContent, err := ioutil.ReadFile(dbFilePath)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", dbFilePath, err)
	}

	//PB: we read but we are not advancing at all.
	//Solution: , see addTask()
	for {

		if len(fileContent) == 0 {
			return nil
		} else if len(fileContent) < 4 {
			return fmt.Errorf("remaining odd %d bytes: %v",)
		}

		var messageLength int64
		if err:=gob.NewDecoder(bytes.NewReader(fileContent[:4])).Decode(&messageLength);err!=nil {
			return fmt.Errorf("could not decode message length: %v",err)
		}

		fileContent = fileContent[4:]

		var task todo.Task
		if err := proto.Unmarshal(fileContent[:messageLength], &task); err == io.EOF {
			return nil
		} else if err != nil {
			return fmt.Errorf("could not read task: %v", err)
		}

		fileContent = fileContent[messageLength:]

		if task.Done {
			fmt.Printf("done")
		} else  {
			fmt.Printf("not done")
		}
		fmt.Printf(" %s\n", task.Text)
	}


	//When to stop ? What is the separator ?
	return nil
}
