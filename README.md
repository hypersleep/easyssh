# easyssh
Golang package for easy remote execution through SSH

So easy to use!

For example `ps ax`

```
package main

import (
	"fmt"
	"github.com/hypersleep/easyssh"
)

func main() {
	ssh := &easyssh.MakeConfig {
		User: "core",
		Server: "core",
		Key: "/.ssh/id_rsa",
	}

	response, err := ssh.ConnectAndRun("ps ax")
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}
```
