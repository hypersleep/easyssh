package easyssh

import (
	"testing"
)

var sshConfig = &MakeConfig{
	User:     "root",
	Server:   "172.30.19.2",
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
		outChannel, errChannel, done, err := sshConfig.Stream(testCase[0], 10)
		if err != nil {
			t.Errorf("Stream failed: %s", err)
		}
		stillGoing := true
		stdout := ""
		stderr := ""
		for stillGoing {
			select {
			case <-done:
				stillGoing = false
			case line := <-outChannel:
				stdout += line
			case line := <-errChannel:
				stderr += line
			}
		}
		if stdout != testCase[1] {
			t.Error("Output didn't match expected: %s,%s", stdout, stderr)
		}
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	commands := []string{
		"echo test", `for i in $(ls); do echo "$i"; done`, "ls",
	}
	for _, cmd := range commands {
		stdout, stderr, istimeout, err := sshConfig.Run(cmd, 10)
		if err != nil {
			t.Errorf("Run failed: %s", err)
		}
		if stdout == "" {
			t.Errorf("Output was empty for command: %s,%s,%s", cmd, stdout, stderr, istimeout)
		}
	}
}
