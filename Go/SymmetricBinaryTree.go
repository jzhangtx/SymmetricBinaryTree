package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func GetArrayFromInput(count int, prompt string) []int {
	if count == 0 {
		return []int{}
	}

	fmt.Print(prompt)

	v := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&v[i])
	}

	return v
}

func ToNode(n int) *Node {
	if n == -1 {
		return nil
	}

	return &Node{n, nil, nil}
}

func ConvertToBT(vec []int, index int, nodes []*Node) {
	if index == len(vec) {
		return
	}

	var v []*Node
	for _, n := range nodes {
		if index == len(vec) {
			break
		}

		n.left = ToNode(vec[index])
		index++
		v = append(v, n.left)

		if index == len(vec) {
			break
		}
		n.right = ToNode(vec[index])
		index++
		v = append(v, n.right)
	}
	ConvertToBT(vec, index, v)
}

func GetBTFromArray(vec []int) *Node {
	pNode := &Node{vec[0], nil, nil}

	nodes := []*Node{pNode}

	ConvertToBT(vec, 1, nodes)
	return pNode
}

func Symmetric(left, right *Node) bool {
	if left == nil {
		if right == nil {
			return true
		}
		return false
	}

	if left.data != right.data {
		return false
	}

	return Symmetric(left.left, right.right) && Symmetric(left.right, right.left)
}

func IsSymmetric(root *Node) bool {
	return Symmetric(root.left, root.right)
}

func main() {
	for {
		fmt.Print("Number of nodes: ")
		var count int
		fmt.Scan(&count)
		if count == 0 {
			break
		}

		vec := GetArrayFromInput(count, "Please enter all the nodes: ")
		pRoot := GetBTFromArray(vec)

		fmt.Print("The binary tree is ")
		if !IsSymmetric(pRoot) {
			fmt.Print("not ")
		}
		fmt.Println("symmetric!")
		fmt.Println("")
	}
}
