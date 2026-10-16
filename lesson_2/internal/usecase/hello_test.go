package usecase_test

import (
	"testing"

	"github.com/alexeevaanst-dev/go_course/lesson_2/internal/usecase"
)

func TestHello(t *testing.T) {
	hello := usecase.NewHello()

	actual := hello.Say()
	expected := "Hello!"

	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
