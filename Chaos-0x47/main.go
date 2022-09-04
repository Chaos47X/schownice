package chaos0x47

import (
	"bufio"
	"fmt"
	"os"
)

func Axbanner() {
	fmt.Println("\033[2J")

	file, err := os.Open("Chaos-0x47/BANNER.txt")

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
