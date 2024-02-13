package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите URL в качестве аргумента")
		return
	}

	url := os.Args[1]
	filename := filepath.Base(url)

	err := download(url, filename)
	if err != nil {
		fmt.Printf("Ошибка при скачивании файла: %s\n", err.Error())
		return
	}

	fmt.Println("Файл успешно скачан")
}

func download(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
