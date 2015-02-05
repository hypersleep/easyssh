package easyssh

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

type MakeConfig struct {
	User   string
	Server string
	Key    string
}

func getKeyFile(keypath string) (ssh.Signer, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	file := usr.HomeDir + keypath
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	pubkey, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, err
	}

	return pubkey, nil
}

func (ssh_conf *MakeConfig) connect() (*ssh.Session, error) {
	pubkey, err := getKeyFile(ssh_conf.Key)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: ssh_conf.User,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(pubkey)},
	}

	client, err := ssh.Dial("tcp", ssh_conf.Server+":22", config)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (ssh_conf *MakeConfig) Run(cmd string) (string, error) {
	session, err := ssh_conf.connect()

	if err != nil {
		return "", err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(cmd)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (ssh_conf *MakeConfig) Scp(sourceFile string) error {
	session, err := ssh_conf.connect()

	if err != nil {
		return err
	}
	defer session.Close()

	targetFile := filepath.Base(sourceFile)

	src, srcErr := os.Open(sourceFile)

	if srcErr != nil {
		return srcErr
	}

	srcStat, statErr := src.Stat()

	if statErr != nil {
		return statErr
	}

	go func() {
		w, _ := session.StdinPipe()

		fmt.Fprintln(w, "C0644", srcStat.Size(), targetFile)

		if srcStat.Size() > 0 {
			io.Copy(w, src)
			fmt.Fprint(w, "\x00")
			w.Close()
		} else {
			fmt.Fprint(w, "\x00")
			w.Close()
		}
	}()

	if err := session.Run(fmt.Sprintf("scp -t %s", targetFile)); err != nil {
		return err
	}

	return nil
}
