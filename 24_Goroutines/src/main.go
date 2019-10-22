package main

import (
        "fmt"
        "time"
       )

func f(str string) {

    for i:=0; i < 3; i++ {
        fmt.Println(str, " : ", i)
    }
}

// A goroutine is a lightweight thread of execution. Kind of thread as in C.
func main() {

    f("direct")

    // To invoke this function in a goroutine, use go f(s).
    // This new goroutine will execute concurrently with the calling one.
    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // You can also start a goroutine for an anonymous function call.
    time.Sleep(time.Second)
    fmt.Println("Done.")
}
