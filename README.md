# present-go: examples building decks with `present`

Josh Finken

## Prerequisites

    go install golang.org/x/tools/cmd/present

## Run it:

    cd $GOPATH/src/path/to/present-go
    present

## Viewing:
   
* After running `present` surf to http://127.0.0.1:3999

    Note http://localhost:3999 will not work due to the embedded .play code which runs 
    over a websocket.  The underlying socket code explicitly checks the IP to prevent execution of 
    arbitrary code from another machine.

* Conveniently, the `-http` flag can for example specify a private IP address thus allowing
one to control the presentation remotely (provided a local network, local FW rules, etc).
Warnings are given for good and obvious reasons:

    $ present -http 10.0.0.123:3999

    2016/01/21 16:22:41
    WARNING!  WARNING!  WARNING!

    The present server appears to be listening on an address that is not localhost.
    Anyone with access to this address and port will have access to this machine as
    the user running present.

    To avoid this message, listen on localhost or run with -play=false.

    If you don't understand this message, hit Control-C to terminate this process.

    WARNING!  WARNING!  WARNING!
    2016/01/21 16:22:41 Open your web browser and visit http://10.0.0.123:3999


If linux, the very easy and convenient tool uncomplicated firewall, `ufw`, is very handy.

    $ sudo ufw status


* Alternatively, talks.godoc.org allows one to view .slide files hosted on Github.com.  Thus:

  * http://talks.godoc.org/github.com/jfinken/present-go/go-talk-duality.slide
  * http://talks.godoc.org/github.com/jfinken/present-go/go-talk-all.slide
