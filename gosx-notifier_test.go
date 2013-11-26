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

//TODO: write some real freakin' tests.... (you got me, I didn't build this using TDD)
