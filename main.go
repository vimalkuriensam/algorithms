/*
Copyright © 2024 vimalkuriensam1992@gmail.com
*/
package main

import (
	"fmt"
	"time"

	"github.com/vimalkuriensam/algorithms/internals/adaptors/controllers"
)

func main() {
	// cmd.Execute()
	t := time.Now()
	ad := controllers.Adaptor{}
	fmt.Println(ad.MyPow(2.00000, 10))
	// fmt.Println(result)
	fmt.Println(time.Since(t))
}
