package numbersorting

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	uintSamlpe    uint    = 15
	uint32Samlpe  uint32  = 32
	uint16Samlpe  uint16  = 16
	uint64Samlpe  uint64  = 64
	uint8Samlpe   uint8   = 8
	int32Samlpe   int32   = -32
	int64Samlpe   int64   = -64
	float32Samlpe float32 = 32.32
	float64Samlpe float64 = 64.64
)

var (
	dataComb        = []interface{}{int32Samlpe, int64Samlpe, 10, 7, uint8Samlpe, 9, uint32Samlpe, 1, 5, uint16Samlpe, uintSamlpe, "D", "A", "c"}
	sortedDataComb  = []interface{}{int64Samlpe, int32Samlpe, 1, 5, 7, uint8Samlpe, 9, 10, uintSamlpe, uint16Samlpe, uint32Samlpe, "A", "D", "c"}
	dataChars       = []interface{}{"D", "A", "Z", "v", "c", struct{}{}}
	sortedDataChars = []interface{}{struct{}{}, "A", "D", "Z", "c", "v"}
	data            = []interface{}{10, uint64Samlpe, 7, uint8Samlpe, 9, float32Samlpe, float64Samlpe, 1, 5}
	sortedData      = []interface{}{1, 5, 7, uint8Samlpe, 9, 10, float32Samlpe, uint64Samlpe, float64Samlpe}
)

func TestMergeSort(t *testing.T) {
	d := make([]interface{}, len(data), len(data))
	copy(d, data)
	MergeSort(d)
	assert.Equal(t, d, sortedData)
	d = make([]interface{}, len(dataChars), len(dataChars))
	copy(d, dataChars)
	MergeSort(d)
	assert.Equal(t, d, sortedDataChars)
	d = make([]interface{}, len(dataComb), len(dataComb))
	copy(d, dataComb)
	MergeSort(d)
	assert.Equal(t, d, sortedDataComb)
}
func TestInsertionSort(t *testing.T) {
	d := make([]interface{}, len(data), len(data))
	copy(d, data)
	InsertionSort(d)
	assert.Equal(t, d, sortedData)
	d = make([]interface{}, len(dataChars), len(dataChars))
	copy(d, dataChars)
	InsertionSort(d)
	assert.Equal(t, d, sortedDataChars)
	d = make([]interface{}, len(dataComb), len(dataComb))
	copy(d, dataComb)
	InsertionSort(d)
	assert.Equal(t, d, sortedDataComb)
}
func TestSelectionSort(t *testing.T) {
	d := make([]interface{}, len(data), len(data))
	copy(d, data)
	SelectionSort(d)
	assert.Equal(t, d, sortedData)
	d = make([]interface{}, len(dataChars), len(dataChars))
	copy(d, dataChars)
	SelectionSort(d)
	assert.Equal(t, d, sortedDataChars)
	d = make([]interface{}, len(dataComb), len(dataComb))
	copy(d, dataComb)
	SelectionSort(d)
	assert.Equal(t, d, sortedDataComb)
}
func TestQuickSort(t *testing.T) {
	d := make([]interface{}, len(data), len(data))
	copy(d, data)
	QuickSort(d)
	assert.Equal(t, d, sortedData)
	d = make([]interface{}, len(dataChars), len(dataChars))
	copy(d, dataChars)
	QuickSort(d)
	assert.Equal(t, d, sortedDataChars)
	d = make([]interface{}, len(dataComb), len(dataComb))
	copy(d, dataComb)
	QuickSort(d)
	assert.Equal(t, d, sortedDataComb)
}
func TestBubbleSort(t *testing.T) {
	d := make([]interface{}, len(data), len(data))
	copy(d, data)
	BubbleSort(d)
	assert.Equal(t, d, sortedData)
	d = make([]interface{}, len(dataChars), len(dataChars))
	copy(d, dataChars)
	BubbleSort(d)
	assert.Equal(t, d, sortedDataChars)
	d = make([]interface{}, len(dataComb), len(dataComb))
	copy(d, dataComb)
	BubbleSort(d)
	assert.Equal(t, d, sortedDataComb)
}

func BenchmarkMergeSort(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		MergeSort(dataChars)
	}
}

func BenchmarkMergeSort10000(b *testing.B) {
	// run the Fib function b.N times
	data := make([]interface{}, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = rand.Int()
	}
	for n := 0; n < b.N; n++ {
		MergeSort(data)
	}
}

func BenchmarkBuiltIn(b *testing.B) {
	// run the Fib function b.N times
	data := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = rand.Int()
	}
	for n := 0; n < b.N; n++ {
		sort.Sort(sort.IntSlice(data))
	}
}
func BenchmarkMergeSortComb(b *testing.B) {
	// run the Fib function b.N times
	d := make([]interface{}, len(dataChars), len(dataComb))
	copy(d, dataComb)
	for n := 0; n < b.N; n++ {
		MergeSort(d)
	}
}
func BenchmarkMergeSortChars(b *testing.B) {
	// run the Fib function b.N times
	d := make([]interface{}, len(dataChars), len(dataChars))
	copy(d, dataChars)
	for n := 0; n < b.N; n++ {
		MergeSort(d)
	}
}
