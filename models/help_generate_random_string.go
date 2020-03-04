package models

import (
	"crypto/rand"
)
/*
	The following two functions will generate a token
*/
func Generate_Random_Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Generate_Random_String(n int, str_chan chan string) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~!@#$%^&*()+"
	bytes, err := Generate_Random_Bytes(n)
	for err != nil {
		bytes, err = Generate_Random_Bytes(n)
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	str_chan <- string(bytes)
}

func Generate_Random_String_Provide_Characters_To_Choose (n int, str_chan chan string) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := Generate_Random_Bytes(n)
	for err != nil {
		bytes, err = Generate_Random_Bytes(n)
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	str_chan <- string(bytes)
}
