package main

// go build 12_functions/function.go
func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "JavaScript", "Java"
}

func processIt(f func(a int) int) {
	fn(1)
}

func main() {
	fn := func(a int) int {
		return 2
	}

	processIt()
	// result := add(3, 5)
	// fmt.Println(result)
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)
}
