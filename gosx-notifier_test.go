package gosxnotifier

import (
	"fmt"
	"go/build"
	"path/filepath"
	"testing"
)

func Test_SendBasicNotification(t *testing.T) {
	note := NewNotification(Info, "Hello")
	err := note.SendNotification()
	if err != nil {
		t.Error("Error not nil, failed to send a basic notification.")
	}
}

func Test_Package(t *testing.T) {

	//fmt.Println(build.Default.GOPATH)

	fp := filepath.Join(build.Default.GOPATH, "github.com/deckarep/gosx-notifier")
	fmt.Println(fp)
}
