#GO workspace
The top level directory is called the workspace.

`GOPATH` environment variable specifies the location of the workspace
```shell
export GOPATH=$HOME/gobook
```

note: Go tool is only compatible with code that resides in a workspace

#Standard directory layout
```
doc/                               # optional, not standard
bin/                               # holds compiled commands (  Each command is named for its source directory, but only using the final element ==> there can be only one binary per package)
    hello                          # command executable
    outyet                         # command executable
pkg/                               # holds installed package objects
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # package object
src/                               # hold source code 
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```            