package main

import (
	"os"
	"time"

	"learn-go-with-tests/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
