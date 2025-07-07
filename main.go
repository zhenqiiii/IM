package main

import "github.com/zhenqiiii/shopping_system/router"

func main() {
	r := router.Router()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
