package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func task1(n int) {
	if n <= 1 {
		fmt.Println("Число не является простым")
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Printf("Число не является простым, найден делитель: %d\n", i)
			return
		}
	}

	fmt.Println("Число является простым")
}

func task2(a, b int) {
	for b != 0 {
		a, b = b, a%b
	}

	fmt.Println("Результат: НОД =", a)
}

func task3(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println("Шаг", i+1, ":", arr)
	}

	fmt.Println("Результат:", arr)
}

func task4() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func task5(n int) {
	var memo = make(map[int]int)

	fib(n, memo)

	for key, value := range memo {
		fmt.Printf("Для n = %v число фибоначи = %v\n", key, value)
	}
}

func fib(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func task6(n int) {
	reversed := 0
	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	fmt.Println("Результат:", reversed)
}

func task7(levels int) {
	triangle := make([][]int, levels)

	for i := 0; i < levels; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		fmt.Println(triangle[i])
	}
}

func task8(n int) {
	original := n
	reversed := 0

	for n != 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	if original == reversed {
		fmt.Println("Результат: число является палиндромом")
	} else {
		fmt.Println("Результат: число не является палиндромом")
	}
}

func task9(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Результат: массив пустой")
		return
	}

	maximum := arr[0]
	minimum := arr[0]

	for _, value := range arr {
		if value > maximum {
			maximum = value
		}
		if value < minimum {
			minimum = value
		}
	}

	fmt.Printf("Результат: максимум = %d, минимум = %d\n", maximum, minimum)
}

func task10() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 10
	var guess int

	fmt.Println("Я загадал число от 1 до 100. У вас есть", attempts, "попыток, чтобы его угадать.")

	for attempts > 0 {
		fmt.Print("Введите вашу попытку: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Больше!")
		} else if guess > target {
			fmt.Println("Меньше!")
		} else {
			fmt.Println("Поздравляю! Вы угадали число:", target)
			return
		}

		attempts--
		fmt.Println("Осталось попыток:", attempts)
	}

	fmt.Println("Вы исчерпали все попытки. Загаданное число было:", target)
}

func task11(n int) {
	original := n
	digits := 0
	sum := 0

	for temp := n; temp != 0; temp /= 10 {
		digits++
	}

	for n != 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		n /= 10
	}

	if original == sum {
		fmt.Println("Результат: число является числом Армстронга")
	} else {
		fmt.Println("Результат: число не является числом Армстронга")
	}
}

func task12(input string) {
	wordCount := make(map[string]int)

	input = strings.ToLower(input)

	var cleanedInput strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsSpace(char) {
			cleanedInput.WriteRune(char)
		}
	}

	words := strings.Fields(cleanedInput.String())

	for _, word := range words {
		wordCount[word]++
	}

	for key, value := range wordCount {
		fmt.Printf("Для n = '%s' число фибоначи = %v\n", key, value)
	}

	fmt.Println("Количество уникальных слов:", len(wordCount))
}

func task13(lifeTimeSeconds int) {
	board := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	fmt.Println("Начальное состояние:")
	printBoard(board)

	for i := 0; i < lifeTimeSeconds; i++ {
		time.Sleep(1 * time.Second)
		board = updateBoard(board)
		fmt.Printf("Шаг %d:\n", i+1)
		printBoard(board)
	}
}

const (
	rows    = 5
	columns = 5
)

func printBoard(board [][]int) {
	fmt.Println("+---+---+---+---+---+")
	for _, row := range board {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("| X ") // Живая клетка
			} else {
				fmt.Print("|   ") // Мертвая клетка
			}
		}
		fmt.Println("|")
		fmt.Println("+---+---+---+---+---+")
	}
	fmt.Println()
}

func updateBoard(board [][]int) [][]int {
	newBoard := make([][]int, rows)
	for i := range newBoard {
		newBoard[i] = make([]int, columns)
		copy(newBoard[i], board[i])
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			liveNeighbors := countLiveNeighbors(board, i, j)

			if board[i][j] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					newBoard[i][j] = 0
				}
			} else {
				if liveNeighbors == 3 {
					newBoard[i][j] = 1
				}
			}
		}
	}

	return newBoard
}

func countLiveNeighbors(board [][]int, x, y int) int {
	liveCount := 0
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, direction := range directions {
		nx, ny := x+direction[0], y+direction[1]
		if nx >= 0 && nx < rows && ny >= 0 && ny < columns {
			liveCount += board[nx][ny]
		}
	}

	return liveCount
}

func task14(n int) {
	fmt.Println("Результат: цифровой корень =", digitalRoot(n))
}

func digitalRoot(n int) int {
	if n < 10 {
		return n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return digitalRoot(sum)
}

func task15(n int) {
	fmt.Println("Результат: римское число =", integerToRoman(n))
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func displayMenu() {
	fmt.Println("Выберите задачу:")
	fmt.Println("1. Проверка на простоту")
	fmt.Println("2. Наибольший общий делитель (НОД)")
	fmt.Println("3. Сортировка пузырьком")
	fmt.Println("4. Таблица умножения в формате матрицы")
	fmt.Println("5. Фибоначчи с мемоизацией")
	fmt.Println("6. Обратные числа")
	fmt.Println("7. Треугольник Паскаля")
	fmt.Println("8. Число палиндром")
	fmt.Println("9. Нахождение максимума и минимума в массиве")
	fmt.Println("10. Игра 'Угадай число'")
	fmt.Println("11. Числа Армстронга")
	fmt.Println("12. Подсчет слов в строке")
	fmt.Println("13. Игра 'Жизнь'")
	fmt.Println("14. Цифровой корень числа")
	fmt.Println("15. Римские цифры")
	fmt.Println("0. Выход")
}

func main() {
	for {
		displayMenu()
		fmt.Print("Ваш выбор: ")
		var choice int
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 0:
			return
		case 1:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task1(n)
		case 2:
			var a, b int
			fmt.Print("Введите два числа: ")
			fmt.Scan(&a, &b)
			task2(a, b)
		case 3:
			var arrSize int
			fmt.Print("Введите размер массива: ")
			fmt.Scan(&arrSize)
			arr := make([]int, arrSize)
			fmt.Print("Введите элементы массива: ")
			for i := 0; i < arrSize; i++ {
				fmt.Scan(&arr[i])
			}
			task3(arr)
		case 4:
			task4()
		case 5:
			var n int
			fmt.Print("Введите число для Фибоначчи: ")
			fmt.Scan(&n)
			task5(n)
		case 6:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task6(n)
		case 7:
			var levels int
			fmt.Print("Введите уровень треугольника Паскаля: ")
			fmt.Scan(&levels)
			task7(levels)
		case 8:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task8(n)
		case 9:
			var arrSize int
			fmt.Print("Введите размер массива: ")
			fmt.Scan(&arrSize)
			arr := make([]int, arrSize)
			fmt.Print("Введите элементы массива: ")
			for i := 0; i < arrSize; i++ {
				fmt.Scan(&arr[i])
			}
			task9(arr)
		case 10:
			task10()
		case 11:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task11(n)
		case 12:
			fmt.Print("Введите строку: ")
			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			task12(text)
		case 13:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task13(n)
		case 14:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task14(n)
		case 15:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			task15(n)
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}

		fmt.Print("\n\n\n")
	}
}
