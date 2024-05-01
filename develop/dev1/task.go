package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal(err.Error())
	}

	curTime := time.Now().Add(response.ClockOffset)
	fmt.Println(curTime)
}
