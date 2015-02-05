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

	err := ssh.Scp("/home/core/zipkin.rb")

	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println("success")

		response, _ := ssh.Run("ls -al zipkin.rb")

		fmt.Println(response)
	}
}
