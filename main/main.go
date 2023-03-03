package main

import (
	"fmt"

	"github.com/EnricoPicci/modb/pkgmodb1"
)

func main() {
	fmt.Println("I am the main program in Module A that calls a function of Module B")
	fmt.Println("The function of Module B does not call any other function from external modules")
	pkgmodb1.ApiThatDoesNotCallExternalApis()
}
