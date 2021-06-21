package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "time"
)

var monDate = map[string]int{
	"1월":  31,
	"2월":  29,
	"3월":  31,
	"4월":  30,
	"5월":  31,
	"6월":  30,
	"7월":  31,
	"8월":  31,
	"9월":  30,
	"10월": 31,
	"11월": 30,
	"12월": 31,
}

// Enter a month in a %d manner to decide date.
// Enter a folder name in a Scanf manner.
//

func main() {
	var monthNum int
	result := monthDate(monthNum)

	basepath := "./"
	baseName := "testing1"
	folderPath := filepath.Join(basepath, baseName)

	for i := 1; i < 13; i++ {

		// Depending on monthDate which user enters.
		for j := 1; j <= result; j++ {

			// Check error for when the folder already has created.
			_, err := os.Stat("test")
			if os.IsNotExist(err) {
				err := os.Mkdir(folderPath, 0755)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

}

// Enter a month to decide how many dates you create in a folder
func monthDate(userMonth int) (result int) {
	fmt.Scanln(&userMonth)
	switch userMonth {
	case 1:
		return 31
	case 2:
		return 28 // 윤달 처리는 하지 않는다.
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	}
	return result
}

// Create a folder based on monthDate

func createFolder() {

}
