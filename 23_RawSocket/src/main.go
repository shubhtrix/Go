package main

import (
    "fmt"
    "reflect"
//    "time"
)

const (
    max_pktlen = 1500
    mac_len    = 6
)

// Just to have a look of crafted packet
func hex_dump(pkt_dump []byte, length int) {

    counter := 0

    for i, val := range pkt_dump {
        if i < length {
           fmt.Printf("\t%#x",val)
        } else {
            continue
        }
        counter++
        if (counter == 8) {
            fmt.Printf("\n")
            counter = 0
        }
    }
    fmt.Printf("\n")
}

// Sockets are network endpoints
func main() {

    offset := 0

    fmt.Println("Type of offset: ", reflect.TypeOf(offset))
    pkt := make ([]byte, max_pktlen)
    srcmac := []byte {0x00, 0x0c, 0x29, 0x6b, 0xa6, 0xc7}
    dstmac := []byte {0x00, 0x0c, 0x29, 0xc0, 0xa2, 0x0f}

    // 00:0c:29:c0:a2:0f with IP 172.16.108.26/24

    // Destination mac
    copy (pkt, dstmac)
    offset += 6
//    copy ((pkt+offset), srcmac)
//    offset += 6

    fmt.Println("Packet length: ", len(pkt))

    hex_dump(pkt, offset)

    //for _, val := range pkt {
    //    if val != nil {
    //        fmt.Println(val)
    //    }
    //}

    // Tickers use a similar mechanism to timers: a channel that is
    // sent values. Here we’ll use the select builtin on the channel
    // to await the values as they arrive every 500ms.

/*    ticker := time.NewTicker(500 * time.Millisecond)
    done   := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
*/
}
