package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func task1(num int) {
	fmt.Println("\nЗадача 1. Сумма цифр")
	fmt.Println("Число: ", num)

	str := strconv.Itoa(num)
	digits := []rune(str)
	sum := 0

	sum += int(digits[0] - '0')
	sum += int(digits[1] - '0')
	sum += int(digits[2] - '0')
	sum += int(digits[3] - '0')

	fmt.Println("Сумма цифр: ", sum)
}

func task2(temp int) {
	fmt.Println("\nЗадача 2. Цельсии в Фаренгейты и обратно")
	fmt.Println("Температура по Цельсию: ", temp)
	far := (float32(temp) * 9 / 5) + 32
	fmt.Println("Температура по Фаренгейту: ", far)
	cel := (far - 32) * 5 / 9
	fmt.Println("Снова температура по Цельсию: ", cel)
}

func task3(arr []int) {
	fmt.Println("\nЗадача 3. Удвоенный массив")
	fmt.Println("Исходный массив: ", arr)

	doubledArr := make([]int, len(arr))
	doubledArr[0] = arr[0] * 2
	doubledArr[1] = arr[1] * 2
	doubledArr[2] = arr[2] * 2
	doubledArr[3] = arr[3] * 2

	fmt.Println("Удвоенный массив: ", doubledArr)
}

func task4(arr []string) {
	fmt.Println("\nЗадача 4. Объединение строк")
	fmt.Println("Исходный массив строк: ", arr)

	str := strings.Join(arr, " ")

	fmt.Println("Объединенная строка: ", str)
}

func task5(x1, y1, x2, y2 int) {
	fmt.Println("\nЗадача 5. Расстояние между точками")
	fmt.Printf("Координаты точки 1: x = %d, y = %d\n", x1, y1)
	fmt.Printf("Координаты точки 2: x = %d, y = %d\n", x2, y2)

	distance := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))

	fmt.Println("Расстояние: ", distance)
}

func task6(num int) {
	fmt.Println("\nЗадача 6. Четность/нечетность")
	fmt.Println("Число: ", num)

	if num%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func task7(year int) {
	fmt.Println("\nЗадача 7. Високосный год")
	fmt.Println("Год: ", year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}
}

func task8(a, b, c int) {
	fmt.Println("\nЗадача 8. Наибольшее из трех чисел")
	fmt.Println("Введенные числа: ", a, b, c)

	maxNum := a
	if b > maxNum {
		maxNum = b
	}
	if c > maxNum {
		maxNum = c
	}

	fmt.Println("Наибольшее число: ", maxNum)
}

func task9(age int) {
	fmt.Println("\nЗадача 9. Определение возрастной группы")
	fmt.Println("Возраст:", age)

	// Возрастные рамки:
	// - Ребенок: 0-12 лет
	// - Подросток: 13-17 лет
	// - Взрослый: 18-64 лет
	// - Пожилой: 65 лет и старше

	var group string
	if age < 13 {
		group = "Ребенок"
	} else if age < 18 {
		group = "Подросток"
	} else if age < 65 {
		group = "Взрослый"
	} else {
		group = "Пожилой"
	}

	fmt.Println("Возрастная группа:", group)
}

func task10(num int) {
	fmt.Println("\nЗадача 10. Проверка делимости на 3 и 5")
	fmt.Println("Число:", num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("Делится одновременно на 3 и 5")
	} else {
		fmt.Println("Не делится одновременно на 3 и 5")
	}
}

func task11(n int) {
	fmt.Println("\nЗадача 11. Вычисление факториала")
	fmt.Println("Число:", n)

	if n < 0 {
		fmt.Println("Факториал не определен для отрицательных чисел")
		return
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Println("Факториал:", factorial)
}

func task12(n int) {
	fmt.Println("\nЗадача 12. Вывод первых n чисел Фибоначчи")
	fmt.Println("Количество чисел:", n)

	if n <= 0 {
		fmt.Println("Некорректное количество чисел")
		return
	}

	a, b := 0, 1
	fmt.Print("Первые ", n, " чисел Фибоначчи: ")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func task13(arr []int) {
	fmt.Println("\nЗадача 13. Переворачивание массива чисел")
	fmt.Println("Исходный массив:", arr)

	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	fmt.Println("Перевернутый массив:", arr)
}

func task14(num int) {
	fmt.Println("\nЗадача 14. Вывод простых чисел до заданного")
	fmt.Println("Число:", num)

	if num <= 1 {
		fmt.Println("Нет простых чисел до 1")
		return
	}

	fmt.Print("Простые числа до ", num, ": ")
	for i := 2; i <= num; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func task15(arr []int) {
	fmt.Println("\nЗадача 15. Cумма чисел в массиве")
	fmt.Println("Массив:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println("Сумма чисел:", sum)
}

func main() {
	task1(1234)
	task2(23)
	task3([]int{1, 2, 3, 4})
	task4([]string{"Hello", "World!"})
	task5(1, 1, 4, 5)
	task6(5)
	task7(2020)
	task8(1, 2, 4)
	task9(15)
	task10(15)
	task11(5)
	task12(7)
	task13([]int{1, 2, 3, 4, 5})
	task14(20)
	task15([]int{1, 2, 3, 4, 5})
}
