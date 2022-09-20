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
		log.Println("ring of healing recharged")
		robotgo.Sleep(env.LifeRing)
	}
}

func rechargeRingOfHealing(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		robotgo.KeyTap("7")
		log.Println("ring of healing recharged")
		robotgo.Sleep(env.RingOfHealing)
	}
}
