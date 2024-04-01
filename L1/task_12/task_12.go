package main

import "fmt"

// contains проверяет, содержится ли элемент e в срезе s.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	var set []string

	for _, s := range sequence {

		if !contains(set, s) {

			set = append(set, s)
		}
	}

	fmt.Println(set)
}
