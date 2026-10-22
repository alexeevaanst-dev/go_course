package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// sentinel errors

var (
	ErrDivisionByZero = errors.New("деление на ноль")
	ErrLessZero       = errors.New("значение меньше нуля")
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	} else if a < 0 || b < 0 {
		return 0, ErrLessZero
	}
	return a / b, nil
}

// ошибка со структурой
// сначала задаем свою структуру с полями
type ValidationError struct {
	Field   string
	Value   string
	Message string
}

// чтобы мы реализовали интерфейс ошибки необходимо реализовать метод Error() string
// принимаем объект структуры по ссылке, чтобы ее не копировать (если структура маленькая можно и не по ссылке)
func (e ValidationError) Error() string {
	return fmt.Sprintf("Ошибка в поле: %s. Введено значение: %s. Детали ошибки: %s", e.Field, e.Value, e.Message)
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Value:   email,
			Message: "неверный формат email: нет @",
		}
	}
	return nil
}

// домашка

var ErrUserNotFound = errors.New("пользователь не найден")

type ValidateError struct {
	Field string
	Value string
}

func (e ValidateError) Error() string {
	return fmt.Sprintf("поле: %s значение: %s", e.Field, e.Value)
}

func activateSubcription(id int, date string) error {
	if id == 0 {
		return fmt.Errorf("Ошибка при активации подписки у пользователя с ID %d: %w", id, ErrUserNotFound)
	}
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return ValidateError{Field: "Дата", Value: date}
	}

	return nil
}

func main() {

	if err := activateSubcription(1, "202-01-01"); err != nil {
		valErr := ValidateError{}
		switch {
		case errors.As(err, &valErr):
			fmt.Println("ошибка с валидацией поля")
			fmt.Printf("%s", valErr)
		case errors.Is(err, ErrUserNotFound):
			fmt.Println("ошибка с userid")
			fmt.Println(err)

		}
	}

	/* sentinel case
	a := 2
	b := -2
	_, err := divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivisionByZero):
			fmt.Printf("деление на ноль")
		case errors.Is(err, ErrLessZero):
			fmt.Printf("одно из значений меньше нуля")
		}
	}
	*/

}
