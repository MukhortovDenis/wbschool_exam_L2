package main

import (
	"github.com/beevik/ntp"
	"testing"
	"fmt"
)

func Test_main(t *testing.T) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time)
}
