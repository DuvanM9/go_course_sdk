package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/DuvanM9/go_course_sdk/course"
)

func main() {

	courseTransport := courseSdk.NewHttpClient("http://localhost:8003", "")
	course, err := courseTransport.Get("4d5344ae-01e1-4f34-b07d-a0906s3209399")

	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("not found: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal server error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(course)
}
