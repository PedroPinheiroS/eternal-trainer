package macro

import (
	"time"

	"github.com/PedroPinheiroS/eternal-trainer/email"
	"github.com/PedroPinheiroS/eternal-trainer/env"
	"github.com/go-vgo/robotgo"
)

func Login() {
	LimparCampoEmail()

	robotgo.TypeStr(env.TibiaUser)
	time.Sleep(1 * time.Second)

	robotgo.KeyTap("tab")
	time.Sleep(1 * time.Second)

	robotgo.TypeStr(env.TibiaPassword)
	time.Sleep(1 * time.Second)

	robotgo.KeyTap("enter")

	token, err := email.RetrieveToken()
	if err != nil {
		for i := 0; i < 60; i++ {
			time.Sleep(5 * time.Second)
			token, err = email.RetrieveToken()

			if err == nil {
				break
			}

		}

		if err != nil {
			panic(err.Error())
		}
	}

	robotgo.TypeStr(token)

	time.Sleep(1 * time.Second)

	robotgo.KeyTap("enter")
	time.Sleep(1 * time.Second)

	for i := 0; i < 6; i++ {
		robotgo.KeyTap("down")
		time.Sleep(300 * time.Millisecond)
	}

	robotgo.KeyTap("enter")
}

func Deslogar() {
	robotgo.KeyDown("lctrl")
	robotgo.KeyDown("q")

	time.Sleep(1 * time.Second)

	robotgo.KeyUp("lctrl")
	robotgo.KeyUp("q")

	robotgo.KeyTap("esc")
}

func LimparCampoEmail() {
	robotgo.KeyDown("lctrl")
	robotgo.KeyDown("a")

	time.Sleep(1 * time.Second)

	robotgo.KeyUp("lctrl")
	robotgo.KeyUp("a")

	robotgo.KeyTap("backspace")
}
