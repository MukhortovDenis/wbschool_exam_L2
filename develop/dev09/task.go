package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var output = flag.String("o", "", "Получить файл с вашим именем")

func response(url string) *http.Response {
	transport := new(http.Transport)
	client := &http.Client{Transport: transport}
	r, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return r
}

func write(file string, r *http.Response) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil{
		fmt.Println(err)
		os.Exit(0)
	}
}

func wget(url string, file string){
	r := response(url)
	write(file, r)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1{
		fmt.Println("Вы не ввели адрес страницы")
		return
	}
	url := args[0]
	if *output == ""{
		wget(url, "new_page.html")
		return
	}
	wget(url, *output)
}
