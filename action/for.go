package action

import (
	"fmt"
	"time"
)

func DieFor() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}
}

func WhileFor() {
	sum := 0
	i := 0
	for {
		i++
		if i >= 100 {
			break
		}
		sum += i
	}

	fmt.Println(sum, i)
}

func DoWhileFor() {
	number := 0
	for {
		number++
		if number > 10 {
			break
		}

		fmt.Println(number)
	}
}

func ArrayFor() {
	array := [3]string{
		"WT", "XMM", "XLH",
	}

	for index, value := range array {
		fmt.Println(index, value)
	}
}

func SliceFor() {
	slices := []string{
		"WT", "XMM", "XLH",
	}

	for index, value := range slices {
		fmt.Println(index, value)
	}
}

func MapFor() {
	maps := map[string]int{
		"WT":  100,
		"XMM": 99,
		"XLH": 98,
	}

	for index, value := range maps {
		fmt.Println(index, value)
	}
}

func ContinueAndBreakFor() {
	maps := map[string]int{
		"WT":  100,
		"XMM": 99,
		"XLH": 98,
	}

	for index, value := range maps {

		if index == "WT" {
			fmt.Println("CONTINUE")
			continue
		}

		if index == "XLH" {
			fmt.Println("BREAK")
			break
		}

		fmt.Println(index, value)
	}
}
