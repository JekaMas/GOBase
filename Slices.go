package main

import (
	"fmt"
	"sort"
)

// jekamas: неудобно иметь больше одного main в одном пакете
func runSliceTasks() {
	// jekamas: каждый раз мы меняем этот слайс и передаем дальше. из-за этого трудно понять верен ли каждый отдельный шаг. лучше бы создавать новый слайс перед каждой функцией.
	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("01")
	task01(s)
	fmt.Println(s)

	fmt.Println("02")
	task02(&s)
	fmt.Println(s)

	fmt.Println("03")
	task03(&s)
	fmt.Println(s)

	fmt.Println("04")
	fmt.Println(task04(&s))
	fmt.Println(s)

	fmt.Println("05")
	fmt.Println(task05(&s))
	fmt.Println(s)

	fmt.Println("06")
	fmt.Println(task06(&s, 2))
	fmt.Println(s)

	fmt.Println("07")
	s2 := []int{8, 2, 7, 4, 1, 6}
	fmt.Println(task07(s, s2))

	fmt.Println("08")
	task08(&s, s2)
	fmt.Println(s)

	// jekamas: разберись почему для заданий с 9 по 12 не нужно передавать *[]int, а достаточно []int
	s2 = []int{1, 2, 3, 4}
	fmt.Println("09")
	task09(s2)
	fmt.Println(s2)

	fmt.Println("10")
	s2 = []int{1, 2, 3, 4}
	task10(s2, 2)
	fmt.Println(s2)

	fmt.Println("11")
	s2 = []int{1, 2, 3, 4}
	task11(s2)
	fmt.Println(s2)

	fmt.Println("12")
	s2 = []int{1, 2, 3, 4}
	task12(s2, 2)
	fmt.Println(s2)

	fmt.Println("13")
	s3 := task13(s2)
	fmt.Println(s3)

	fmt.Println("14")
	task14(s3)
	fmt.Println(s3)

	fmt.Println("15")
	task15(s3)
	s4 := []string{"ZAR", "AZs", "ABC", "AVC", "SAR"}
	sort.Strings(s4)
	fmt.Println(s4)
}

// упражнения:
//К каждому элементу []int прибавить 1
func task01(s []int) {
	for i := range s {
		s[i]++
	}
}

//Добавить в конец слайса число 5
func task02(s *[]int) {
	*s = append(*s, 5)
}

//Добавить в начало слайса число 5
func task03(s *[]int) {
	*s = append([]int{5}, *s...)
}

//Взять последнее число слайса, вернуть его пользователю, а из слайса этот элемент удалить
func task04(s *[]int) (num int) {
	num, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return
}

//Взять первое число слайса, вернуть его пользователю, а из слайса этот элемент удалить
func task05(s *[]int) (num int) {
	// если уж проверяешь везьде длину слайса, то и здесь это надо сделать.
	num, *s = (*s)[0], (*s)[1:]
	return
}

//Взять i-е число слайса, вернуть его пользователю, а из слайса этот элемент удалить. Число i передает пользователь в функцию
func task06(s *[]int, i int) (num int) {
	if i >= len(*s) || i < 0 {
		fmt.Println("Error")
		return -1
	}

	num = (*s)[i]
	*s = append((*s)[:i], (*s)[i+1:]...)
	return
}

//Объединить два слайса и вернуть новый со всеми элементами первого и второго
func task07(s1, s2 []int) []int {
	return append(s1, s2...)
}

//Из первого слайса удалить все числа, которые есть во втором
func task08(s1 *[]int, s2 []int) {
	sort.Ints(s2)

	for i := 0; i < len(*s1); i++ {
		if k := sort.SearchInts(s2, (*s1)[i]); k < len(s2) && s2[k] == (*s1)[i] {
			(*s1)[i] = (*s1)[len(*s1)-1]
			*s1 = (*s1)[:len(*s1)-1]
			i--
		}
	}
}

//Сдвинуть все элементы слайса на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним.
func task09(s []int) {
	// jekamas: переиспользуй свой же код.
	task10(s, 1)
}

//Тоже, но сдвиг на заданное пользователем i
func task10(s []int, offset int) {
	offset %= len(s)
	if len(s) == 0 || offset == 0 {
		return
	}

	leftSide := append([]int{}, s[:offset]...)
	rightSide := s[offset:]

	copy(s, rightSide)
	copy(s[len(rightSide):], leftSide)
}

//Сдвиг на 1 вправо
func task11(s []int) {
	task10(s, len(s)-1) // task12(s, 1)
}

//Тоже, но сдвиг на i вправо
func task12(s []int, i int) {
	i %= len(s)
	// jekamas: сдвиг вправо это тот же сдвиг влево, но len(s)-i - посмотри про алгебру часов или модульная арифметика, в курсе Общей алгебры - https://ru.wikipedia.org/wiki/%D0%A1%D1%80%D0%B0%D0%B2%D0%BD%D0%B5%D0%BD%D0%B8%D0%B5_%D0%BF%D0%BE_%D0%BC%D0%BE%D0%B4%D1%83%D0%BB%D1%8E
	i = len(s)-i
	task10(s, i)
}

//Вернуть пользователю копию переданного слайса
// jekamas: всё верно
func task13(s []int) (this []int) {
	return append(this, s...)
}

//В слайсе поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
func task14(s []int) {
	for i := 0; i < len(s)-1; i += 2 {
		s[i], s[i+1] = s[i+1], s[i]
	}
}

//Упорядочить слайс в порядке: прямом, обратном, лексикографическом.
func task15(s []int) {
	sort.Ints(s)
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
