package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Итак, найдем НОД номиналов всех монет.
Если необходимая сумма не делится на НОД нацело,
то сразу выводим "-" (мы никак не можем набрать 5 рублей монетками
в 2 и 4 рубля! Монеты в 4 рубля - моя идея.).
Если же необходимая сумма делится на НОД без остатка,
то придется действовать хитро:

Возьмем монету с наибольшим номиналом,
возьмем наибольшее количество монет, так чтобы их
сумма не превышала искомую сумму.

Теперь посчитаем НОД оставшихся монет.
Если остаток искомой суммы от тех денег,
которые мы собрали, не делится на НОД,

то уменьшим количество монет с наибольшим номиналом на 1
и снова проверим, делится ли остаток на НОД.
Если делится - то переходим к следующей монете,

если нет - то опять уменьшаем на 1 количество монет.
Если в какой то момент количество монет стало отрицательным,
то выводим "-" и выходим.

Если вся сумма выразилась в определенном количестве монет разных номиналов,
то выводим "+" и количество монет разного номинала (для хранения количества
монет удобно организовать массив).
*/

func main() {
	inputs := GetInputs()
	nominals, m := inputs[1:len(inputs)-1], inputs[len(inputs)-1]
	start := time.Now()

	combArray := combinationSum(nominals, m)

	// Если массив пустой
	if len(combArray) == 0 {
		fmt.Fprintln(os.Stdout, "IMPOSSIBLE")
		return
	}
	// Если элемент массива[0] пустой
	if len(combArray[0]) == 0 {
		fmt.Fprintln(os.Stdout, "IMPOSSIBLE")
		return
	}

	// Нахождение минимального len(массива)
	min := len(combArray[0])
	for _, v := range combArray {
		if len(v) < min {
			min = len(v)
		}
	}

	fmt.Fprintln(os.Stdout, min)
	fmt.Printf("Script takes : %v\n", time.Since(start))

}

// Рекурсивная функция подбирающая все комбинации слагаемых
// в сумме дающих target
// Возвращает 10 последовательностей
// номиналов в подмассивах в массиве.
// Либо возвращает [] если не может построить ни одной,
//  и [[]] если target == 0

func combinationSum(nominals []int, target int) [][]int {

	// оригинал
	// массив сортируется
	//sort.Ints(nominals)

	// моё изменение
	// массив сортируется и реверсируется
	sort.Sort(sort.Reverse(sort.IntSlice(nominals)))

	var helper func(nominal, share []int, target int)
	var r [][]int
	helper = func(c, share []int, target int) {
		if target == 0 {
			r = append(r, share)
			return
		}

		// оригинал
		// второе условие target меньше первого элемента
		// if len(c) == 0 || target < c[0] {
		//	return
		// }

		// моё изменение
		// второе условие target меньше последнего элемента
		if len(c) == 0 || target < c[len(c)-1] {
			return
		}
		share = share[:len(share):len(share)]

		// здесь функция остановится и вернёт 10 последовательностей
		if len(r) == 10 {
			return
		}

		helper(c, append(share, c[0]), target-c[0])
		helper(c[1:], share, target)
	}
	helper(nominals, []int{}, target)
	return r
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
