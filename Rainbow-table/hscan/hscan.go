package hscan

import (
	"bufio"
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"hash"

	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"

	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// reads from file and sends passwords to the 'read' chanel one by one (line by line)
func reader(read chan string, scanner *bufio.Scanner, wg *sync.WaitGroup) {
	defer close(read)
	for scanner.Scan() {
		read <- scanner.Text()
	}
}

// originally writer needed a mutex lock but now that I'm only using one there shouldn't be any issues.
func writer(ch chan string, file *os.File, wg *sync.WaitGroup) {
	i := 0
	str := ""
	for {
		data, ok := <-ch
		//when we reach the end of the file ch will close and this will write whatever is leftover to the file.
		if !ok {
			if _, err := file.WriteString(str); err != nil {
				panic(err)
			}
			return
			//write to file
		} else if i == 10 {
			if _, err := file.WriteString(str); err != nil {
				panic(err)
			}
			i = 0
			str = data
			//fill str but don't do I/O
		} else {
			i++
			str += data
		}
	}
}

func control(write chan string, read chan string, hsh hash.Hash) {
	for data := range read {
		hsh.Write([]byte(data))
		sum := hsh.Sum(nil)
		hsh.Reset()
		write <- fmt.Sprintf("%x %s\n", sum, data)
	}
	wg.Done()
	wg.Wait()
}

func CreateHashFiles(filename string, hashes int) {
	var newfile *os.File
	//open 'filename'
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("")
		log.Fatalln(err)
	}
	defer file.Close()
	//creates a new file to store the rainbow table
	var hindex = [18]string{"", "md5", "sha1", "sha224", "sha256", "sha384", "sha512", "BLAKE2b_256", "BLAKE2b_384", "BLAKE2b_512", "RIPEMD160", "SHA3_224",
		"SHA3_256", "SHA3_256", "SHA3_384", "SHA3_512", "SHA512_224", "MD4"}
	if hashes == 0 {
		return
	} else {
		newfile, err = os.Create(filename + "-" + hindex[hashes] + ".txt")
		if err != nil {
			log.Fatal(err)
		}

	}
	//create I/O channels, read is closed by the reader function. write is closed by control.
	write := make(chan string, 100)
	read := make(chan string, 100)

	scanner := bufio.NewScanner(file)
	go reader(read, scanner, &wg)
	go writer(write, newfile, &wg)

	//use the needed hashing function
	var hsh crypto.Hash
	switch hashes {
	case 1:
		hsh = crypto.MD5
	case 2:
		hsh = crypto.SHA1
	case 3:
		hsh = crypto.SHA224
	case 4:
		hsh = crypto.SHA256
	case 5:
		hsh = crypto.SHA384
	case 6:
		hsh = crypto.SHA512
	case 7:
		hsh = crypto.BLAKE2b_256
	case 8:
		hsh = crypto.BLAKE2b_384
	case 9:
		hsh = crypto.BLAKE2b_512
	case 10:
		hsh = crypto.RIPEMD160
	case 11:
		hsh = crypto.SHA3_224
	case 12:
		hsh = crypto.SHA3_256
	case 13:
		hsh = crypto.SHA3_384
	case 14:
		hsh = crypto.SHA3_512
	case 15:
		hsh = crypto.SHA512_224
	case 16:
		hsh = crypto.SHA512_256
	case 17:
		hsh = crypto.MD4
	}

	wg.Add(1)
	go func() {
		defer close(write)
		control(write, read, hsh.New())
	}()
	for i := 0; i < runtime.NumCPU()-1; i++ {
		wg.Add(1)
		go func() {
			control(write, read, hsh.New())
		}()
	}
	wg.Wait()
}

// used when finding a single hash
func GuessSingle(hash string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("")
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//writer only needed if we are guessing multi
	read := make(chan string, 100)
	go reader(read, scanner, &wg)

	var found bool = false
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			str := worker(read, hash, &found)
			if str != "" {
				fmt.Println("found password for: " + hash + " " + str)
				found = true
			}
		}()
	}
	wg.Wait()
	if !found {
		fmt.Println("password not found for: " + hash)
	}
}

func worker(ch chan string, hash string, found *bool) string {
	var splitdata []string
	for data := range ch {
		if *found {
			wg.Done()
			return ""
		}
		splitdata = strings.Split(data, " ")
		if splitdata[0] == hash {
			*found = true
			wg.Done()
			return splitdata[1]
		}
	}
	wg.Done()
	//nothing was found
	return ""
}

// used when finding a file with hashes
func GuessMultiple(hashes string, filename string) {
	file, err := os.Open(hashes)
	if err != nil {
		fmt.Println("")
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	read := make(chan string, 100)
	go reader(read, scanner, &wg)

	//do this one by one to avoid flooding system with requests.
	for data := range read {
		GuessSingle(data, filename)
	}

}
