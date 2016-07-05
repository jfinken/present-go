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

* Alternatively, talks.godoc.org allows one to view .slide files hosted on Github.com.  Thus:

  * http://talks.godoc.org/github.com/jfinken/present-go/go-talk-duality.slide
  * http://talks.godoc.org/github.com/jfinken/present-go/go-talk-all.slide
