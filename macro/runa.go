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
		log.Println("gfb done")
		robotgo.Sleep(env.Rune)
	}
}
