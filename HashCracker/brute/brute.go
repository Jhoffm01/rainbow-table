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
	pipe    *chan string
}
type inter interface {
	//converts []rune from passTest struct into a string, returns string.
	toString() string
	//will get the next password to test for passTest
	getNext(*[]rune)
	//this is ti initiate passTest structures.
	pTinit(int, int, int, *[]rune, *chan string) PassTest
	//this is for guessing a single hash.
	GuessSingle(string) string
}

func (x PassTest) produce(passwords int, rns *[]rune) {
	for i := 0; i < passwords; i++ {
		*x.pipe <- x.toString()
		x.getNext(rns)
		//needs removed
		fmt.Println(<-*x.pipe)
	}
}

func (x PassTest) getNext(rns *[]rune) {
	max := len(*rns)
	if (x.nums[0]+x.inc)/max == 1 {
		//this is so we don't have to add 1 to the if statment each loop
		max -= 1
		for i := range x.nums[1:] {
			if x.nums[i+1] == max {
				x.nums[i+1] = 0
				x.current[i+1] = (*rns)[0]
			} else {
				x.nums[i+1] += 1
				x.current[i+1] = (*rns)[x.nums[i+1]]
				break
			}
		}
		//readjusting max fomr the -= statment
		max += 1
	}
	x.nums[0] = (x.nums[0] + x.inc) % max
	x.current[0] = (*rns)[x.nums[0]]
}

func (x PassTest) toString() string {
	return string(x.current)
}

func (z *PassTest) pTinit(sz, spot, threads int, rns *[]rune, pipe *chan string) {
	x := PassTest{threads, make([]int, sz), make([]rune, sz), pipe}
	for i := range x.current {
		x.current[i] = (*rns)[0]
		x.nums[i] = 0
	}
	x.current[0] = (*rns)[spot]
	x.nums[0] = spot
	z.current = x.current
	z.inc = x.inc
	z.nums = x.nums
	z.pipe = x.pipe
}

func GuessSingle(hash string, possibleRunes []rune) string {
	//create a pipe to pass hashes to our workers
	pipe := make(chan string, 100)
	found := ""
	//create the function that will generate plaintext passwords
	func() {
		defer close(pipe)
		passSize := 4
		possibleRunes := []rune{'a', 'b', 'c', 'd'}
		workload := int(math.Pow(float64(len(possibleRunes)), float64(passSize)) / 3)
		var generator PassTest
		var generator2 PassTest
		var generator3 PassTest
		generator.pTinit(passSize, 0, 3, &possibleRunes, &pipe)
		generator2.pTinit(passSize, 1, 3, &possibleRunes, &pipe)
		generator3.pTinit(passSize, 2, 3, &possibleRunes, &pipe)
		generator.produce(workload, &possibleRunes)
		generator2.produce(workload, &possibleRunes)
		generator3.produce(workload, &possibleRunes)
		fmt.Println("Hello, World!")

	}()
	//for x threads spwan x-1 workers
	//might change to x-2 workers but for now we will just do 1
	go func() {
		ret := worker(hash, pipe)
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
