package aes

import (
	"fmt"
)

var (
	roundConstants []byte = []byte{0x01,0x02,0x04,0x08,0x10,0x20,0x40,0x80,0x1B,0x36}
)

/*
keySchedule seperates a 128-bit key into a number of separate round keys
*/
func KeySchedule(key []byte) []byte{

	key_expanded := make([]byte, len(key)*11)
	
	offset := copy(key_expanded, key) //copy first 128 bits to generate K_0
	if offset != 16{
		fmt.Println("Incorrect key size")
		return nil
	}

	tempword := make([]byte, 4)
	rconInd := 0
	//generate next 10 round keys
	for i:=offset;i<16*11;i+=4{
		copy(tempword, key_expanded[i-4:])
		if i%16 == 0{
			// fmt.Printf("Round: %v \n", rconInd)
			RotWord(tempword)
			// fmt.Printf("Cipher Key: %X \n",tempword)
			SubWord(tempword)
			// fmt.Printf("Cipher Key: %X \n",tempword)
			Rcon(tempword, rconInd)
			rconInd++
			// fmt.Printf("Cipher Key: %X \n",tempword)
			copy(key_expanded[i:], tempword)
		}
		//wi = wi - nk ^ wi-1
		for j:=0; j<4; j++{
			key_expanded[i+j] = tempword[j] ^ key_expanded[i+j-16]
		}
	}

	// for i:= range key_expanded{
	// 	if i%16 == 0{
	// 		fmt.Printf("Key i: %v, %x\n",i/16,key_expanded[i:i+16])
	// 	}
	// }
	
	return key_expanded
}

/*
Rotates a 4 byte word left by a single byte 
*/
func RotWord(word []byte){
	temp := word[0]
	for i:=0;i<3;i++{
		word[i] = word[i+1]
	}
	word[3] = temp
}

/*
Preforms S-box substitution on each byte in a word
*/
func SubWord(word []byte){
	for i:=0;i<len(word);i++{
		word[i] = SubByte(word[i])
	}
}

/*
Adds round constant to first byte of word
*/
func Rcon(word []byte, roundIndex int){
	word[0] = word[0] ^ roundConstants[roundIndex]
}
	