package utils

import (
	"strings"
	"time"

	"github.com/PedroPinheiroS/eternal-trainer/env"
)

func removeSpaces(code string) string {
	return strings.ReplaceAll(code, " ", "")
}

func removeBreakLines(code string) string {
	return strings.ReplaceAll(code, "\n", "")
}

func GetToken(body string) string {
	inicio := strings.Index(body, ":")
	fim := strings.Index(body, "Note")

	return body[inicio+1 : fim]
}

func CleanToken(token string) string {
	token = removeSpaces(token)
	token = removeBreakLines(token)

	return token
}

func StopAllMacros() bool {
	return time.Now().Hour() == env.HoraDeslogar && time.Now().Minute() == env.MinutoDeslogar
}

func InitAllMacros() bool {
	return time.Now().Hour() == env.HoraLogar && time.Now().Minute() == env.MinutoLogar
}
