package main

import (
	"fmt"

	"github.com/hypersleep/easyssh"
)

func main() {
	ssh := &easyssh.MakeConfig{
		User:   "core",
		Server: "core",
		Key:    "/.ssh/id_rsa",
	}

	response, err := ssh.Run("ps ax")
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}
