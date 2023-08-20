package sort

import "math/rand"

func _quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	low_part := make([]int, 0, len(arr))
	high_part := make([]int, 0, len(arr))
	middle_part := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			low_part = append(low_part, item)
		case item == median:
			middle_part = append(middle_part, item)
		case item > median:
			high_part = append(high_part, item)
		}
	}

	low_part = _quickSort(low_part)
	high_part = _quickSort(high_part)

	low_part = append(low_part, middle_part...)
	low_part = append(low_part, high_part...)

	return low_part
}



func _bubbleSort(arr []int) []int{
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}


func sift(arr []int, i int, arrLen int) []int {
	done := false

	tmp       := 0
	maxChild  := 0

	for (i * 2 + 1 < arrLen) && (!done) {
		if i * 2 + 1 == arrLen - 1 {
			maxChild = i * 2 + 1
		} else if (arr[i * 2 + 1] > arr[i * 2 + 2]) {
			maxChild = i * 2 + 1
		} else {
			maxChild = i * 2 + 2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}

func _heapSort(arr []int) {
	i := 0
	tmp := 0

	for i = len(arr) / 2 - 1; i >= 0; i-- {
		arr = sift(arr, i, len(arr))
	}

	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0];
		arr[0] = arr[i];
		arr[i] = tmp;
		arr = sift(arr, 0, i);
	}
}



func _gnomeSort(arr []int) {
	i := 1
	tmp := 0
	for ; i < len(arr) ; {
		if arr[i] >= arr[i - 1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i - 1]
			arr[i - 1] = tmp

			if i > 1 {
				i--
			}
		}
	}
}



func merge(left, right []int) []int {
	result := make([]int, 0, len(left) + len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func _mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := _mergeSort(arr[:middle])
	right := _mergeSort(arr[middle:])

	return merge(left, right)
}


// tree sort

type node struct {
	val   int
	left  *node
	right *node
}

// definition of a node
type btree struct {
	root *node
}

// allocating a new node
func newNode(val int) *node {
	return &node{val, nil, nil}
}

// insert nodes into a binary search tree
func insert(root *node, val int) *node {
	if root == nil {
		return newNode(val)
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

// inorder traversal algorithm
// Copies the elements of the bst to the array in sorted orm
func inorderCopy(n *node, array []int, index *int) {
	if n != nil {
		inorderCopy(n.left, array, index)
		array[*index] = n.val
		*index++
		inorderCopy(n.right, array, index)
	}
}

func _treeSort(array []int, tree *btree) {
	// build the binary search tree
	for _, element := range array {
		tree.root = insert(tree.root, element)
	}
	index := 0
	// perform inorder traversal to get the elements in sorted orm
	inorderCopy(tree.root, array, &index)
}