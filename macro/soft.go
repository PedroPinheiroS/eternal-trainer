package macro

import (
	"log"
	"sync"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func rechargeSoftBoots(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		robotgo.KeyTap("0")
		log.Println("soft recharged")
		robotgo.Sleep(env.SoftBoots)
	}
}
