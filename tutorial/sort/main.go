package sort

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var (
	DEAFAULT = []string{"quick", "bubble", "heap", "gnome", "merge", "tree"}
)

const (
	maxIntRange = 1000000
)

func New(name string, size int) (sort, error) {

	var sortEx sort
	arr := createRandomArr(size, maxIntRange)
	switch name {
	case "quick":
		sortEx = quickSort{
			size: size,
			arr : arr,
			elapseTime: make(chan time.Duration),
		}
	case "bubble":
		sortEx = bubbleSort{
			size: size,
			arr : arr,
			elapseTime: make(chan time.Duration),
		}
	case "heap":
		sortEx = heapSort{
			size: size,
			arr : arr,
			elapseTime: make(chan time.Duration),
		}
	case "gnome":
		sortEx = gnomeSort{
			size: size,
			arr : arr,
			elapseTime: make(chan time.Duration),
		}
	case "merge":
		sortEx = mergeSort{
			size: size,
			arr : arr,
			elapseTime: make(chan time.Duration),
		}
	case "tree":
		sortEx = treeSort{
			size: size,
			arr : arr,
			elapseTime: make(chan time.Duration),
		}
	default:
		return nil, fmt.Errorf("`%s` is not supported sort algorithm.", name)

	}
	return sortEx, nil


}

func getMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc

}
func createRandomArr(size int, max int) []int {

	var randomIntArr []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		randomIntArr = append(randomIntArr, rand.Intn(max))
	}
	return randomIntArr
}


func setDummy(arr []int) {
	arr = []int{}
}

type sort interface {
	do()
	getName() string
	getElapseTime() time.Duration
	getSize() int
	getMemoryUsage() uint64

}

type quickSort struct {
	size int
	arr []int
	elapseTime chan time.Duration
	result []int
	memoryUsage uint64
}

func (q quickSort) getName() string{
	return "quick"
}

func (q quickSort) getElapseTime() time.Duration {
	return <- q.elapseTime

}

func (q quickSort) getSize() int {
	return q.size
}

func (q quickSort) getMemoryUsage() uint64 {
	return q.memoryUsage
}

func (q quickSort) do() {
	start := time.Now()
	result := _quickSort(q.arr)
	setDummy(result)
	q.memoryUsage = getMemoryUsage()
	elapased := time.Now().Sub(start)
	q.elapseTime <- elapased
}



type bubbleSort struct {
	size int
	arr []int
	elapseTime chan time.Duration
	result []int
	memoryUsage uint64

}


func (b bubbleSort) do() {
	start := time.Now()
	result := _bubbleSort(b.arr)
	setDummy(result)
	b.memoryUsage = getMemoryUsage()

	elapased := time.Now().Sub(start)
	b.elapseTime <- elapased
}


func (b bubbleSort) getName() string{
	return "bubble"
}

func (b bubbleSort) getElapseTime() time.Duration {
	return <- b.elapseTime
}

func (b bubbleSort) getSize() int {
	return b.size
}

func (b bubbleSort) getMemoryUsage() uint64 {
	return b.memoryUsage
}

type heapSort struct {
	size int
	arr []int
	elapseTime chan time.Duration
	result []int
	memoryUsage uint64

}


func (h heapSort) do() {
	start := time.Now()
	_heapSort(h.arr)
	h.memoryUsage = getMemoryUsage()
	elapased := time.Now().Sub(start)
	h.elapseTime <- elapased
}


func (h heapSort) getName() string{
	return "heap"
}

func (h heapSort) getElapseTime() time.Duration {
	return <- h.elapseTime
}

func (h heapSort) getSize() int {
	return h.size
}

func (h heapSort) getMemoryUsage() uint64 {
	return h.memoryUsage
}

type gnomeSort struct {
	size int
	arr []int
	elapseTime chan time.Duration
	result []int
	memoryUsage uint64

}


func (h gnomeSort) do() {
	start := time.Now()
	_gnomeSort(h.arr)
	h.memoryUsage = getMemoryUsage()
	elapased := time.Now().Sub(start)
	h.elapseTime <- elapased
}


func (h gnomeSort) getName() string{
	return "gnome"
}

func (h gnomeSort) getElapseTime() time.Duration {
	return <- h.elapseTime
}

func (h gnomeSort) getSize() int {
	return h.size
}

func (h gnomeSort) getMemoryUsage() uint64 {
	return h.memoryUsage
}

type mergeSort struct {
	size int
	arr []int
	elapseTime chan time.Duration
	result []int
	memoryUsage uint64

}


func (h mergeSort) do() {
	start := time.Now()
	_mergeSort(h.arr)
	h.memoryUsage = getMemoryUsage()
	elapased := time.Now().Sub(start)
	h.elapseTime <- elapased
}


func (h mergeSort) getName() string{
	return "merge"
}

func (h mergeSort) getElapseTime() time.Duration {
	return <- h.elapseTime
}

func (h mergeSort) getSize() int {
	return h.size
}

func (h mergeSort) getMemoryUsage() uint64 {
	return h.memoryUsage
}

type treeSort struct {
	size int
	arr []int
	elapseTime chan time.Duration
	result []int
	memoryUsage uint64

}


func (h treeSort) do() {
	start := time.Now()
	tree := &btree{nil}
	_treeSort(h.arr, tree)
	h.memoryUsage = getMemoryUsage()
	elapased := time.Now().Sub(start)
	h.elapseTime <- elapased
}


func (h treeSort) getName() string{
	return "tree"
}

func (h treeSort) getElapseTime() time.Duration {
	return <- h.elapseTime
}

func (h treeSort) getSize() int {
	return h.size
}

func (h treeSort) getMemoryUsage() uint64 {
	return h.memoryUsage
}


func DoSort(s sort)  {
	go s.do()
	elapseTime := s.getElapseTime()
	fmt.Println("name:", s.getName(), "  size:", s.getSize(), "  elapseTime:", elapseTime)
}


