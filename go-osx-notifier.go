package main

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	binaryPath = "../osx/terminal-notifier-{type}.app/Contents/MacOS/terminal-notifier"
)

const (
	Fail = "fail"
	Info = "info"
	Pass = "pass"
)

//how to do sounds enumeration????
const (
	Basso  = "Basso"
	Blow   = "Blow"
	Bottle = "Bottle"
)

type NotificationOptions struct {
	Type     string
	Message  string
	Title    string //optional
	Subtitle string //optional
}

func main() {

	//TODO: need to add the commands to a command slice
	//then call the exec.Command with the slice separated ...(like a params object)

	notify(&NotificationOptions{Type: Fail, Message: "Program stopped", Title: "Hello"})
	time.Sleep(1 * time.Second)

	notify(&NotificationOptions{Type: Info, Message: "Check your Apple Stock: 💰"})
	time.Sleep(1 * time.Second)

	notify(&NotificationOptions{Type: Pass, Message: "It's done!"})
	time.Sleep(1 * time.Second)
}

func notify(notificationOptions *NotificationOptions) {

	bPath := strings.Replace(binaryPath, "{type}", notificationOptions.Type, -1)

	var err error = nil
	if notificationOptions.Title != "" {
		_, err = exec.Command(bPath, "-message", notificationOptions.Message, "-title", notificationOptions.Title).Output()
	} else {
		_, err = exec.Command(bPath, "-message", notificationOptions.Message).Output()
	}

	if err != nil {
		log.Fatal(err)
	}

}
