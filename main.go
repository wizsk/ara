package main

import (
	"fmt"
	"os"
	"os/exec"
)

const dataSize = 1024 * 512 // 0.5mb

func main() {
	in := ""
	if isPiped(os.Stdin) {
		data := make([]byte, dataSize)
		i, err := os.Stdin.Read(data)
		if err != nil {
			fmt.Fprintln(os.Stderr, "something unexpected happend :)")
			os.Exit(1)
		}
		in = string(data[:i])
	} else {
		if len(os.Args) < 2 {
			fmt.Fprintf(os.Stderr, "%s: Usages %s <input_text>", os.Args[0], os.Args[0])
			os.Exit(1)
		}
		in = os.Args[1]
	}

	r := removeHarakats(in)
	if !isPiped(os.Stdout) {
		r += "\n"
	}
	fmt.Print(r)

	cmd := exec.Command("sh", "-c", "echo "+removeHarakats(in)+" | wl-copy")
	cmd.Start()
	cmd.Wait()
}

func isPiped(f *os.File) bool {
	fs, _ := os.Stdin.Stat()
	if (fs.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
