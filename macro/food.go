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
		log.Println("Food ate")
		robotgo.Sleep(30)
	}
}
