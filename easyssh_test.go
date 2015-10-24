package easyssh

import (
	"testing"
)

var sshConfig = &MakeConfig{
	User:     "username",
	Server:   "example.com",
	Password: "password",
	//Key:  "/.ssh/id_rsa",
	Port: "22",
}

func TestStream(t *testing.T) {
	t.Parallel()
	// input command/output string pairs
	testCases := [][]string{
		{`for i in $(seq 1 5); do echo "$i"; done`, "12345"},
		{`echo "test"`, "test"},
	}
	for _, testCase := range testCases {
		channel, done, err := sshConfig.Stream(testCase[0])
		if err != nil {
			t.Errorf("Stream failed: %s", err)
		}
		stillGoing := true
		output := ""
		for stillGoing {
			select {
			case <-done:
				stillGoing = false
			case line := <-channel:
				output += line
			}
		}
		if output != testCase[1] {
			t.Error("Output didn't match expected: %s", output)
		}
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	commands := []string{
		"echo test", `for i in $(ls); do echo "$i"; done`, "ls",
	}
	for _, cmd := range commands {
		out, err := sshConfig.Run(cmd)
		if err != nil {
			t.Errorf("Run failed: %s", err)
		}
		if out == "" {
			t.Errorf("Output was empty for command: %s", cmd)
		}
	}
}
