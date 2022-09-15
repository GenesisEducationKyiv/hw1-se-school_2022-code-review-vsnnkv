package repository

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type EmailFile struct {
}

func (*EmailFile) SaveEmailToFile(email string) int {
	const (
		fileModeFlags       = os.O_APPEND | os.O_CREATE | os.O_RDWR
		fileModePermutation = 0644
	)

	file, err := os.OpenFile("email", fileModeFlags, fileModePermutation)

	if err != nil {
		return 500
	}

	defer safelyClose(file)

	scanner := bufio.NewScanner(file)
	if !lookIfIsEmailInTheList(scanner, email) {
		_, err := file.WriteString(email + "\n")
		log.Printf("error: %v", err)
		return 200
	}

	return 400

}

func safelyClose(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("Problem with closing a data file.")
	}
}

func lookIfIsEmailInTheList(scanner *bufio.Scanner, email string) bool {
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		print(scanner.Text())
		if scanner.Text() == email {
			return true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}
