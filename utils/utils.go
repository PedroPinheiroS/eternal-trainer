package utils

import "strings"

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
