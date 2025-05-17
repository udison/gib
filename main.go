package main

import (
	"fmt"
	"io"
	"os/exec"
)

func WriteToClipboard(content string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		// Fallback to xsel
		cmd = exec.Command("xsel", "--clipboard", "--input")
		in, err = cmd.StdinPipe()
		if err != nil {
			return err
		}
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	_, err = io.WriteString(in, content)
	if err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

func main() {
	fmt.Println("Hello, World!")

	result := "ls -la"
	WriteToClipboard(result)
}
