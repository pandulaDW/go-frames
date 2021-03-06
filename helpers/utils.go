package helpers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pandulaDW/go-frames/errors"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

//GetRealSizeOf returns the number of bytes occupied by a given interface
func GetRealSizeOf(v interface{}) int {
	b := new(bytes.Buffer)
	gob.Register(time.Time{})
	err := gob.NewEncoder(b).Encode(v)
	if err != nil {
		panic(err)
	}
	return b.Len()
}

//ConvertSizeToString will convert a number of bytes into a KB, MG, and GB string format
// respectively given the number of bytes.
func ConvertSizeToString(size int) string {
	if size/(1024*1024*1024) > 1 {
		return fmt.Sprintf("%.2f GB", float32(size)/(1024*1024*1024))
	}

	if size/(1024*1024) > 1 {
		return fmt.Sprintf("%.2f MB", float32(size)/(1024*1024))
	}

	if size/(1024) > 1 {
		return fmt.Sprintf("%.2f KB", float32(size)/(1024))
	}

	return fmt.Sprintf("%d bytes", size)
}

//RepeatIntoSlice takes in any empty interface element and a number of times it needs to be repeated.
// It will return a slice of nil interfaces instead of a compact string, as the Repeat method in strings package.
func RepeatIntoSlice(s interface{}, n int) []interface{} {
	repeats := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		repeats = append(repeats, s)
	}
	return repeats
}

// ConvertToFloat will convert the given empty interface value to float64. Will consider int types
// also. If there's an error, will return a boolean confirming the result.
func ConvertToFloat(val interface{}) (*float64, bool) {
	assertedVal, ok := val.(float64)
	if !ok {
		intVal, ok := val.(int)
		if !ok {
			return nil, false
		}
		assertedVal = float64(intVal)
	}
	return &assertedVal, true
}

// Range is a helper function which generates a sequence of numbers starting from
// the given start integer to the stop integer by the increment of the step integer.
//
// Panics if the step value is incorrectly specified.
func Range(low, high, step int) []int {
	rangeSlice := make([]int, 0)

	if high < low && step > 0 {
		panic(errors.CustomError("step should be a negative value when high is lower than low"))
	}

	if high > low && step < 0 {
		panic(errors.CustomError("step should be a positive value when high is higher than low"))
	}

	if high > low {
		for i := low; i < high; i += step {
			rangeSlice = append(rangeSlice, i)
		}
	}

	if high < low {
		for i := low; i > high; i += step {
			rangeSlice = append(rangeSlice, i)
		}
	}

	return rangeSlice
}

// GetFunctionName will return the function name of the provided function
func GetFunctionName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	split := strings.Split(name, ".")
	return split[len(split)-1]
}

// FunctionNameWrapper returns a wrapper name with a given function name and a column name
func FunctionNameWrapper(funcName, colName string) string {
	return fmt.Sprintf("%s(%s)", funcName, colName)
}

// GenerateRandomSeries will generate n number of random integer values where n is in between [0, range_] using
// the seed value provided.
//
// If withReplacement is set to true, the sample will contain duplicate values. And if withReplacement is set to
// false and if the range_ is lower than n, then the returned sample will contain number of elements equal to
// the range_.
func GenerateRandomSeries(n uint64, range_ uint64, seed int64, withReplacement bool) []int {
	rand.Seed(seed)
	m := make(map[int]int)
	seq := make([]int, 0, n)

	if withReplacement {
		for i := 0; i < int(n); i++ {
			seq = append(seq, rand.Intn(int(range_)))
		}
		return seq
	}

	if n < range_ {
		for len(seq) < int(n) {
			val := rand.Intn(int(range_))
			if _, ok := m[val]; !ok {
				m[val] = 0
				seq = append(seq, val)
			}
		}
	} else {
		for len(seq) < int(range_) {
			val := rand.Intn(int(range_))
			if _, ok := m[val]; !ok {
				m[val] = 0
				seq = append(seq, val)
			}
		}
	}

	return seq
}
