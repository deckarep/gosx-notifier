gosx-notifier
===========================
A [Go](http://golang.org) lib for sending notifications to OSX Mountain Lion's (10.8 or higher REQUIRED)
[Notification Center](http://www.macworld.com/article/1165411/mountain_lion_hands_on_with_notification_center.html).

Synopsis
--------
OSX Mountain Lion comes packaged with a built-in notification center. For whatever reason, [Apple sandboxed the
notification center API](http://forums.macrumors.com/showthread.php?t=1403807) to apps hosted in its App Store. The end
result? A potentially useful API shackled to Apple's ecosystem.

Thankfully, [Eloy Durán](https://github.com/alloy) put together a
[set of sweet osx apps](https://github.com/alloy/terminal-notifier) that allow terminal access to the sandboxed API.
**gosx-notifier** wraps these apps with a simple interface to the closed API.

It's not perfect, and the implementor will quickly notice its limitations. However, it's a start and any pull requests are accepted and encouraged!

Go Version Inspired By:
--------------------
[node-osx-notifier](https://github.com/azoff/node-osx-notifier) This version is designed to be a node.js server where you can fire notifications by making REST requests.  The Go version is a static API for Go designed to be used within a Go app as needed.


Installation and Requirements
-----------------------------
The following command will install the notification api for Go along with the binaries.  Also, utilizing this lib requires OSX 10.8 or higher. It will simply not work on lower versions of OSX.

```sh
go get github.com/deckarep/gosx-notifier
```

Using the Code
------------------
It's a pretty straightforward API:

```Go
	
//At a minimum specifiy a message to display to end-user.
note := NewNotification("Check your Apple Stock!")

//Optionally, set a title
note.Title = "It's money making time 💰"

//Optionally, set a subtitle
note.Subtitle = "My subtitle" 

//Optionally, set a sound from a predefined set.
note.Sound = Basso

//Optionally, set a sender (Notification will now use the Safari icon)
note.Sender = "com.apple.Safari"

//Optionally, specifiy a url or bundleid to open should the notification be clicked.
note.Link = "http://www.yahoo.com" //or BundleID like: com.apple.Terminal

//Then, fire off the notification
err := note.SendNotification()

//If necessary, check error
if err != nil {
	log.Println("Uh oh!")
}
```

Sample App: Desktop Pinger Notification - monitors your websites and will notifiy you when a website is down.
```Go
package main

import (
	"github.com/deckarep/gosx-notifier"
	"net/http"
	"strings"
	"time"
)

//a slice of string sites that you are interested in watching
var sites []string = []string{
	"http://www.yahoo.com",
	"http://www.google.com",
	"http://www.bing.com"}

func main() {
	ch := make(chan string)

	for _, s := range sites {
		go pinger(ch, s)
	}

	for {
		select {
		case result := <-ch:
			if strings.HasPrefix(result, "-") {
				s := strings.Trim(result, "-")
				showNotification("Urgent, can't ping website: " + s)
			}
		}
	}
}

func showNotification(message string) {

	note := gosxnotifier.NewNotification(message)
	note.Title = "Site Down"
	note.Sound = gosxnotifier.Default

	note.SendNotification()
}

//Prefixing a site with a + means it's up, while - means it's down
func pinger(ch chan string, site string) {
	for {
		res, err := http.Get(site)

		if err != nil {
			ch <- "-" + site
		}

		if res != nil && res.Body != nil {
			defer res.Body.Close()

			if res.StatusCode != 200 {
				ch <- "-" + site
			} else {
				ch <- "+" + site
			}
		}

		time.Sleep(30 * time.Second)
	}
}
```

Coming Soon
-----------
* Group ID
* Remove ID


Licence
-------
This project is dual licensed under [any licensing defined by the underlying apps](https://github.com/alloy/terminal-notifier) and MIT licensed for this version written in Go.
