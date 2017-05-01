package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

/* Adaptive hashing algorithm: slow it down with salt, so the brute
force attack can't crack it in hours.
*/

const usage = `
usage:
	bcrypt hash|verify <password> [<cost>] [<pass-hash>]

<password> is required for both 'hash' and 'verify'
<cost> is required only for 'hash'
<pass-hash> is required only for 'verify'
`

func main() {
	if len(os.Args) < 4 ||
		(os.Args[1] != "hash" && os.Args[1] != "verify") {
		fmt.Println(usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	password := []byte(os.Args[2])

	switch cmd {
	case "hash":
		// needs to be an integer
		cost, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("cost must be an integer")
			os.Exit(1)
		}
		passhash, err := bcrypt.GenerateFromPassword(password, cost)

		if err != nil {
			fmt.Printf("error hashing password: %v", err)
			os.Exit(1)
		}
		fmt.Println(string(passhash))

	case "verify":
		passhash := []byte(os.Args[3])
		err := bcrypt.CompareHashAndPassword(passhash, password)
		if err != nil {
			fmt.Println("INVALID PASSWORD!!!!!")
		} else {
			fmt.Println("valid password")
		}
	}
}
