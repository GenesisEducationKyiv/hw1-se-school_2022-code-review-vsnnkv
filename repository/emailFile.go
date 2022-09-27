package repository

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FileRepository interface {
	SaveEmailToFile(email string) error
	IsExists(email string) (bool, error)
	GetEmails() []string
}

type EmailFile struct {
}

const (
	fileModeFlags       = os.O_APPEND | os.O_CREATE | os.O_RDWR
	fileModePermutation = 0644
)

func (*EmailFile) SaveEmailToFile(email string) error {

	file, err := os.OpenFile("emails", fileModeFlags, fileModePermutation)

	if err != nil {
		return err
	}

	defer safelyClose(file)

	_, err = file.WriteString(email + "\n")
	log.Printf("error: %v", err)
	return err

}

func (*EmailFile) IsExists(email string) (bool, error) {
	file, err := os.OpenFile("emails", fileModeFlags, fileModePermutation)
	if err != nil {
		log.Fatal("Problems with opening a file.")
		return false, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		print(scanner.Text())
		if scanner.Text() == email {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return false, err
	}

	return false, nil
}

func (*EmailFile) GetEmails() []string {
	file, err := os.Open("emails")
	if err != nil {
		return nil
	}
	defer safelyClose(file)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func safelyClose(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("Problem with closing a data file.")
	}
}
