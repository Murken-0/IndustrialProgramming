package main

import (
	"fmt"
	"math"
)

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

func task1(height, base float64) {
	square := triangleArea(base, height)
	fmt.Printf("Площадь треугольника: %f\n", square)
}

func sortArray(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return arr
}

func task2(arr []int) {
	fmt.Println("Исходный массив:", arr)

	sortArr := sortArray(arr)

	fmt.Println("Отсортированный массив:", sortArr)
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sum += i * i
		}
	}
	return sum
}

func task3(num int) {
	fmt.Printf("Сумма квадратов всех чётных чисел от 1 до %d: %d\n", num, sumOfSquares(num))
}

func isPalindrome(s string) bool {
	length := len([]rune(s))
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func task4(input string) {
	if isPalindrome(input) {
		fmt.Printf("Строка %s является палиндромом\n", input)
	} else {
		fmt.Printf("Строка %s не является палиндромом\n", input)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func task5(n int) {
	if isPrime(n) {
		fmt.Printf("Число %d является простым\n", n)
	} else {
		fmt.Printf("Число %d не является простым\n", n)
	}
}

func generatePrimes(limit int) []int {
	var arr []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			arr = append(arr, i)
		}
	}
	return arr
}

func task6(n int) {
	result := generatePrimes(n)
	fmt.Printf("Простые числа до %d: %d\n", n, result)
}

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		bit := n % 2
		binary = fmt.Sprintf("%d%s", bit, binary)
		n /= 2
	}
	return binary
}

func task7(n int) {
	result := toBinary(n)
	fmt.Printf("Число %d в двоичной системе: %s\n", n, result)
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := math.MinInt

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

func task8(arr []int) {
	fmt.Println("Исходный массив:", arr)

	max := findMax(arr)

	fmt.Println("Максимальный элемент: ", max)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func task9(a int, b int) {
	result := gcd(a, b)
	fmt.Printf("НОД(%d, %d) = %d\n", a, b, result)
}

func sumArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func task10(arr []int) {
	fmt.Println("Исходный массив:", arr)

	sum := sumArray(arr)

	fmt.Println("Сумма элементов массива: ", sum)
}

func main() {
	task1(2.5, 5)
	task2([]int{64, 34, 25, 12, 22, 11, 90})
	task3(15)
	task4("sas")
	task5(25)
	task6(52)
	task7(25)
	task8([]int{64, 34, 25, 12, 22, 11, 90})
	task9(12, 21)
	task10([]int{64, 34, 25, 12, 22, 11, 90})
}
