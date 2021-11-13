package main

import (
	router "github.com/Le-MaliX/ACADEMY-GO-Q42021/infrastructure/router"
)

func main() {
	r := router.Routes()

	r.Run()
}
