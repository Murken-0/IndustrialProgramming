package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func task1(number string, sourceBase int, targetBase int) {
	fmt.Println("\nЗадача 1. Перевод чисел из одной системы счисления в другую")
	fmt.Printf("Число: %s\n", number)
	fmt.Printf("Исходная система счисления: %d\n", sourceBase)
	fmt.Printf("Конечная система счисления: %d\n", targetBase)

	decimal, err := strconv.ParseInt(number, sourceBase, 64)
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа:", err)
		return
	}

	result := strconv.FormatInt(decimal, targetBase)
	fmt.Printf("Результат: %s\n", result)
}

func task2(a, b, c float64) {
	fmt.Println("\nЗадача 2. Решение квадратного уравнения")
	fmt.Printf("Коэффициенты: a = %.2f, b = %.2f, c = %.2f\n", a, b, c)

	discriminant := b*b - 4*a*c

	if discriminant >= 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Корни: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else {
		realPart := -b / (2 * a)
		imagPart := math.Sqrt(math.Abs(discriminant)) / (2 * a)
		x1 := complex(realPart, imagPart)
		x2 := complex(realPart, -imagPart)
		fmt.Printf("Корни: x1 = %.2f + %.2fi, x2 = %.2f - %.2fi\n", real(x1), imag(x1), real(x2), imag(x2))
	}
}

func task3(nums []int) {
	fmt.Println("\nЗадача 3. Сортировка чисел по модулю")
	fmt.Printf("Исходный массив: %v\n", nums)

	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})

	fmt.Printf("Отсортированный массив: %v\n", nums)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func task4(arr1, arr2 []int) {
	fmt.Println("\nЗадача 4. Слияние двух отсортированных массивов")
	fmt.Printf("Первый массив: %v\n", arr1)
	fmt.Printf("Второй массив: %v\n", arr2)

	merged := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	fmt.Printf("Объединенный массив: %v\n", merged)
}

func task5(text, substring string) {
	fmt.Println("\nЗадача 5. Нахождение подстроки в строке без использования встроенных функций")
	fmt.Printf("Текст: %s\n", text)
	fmt.Printf("Подстрока: %s\n", substring)

	if len(substring) > len(text) {
		fmt.Println("Невозможно! Подстрока длинее строки")
	}

	for i := 0; i <= len(text)-len(substring); i++ {
		match := true
		for j := 0; j < len(substring); j++ {
			if text[i+j] != substring[j] {
				match = false
				break
			}
		}
		if match {
			fmt.Println("Надено по индексу ", i)
		}
	}
	fmt.Println("Не надено")
}

func task6(str string) {
	fmt.Println("\nЗадача 6. Калькулятор с расширенными операциями")

	splitStr := strings.Split(str, " ")

	num1, err := strconv.ParseFloat(splitStr[0], 64)
	if err != nil {
		fmt.Println("Ошибка при преобразовании первого числа:", err)
		return
	}

	num2, err := strconv.ParseFloat(splitStr[2], 64)
	if err != nil {
		fmt.Println("Ошибка при преобразовании второго числа:", err)
		return
	}

	operator := splitStr[1]

	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Деление на ноль недопустимо!")
			return
		}
		result = num1 / num2
	case "^":
		result = math.Pow(num1, num2)
	case "%":
		if num2 == 0 {
			fmt.Println("Деление на ноль недопустимо!")
			return
		}
		result = math.Mod(num1, num2)
	default:
		fmt.Println("Недопустимая операция!")
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}

func task7(text string) {
	fmt.Println("\nЗадача 7. Проверка палиндрома")
	fmt.Printf("Строка: %s\n", text)

	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")

	runes := []rune(text)

	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-i-1] {
			fmt.Println("Не палиндром")
			return
		}
	}
	fmt.Println("Палиндром")
}

func task8(a1, a2, b1, b2, c1, c2 int) {
	fmt.Println("\nЗадача 8. Нахождение пересечения трех отрезков")
	fmt.Printf("Отрезок 1: [%d, %d]\n", a1, a2)
	fmt.Printf("Отрезок 2: [%d, %d]\n", b1, b2)
	fmt.Printf("Отрезок 3: [%d, %d]\n", c1, c2)

	if (a1 <= b2 && a2 >= b1) || (b1 <= a2 && b2 >= a1) {
		if (a1 <= c2 && a2 >= c1) || (c1 <= a2 && c2 >= a1) && (b1 <= c2 && b2 >= c1) || (c1 <= b2 && c2 >= b1) {
			fmt.Println("Есть")
			return
		}
	}
	fmt.Println("Нету")
}

func task9(sentence string) string {
	fmt.Println("\nЗадача 9. Выбор самого длинного слова в предложении")
	fmt.Printf("Предложение: %s\n", sentence)

	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	fmt.Printf("Самое длинное слово: %s\n", longestWord)
	return longestWord
}

func task10(year int) {
	fmt.Println("\nЗадача 10. Проверка високосного года")
	fmt.Printf("Год: %d\n", year)
	var isLeapYear bool
	if year%4 != 0 {
		isLeapYear = false
	} else if year%100 == 0 && year%400 != 0 {
		isLeapYear = false
	} else {
		isLeapYear = true
	}
	fmt.Printf("Високосный год: %t\n", isLeapYear)
}

func task11(n int) {
	fmt.Println("\nЗадача 11. Числа Фибоначчи до определенного числа")
	fmt.Printf("Число: %d\n", n)

	a, b := 0, 1
	for a <= n {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func task12(start, end int) {
	fmt.Println("\nЗадача 12. Определение простых чисел в диапазоне")
	fmt.Printf("Начало диапазона: %d\n", start)
	fmt.Printf("Конец диапазона: %d\n", end)

	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func task13(start, end int) {
	fmt.Println("\nЗадача 13. Числа Армстронга в заданном диапазоне")
	fmt.Printf("Начало диапазона: %d\n", start)
	fmt.Printf("Конец диапазона: %d\n", end)

	for i := start; i <= end; i++ {
		if isArmstrong(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func isArmstrong(num int) bool {
	numStr := strconv.Itoa(num)
	n := len([]rune(numStr))
	sum := 0
	for i := 0; i < n; i++ {
		sum += int(numStr[i] - '0')
	}
	return math.Pow(float64(sum), float64(n)) == float64(num)
}

func task14(text string) {
	fmt.Println("\nЗадача 14. Реверс строки")
	fmt.Printf("Строка: %s\n", text)

	reversed := ""
	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}

	fmt.Printf("Перевернутая строка: %s\n", reversed)
}

func task15(a, b int) {
	fmt.Println("\nЗадача 15. Нахождение наибольшего общего делителя (НОД)")
	fmt.Printf("Число 1: %d\n", a)
	fmt.Printf("Число 2: %d\n", b)

	for b != 0 {
		a, b = b, a%b
	}
	fmt.Printf("НОД: %d\n", a)
}

func main() {
	task1("1011", 2, 10)
	task2(1, -5, 6)
	task3([]int{5, -2, 1, -4, 3})
	task4([]int{1, 3, 5, 7}, []int{2, 4, 6, 8})
	task5("Hello, world!", "world")
	task6("5 + 10")
	task7("А роза упала на лапу Азора")
	task8(1, 5, 3, 7, 2, 6)
	task9("Это, предложение с, разными словами!")
	task10(2024)
	task11(10)
	task12(2, 20)
	task13(1, 1000)
	task14("Hello, world!")
	task15(24, 36)
}
