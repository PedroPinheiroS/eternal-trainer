package macro

import (
	"log"
	"sync"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func castCreateRune(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		robotgo.KeyTap("8")
		robotgo.Sleep(env.Rune)

		if !isOnline {
			log.Println("Stopping Food")
			break
		}
	}
}
