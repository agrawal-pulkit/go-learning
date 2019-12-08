package main

import "fmt"

func main() {
	//first way to create map
	colors := map[string]string{
		"red":   "ff0000",
		"black": "000000",
		"white": "ffffff",
	}
	fmt.Println(colors)

	//second way to create map
	var colors2 map[string]string
	fmt.Println(colors2)

	//third way to create map
	colors3 := make(map[string]string)
	colors3["white"] = "ffffff"
	fmt.Println(colors3)

	//built-in delete function for delete a key in map
	delete(colors3, "white")
	fmt.Println(colors3)

	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
