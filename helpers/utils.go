package helpers

import (
	"bytes"
	"encoding/gob"
	"math"
)

//GetRealSizeOf returns the number of bytes occupied by a given interface
func GetRealSizeOf(v interface{}) int {
	b := new(bytes.Buffer)
	err := gob.NewEncoder(b).Encode(v)
	if err != nil {
		panic(err)
	}
	return b.Len()
}

// ValueCounts will find the duplicate elements from a given array of
// empty interfaces and will return a map containing a value count for each unique elements
func ValueCounts(arr []interface{}) map[interface{}]int {
	valueCounts := map[interface{}]int{}

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if _, ok := valueCounts[val]; !ok {
			valueCounts[val] = 1
		} else {
			valueCounts[val]++
		}
	}

	return valueCounts
}

// MaxIntSlice returns the maximum element of a given int array
func MaxIntSlice(nums []int) int {
	max := math.MinInt32
	for _, val := range nums {
		if max < val {
			max = val
		}
	}
	return max
}

//RepeatStringSlice takes in a string and a number of times it needs to be repeated.
// It will return a slice of nil interfaces instead of a compact string, as the Repeat method in strings package.
func RepeatStringSlice(s string, n int) []interface{} {
	repeats := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		repeats = append(repeats, s)
	}
	return repeats
}
