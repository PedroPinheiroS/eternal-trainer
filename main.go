package main

import (
	"time"

	"github.com/PedroPinheiroS/eternal-trainer/email"
	"github.com/PedroPinheiroS/eternal-trainer/macro"
	"github.com/go-vgo/robotgo"
)

func main() {
	time.Sleep(2 * time.Second)

	macro.Login()

	token, err := email.RetrieveToken()
	if err != nil {
		for i := 0; i < 60; i++ {
			time.Sleep(5 * time.Second)
			token, err = email.RetrieveToken()

			if err == nil {
				break
			}

		}

		if err != nil {
			panic(err.Error())
		}
	}

	robotgo.TypeStr(token)

	time.Sleep(1 * time.Second)

	robotgo.KeyTap("enter")
	time.Sleep(1 * time.Second)

	for i := 0; i < 6; i++ {
		robotgo.KeyTap("down")
		time.Sleep(100 * time.Millisecond)
	}

	robotgo.KeyTap("enter")
}
