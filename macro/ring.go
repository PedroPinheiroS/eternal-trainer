package macro

import (
	"log"
	"sync"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func rechargeLifeRing(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		robotgo.KeyTap("7")
		robotgo.Sleep(env.LifeRing)

		if !isOnline {
			log.Println("Stopping Food")
			break
		}
	}
}
