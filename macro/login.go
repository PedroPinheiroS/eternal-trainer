package macro

import (
	"time"

	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func Login() {
	robotgo.TypeStr(env.TibiaUser)
	time.Sleep(1 * time.Second)

	robotgo.KeyTap("tab")
	time.Sleep(1 * time.Second)

	robotgo.TypeStr(env.TibiaPassword)
	time.Sleep(1 * time.Second)

	robotgo.KeyTap("enter")
}
