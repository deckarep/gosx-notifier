package gosxnotifier

import (
	"errors"
	"net/url"
	"os/exec"
	"strings"
)

const (
	binaryPath = "osx/terminal-notifier-{type}.app/Contents/MacOS/terminal-notifier"
)

type MessageType string

const (
	Fail MessageType = "fail"
	Info MessageType = "info"
	Pass MessageType = "pass"
)

type Sound string

const (
	Default Sound = "default"
	Basso   Sound = "Basso"
	Blow    Sound = "Blow"
	Bottle  Sound = "Bottle"
	Frog    Sound = "Frog"
	Funk    Sound = "Funk"
	Glass   Sound = "Glass"
	Hero    Sound = "Hero"
	Morse   Sound = "Morse"
	Ping    Sound = "Ping"
	Pop     Sound = "Pop"
	Purr    Sound = "Purr"
	Sosumi  Sound = "Sosumi"
	Tink    Sound = "Tink"
)

type Notification struct {
	Type     MessageType //required
	Message  string      //required
	Title    string      //optional
	Subtitle string      //optional
	Sound    Sound       //optional
	Link     string      //optional
}

func NewNotification(messageType MessageType, message string) *Notification {
	n := &Notification{Type: messageType, Message: message}
	return n
}

func (n *Notification) SendNotification() error {

	commandTuples := make([]string, 0)

	//check required commands
	if n.Message == "" {
		return errors.New("Please specifiy a proper message argument.")
	} else {
		commandTuples = append(commandTuples, []string{"-message", n.Message}...)
	}

	//add title if found
	if n.Title != "" {
		commandTuples = append(commandTuples, []string{"-title", n.Title}...)
	}

	//add subtitle if found
	if n.Subtitle != "" {
		commandTuples = append(commandTuples, []string{"-subtitle", n.Subtitle}...)
	}

	//add sound if specified
	if n.Sound != "" {
		commandTuples = append(commandTuples, []string{"-sound", string(n.Sound)}...)
	}

	//add url if specified
	url, err := url.Parse(n.Link)
	if err != nil {
		n.Link = ""
	}
	if url != nil {
		commandTuples = append(commandTuples, []string{"-open", n.Link}...)
	}

	//add bundle id if specified
	if strings.HasPrefix(strings.ToLower(n.Link), "com.") {
		commandTuples = append(commandTuples, []string{"-activate", n.Link}...)
	}

	if len(commandTuples) == 0 {
		return errors.New("Please provide a Message and Type at a minimum.")
	}

	bPath := strings.Replace(binaryPath, "{type}", string(n.Type), -1)

	_, err = exec.Command(bPath, commandTuples...).Output()
	if err != nil {
		return err
	}

	return nil
}
