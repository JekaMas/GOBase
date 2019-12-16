package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"runtime"
	"sort"
	"time"
)

func main() {

	fmt.Println("01")
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
	text := "some text, some words, some letters. End of text"

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
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(length / 2)
	}

	m := make(map[int]struct{}, length/2)
	fmt.Printf("Here: ")
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
			m[v] = struct{ inBoth bool }{}
		}
	}

	fmt.Println("InBoth: ")
	for _, v := range slice2 {
		if !m[v].inBoth {
			m[v] = struct{ inBoth bool }{inBoth: true}
			fmt.Printf("%v ", v)
		}
	}
	memoryAfter := currentMemory()
	fmt.Printf("\nLen, memory: %v, %v\n", len(m), memoryAfter-memoryBefore)
}

//Сделать Фибоначчи с мемоизацией
func map04() {
	m := map[int]int{
		0: 1,
		1: 2,
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
	if n == 0 {
		return m[0]
	}
	if n == 1 {
		return m[1]
	}
	if m[n] == 0 {
		m[n] = Fibonacci(n-2, m) + Fibonacci(n-1, m)
	}
	return m[n]
}

func currentMemory() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}
