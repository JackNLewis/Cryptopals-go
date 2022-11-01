package aes

import (
	"fmt"
)

var (

	/*
	State holds the 2d array of a 128 bit block with column major ordering
	*/
	state [][]byte
	
	/*
	Key Length (Nk words)
	*/
	Nk = 4

	/*
	Block Size (Nb words)
	*/
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
	printState(state)
	AddRoundKey(state, round_key)
	printState(state)
	SubBytesState(state)
	printState(state)
	ShiftRows(state)
	printState(state)
	return plaintext
}



/*
Shift Rows
*/
func ShiftRows(state [][]byte){
	//state is 2d array of dimensions m x n
	m := len(state)
	n := len(state[0])
	offset := 1
	temp := make([]byte,m)
	for row:=1;row<m;row++{
		for i:=0;i<offset;i++{
			temp[i] = state[row][i]
		}
		for i:=0;i<n-offset;i++{
			state[row][i] = state[row][i+offset]
		}
		for i:=0;i<offset;i++{
			state[row][(n-offset)+i] = temp[i]
		}
		offset++
	}
}

/*
Mix Rows
*/
func MixRows(state [][]byte){

}


/*
SubBytesState applies SubByte() to each byte in the state
*/
func SubBytesState(state [][]byte){
	for i:=0;i<len(state);i++{
		for j:=0;j<len(state);j++{
			state[i][j] = SubByte(state[i][j])
		}
	}
}

/*
AddRoundKey adds the round key to the state
*/
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