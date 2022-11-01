package aes

import (
	"fmt"
)

var (

	state [][]byte
	
	Nb = 4
)

/*
EncryptBlock is used encrypting a 128 bit block using AES-128. 
*/
func EncryptBlock(plaintext, key []byte) []byte{
	pb := []byte(plaintext)
    if len(key) == 16{ // 128 bit string
		state = [][]byte{
			{pb[0],pb[4],pb[8],pb[12]},
			{pb[1],pb[5],pb[9],pb[13]},
			{pb[2],pb[6],pb[10],pb[14]},
			{pb[3],pb[7],pb[11],pb[15]},
		}
	}

	key_expanded := KeySchedule([]byte(key))

	round_key := key_expanded[0:16]
	fmt.Printf("%x\n",round_key)
	printState(state)
	AddRoundKey(state, round_key)
	printState(state)
	return plaintext
}

func AddRoundKey(state [][]byte, roundkey []byte){
	for i:=0; i<len(state); i++{ 
		for j:=0;j<len(state[0]);j++{
			state[j][i] = state[j][i] ^ roundkey[(i*Nb)+j]
		}
	} 
}

func printState(state [][]byte){
	for i:=0; i<4; i++{ 
		fmt.Printf("{%x, %x, %x, %x}\n",state[i][0],state[i][1],state[i][2],state[i][3])
	}
	fmt.Println()
}