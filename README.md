# easyssh

## Description

Package easyssh provides a simple implementation of some SSH protocol features in Go.
You can simply run command on remote server or get a file even simple than native console SSH client.
Do not need to think about Dials, sessions, defers and public keys...Let easyssh will be think about it!

## So easy to use!

Run a command on remote server and get STDOUT output:

example/run.go

```
package main

import (
	"fmt"

	"github.com/hypersleep/easyssh"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "john",
		Server: "example.com",
		Key:    "/.ssh/id_rsa",
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run("ps ax")
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}
```

Upload a file to remote server:

example/scp.go

```
package main

import (
	"fmt"

	"github.com/hypersleep/easyssh"
)

func main() {
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "john",
		Server: "example.com",
		Key:    "/.ssh/id_rsa",
	}

	// Call Scp method with file you want to upload to remote server.
	err := ssh.Scp("/home/core/zipkin.rb")

	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("success")

		response, _ := ssh.Run("ls -al zipkin.rb")

		fmt.Println(response)
	}
}
```
