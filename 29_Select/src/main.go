package main

// Go’s select lets you wait on multiple channel operations. 
// Combining goroutines and channels with select is a powerful feature of Go.

import (
    "fmt"
    "time"
)

func main() {

    ch1 := make(chan string)
    ch2 := make(chan string)

    // Each channel will receive a value after some amount of time, 
    // to simulate e.g. blocking RPC operations executing in concurrent goroutines.

    go func() {
        time.Sleep(time.Second * 1)
        ch1 <- "One"
    } ()

    go func() {
        time.Sleep(time.Second * 2)
        ch2 <- "Two"
    } ()

    // We’ll use select to await both of these values simultaneously, 
    // printing each one as it arrives
    for i:=0; i<2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Recieved : ", msg1)
        case msg2 := <-ch2:
            fmt.Println("Recieved : ", msg2)
        }
    }
}
