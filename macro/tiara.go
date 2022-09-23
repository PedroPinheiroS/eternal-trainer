package macro

import (
	"log"
	"sync"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func rechargeTiara(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		robotgo.KeyTap("6")
		robotgo.Sleep(env.Tiara)

		if !isOnline {
			log.Println("Stopping Food")
			break
		}
	}
}
