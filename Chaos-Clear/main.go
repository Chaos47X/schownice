package claer

import (
	"os"
	"os/exec"
)

func Clear(OS string) {
	if OS == "windows" {
		func() {
			cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	} else {
		func() {
			cmd := exec.Command("clear") //Linux example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	}
}
