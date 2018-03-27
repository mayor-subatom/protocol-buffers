#The problem
```
src/                               # hold source code (usually src/github.com/<username>/<project_name>)
    github.com/golang/example/
	outyet/
	    main.go                # command source
	    main_test.go           # test source
```      
This has 2 drawbacks:
* only one binary allowed in a package. If you have more than one command in your application ....
* application cannot be used as a library

#The solution
In workspace/src/github.com/<username>/<appname>, use a cmd directory, having one subdirectory for each binary.
```
cmd/
    command1/
             main.go
    ...
    commandN/
             main.go
```

Subdirectory's name should match the name of the command you want to have in /bin 

Ideally subdirectory has a file `main.go` (or `<subdirectory_name>.go`) containing the main function.

Install with by example :
```go
go get github.com/johndoe/my_app/...
```

In your command source you can import your application logic (package) located in /pkg 

#Example
https://github.com/go-kit/kit/tree/master/examples/addsvc
