package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// for {

	// 	// Enter a password and generate a salted hash
	// 	pwd := getPwd()
	// 	hash := hashAndSalt(pwd)

	// 	// Enter the same password again and compare it with the
	// 	// first password entered
	// 	pwd2 := getPwd()
	// 	pwdMatch := comparePasswords(hash, pwd2)
	// 	fmt.Println("Passwords Match?", pwd, pwd2)
	// 	fmt.Println(pwdMatch)
	// }
	dbPass := "$2a$04$poDNCJCSWmLirw54e0Eth.oSCrgIlQGnP8b71dXApmznjx.Pl7.Ie"
	// dbPassByte := []byte(dbPass)

	loginPass := "1river1"

	isAuthValid := ComparePasswords(dbPass, loginPass)
	fmt.Println(isAuthValid)

	passForDbSave := HashAndSalt(loginPass)
	fmt.Println(passForDbSave)

}
func getPwd() []byte {
	// Prompt the user to enter a password
	fmt.Println("Enter a password")
	// We will use this to store the users input
	var pwd string
	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// ComparePasswords will compare the db password with the password entered at login
func ComparePasswords(dbPwd, loginPwd string) bool {

	// We will need to convert both pass word to byte arrays
	byteHash := []byte(dbPwd)
	plainBytePwd := []byte(loginPwd)

	// Compare
	err := bcrypt.CompareHashAndPassword(byteHash, plainBytePwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// HashAndSalt will encrypt a password and return a string to store in the db
func HashAndSalt(pwd string) string {

	bytePwd := []byte(pwd)

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func getBytePwd() []byte {
	// Prompt the user to enter a password
	fmt.Println("Enter a password")
	// We will use this to store the users input
	var pwd string
	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)
}
