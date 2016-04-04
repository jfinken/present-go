package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    //ch <- 3
    
    v := <-ch
    fmt.Println(v)
    fmt.Println(<-ch)
}
