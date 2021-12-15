package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func shell(str string) error {
	var cmd *exec.Cmd
	str = strings.TrimSuffix(str, "\n")
	slice := strings.Fields(str)
	switch slice[0] {
	case "cd":
		if len(slice) != 2 {
			return errors.New("неправильный путь")
		}
		return os.Chdir(slice[1])
	case "pwd":
		if len(slice) != 1 {
			return errors.New("лишние аргументы")
		}
		path, _ := os.Getwd()
		fmt.Println(path)
		return nil
	case "echo":
		s := strings.Join(slice[1:], " ")
		_, err := io.WriteString(os.Stdout, s)
		if err != nil {
			return err
		}
		return nil
	case "ps":
		cmd = exec.Command("powershell", "ps")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	case "kill":
		cmd = exec.Command("powershell", "kill", slice[1])
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	case "netcat":
		return netcat(fmt.Sprintf("%s:%s", slice[1], slice[2]))
	case "exit":
		os.Exit(0)
	}
	if !strings.Contains(str, "|") {
		cmd = exec.Command(slice[0], slice[1:]...)
	} else {
		cmd = exec.Command("powershell", `.\` + str)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func netcat(addr string) error {
	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		conn, err = net.Dial("udp", addr)
		if err != nil {
			fmt.Printf("Can't connect to server: %s\n", err)
			return err
		}
	}
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
		return err
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		dir, err := os.Getwd()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		fmt.Printf("%v>", dir)
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		err = shell(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}
