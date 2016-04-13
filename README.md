# present-go: an example of building a deck using `present`

Josh Finken

## Prerequisites

    go install golang.org/x/tools/cmd/present

## Run it!

    cd $GOPATH/src/path/to/present-go
    present

## View it:
   
    Surf to http://127.0.0.1:3999
    Note http://localhost:3999 will not work due to the embedded .play code which runs 
    over a websocket.  The underlying socket code explicitly checks the IP to prevent execution of 
    arbitrary code from another machine.
