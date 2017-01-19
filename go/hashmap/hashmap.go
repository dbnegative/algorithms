//based on https://github.com/prakhar1989/go-hashmap/blob/master/hashmap.go
package main

import (
	"errors"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const (
	fnvPrime    = 16777619
	offSetBasis = 2166136261
)

type Cell struct {
	Key   string
	Value interface{}
}

type HashMap struct {
	size  int
	count int
	cell  [][]Cell
}

//Set data at hash index and increase cell count
func (h *HashMap) Set(key string, value interface{}) {
	index := h.getIndex(key)
	data := &Cell{Key: key, Value: value}
	if len(h.cell[index]) >= 1 {
		collisions++
	}
	h.cell[index] = append(h.cell[index], *data)
	h.count++
}

//Get data from hash map based on key return error if not found
func (h *HashMap) Get(key string) (*Cell, error) {
	index := h.getIndex(key)
	if len(h.cell[index]) > 0 { // is this cell even used
		// takes n time as it iterates through cell slice
		for i := range h.cell[index] {
			if h.cell[index][i].Key == key {
				return &Cell{Key: h.cell[index][i].Key, Value: h.cell[index][i].Value}, nil
			}
		}
	}
	return nil, errors.New("Could not find key")
}

//NewHashMap builds a new hash map
func NewHashMap(size int) (*HashMap, error) {
	hm := new(HashMap)
	hm.count = 0
	hm.size = size
	if size < 0 {
		return nil, errors.New("Cannot be zero")
	}
	hm.cell = make([][]Cell, size)
	for i := 0; i < size; i++ {
		hm.cell[i] = make([]Cell, 0)
	}
	return hm, nil
}

//helper func to get index
func (h *HashMap) getIndex(key string) int {
	return int(hash(key)) % h.size
}

//generate a hash based on FNV
func hash(key string) uint32 {
	hash := uint32(offSetBasis)
	for _, x := range []byte(key) {
		hash ^= uint32(x)
		hash *= fnvPrime
	}
	return uint32(hash)
}

func (h *HashMap) printDistribution() {
	high := 0
	empty := 0
	for _, v := range h.cell {
		if len(v) > 3 {
			high++
		}
		if len(v) < 1 {
			empty++
		}
	}
	fmt.Printf("Empty Cell's: %d or %.2f percent, Overpoulated Cell's: %d or %.2f percent, Collisions: %d \n ", //
		empty, float32(empty)/float32(hashmapsize)*100.0, high, float32(high)/float32(hashmapsize)*100.0, collisions)
}

var (
	testsize    = 900
	hashmapsize = 1000
	collisions  = 0
)

func main() {

	//begin test
	//new hash map
	test, _ := NewHashMap(hashmapsize)
	//generate test data
	testdata := make(map[string]interface{}, testsize)
	for i := 0; i < testsize; i++ {
		testdata[randomdata.SillyName()] = randomdata.Address()
	}
	//generate index so can be referenced easily
	testindex := make([]string, 0, len(testdata))
	for k := range testdata {
		testindex = append(testindex, k)
	}

	//populate hashmap
	for k, v := range testdata {
		test.Set(k, v)
	}

	fmt.Printf("Hashmap Size: %d, Hashmap Item Count: %d, Load factor: %.2f \n", test.size, test.count, float32(test.count)/float32(test.size))

	//check if data can be found
	_, err := test.Get(testindex[5])
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	//fmt.Printf("Data Found: %v \n", test2)
	test.printDistribution()
}
