Markdown
# 🚀 Backend Journey: Python & Go Practice

Репозиторий для отработки фундаментальных навыков бэкенд-разработки, работы с командной строкой (CLI), аргументами, вводом-выводом и структурированием кода на **Python** и **Go**.

---

## 📂 Структура репозитория

```text
.
├── python/
│   └── calculator.py    # Консольный калькулятор (интерактивный + CLI-аргументы)
├── go/
│   └── main.go          # Утилита анализа текстовых файлов (аналог wc)
└── README.md
🛠️ Описание проектов
1. Python CLI Calculator
Консольный калькулятор с поддержкой базовых арифметических операций.

Особенности:

Работает как в интерактивном режиме, так и через флаги командной строки (argparse).

Обработка ошибок (деление на ноль, некорректный ввод).

Запуск:

Bash
# Интерактивный режим
python python/calculator.py

# Передача аргументов через флаги
python python/calculator.py -a 10 -b 5 -op mul
2. Go Text Utility (wc Clone)
Быстрая консольная утилита для подсчета статистики текстового файла.

Особенности:

Использование пакета flag для парсинга аргументов.

Построчное чтение файла через bufio.Scanner и обработка UTF-8 символов.

Запуск:

Bash
# Вся статистика (строки, слова, символы)
go run go/main.go path/to/file.txt

# Только количество строк
go run go/main.go -l path/to/file.txt
⚙️ Стек технологий
Languages: Python, Go (Golang)

Tools & Libraries: argparse (Python), flag, bufio, unicode/utf8 (Go standard library), Git, VS Code

📌 Автор
GitHub: @keenztechnique