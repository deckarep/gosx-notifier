node-osx-notifier
===========================
A [NodeJS](http://nodejs.org) Server for sending notifications to OSX Mountain Lion's
[Notification Center](http://www.macworld.com/article/1165411/mountain_lion_hands_on_with_notification_center.html).

Synopsis
--------
OSX Mountain Lion comes packaged with a built-in notification center. For whatever reason, [Apple sandboxed the
notification center API](http://forums.macrumors.com/showthread.php?t=1403807) to apps hosted in its App Store. The end
result? A potentially useful API shackled to Apple's ecosystem.

Thankfully, [Eloy Durán](https://github.com/alloy) put together a
[set of sweet osx apps](https://github.com/alloy/terminal-notifier) that allow terminal access to the sandboxed API.
**node-osx-notifier** wraps these apps with a simple [express](https://github.com/visionmedia/express) server, exposing
an HTTP interface to the closed API.

It's not perfect, and the implementor will quickly notice its limitations. However, it's a start and any pull requests
are accepted and encouraged!

Installation
------------
The following command will install the notification server. Use `-g` to install the server as a global binary.

```sh
[sudo] npm install [-g] node-osx-notifier
```

Running The Server
------------------
Running the server is easy peasy. If you installed the server globally, then starting the server is as easy as:

```sh
node-osx-notifier [port] [host]
```

The port and host will default to `1337` and `localhost` respectively.

Testing The Server
------------------
You can then test that the server is running correctly by making a request to it. The simplest request uses the
querystring over HTTP GET:

```sh
curl "http://localhost:1337/info?message=test"
```

```json
{
  "status": "* Notification delivered."
}
```

You can also use the JSON interface to POST the same content to the server:

```sh
curl -H "Content-Type: application/json" -X POST -d '{"message":"test"}' "http://localhost:1338/info"
```

```json
{
  "status": "* Notification delivered."
}
```

The HTTP API
------------
For starters, you want to pick from one of three notification types. The notification types are designated by the path
of the server request. Each notification type changes the icon and section in the notification center:

* _info_ `http://localhost:1337/info` used for basic notifications
   * ![info screenshot](http://f.cl.ly/items/0P3i301J281x1A0Q1L46/Screen%20Shot%202012-08-30%20at%201.19.18%20AM.png)
* _pass_ `http://localhost:1337/pass` used for showing that a job has passed
   * ![info screenshot](http://f.cl.ly/items/0Q0A3b2S0v0Q2E3l0B0q/Screen%20Shot%202012-08-30%20at%201.21.03%20AM.png)
* _fail_ `http://localhost:1337/fail` used for showing that a job has failed
   * ![info screenshot](http://f.cl.ly/items/1H3v2H173A0r3a2F3l0x/Screen%20Shot%202012-08-30%20at%201.19.57%20AM.png)

In addition, you will also need to pass parameters (as JSON POST-data or a querystring) that tells the server what to
do for a given notification type. Since the server acts as a wrapper, these parameters match
[the command-line options](https://github.com/alloy/terminal-notifier#options) defined by the underlying apps. For
completeness, those parameters are outlined below:

At a minimum, you have to specify either the `-message` , the `-remove`
option or the `-list` option.

-------------------------------------------------------------------------------

`-message VALUE`  **[required]**

The message body of the notification.

-------------------------------------------------------------------------------

`-title VALUE`

The title of the notification. This defaults to ‘Terminal’.

-------------------------------------------------------------------------------

`-subtitle VALUE`

The subtitle of the notification.

-------------------------------------------------------------------------------

`-group ID`

Specifies the ‘group’ a notification belongs to. For any ‘group’ only _one_
notification will ever be shown, replacing previously posted notifications.

A notification can be explicitely removed with the `-remove` option, describe
below.

Examples are:

* The sender’s name to scope the notifications by tool.
* The sender’s process ID to scope the notifications by a unique process.
* The current working directory to scope notifications by project.

-------------------------------------------------------------------------------

`-remove ID`  **[required]**

Removes a notification that was previously sent with the specified ‘group’ ID,
if one exists. If used with the special group "ALL", all message are removed.

-------------------------------------------------------------------------------

`-list ID` **[required]**

Lists details about the specified ‘group’ ID. If used with the special group
"ALL", details about all currently active  messages are displayed.

The output of this command is tab-separated, which makes it easy to parse.

-------------------------------------------------------------------------------

`-activate ID`

Specifies which application should be activated when the user clicks the
notification.

You can find the bundle identifier of an application in its `Info.plist` file
_inside_ the application bundle.

Examples are:

* `com.apple.Terminal` to activate Terminal.app
* `com.apple.Safari` to activate Safari.app

-------------------------------------------------------------------------------

`-open URL`

Specifies a resource to be opened when the user clicks the notification. This
can be a web or file URL, or any custom URL scheme.

-------------------------------------------------------------------------------

`-execute COMMAND`

Specifies a shell command to run when the user clicks the notification.

Licence
-------
This project is dual licensed under the [MIT](https://github.com/azoff/node-osx-notifier/blob/master/LICENSE-MIT)
license and defers to [any licensing defined by the underlying apps](https://github.com/alloy/terminal-notifier).