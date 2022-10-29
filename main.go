package main

import (
	"flag"
	"Cryptopals-go/task1"
)

func main() {

	// Declare flags
	// var encryptFlag = flag.Bool("encrypt", true, "flag for encryption") 
	// var decryptFlag = flag.Bool("decrypt", false, "flag for decryption") 
	var aesFlag = flag.Bool("aes", false, "flag to use AES cipher") 
	// var input = flag.String("input", "", "flag to use AES cipher") 
	

	flag.Parse()

	cipherKey := []byte{0x2b,0x7e,0x15,0x16,0x28,0xae,0xd2,0xa6,0xab,0xf7,0x15,0x88,0x09,0xcf,0x4f,0x3c}

    // use arg1 to decide which packages to call
    if *aesFlag{
        task1.KeySchedule(cipherKey)
    }

}