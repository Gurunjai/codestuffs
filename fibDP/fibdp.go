package fibdp

import "fmt"

const (
	// TableSize The size of the Hash Table
	TableSize = (1 << 9)

	// TableMask Mask Value for the Hash Table
	TableMask = TableSize - 1
)

var (
	// Cache - for the new cache for the fibonacciy dynamic programming
	Cache *FibTable
)

// FibVal Node representing the {key: value} pair
type FibVal struct {
	key int
	val uint64
	height int32
	left *FibVal
	right *FibVal
}

// NewFibVal returns a new Node of type FibVal for the given {key: value} pair
func NewFibVal(key int, val uint64) *FibVal {
	return &FibVal {
		key: key,
		val: val,
		height: -1,
		left: nil,
		right: nil,
	}
}

// Find to get the key and value for the provided key
func Find(root *FibVal, key int) *FibVal {
	if root == nil || root.key == key {
		return root
	}

	if key > root.key {
		return Find(root.right, key)
	}

	return Find(root.left, key)	
}

// Insert insert a new {key: value} pair to the Node
func Insert(root *FibVal, key int, val uint64) *FibVal {
	if root == nil {
		root = NewFibVal(key, val)
		return root
	} 
	
	if key > root.key {
		root.right = Insert(root.right, key, val)
	} else {
		root.left = Insert(root.left, key, val)
	}

	return root
}

// InOrder dump the tree in the form of InOrder (Left, Root & Right)
func InOrder(root *FibVal) {
	if root == nil {
		return;
	}

	InOrder(root.left)
	fmt.Printf("Key: %v, Value: %v\t", root.key, root.val)
	InOrder(root.right)
}

// FibTable the in-memory table representing the Computed Fibonacci values
// collision is handled with a binary search tree
type FibTable struct {
	slot []*FibVal
	counter map[int]int
}

// NewFibTable constructs a pointer to the new table
func NewFibTable(size int) *FibTable {
	return &FibTable {
		slot: make([]*FibVal, size),
		counter: make(map[int]int),
	}
}

func (ft *FibTable) hashFunction(key int) (idx int) {
	idx ^= key
	idx |= idx >> (1 << 26)
	idx *= idx ^ (1 << 8)
	idx |= idx << (1 << 14)
	idx ^= TableSize

	return idx & TableMask
}

func (ft *FibTable) add(key int, val uint64) {
	idx := ft.hashFunction(key)
	ft.slot[idx] = Insert(ft.slot[idx], key, val)
	ft.counter[idx]++

}

func (ft *FibTable) get(key int) uint64 {
	idx := ft.hashFunction(key)
	node := Find(ft.slot[idx], key)

	if node == nil {
		return 0
	}

	return node.val	
}

func (ft *FibTable) dump() {
	for idx, c := range(ft.counter) {
		fmt.Printf("Dumping the data for Index: %v, Counter: %v\n\t", idx, c)
		InOrder(ft.slot[idx])
		fmt.Println()
	}	
}

// Fibonacci used to provide the fibonacci value for the provided input
func Fibonacci(n int) uint64 {
	if v := Cache.get(n); v != 0 {
		return v
	}

	v := Fibonacci(n - 1) + Fibonacci(n - 2)
	Cache.add(n, v)

	return v
}

func fib(n int) int {
	if 0 >= n {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return fib(n - 2) + fib (n - 1)
}

// Initialize used to define the cache and set the initial parameters
func init() {
	Cache = NewFibTable(TableSize)

	Cache.add(0, 1)
	Cache.add(1, 1)
	Cache.add(2, 1)
}