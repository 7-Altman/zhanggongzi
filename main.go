package main

import (
	"fmt"
	"os"
	"zhanggongzi/gateway"
)

func main()  {
	client := new(gateway.GTClient)
	runErr := client.Run()

	if runErr != nil {

		fmt.Println("something was wrong check your project", runErr.Error())
	} else {
		fmt.Println("xxx")
	}
	os.Exit(1)
}
