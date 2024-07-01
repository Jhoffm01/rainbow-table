package brute

import (
	"crypto"
	"fmt"
	"math"
)

type PassTest struct {
	inc     int
	nums    []int
	current []rune
}
type inter interface {
	//converts []rune from passTest struct into a string, returns string.
	toString() string
	//will get the next password to test for passTest
	getNext([]rune)
	//this is ti initiate passTest structures.
	pTinit(int, int, int, []rune) PassTest
	//this is for guessing a single hash.
	GuessSingle(string) string
}

func (x PassTest) getNext(rns []rune) {
	max := len(rns)
	for i := range x.nums {
		cur := x.nums[i] + x.inc
		if cur < max {
			x.nums[i] = cur
			x.current[i] = rns[x.nums[i]]
			break
		} else {
			x.nums[i] = cur % max
			x.current[i] = rns[cur%max]
		}
	}
}

func getNextHelper() {

}

func (x PassTest) toString() string {
	return string(x.current)
}

func (z *PassTest) pTinit(sz, spot, threads int, rns []rune) {
	x := PassTest{threads, make([]int, sz), make([]rune, sz)}
	for i := range x.current {
		x.current[i] = rns[0]
		x.nums[i] = 0
	}
	x.current[0] = rns[spot]
	x.nums[0] = spot
	z.current = x.current
	z.inc = x.inc
	z.nums = x.nums
}

func GuessSingle(hash string, possibleRunes []rune) string {
	//create a pipe to pass hashes to our workers
	test := make(chan string, 100)
	found := ""
	//create the function that will generate plaintext passwords
	go func() {
		defer close(test)
		passSize := 6
		possibleRunes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
		var worker PassTest
		worker.pTinit(passSize, 0, 1, possibleRunes)
		for i := math.Pow(float64(len(possibleRunes)), float64(passSize)); i > 0; i -= 1 {
			worker.getNext(possibleRunes)
		}
		fmt.Println("Hello, World!")

	}()
	//for x threads spwan x-1 workers
	//might change to x-2 workers but for now we will just do 1
	go func() {
		ret := worker(hash, test)
		if ret != "" {
			found = ret
		}
	}()
	if found == "" {
		return "password not found"
	} else {
		return "password found: " + found
	}
}

func worker(hash string, test chan string) string {
	//hash input from test
	hsh := crypto.MD5.New()
	for elem := range test {
		hsh.Write([]byte(elem))
		txt := hsh.Sum(nil)
		//test if the hashed input from test matches the hash we collected
		if string(txt) == hash {
			return string(txt)
		}
	}
	return ""
}
