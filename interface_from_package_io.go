// Interfaces from the io package:

// Reads up to len(p) bytes into p, returning the number of bytes read and any error encountered.
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writes len(p) bytes from p to the underlying data stream
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Interface that groups the basic Read and Write methods
type ReadWriter interface {
    Reader
    Writer
}

