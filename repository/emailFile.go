package repository

import (
	"bufio"
	"log"
	"os"
)

type EmailFile struct {
}

func (*EmailFile) SaveEmailToFile(email string) error {
	const (
		fileModeFlags       = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		fileModePermutation = 0o644
	)

	file, err := os.OpenFile("email.txt", fileModeFlags, fileModePermutation)

	if err != nil {
		return err
	}

	defer safelyClose(file)

	scanner := bufio.NewScanner(file)
	if !lookIfIsEmailInTheList(scanner, email) {
		_, err := file.WriteString(email + "\n")
		return err
	}
	return err

}

func safelyClose(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("Problem with closing a data file.")
	}
}

func lookIfIsEmailInTheList(scanner *bufio.Scanner, email string) bool {
	for scanner.Scan() {
		if scanner.Text() == email {
			return true
		}
	}
	return false
}
