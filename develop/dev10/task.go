package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// Config адреса
type Config struct {
	host    string
	port    string
	timeout time.Duration
}

// GetFlag получение флага
func GetFlag(cfg *Config) {
	flag.DurationVar(&cfg.timeout, "timeout", time.Second*10, "timeout")
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("Ввели не все данные")
		os.Exit(1)
	}
	cfg.host = flag.Args()[0]
	cfg.port = flag.Args()[1]
}

func writeConn(conn net.Conn, ctx context.CancelFunc) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		if sc.Scan() {
			s := sc.Text()
			slice := strings.Split(fmt.Sprintf("% x", s), " ")
			for _, ctrlD := range slice {
				if ctrlD == "04" {
					fmt.Println("telnet клиент закрывается!")
					ctx()
					return
				}
			}
			_, err := conn.Write([]byte(fmt.Sprintln(s)))
			if err != nil {
				fmt.Println("ошибка записи :", err)
				ctx()
				return
			}
		} else {
			fmt.Println("Ошибка чтения stdin")
			ctx()
			return
		}
	}
}

func readConn(conn net.Conn, ctx context.CancelFunc) {
	sc := bufio.NewScanner(conn)
	for {
		if sc.Scan() {
			line := sc.Text()
			_, _ = fmt.Fprintln(os.Stdout, line)
		} else {
			ctx()
			return
		}
	}
}

func main() {
	cfg := new(Config)
	GetFlag(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(cfg.host, cfg.port), cfg.timeout)
	if err != nil {
		log.Fatal(err)
	}
	go writeConn(conn, cancel)
	go readConn(conn, cancel)
	defer conn.Close()
	<-ctx.Done()
}
