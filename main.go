package main

import (
	"command-line-app/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("--- Starting ---")
	application := app.Create()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
