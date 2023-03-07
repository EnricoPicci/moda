package main

/****************************************************************************************************************
*******************              "moda" depends only on a package of "modb"       *******************************
****************************************************************************************************************/

import (
	"fmt"

	"github.com/EnricoPicci/modb/pkgmodb1"
)

func main() {
	fmt.Println("I am the main program in Module A that calls a function of Module B")
	fmt.Println("The function of Module B does not call any function external to Module B")
	pkgmodb1.ApiThatDoesNotCallExternalApis()
}

/****************************************************************************************************************
*******************       "moda" depends on a package of "modb" which depends on a package of "modc"     ********
****************************************************************************************************************/

// import (
// 	"fmt"

// 	"github.com/EnricoPicci/modb/pkgmodb2"
// )

// func main() {
// 	fmt.Println("I am the main program in Module A that calls a function of Module B")
// 	fmt.Println("The function of Module B calls a function from external module Module C")
// 	pkgmodb2.ApiThatCallsModuleCApiThatCallRemainsInteral()
// }

/***********************************************************************************************************************************************
********     "moda" depends on a package of "modb" which depends on a package of "modc" which depends on a package of "modd"     ***************
***********************************************************************************************************************************************/

// import (
// 	"fmt"

// 	"github.com/EnricoPicci/modb/pkgmodb3"
// )

// func main() {
// 	fmt.Println("I am the main program in Module A that calls a function of Module B")
// 	fmt.Println("The function of Module B calls a function from external module Module C which calls a function from Module D")
// 	pkgmodb3.ApiThatCallsModuleCApiThatCallsModuleDApi()
// }
