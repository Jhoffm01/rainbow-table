package main

import (
	brute "HashCracker/brute"
	rainbow "HashCracker/rainbow"
	"fmt"
	"os"
	"strings"
	//"log"
	//"os"
)

func main() {

	if os.Args[1] == "b" {
		fmt.Println(brute.GuessSingle("ccc", []rune{'a', 'b', 'c'}))
	} else if os.Args[1] == "help" {
		fmt.Println("You can use this program for two things, generate a rainbow road and searching passwords by their hash")
		fmt.Println("generate a rainbow road: \"main.exe generate (hash type EX: md5) (passworddump.txt)\"")
		fmt.Println("Find passwords: \"main.ext fine (hash or file containing hashes) (rainbowroad.txt)\"")
		fmt.Println("you can also use g insted of generate and f insted of find")
		fmt.Println("")
		fmt.Println("supported hashes:")
		fmt.Println("MD4,MD5,SHA1,SHA224,SHA256,SHA384,SHA512,BLAKE2b_256,BLAKE2b_384,BLAKE2b_512,")
		fmt.Println("RIPEMD160,SHA3_224,SHA3_256,SHA3_384,SHA3_512,SHA512_224,SHA512_256")
		fmt.Println("")
		fmt.Println("when generating a rainbow road the .txt generated will be called \"(password dump name)-(hash type).txt\"")
		fmt.Println("if you are looking for a password dump to start with weakpass.com is a good place to start")
	} else if os.Args[1] == "generate" || os.Args[1] == "g" {
		switch os.Args[2] {
		case "MD5", "md5":
			rainbow.CreateHashFiles(os.Args[3], 1)
		case "SHA1", "sha1":
			rainbow.CreateHashFiles(os.Args[3], 2)
		case "SHA224", "sha224":
			rainbow.CreateHashFiles(os.Args[3], 3)
		case "SHA256", "sha256":
			rainbow.CreateHashFiles(os.Args[3], 4)
		case "SHA384", "sha384":
			rainbow.CreateHashFiles(os.Args[3], 5)
		case "SHA512", "sha512":
			rainbow.CreateHashFiles(os.Args[3], 6)
		case "BLAKE2b_256", "blake2b_256":
			rainbow.CreateHashFiles(os.Args[3], 7)
		case "BLAKE2b_384", "blake2b_384":
			rainbow.CreateHashFiles(os.Args[3], 8)
		case "BLAKE2b_512", "blake2b_512":
			rainbow.CreateHashFiles(os.Args[3], 9)
		case "RIPEMD160", "ripemd160":
			rainbow.CreateHashFiles(os.Args[3], 10)
		case "SHA3_224", "sha3_224":
			rainbow.CreateHashFiles(os.Args[3], 11)
		case "SHA3_256", "sha3_256":
			rainbow.CreateHashFiles(os.Args[3], 12)
		case "SHA3_384", "sha3_384":
			rainbow.CreateHashFiles(os.Args[3], 13)
		case "SHA3_512", "sha_512":
			rainbow.CreateHashFiles(os.Args[3], 14)
		case "SHA512_224", "sha512_224":
			rainbow.CreateHashFiles(os.Args[3], 15)
		case "SHA512_256", "sha512_256":
			rainbow.CreateHashFiles(os.Args[3], 16)
		case "MD4", "md4":
			rainbow.CreateHashFiles(os.Args[3], 17)
		default:
			fmt.Println("hash not recognised, use help to see recognised hashes.")
		}
	} else if os.Args[1] == "find" || os.Args[1] == "f" {
		if strings.Contains(os.Args[2], ".txt") {
			rainbow.GuessMultiple(os.Args[2], os.Args[3])
		} else {
			rainbow.GuessSingle(os.Args[2], os.Args[3])
		}
	}

	//https://weakpass.com/wordlist/tiny

}
