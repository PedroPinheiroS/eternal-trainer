package macro

import (
	"log"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	isOnline = true
)

func Macro() {

	for {
		time.Sleep(3 * time.Second)
		check()

		robotgo.KeySleep = 1
		var wg sync.WaitGroup

		wg.Add(1)
		go eatFood(&wg)

		wg.Add(1)
		time.Sleep(1 * time.Second)
		go rechargeLifeRing(&wg)

		wg.Add(1)
		time.Sleep(1 * time.Second)
		go rechargeSoftBoots(&wg)

		wg.Add(1)
		time.Sleep(1 * time.Second)
		go castCreateRune(&wg)

		wg.Add(1)
		go Online(&wg)

		wg.Wait()
		log.Println("Routine ended")
	}
}
