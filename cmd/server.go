package main

import (
	"fmt"
	"github.com/lynndotconfig/gkgo2/pkg/boostrap"
)


func main() {
	err := boostrap.LoadConf()
	if err != nil {
		fmt.Printf("start fail, err=%+v", err)
	}
	// conf := boostrap.GetConf()
	// fmt.Printf("conf is=%s",conf)
}
