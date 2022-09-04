package main

import (
	Ch47 "Chaos-Ax47/Destroy/Chaos-0x47"

	ChAtk "Chaos-Ax47/Destroy/Chaos-ATK"
	CC "Chaos-Ax47/Destroy/Chaos-Clear"

	"os"
	"runtime"
)

//limit threads กันเองไแสัส

func main() {
	CC.Clear(runtime.GOOS)
	Ch47.Axbanner()
	ChAtk.ATKTCP(os.Args[1])

}
