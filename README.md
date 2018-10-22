# go-keylogger
Keylogger written in Go

[![Build Status](https://travis-ci.org/mtojek/go-keylogger.svg?branch=master)](https://travis-ci.org/mtojek/go-keylogger)

Status: **Done**

Record keystrokes in the Linux environment. Keylogger listen for input events on selected input device and stores decoded hits in a specified log file. The application requires root permissions.

## Features

* List available input devices
* Record keystrokes sent to observed input devices (e.g. keyboards)
* Collect logged input data in log files

## Usage

~~~
$ keylogger 
Usage: keylogger [--version] [--help] <command> [<args>]

Available commands are:
    devices    Lists available input devices
    record     Records any keys pressed on the selected device
    version    Prints the application version
~~~

### Examples

List available input devices:

~~~
# keylogger devices
Available event devices:
  event0 (name: "AT Translated Set 2 keyboard", path: /dev/input/event0)
  event1 (name: "Power Button", path: /dev/input/event1)
  event2 (name: "Sleep Button", path: /dev/input/event2)
  event3 (name: "VirtualBox mouse integration", path: /dev/input/event3)
  event4 (name: "ImExPS/2 Generic Explorer Mouse", path: /dev/input/event4)
~~~

Start recording input events:

~~~
# keylogger record --eventPath=/dev/input/event0 --logPath=/tmp/keylogger.log
Start recording...
~~~

You can see ```/tmp/keylogger.log``` content:

~~~
# cat /tmp/keylogger.log 
HELLO
<R_SHIFT>MARCIN <R_SHIFT>TOJEK<L_CTRL><L_CTRL><L_CTRL><L_CTRL><L_CTRL><L_CTRL>C
~~~

## License

**Apache License 2.0**

A permissive license whose main conditions require preservation of copyright and license notices. Contributors provide an express grant of patent rights. Licensed works, modifications, and larger works may be distributed under different terms and without source code.
