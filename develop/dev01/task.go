package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

// ntpServer - константа, хранящая адрес NTP сервера.
const ntpServer = "pool.ntp.org"

// GetTime - функция для получения времени с NTP сервера.
func GetTime() (*ntp.Response, error) {
	return ntp.Query(ntpServer)
}

// GetTimeNow - функция для получения текущего системного времени.
func GetTimeNow() time.Time {
	return time.Now()
}

func main() {
	// Попытка запроса времени с NTP сервера.
	resp, err := GetTime()
	// Обработка возможных ошибок запроса.
	if err != nil {
		// Вывод ошибки в стандартный поток ошибок и завершение программы с кодом ошибки  1.
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	// Получение текущего системного времени.
	currentTime := GetTimeNow()

	// Вывод текущего системного времени.
	fmt.Printf("current time: %s\n", currentTime)
	// Вывод времени полученного с NTP сервера.
	fmt.Printf("response time: %s\n", resp.Time)
}
