package main

import (
	"fmt"

	"github.com/gocodecraving/go-dotenv-tutorial/packages/env"
)

func main() {

	envVal := env.Get("BIG_SECRET_KEY")

	fmt.Println(envVal)

}
