package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

const usage = `
usage:
	hmac sign|verify <key> <value>
`

func main() {
	if len(os.Args) < 4 ||
		(os.Args[1] != "sign" && os.Args[1] != "verify") {
		fmt.Println(usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	key := os.Args[2]
	value := os.Args[3]

	switch cmd {
	case "sign":
		v := []byte(value)
		// Pointer to a function, key(slice of byte)
		h := hmac.New(sha256.New, []byte(key))
		h.Write(v)

		// The actual signature
		sig := h.Sum(nil)

		// New byte slice --> combining
		buf := make([]byte, len(value)+len(sig))

		// put value into buffer
		copy(buf, v)

		// copy sig into last part of buffer
		copy(buf[len(v):], sig)

		// using URLEncoding doesn't use + and / so we don't mess up the URL
		fmt.Println(base64.URLEncoding.EncodeToString(buf))

	case "verify":
		buf, err := base64.URLEncoding.DecodeString(value)
		if err != nil {
			fmt.Printf("error decoding: %v\n", err)

			// only use for command line. Don't use it on the web server or the entire
			// server will be stopped if you hit an error
			os.Exit(1)
		}

		v := buf[:len(buf)-sha256.Size]
		sig := buf[len(buf)-sha256.Size:]

		h := hmac.New(sha256.New, []byte(key))
		h.Write(v)
		sig2 := h.Sum(nil)

		// if the two signatures are the same, then it's legit
		// this is constant time, giving no clue to the attackers how many bytes are correct,
		// and how many are wrong
		if hmac.Equal(sig, sig2) {
			fmt.Println("signature is valid!")
		} else {
			fmt.Println("NOT VALID!")
		}

	}

}
