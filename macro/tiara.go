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
		log.Println("tiara recharged")
		robotgo.Sleep(env.Tiara)
	}
}
