package main

import (
	"fmt"
	"time"
)

func main() {

	timeWhenIsCompiled := time.Now().UTC().Unix()
	fmt.Println(timeWhenIsCompiled)

	timeExpiration := time.Now().UTC().Add(60 * time.Second).Unix()
	fmt.Println(timeExpiration)

	if timeWhenIsCompiled > timeExpiration {
		fmt.Println("the ticket is not expired")
	} else {
		fmt.Println("the ticket is expired")
	}

}
