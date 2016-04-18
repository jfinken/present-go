package main

import (
    "os"
    "io"
    "strings"
    "log"
)
func main() {
    var thingy io.ReadWriter
    // os.Stdout implements Writer (and Reader), thus this assignment is not a compile error
    thingy = os.Stdout
    
    r1 := strings.NewReader("first reader\n")
    r2 := strings.NewReader("second reader")
    
    // Returns a Reader (in this case a logical concat of the input Readers)
    r := io.MultiReader(r1, r2)

    // io.Copy takes a Writer (and a Reader) thus our 'thingy' satifies that interface...
    if _, err := io.Copy(thingy, r); err != nil {
        log.Fatal(err)
    }
}
