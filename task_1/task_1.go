package main

import (
	"bufio"
	"fmt"
	"os"
)

// 4 корабля по 1 клеточке
// 3 корабля по 2 клеточки
// 2 корабля по 3 клеточки
// 1 корабль по 4 клеточкам
// Суммарно занимает: 4 * 1 + 3 * 2 + 2 * 3 + 4 * 1 =  4 + 6 + 6 + 4 = 20;

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

	var result = make([]string, rowCount)

	for i := 0; i < rowCount; i++ {
		var numbers = make([]int, 10)

		// 1 вариант.
		// Для каждого вида корабля подсчитываем количество вхождений в строке.
		// Далее делаем проверку на количество кораблей каждого вида.

		// var boards = map[int]int{}

		// for j := 0; j < 10; j++ {
		// 	fmt.Fscan(in, &numbers[j])

		// 	boards[numbers[j]] = boards[numbers[j]] + 1
		// }

		// if boards[1] == 4 && boards[2] == 3 && boards[3] == 2 && boards[4] == 1 {
		// 	result[i] = "YES"
		// } else {
		// 	result[i] = "NO"
		// }

		// 2 вариант.
		// Суммируем все числа в строке. Если сумма == 20 - то значит все верно.
		var sum int = 0

		for j := 0; j < 10; j++ {
			fmt.Fscan(in, &numbers[j])

			sum += numbers[j]
		}

		if sum == 20 {
			result[i] = "YES"
		} else {
			result[i] = "NO"
		}
	}

	var resultFile, resultFileErr = os.Open(os.Args[2])

	if resultFileErr != nil {
		fmt.Print(err)
	}

	defer resultFile.Close()

	scanner := bufio.NewScanner(resultFile)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < rowCount; i++ {
		scanner.Scan()
		var correct = scanner.Text()

		if correct == result[i] {
			fmt.Println(result[i], "\t", correct, "\t", "OK")
		} else {
			fmt.Println(result[i], "\t", correct, "\t", "BAD")
		}
	}
}
