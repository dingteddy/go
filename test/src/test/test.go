package test

import (
	"flag"
	"fmt"
)

var (
	redisServer   = flag.String("s", ":6379", "")
	redisPassword = flag.String("p", "", "")
)

func Test() {
	//flag.Parse()
	fmt.Printf("aaaa:%s %s\n", *redisServer, *redisPassword)
}
