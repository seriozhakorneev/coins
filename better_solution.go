package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
алгоритм Dynamic Programming Bottom Up  O(m * d)

nominals - 1, 3, 4; m - 10;

				сразу минимальные варианты
dp[0] = 0;
dp[1] = 1; nominals(1)
dp[2] = 2; 2 - nominals(1) == 1 (nominals(1) + dp[1])
dp[3] = 1; nominals(3)
dp[4] = 1; nominals(4)
dp[5] = 2; 5 - nominals(1) == 4 (nominals(1) + dp[4])
dp[6] = 2; 6 - nominals(3) == 3 (nominals(3) + dp[3])
dp[7] = 2; 7 - nominals(3) == 4 (nominals(3) + dp[4])
dp[8] = 2; 8 - nominals(4) == 4 (nominals(4) + dp[4])
dp[9] = 3; 9 - nominals(1) == 8 (nominals(1) + dp[8])
dp[10] = 3; 10 - nominals(3) == 7 (nominals(3) + dp[7])

	dp[m] = 3
*/

func main() {
	inputs := GetInputs()
	d, nominals, m := inputs[0], inputs[1:len(inputs)-1], inputs[len(inputs)-1]

	if m < 1 || m > powInt(10, 4) {
		fmt.Fprintln(os.Stdout, "Не проходит ограничения", m)
		return
	}
	if d < 1 || d > powInt(10, 2) {
		fmt.Fprintln(os.Stdout, "Не проходит ограничения", d)
		return
	}

	start := time.Now()

	// массив dpArray размером m + 1
	// dpArray[0] = 0; dpArray[1:] = m + 1, m + 1,....m + 1
	dpArray := make([]int, m+1)
	dpArray[0] = 0
	for i := 1; i < len(dpArray); i++ {
		dpArray[i] = m + 1
	}

	// На каждое натуральное число от 1 до m
	// на каждое н.число номиналов раз
	// если из номиналов можно получить н.число
	for i := 1; i <= m; i++ {
		for _, n := range nominals {
			if i-n >= 0 {
				// dpArray[i-n]+1
				// н.число - номинал + 1(шаг до суммы) == минимальное кол-во шагов
				dpArray[i] = min(dpArray[i], dpArray[i-n]+1)
			}
		}
	}
	// если не превышает сумму
	if dpArray[m] != m+1 {
		fmt.Fprintln(os.Stdout, dpArray[m])
	} else {
		fmt.Fprintln(os.Stdout, "IMPOSSIBLE")
	}

	fmt.Printf("Script takes : %v\n", time.Since(start))
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// Собираю инпуты
func GetInputs() []int {
	var str string
	var inputsArray []int

	scanner := bufio.NewScanner(os.Stdin)

	for i := 1; i < 4; i++ {
		fmt.Printf("Input %v: ", i)

		scanner.Scan()
		input := scanner.Text()
		str += " " + input
	}
	// Конвертирую строки в инт
	inputsArray, err := parseInts(str)
	if err != nil {
		panic(err)
	}
	return inputsArray
}

func parseInts(s string) ([]int, error) {
	var (
		fields = strings.Fields(s)
		ints   = make([]int, len(fields))
		err    error
	)
	for i, f := range fields {
		ints[i], err = strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
