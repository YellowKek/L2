package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		cmd := exec.Command("pwd")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("что-то пошло не так")
		}
		path := strings.TrimSpace(string(out))

		fmt.Print(path, "$ ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		commands := strings.Split(line, "|")
		var prevOut io.Reader = nil

		for _, cmdStr := range commands {
			cmdStr = strings.TrimSpace(cmdStr)
			args := strings.Split(cmdStr, " ")

			switch args[0] {
			case "/quit":
				os.Exit(0)
			case "pwd":
				fmt.Println(path)
				prevOut = strings.NewReader(path)
			case "cd":
				if len(args) != 2 || args[1] == "" {
					fmt.Println("Неверный формат")
					continue
				}
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println(err)
				}
			case "ps":
				if len(args) != 1 {
					fmt.Println("Неверный формат")
					continue
				}
				cmd := exec.Command(args[0])
				output, err := cmd.Output()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(string(output))
				}
			case "kill":
				if len(args) < 2 {
					fmt.Println("Неверный формат")
					continue
				}
				cmd := exec.Command(args[0], args[1:]...)
				err := cmd.Start()
				if err != nil {
					fmt.Println(err.Error())
				}
			case "echo":
				cmd := exec.Command(args[0], getEchoArgs(args)...)
				if prevOut != nil {
					cmd.Stdin = prevOut
				}

				output, err := cmd.Output()
				if err != nil {
					fmt.Println(err.Error())
				} else {
					prevOut = strings.NewReader(string(output))
				}
			default:
				cmd := exec.Command(args[0], args[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					fmt.Println("Команда не найдена!")
					fmt.Println(err.Error())
				}
			}
		}
		if prevOut != nil {
			io.Copy(os.Stdout, prevOut)
		}
	}
}

func getEchoArgs(args []string) []string {
	var echoArgs []string
	for _, v := range args[1:] {
		if strings.Contains(v, "$") {
			val := os.Getenv(v[1:])
			if len(val) > 0 && val != "" && val != "\n" {
				echoArgs = append(echoArgs, os.Getenv(v[1:]))
			} else {
				echoArgs = append(echoArgs, v)
			}
		} else {
			echoArgs = append(echoArgs, v)
		}
	}
	return echoArgs
}
