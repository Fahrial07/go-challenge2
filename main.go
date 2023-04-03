package main

import "fmt"

func main() {
	i := 0
	j := 0

	for {
		fmt.Println("Nilai i = ", i)
		i++
		if i == 5 {
			for j <= 10 {
				if j == 5 {
					input := "САШАРВО"
					for index, value := range input {
						fmt.Printf("character %U 'C' starts at byte position %d\n", value, index)
					}
				} else {
					fmt.Println("Nilai j = ", j)
				}
				j++
			}
			break
		} else {
			continue
		}
	}

}
