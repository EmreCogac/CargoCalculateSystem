package main

import "Cargo/Cargo/router"

func main() {

	r := router.SetupRouter()

	r.Run(":8080")
}
