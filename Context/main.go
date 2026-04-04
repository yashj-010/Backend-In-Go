package main

import "fmt"

func main() {
	fmt.Println("---- Starting Program ----")

	fmt.Println("\n Running ctx_api")
	ctx_api()

	fmt.Println("\n Running ctx_cancel")
	ctx_cancel()

	fmt.Println("\n Running ctx_data")
	ctx_data()

	fmt.Println("\n Running ctx_done")
	ctx_done()

	fmt.Println("\n Running ctx_timeout")
	ctx_timeout()

	fmt.Println("\n Running mini_project")
	mini_project()

	fmt.Println("\n---- Program Finished ----")
}