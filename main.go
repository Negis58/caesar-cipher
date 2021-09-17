package main

import (
	"bufio"
	"fmt"
	"github.com/Negis58/caesar-cipher.git/cipher"
	"os"
)

func main() {
	caesarCipher := cipher.NewCaesarCipher()
	var x string
	var key int
	fmt.Print("Введите encrypt или decrypt\n")
	fmt.Scanf("%s\n", &x)
	fmt.Print("Введите исходную строку символов\n")
	inputString := bufio.NewScanner(os.Stdin)
	inputString.Scan()
	line := inputString.Text()
	fmt.Print("Введите ключ шифрования\n")
	fmt.Scanf("%d\n", &key)
	if x == "encrypt" {
		output := caesarCipher.Encrypt(line, key)
		fmt.Printf("Encrypt %s\n", output)
	} else if x == "decrypt" {
		output := caesarCipher.Decrypt(line, key)
		fmt.Printf("Decrypt %s\n", output)
	} else {
		fmt.Printf("Error string")
	}
}
