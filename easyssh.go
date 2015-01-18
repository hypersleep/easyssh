package easyssh

import(
	"golang.org/x/crypto/ssh"
	"os/user"
	"io/ioutil"
	"bytes"
)

type MakeConfig struct {
	User string
	Server string
	Key string
}

func getKeyFile(keypath string) (ssh.Signer, error){
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

func (ssh_conf *MakeConfig) ConnectAndRun(cmd string) (string, error) {
	pubkey, err := getKeyFile(ssh_conf.Key)
	if err != nil {
		return "", err
	}

	config := &ssh.ClientConfig{
			User: ssh_conf.User,
			Auth: []ssh.AuthMethod{ssh.PublicKeys(pubkey)},
		}
	
	client, err := ssh.Dial("tcp", ssh_conf.Server+":22", config)
	if err != nil {
		return "", err
	}

	session, err := client.NewSession()
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
	defer session.Close()

	return b.String(), nil
}
