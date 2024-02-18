package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var in *bufio.Reader
	var inputFile, err = os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	defer inputFile.Close()

	in = bufio.NewReader(inputFile)

	var rowCount int
	fmt.Fscan(in, &rowCount)
	fmt.Fscanln(in)

	var result = make([]string, rowCount)

	for i := 0; i < rowCount; i++ {
		var request string
		fmt.Fscanln(in, &request)

		/*
			Разбираем строку с помощью регулярного выражения, которое соответствует номеру 1 или 2 типа
			и выполянем замену соответствия на пустую строку.
			Если в результате прохода по всей строке остается пустая строка - то все номера в строке правильных форматов
			Иначе - есть номера с некорректными номерами.
		*/
		matched, _ := regexp.Compile(`\D{1}\d{1}\d{1}\D{1}\D{1}|\D{1}\d{1}\D{1}\D{1}`)
		res := matched.ReplaceAllString(request, "")

		if res != "" {
			result[i] = "-"
			continue
		}

		result[i] = strings.Join(matched.FindAllString(request, -1), " ")
	}

	var resultFile, resultFileErr = os.Open(os.Args[2])

	if resultFileErr != nil {
		fmt.Print(err)
	}

	defer resultFile.Close()

	scanner := bufio.NewScanner(resultFile)
	scanner.Split(bufio.ScanLines)

	for i := 0; i < rowCount; i++ {
		scanner.Scan()
		var correct = scanner.Text()

		fmt.Println("Result: ", result[i])
		fmt.Println("Answer: ", correct)
		fmt.Println()
	}
}
