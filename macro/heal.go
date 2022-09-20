package macro

import (
	"sync"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func castUltimateHealing(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		robotgo.TypeStr("exura vita")
		robotgo.KeyTap("enter")
		robotgo.Sleep(env.Heal)
	}
}
