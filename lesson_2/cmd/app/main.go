package main

import (
	"fmt"

	"github.com/alexeevaanst-dev/go_course/lesson_2/internal/usecase"
	"github.com/google/uuid"
	_ "go.uber.org/automaxprocs"
)

func main() {
	hello := usecase.NewHello()
	fmt.Println(hello.Say())

	fmt.Println("UUID:", uuid.New().String())
}
