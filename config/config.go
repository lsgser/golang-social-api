package config

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"io"
	"crypto/rand"
	"fmt"
)

// MakeTimestamp function that works for MYSQL,it works on the datetime MySQL type
func MakeTimeStamp() string {
	t := time.Now()
	return t.Format("2006-01-02 15:02:16")
}

// MakeDate function that works for MYSQL,it works on the date MySQL type
func MakeDate() string {
	t := time.Now()
	return t.Format("2020-01-01")
}

// Hash a user password
func HashPassword(password string) ([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
}


//These header is for when a request
func AddSafeHeaders(w *http.ResponseWriter){
	(*w).Header().Set("Content-Type","application/json")
	(*w).Header().Set("Access-Control-Allow-Origin","*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
}

func AddHeadersNoJson(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin","*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
}

// newUUID generates a random UUID according to RFC 4122
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

//GenerateID creates a random unique ID/identifier string for a specific file
func GenerateID(prefix string,length int) string{
	//Source String used when generating a random identifier.
	const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Save the length in a constant so we don't look it up each time.
	const idSourceLen = byte(len(idSource))
	// Create an array with the correct capacity
	id := make([]byte, length)
	// Fill our array with random numbers
	rand.Read(id)
	// Replace each random number with an alphanumeric value
	for i, b := range id {
		id[i] = idSource[b%idSourceLen]
	}
	// Return the formatted id
	return fmt.Sprintf("%s_%s", prefix, string(id))
}
