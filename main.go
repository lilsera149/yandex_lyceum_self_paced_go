package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Чтобы числа были разными при каждом запуске
	rand.Seed(time.Now().UnixNano())

	// Загадываем число от 1 до 100
	secret := rand.Intn(100) + 1
	attempts := 0

	fmt.Println("🎮 Я загадал число от 1 до 100. Попробуй угадать!")

	for {
		fmt.Print("Твой вариант: ")

		var input string
		fmt.Scanln(&input)

		// Убираем лишние пробелы и переводим в нижний регистр на случай «выход»
		cleanInput := strings.TrimSpace(strings.ToLower(input))
		if cleanInput == "выход" || cleanInput == "quit" {
			fmt.Println("Ладно, выходим. Загаданное число было:", secret)
			return
		}

		num, err := strconv.Atoi(cleanInput)
		if err != nil {
			fmt.Println("❌ Это не число! Напиши число или «выход».")
			continue
		}

		attempts++

		if num < secret {
			fmt.Println("📉 Слишком мало, пробуй выше!")
		} else if num > secret {
			fmt.Println("📈 Слишком много, пробуй ниже!")
		} else {
			fmt.Printf("🎉 Ура! Ты угадал число %d за %d попыток!\n", secret, attempts)
			break
		}
	}
}
