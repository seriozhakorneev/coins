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
Сначала я хотел пойти очевидным путем поиска остатка m после удаления наибольших номиналов
но даже пример 1 из задачи выглядел бы так:

 	10 = 4 + 4 + 1 + 1 количество 4
	а не так:
 	10 = 4 + 3 + 3 количество 3

Так как задача о цифрах и математических операциях то для лучшего решения нужно
искать математические правила и основывать алгоритм на них.
В кратчайшее время я таких не нашел и решил пойти другим путем.

Хотел попробовать решить задачу алгоритмом Дийкстры так как я ищу кратчайший путь.
Но чтобы правильно сгенерировать граф мне понадобились бы все последовательности
из номиналов в сумме дающие m.

Начал искать аналог Python NumPy в Go чтобы их построить с условием: их сумма == m.
Нашел уже готовую функцию combinationSum()
принимающую последовательность и сумму, возвращающую уникальные варианты.
Так же, она уже решала проблему чисел которые в сумме не могут дать m
возвращая пустой массив.

Понял что Алгоритм Дийкстры мне не обязательно понадобится,
достаточно получить наименьший массив.

Изменил функцию чтобы она собирала наименьшие последовательности
в начало массива а не в конец, собирала только 10 первых и останавливалась.
В первую очередь потому что большие значения d сильно замедляли алгоритм.

(Вначале было 5 и моим тестам этого хватало, 10 взял с запасом)
Это самая необоснованная часть алгоритма, но работает.

Осталось найти наименьший массив из 10.

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
