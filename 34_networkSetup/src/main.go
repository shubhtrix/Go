
package main

import (
	"encoding/binary"
    "os"
	"fmt"
    "log"
	"net"
	"time"
    "os/exec"
//    "strings"
)

func readMode (tun Interface) {

	rBuf := make([]byte, 1500)

    for {
        n, err := tun.Read(rBuf)
        //_, err := tun.Read(rBuf)
        if err == ErrNotReady {
		    fmt.Println("Warning: tun: interface not ready, waiting 1s...")
		    time.Sleep(1 * time.Second)
		    continue
	    } else if err != nil {
	        fmt.Println("error: read:", err)
		    break
	    }

        fmt.Print("\n-- Packet Received --\n\n")

        switch rBuf[0] & 0xF0 {
        case 0x40:
            fmt.Println("IP Version: 4")

            fmt.Printf("Length: %d\n", binary.BigEndian.Uint16(rBuf[2:4]))
            fmt.Printf("Protocol: %d (1=ICMP, 6=TCP, 17=UDP)\n", rBuf[9])
            fmt.Printf("Source IP: %s\n", net.IP(rBuf[12:16]))
            fmt.Printf("Destination IP: %s\n", net.IP(rBuf[16:20]))

            ihl := (rBuf[0] & 0x0F) * 4
            fmt.Printf("Payload: [%x]\n", rBuf[ihl:n])

        case 0x60:
            fmt.Println("IP Version: 6")

            fmt.Printf("Length: %d\n", binary.BigEndian.Uint16(rBuf[4:6]))
            fmt.Printf("Protocol: %d (1=ICMP, 6=TCP, 17=UDP)\n", rBuf[7])
            fmt.Printf("Source IP: %s\n", net.IP(rBuf[8:24]))
            fmt.Printf("Destination IP: %s\n", net.IP(rBuf[24:40]))

            fmt.Printf("Payload: [%x]\n", rBuf[40:n])


        }
        _, err = tun.Write(rBuf)
        //_, err := tun.Read(rBuf)
        if err == ErrNotReady {
		    fmt.Println("Warning: tun: interface not ready, waiting 1s...")
		    time.Sleep(1 * time.Second)
		    continue
	    } else if err != nil {
	        fmt.Println("error: read:", err)
		    break
	    }
    }
}

func TunSetup(thrdId uint64, ipAddr string) (int, error) {

    // Create tun name here
    intfName := fmt.Sprintf("upf_tun%d",uint64(thrdId))
    fmt.Printf("Tun interface is : %s", intfName)

	tunFd, tun, err := Tun(intfName)
	if err != nil {
		fmt.Println("error: tun:", err)
		return -1, err
	}

    //defer tun.Close()
    tun = tun
    // Enable the tun interface
    cmdEnable := exec.Command("ip", "link", "set", "dev", intfName, "up")
    cmdEnable.Stdout = os.Stdout
    cmdEnable.Stderr = os.Stderr
    err = cmdEnable.Run()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
        return tunFd, err
    }

    // Assign IP Address to the tunnel interface
    cmd := exec.Command("ip", "a", "a", ipAddr, "dev", intfName)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
        return tunFd, err
    }

    go readMode(tun)
	time.Sleep(5 * time.Second)
    //tun.Close()

    for {
    }

    return tunFd, nil
}

func main() {

    thrdId  := uint64(0)
    ipAddr  := "12.0.0.34/24"

    // TunSetup function
    tunFd, err := TunSetup (thrdId, ipAddr)
	if err != nil {
		fmt.Println("error: TunSetup failed.", err)
		return
    }

    fmt.Printf("Tunnel file descriptor is : %d", tunFd);
    // TunCleanup function
    /*err := TunCleanup (intfName, ipAddr)
	if err != nil {
		fmt.Println("error: TunSetup failed.", err)
		return
    */

}
