package main

import (
	"fmt"

	"github.com/weekface/easyssh"
)

func main() {
	ssh := &easyssh.MakeConfig{
		User:   "cjc",
		Server: "localhost",
		Key:    "/.ssh/id_rsa",
	}

	response, err := ssh.Run("ls")
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}
