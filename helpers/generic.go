package helpers

import (
	"fmt"
	"os"
)

func Must(err error) {
	if err != nil {
		fmt.Println("error: ", err.Error())
		os.Exit(1)
	}
}
