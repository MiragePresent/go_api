package main

import (
	"fmt"
)

// User is a descriontion of system member
type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

// Session it's a group of rounds
type Session struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	SessionKey string        `json:"sessionKey"`
	Rounds     []Round       `json:"rounds"`
	Users      []SessionUser `json:"users"`
}

// SessionUser it's relation between session and its members
type SessionUser struct {
	SessionID string `json:"sessionID"`
	UserID    string `json:"userID"`
	IsManager bool   `json:"isManager"`
}

// Round object
type Round struct {
	SessionID string `json:"sessionID"`
	Number    string `json:"number"`
	StartTime string `json:"startTime"`
	Winner    string `json:"winner"`
}

// Attempt to win an round
type Attempt struct {
	RoundID string `json:"roundID"`
	UserID  string `json:"userID"`
	Time    string `json:"time"`
	Passed  bool   `json:"passed"`
}

func main() {
	fmt.Println("Hello from API")
}
