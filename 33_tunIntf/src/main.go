package main

import (
	"fmt"
	"net"
	//"errors"
	"strings"
	"os/exec"
)

func TunCleanup() error {

    ifaces, err := net.Interfaces()
    if err != nil {
        fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
        return err
    }

    for _, i := range ifaces {
	exists := strings.Contains(i.Name, "upf_tun")
	if exists {
		fmt.Println("tun already exists : ", i.Name)
		cmd := exec.Command("ip", "tuntap", "del",  "mode", "tun", "dev", i.Name)
		err = cmd.Run()
		if err != nil {
			fmt.Print(fmt.Errorf("Exec command failed : %+v\n", err.Error()))
			return err
		}
	}
    }

    return nil
}

func main() {

	err := TunCleanup()
	if err != nil {
		fmt.Print(fmt.Errorf("Exec command failed : %+v\n", err.Error()))
		return
	}
}
