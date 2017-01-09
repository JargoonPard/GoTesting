package main

import mystuff "github.com/jargoonpard/GoTesting/MyStuff"
import "fmt"

var mystuffAddName = mystuff.AddName

func main() {
	fmt.Println("calling inner")
	inner()
	fmt.Println("calling test")
	test()
	fmt.Println("calling inner")
	inner()
}

func inner() {
	count, namelist := mystuffAddName("Larry")

	fmt.Printf("count: %v", count)
	fmt.Println()
	fmt.Printf("namelist: %v", namelist)
	fmt.Println()
}

func test() {
	oldAddName := mystuffAddName

	defer func() { mystuffAddName = oldAddName }()

	mystuffAddName = func(name string) (int, []string) {
		return 0, nil
	}

	inner()
}
