package macro

import (
	"log"
	"sync"

	"github.com/go-vgo/robotgo"
)

func eatFood(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		robotgo.KeyTap("9")
		robotgo.Sleep(10)

		if !isOnline {
			log.Println("Stopping Food")
			break
		}
	}
}
