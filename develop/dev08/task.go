package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		var cmd *exec.Cmd

		if len(parts) == 1 {
			args := strings.Fields(line)
			if len(args) == 0 {
				continue
			}

			switch args[0] {
			case "cd":
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println("Error:", err)
				}

			case "pwd":
				dir, _ := os.Getwd()
				fmt.Println(dir)

			case "echo":
				fmt.Println(strings.Join(args[1:], " "))

			case "kill":
				fmt.Println("Killing", args[1])

			case "ps":
				fmt.Println("Process list")

			default:
				cmd = exec.Command(args[0], args[1:]...)
			}
		} else {
			cmd = exec.Command("bash", "-c", line)
		}

		if cmd != nil {
			err := runCmd(cmd, os.Stdout)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err())
	}
}

func runCmd(cmd *exec.Cmd, outStream *os.File) error {
	cmd.Stdout = outStream
	cmd.Stderr = outStream

	return cmd.Run()
}
