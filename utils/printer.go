package utils

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func PrintSuccess(msg string) {
	fmt.Println(Green + msg + Reset)
}

func PrintError(msg string) {
	fmt.Println(Red + msg + Reset)
}
