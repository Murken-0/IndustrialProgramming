package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

func main() {
	task1("1011", 2, 10)
	task2(1, -5, 6)
	task3([]int{5, -2, 1, -4, 3})
	task4([]int{1, 3, 5, 7}, []int{2, 4, 6, 8})
	task5("Hello, world!", "world")
	task6("5 + 10")
	task7("А роза упала на лапу Азора")
	task8(1, 5, 3, 7, 2, 6)

}
