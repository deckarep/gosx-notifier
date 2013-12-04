package gosxnotifier

import (
	"log"
	"os"
	"testing"
)

func exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func Test_Install(t *testing.T) {
	//assert file exists

	if !exists(FinalPath) {
		t.Error("Test_Install failed to install the terminal-notifier.app bundle")
	} else {
		log.Println("terminal-notifier.app bundle installed successfully at: ", FinalPath)
	}
}

func Test_NewNotifier(t *testing.T) {
	n := NewNotification("Hello")

	//assert defaults
	if n.Message != "Hello" {
		t.Error("NewNotification doesn't have a Message specified")
	}
}

func Test_Push(t *testing.T) {
	n := NewNotification("Testing Push")
	err := n.Push()

	if err != nil {
		t.Error("Test_Push failed with error: ", err)
	}
}

func Test_Title(t *testing.T) {
	n := NewNotification("Testing Title")
	n.Title = "gosx-notifier is amazing!"
	err := n.Push()

	if err != nil {
		t.Error("Test_Title failed with error: ", err)
	}
}

func Test_Subtitle(t *testing.T) {
	n := NewNotification("Testing Subtitle")
	n.Subtitle = "gosx-notifier rocks!"

	err := n.Push()

	if err != nil {
		t.Error("Test_Subtitle failed with error: ", err)
	}
}

func Test_Sender(t *testing.T) {

	for _, s := range []string{"com.apple.Safari", "com.apple.iTunes"} {

		n := NewNotification("Testing Icon")
		n.Title = s
		n.Sender = s

		err := n.Push()

		if err != nil {
			t.Error("Test_Sender failed with error: ", err)
		}
	}
}

func Test_Group(t *testing.T) {
	const app_id string = "github.com/deckarep/gosx-notifier"

	for i := 0; i < 3; i++ {
		n := NewNotification("Testing Group Functionality...")
		n.Group = app_id

		err := n.Push()

		if err != nil {
			t.Error("Test_Group failed with error: ", err)
		}

	}
}

/*
	Not an easy way to verify the tests below actually work as designed, but here for completion.
*/

func Test_Sound(t *testing.T) {
	n := NewNotification("Testing Sound")
	n.Sound = Default
	err := n.Push()

	if err != nil {
		t.Error("Test_Sound failed with error: ", err)
	}
}

func Test_Link_Url(t *testing.T) {
	n := NewNotification("Testing Link Url")
	n.Link = "http://www.yahoo.com"
	err := n.Push()

	if err != nil {
		t.Error("Test_Link failed with error: ", err)
	}
}

func Test_Link_App_Bundle(t *testing.T) {
	n := NewNotification("Testing Link Terminal")
	n.Link = "com.apple.Safari"
	err := n.Push()

	if err != nil {
		t.Error("Test_Link failed with error: ", err)
	}
}
