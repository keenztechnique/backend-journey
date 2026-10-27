package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	// Определение флагов командной строки
	showLines := flag.Bool("l", false, "Показать количество строк")
	showWords := flag.Bool("w", false, "Показать количество слов")
	showChars := flag.Bool("c", false, "Показать количество символов")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Использование: go run main.go [флаги] <путь_к_файлу>")
		fmt.Println("Флаги: -l (строки), -w (слова), -c (символы)")
		os.Exit(1)
	}

	filePath := args[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при открытии файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines, words, chars int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
		line := scanner.Text()
		chars += int64(utf8.RuneCountInString(line) + 1) // +1 учитывает перевод строки

		// Подсчет слов через Scanner (по умолчанию разбивает по пробелам)
		wordScanner := bufio.NewScanner(file) // упрощенный подсчет слов по строке
		_ = wordScanner
		// Альтернативный подсчет слов в строке:
		inWord := false
		for _, r := range line {
			if r == ' ' || r == '\t' || r == '\n' {
				inWord = false
			} else if !inWord {
				inWord = true
				words++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}

	// Если флаги не заданы, выводим всю статистику
	all := !*showLines && !*showWords && !*showChars

	if all || *showLines {
		fmt.Printf("Строк: %d\n", lines)
	}
	if all || *showWords {
		fmt.Printf("Слов: %d\n", words)
	}
	if all || *showChars {
		fmt.Printf("Символов: %d\n", chars)
	}
}