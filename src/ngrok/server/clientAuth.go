package server

import (
	"bufio"
	"fmt"
	"log"
	"ngrok/msg"
	"os"
	"strings"
)

var authTokens []string

//Custom client authorization
type ClientAuth interface {
	IsValid() bool
}

//Default use file to store username and password and to check client auth info
type fileAuthDb struct {
	Auth *msg.Auth
}

func NewClientAuth(auth *msg.Auth) ClientAuth {
	return &fileAuthDb{Auth: auth}
}

func (fa *fileAuthDb) IsValid() bool {
	if fa.Auth.User == "" {
		return false
	}
	if len(authTokens) == 0 {
		tokens, err := fa.readLines("./tokens.txt")
		if err != nil {
			log.Fatalf("failed opening directory: %s", err)
		} else {
			authTokens = tokens
			if len(authTokens) == 0 {
				return false
			}
		}
	}

	for _, token := range authTokens {
		// fmt.Println("======")
		// fmt.Println(fa.Auth.User)
		// fmt.Println(token)
		if strings.TrimSpace(token) == fmt.Sprintf("%s", fa.Auth.User) {
			return true
		}
	}
	return false
}

func (fa *fileAuthDb) readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
