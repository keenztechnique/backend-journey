import argparse
import sys


def calculate(a: float, b: float, operation: str) -> float:
  """Выполняет базовые арифметические операции."""
  if operation == "add":
    return a + b
  elif operation == "sub":
    return a - b
  elif operation == "mul":
    return a * b
  elif operation == "div":
    if b == 0:
      raise ValueError("Деление на ноль невозможно.")
    return a / b
  else:
    raise ValueError(f"Неизвестная операция: {operation}")


def main():
  parser = argparse.ArgumentParser(
      description="Консольный калькулятор на Python"
  )
  parser.add_argument(
      "-a", type=float, help="Первое число", default=None
  )
  parser.add_argument(
      "-b", type=float, help="Второе число", default=None
  )
  parser.add_argument(
      "-op",
      choices=["add", "sub", "mul", "div"],
      help="Операция",
      default=None,
  )

  args = parser.parse_args()

  # Если аргументы не переданы через флаги, запускаем интерактивный режим
  if args.a is None or args.b is None or args.op is None:
    print("--- Интерактивный режим калькулятора ---")
    try:
      a = float(input("Введите первое число: "))
      op = input("Введите операцию (add, sub, mul, div): ").strip()
      b = float(input("Введите второе число: "))
      result = calculate(a, b, op)
      print(f"Результат: {result}")
    except ValueError as e:
      print(f"Ошибка ввода: {e}", file=sys.stderr)
      sys.exit(1)
  else:
    try:
      result = calculate(args.a, args.b, args.op)
      print(f"Результат: {result}")
    except ValueError as e:
      print(f"Ошибка: {e}", file=sys.stderr)
      sys.exit(1)


if __name__ == "__main__":
  main()