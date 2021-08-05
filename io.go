package main

import "fmt"

const ()

func GetInput(question string) string {
	input := ""
	fmt.Printf("%v\033[36m", question)
	fmt.Scanln(&input)
	return input
}
func ResetStyle() {
	fmt.Print("\033[0m")
}
func CleanLine(line int32) {
	var i int32
	for i = 0; i < line; i++ {
		fmt.Print("\033[1F\033[K")
	}
}
func PrintSucess(msg string) {
	fmt.Printf("\033[92m%v\033[0m", msg)
}
func PrintWarning(msg string) {
	fmt.Printf("\033[91m%v\033[0m", msg)
}
func PrintBold(msg string) {
	fmt.Printf("\033[1m%v\033[0m", msg)
}
func PrintUnderline(msg string) {
	fmt.Printf("\033[4m%v\033[0m", msg)
}
func Bold(msg string) string {
	return fmt.Sprintf("\033[1m%v\033[0m", msg)
}
func Yellow(msg string) string {
	return fmt.Sprintf("\033[93m%v\033[0m", msg)
}
