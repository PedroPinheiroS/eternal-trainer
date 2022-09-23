package macro

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/PedroPinheiroS/eternal-trainer/tibiadata"
)

func Online(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(30 * time.Second)
		online, _ := tibiadata.GetOnline()
		if !online {
			isOnline = false
			break
		}
	}
}

func check() {
	online, _ := tibiadata.GetOnline()
	log.Println("Online: " + fmt.Sprint(online))
	if online {
		isOnline = true
	} else {
		isOnline = false
		Login()
	}

	for !isOnline {
		time.Sleep(30 * time.Second)
		online, _ = tibiadata.GetOnline()
		log.Println(online)
		if online {
			isOnline = true
		}
	}
}
