package testutil

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// BasicSort -
func BasicSort(t *testing.T, f func(a []int) []int) {
	inp := []int{3, 0, 5, 2, 1, 4}
	out := []int{0, 1, 2, 3, 4, 5}

	r := f(inp)

	if !reflect.DeepEqual(r, out) {
		t.Errorf(fmt.Sprintf("input: %v", inp) + fmt.Sprintf("output: %v", r))
	}
}

// Sort -
func Sort(t *testing.T, f func([]int) []int) {
	arrSize := rand.Intn(100)
	arr := make([]int, arrSize, arrSize)
	expArr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	copy(expArr, arr)
	//insertionSort(expArr)
	f(arr)
	if !reflect.DeepEqual(arr, expArr) {
		t.Log(fmt.Sprintf("expect:\n%v\n", expArr) + fmt.Sprintf("but get:\n%v\n", arr))
		t.Fail()
	}
}

// Bench -
func Bench(b *testing.B, f func(a []int) []int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrSize := 10000
		arr := make([]int, arrSize, arrSize)
		for j := range arr {
			arr[j] = rand.Intn(10000) + 1
		}
		b.StartTimer()
		f(arr)
	}
}
