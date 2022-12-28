package helper

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateUserID() string {
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("user-%s", id)
}

func GenerateGoodID() string {
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("good-%s", id)
}