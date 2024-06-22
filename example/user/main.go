package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/DuvanM9/go_course_sdk/user"
)

func main() {

	userTransport := userSdk.NewHttpClient("http://localhost:8002", "")
	user, err := userTransport.Get("c4c9855b-4d17-4f6a-a920-3b19x4c43f98e")

	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("not found: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal server error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
