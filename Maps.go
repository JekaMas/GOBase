package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"runtime"
	"sort"
	"time"
)

func runMapTasks() {
// Общие комментарии
	// jekamas: Функция должна быть полезна для внешнего мира и не делать всю работу в себе: создавать слайс или текст, инициализировать что-то, затем считать, печатать на экран и ничего не возвращать. Это всё разные задачи и они должны делаться не внутри одной функции.
	// jekamas: Подумай над лучшим названием для функций, отращиющими ответственность функций.

	fmt.Println("01")
	// jekamas: лучше сделать входящим параметром текст в виде строки, а исходящим - готовый map.
	map01()

	fmt.Println("02")
	map02()

	fmt.Println("03")
	map03()

	fmt.Println("04")
	map04()
}

//Есть текст, надо посчитать сколько раз каждое слова встречается.
func map01() {

	m := make(map[string]int)
	// jekamas: я немного поменял строку, проверь верно ли она считает теперь
	text := "Some text, SOME words, some letters. End of text,\nhere is the true end."

	// а если другие знаки препинания или перевод строки, табуляция?
	re := regexp.MustCompile(`[.]*[,]*[ ]`)
	for _, s := range re.Split(text, -1) {
		m[s]++
	}
	fmt.Println(m)
}

//Есть очень большой массив(слайс) целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
func map02() {
	const length = 10000

	slice := make([]int, length)

	memoryBefore := currentMemory()
	// jekamas: не имеет смысла многократно менять rand.Seed
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(length / 2)
	}

	m := make(map[int]struct{}, length/2)
	fmt.Printf("Here: ")

	// jekamas: лучше for range по входящему параметру slice
	for i := 0; i < length; i++ {
		if _, ok := m[slice[i]]; !ok {
			m[slice[i]] = struct{}{}
			fmt.Print(slice[i], " ")
		}
	}
	memoryAfter := currentMemory()
	fmt.Printf("\nLen, memory: %v, %v\n", len(m), memoryAfter-memoryBefore)
}

type InBoth struct {
	inBoth bool
}

//Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
func map03() {
	const length = 10000

	slice1 := make([]int, length)
	slice2 := make([]int, length)

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < length; i++ {
		slice1[i] = rand.Intn(length / 2)
	}

	rand.Seed(int64(time.Now().Nanosecond() + 1))
	for i := 0; i < length; i++ {
		slice2[i] = rand.Intn(length / 2)
	}

	memoryBefore := currentMemory()
	m := make(map[int]InBoth, length/2)

	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			// jekamas: ты тут создавал новый анонимный тип, отличный от InBoth
			m[v] = InBoth{}
		}
	}

	// jekamas: просто любопытная штука: на твои измерения памяти сильно будет влиять вот этот и вообще все вызовы fmt. Из-за этого могут получаться разные данные не от того, что ты поменял алгоритм, а из-за того, что добавил еще пояснений в вывод.
	fmt.Println("InBoth: ")
	for _, v := range slice2 {
		mapCounter := m[v]
		if !mapCounter.inBoth {
			mapCounter.inBoth = true
			fmt.Printf("%v ", v)
		}
	}
	memoryAfter := currentMemory()
	fmt.Printf("\nLen, memory: %v, %v\n", len(m), memoryAfter-memoryBefore)
}

//Сделать Фибоначчи с мемоизацией
func map04() {
	m := map[int]int{
		0: 0,
		1: 1,
	}
	Fibonacci(8, m)

	Fibonacci(10, m)

	s := make([]int, len(m))
	for i := range m {
		s[i] = m[i]
	}
	sort.Ints(s)
	fmt.Println(s)
}

func Fibonacci(n int, m map[int]int) int {
	// jekamas: тут можно проще
	_, ok := m[n]
	if !ok {
		m[n] = Fibonacci(n-2, m) + Fibonacci(n-1, m)
	}
	return m[n]
}

func currentMemory() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
