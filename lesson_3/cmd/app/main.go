package main

import (
	"errors"
	"fmt"
	"slices"
)

var ErrUserNotFound = errors.New("user not found")
var ErrDivisionByZero = errors.New("division by zero")

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func removeDuplicates(s []string) []string {
	result := make([]string, 0, len(s))
	for _, val := range s {
		if slices.Contains(result, val) {
			continue
		}
		result = append(result, val)
	}
	return result
}

// Есть слайс рейтингов фильмов: []float64{7.2, 5.1, 8.9, 6.0, 9.1, 4.3, 7.8}.
// Напиши функцию filterByRating(ratings []float64, min float64) []float64, которая возвращает только те рейтинги, которые >= min.
// Используй make с cap для оптимизации.

func filterByRating(ratings []float64, min float64) []float64 {
	if ratings == nil {
		return nil
	}

	filteredRatings := make([]float64, 0, len(ratings))
	for _, val := range ratings {
		if val >= min {
			filteredRatings = append(filteredRatings, val)
		}
	}

	return filteredRatings
}

func dropDuplicates(s []string) []string {
	strokes := make(map[string]struct{}, len(s))
	noDuplicates := make([]string, 0, len(s))
	for _, val := range s {
		if _, ok := strokes[val]; !ok {
			strokes[val] = struct{}{}
			noDuplicates = append(noDuplicates, val)
		}
	}
	return noDuplicates

}

func countChars(s string) map[rune]int {
	countChar := make(map[rune]int, len(s))
	for _, symb := range s {
		countChar[symb] += 1
	}
	delete(countChar, rune(' '))
	return countChar
}

// функции для мапы каталога

// добавляет фильм в каталог, если такой фильм уже есть в каталоге ошибка
func addFilm(catalog map[string]float64, title string, rating float64) error {
	if rating < 0 || rating > 10 {
		return fmt.Errorf("рейтинг должен быть в пределах от 0 до 10")
	}
	if title == "" {
		return fmt.Errorf("пустое название фильма")
	}
	if _, ok := catalog[title]; ok {
		return fmt.Errorf("такой фильм уже есть")
	}

	catalog[title] = rating
	return nil
}

// возвращает слайс названий фильмов из каталога, которые имеют рейтинг выше заданного
func topFilms(catalog map[string]float64, minRating float64) ([]string, error) {
	if minRating < 0 || minRating > 10 {
		return nil, fmt.Errorf("рейтинг должен быть в пределах от 0 до 10")
	}

	topFilms := make([]string, 0, len(catalog))
	for title, rating := range catalog {
		if rating >= minRating {
			topFilms = append(topFilms, title)
		}
	}

	return topFilms, nil
}

func findUserById(id int) (bool, error) {
	if id < 0 {
		return false, fmt.Errorf("некорретный ID: %d", id)
	}
	if id%2 == 0 {
		return true, nil
	}

	return false, nil
}

func activateSubscription(id int, date string) error {

}

func main() {
	var userRating *float64
	rating := 8.5
	userRating = &rating

	if userRating != nil {
		fmt.Printf("%d %f", userRating, *userRating)
	}

}
