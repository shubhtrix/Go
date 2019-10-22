package main

import (
    "fmt"
    "time"
)

// This is the function we’ll run in a goroutine. 
// The done channel will be used to notify another goroutine 
// that this function’s work is done.
func routine1 (bln chan bool) {

    fmt.Print("11 Working...")
    time.Sleep(time.Second)
    fmt.Println("11 Done!!!")

    bln <- true
}

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine
// to finish. When waiting for multiple goroutines to finish, 
// you may prefer to use a WaitGroup.
func main() {

    // Start a worker goroutine, giving it the channel to notify on.
    done := make(chan bool, 1)
    go routine1 (done)

    // Block until we receive a notification from the worker on the channel.
    <-done

    fmt.Println("Final exit from main.")
}
