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
	if (x.nums[0]+x.inc)/max == 1 {
		//this is so we don't have to add 1 to the if statment each loop
		max -= 1
		for i := range x.nums[1:] {
			if x.nums[i+1] == max {
				x.nums[i+1] = 0
				x.current[i+1] = rns[0]
			} else {
				x.nums[i+1] += 1
				x.current[i+1] = rns[x.nums[i+1]]
				break
			}
		}
		//readjusting max fomr the -= statment
		max += 1
	}
	x.nums[0] = (x.nums[0] + x.inc) % max
	x.current[0] = rns[x.nums[0]]
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
	func() {
		defer close(test)
		passSize := 3
		possibleRunes := []rune{'a', 'b', 'c'}
		var worker PassTest
		worker.pTinit(passSize, 0, 3, possibleRunes)
		var worker2 PassTest
		worker2.pTinit(passSize, 1, 3, possibleRunes)
		var worker3 PassTest
		worker3.pTinit(passSize, 2, 3, possibleRunes)
		for i := math.Pow(float64(len(possibleRunes)), float64(passSize)); i > 0; i -= 1 {
			fmt.Println(worker.toString())
			worker.getNext(possibleRunes)
			fmt.Println(worker2.toString())
			worker2.getNext(possibleRunes)
			fmt.Println(worker3.toString())
			worker3.getNext(possibleRunes)
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
