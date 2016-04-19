package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum                // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c) // [7, 2, 8]
    go sum(s[len(s)/2:], c) // [-9, 4, 0]
    x, y := <-c, <-c        // receive from c

    fmt.Printf("x=%d, y=%d, x+y=%d", x, y, x+y)
}
