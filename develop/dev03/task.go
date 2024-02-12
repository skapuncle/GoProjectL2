package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Флаги командной строки
	columnNum := flag.Int("k", 0, "column number to sort")
	sortNumeric := flag.Bool("n", false, "sort numerically")
	reverse := flag.Bool("r", false, "sort in reverse order")
	unique := flag.Bool("u", false, "do not print duplicate lines")

	monthSort := flag.Bool("M", false, "sort by month name")
	ignoreTrailingSpace := flag.Bool("b", false, "ignore trailing whitespace")
	checkSorted := flag.Bool("c", false, "check if the data is sorted")
	suffixSort := flag.Bool("h", false, "sort numeric values with suffix")

	flag.Parse()

	// Проверка и обработка флагов
	if *columnNum < 0 {
		log.Fatal("invalid column number")
	}

	if *suffixSort && !*sortNumeric {
		log.Fatal("the -h flag can only be used with the -n flag")
	}

	if *checkSorted {
		if err := checkDataSorted(flag.Arg(0), *ignoreTrailingSpace); err != nil {
			log.Fatal(err)
		}
		return
	}

	// Чтение файла
	lines, err := readLines(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	// Сортировка строк
	sortMethod := sort.Strings
	if *sortNumeric {
		sortMethod = sortSliceNumeric
	}
	if *monthSort {
		sortMethod = sortSliceByMonth
	}
	if *suffixSort {
		sortMethod = sortSliceWithSuffix
	}
	sortMethod(lines)

	// Удаление повторяющихся строк
	if *unique {
		lines = removeDuplicates(lines)
	}

	// Обратный порядок
	if *reverse {
		reverseSlice(lines)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Функция для чтения строк из файла
func readLines(filename string) ([]string, error) {
	file, err := os.Open("/home/safetyduck/WBL2/develop/dev03/filename")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		lines = append(lines, strings.TrimRight(line, "\n"))
	}

	return lines, nil
}

// Функция для сортировки строк с числовыми значениями
func sortSliceNumeric(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		num1, err1 := strconv.Atoi(arr[i])
		num2, err2 := strconv.Atoi(arr[j])

		if err1 == nil && err2 == nil {
			return num1 < num2
		}

		return arr[i] < arr[j]
	})
}

// Функция для сортировки строк по названию месяца
func sortSliceByMonth(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		date1, err1 := time.Parse("Jan", arr[i])
		date2, err2 := time.Parse("Jan", arr[j])

		if err1 == nil && err2 == nil {
			return date1.Before(date2)
		}

		return arr[i] < arr[j]
	})
}

// Функция для сортировки строк с учетом суффиксов
func sortSliceWithSuffix(arr []string) {
	sort.Slice(arr, func(i, j int) bool {
		num1, suffix1 := parseNumberWithSuffix(arr[i])
		num2, suffix2 := parseNumberWithSuffix(arr[j])

		if num1 == num2 {
			return suffix1 < suffix2
		}

		return num1 < num2
	})
}

// Функция для разбора числа с суффиксом
func parseNumberWithSuffix(str string) (int, string) {
	suffix := ""
	numStr := str

	for i := len(str) - 1; i >= 0; i-- {
		if !isDigit(str[i]) {
			suffix = str[i+1:]
			numStr = str[:i+1]
			break
		}
	}

	num, _ := strconv.Atoi(numStr)
	return num, suffix
}

// Функция для проверки, является ли символ цифрой
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Функция для удаления повторяющихся строк
func removeDuplicates(arr []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0)

	for _, str := range arr {
		if _, ok := seen[str]; !ok {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}

// Функция для обращения порядка элементов в срезе
func reverseSlice(arr []string) {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// Функция для проверки отсортированности данных
func checkDataSorted(filename string, ignoreTrailingSpace bool) error {
	lines, err := readLines(filename)
	if err != nil {
		return err
	}

	for i := 1; i < len(lines); i++ {
		prevLine := lines[i-1]
		currLine := lines[i]

		if ignoreTrailingSpace {
			prevLine = strings.TrimRight(prevLine, " ")
			currLine = strings.TrimRight(currLine, " ")
		}

		if currLine < prevLine {
			return fmt.Errorf("data is not sorted")
		}
	}

	fmt.Println("data is sorted")

	return nil
}
