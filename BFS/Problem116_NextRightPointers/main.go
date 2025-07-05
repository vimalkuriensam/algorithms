package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func enqueue(queue *[]*Node, node *Node) {
	(*queue) = append((*queue), node)
}

func dequeue(queue *[]*Node) *Node {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.Next = nil
	queue := []*Node{root}
	for len(queue) > 0 {
		level_size := len(queue)
		for i := 0; i < level_size; i++ {
			current_node := dequeue(&queue)
			if current_node.Left != nil {
				current_node.Left.Next = current_node.Right
				if current_node.Next != nil {
					current_node.Right.Next = current_node.Next.Left
				} else if current_node.Next == nil && current_node.Right != nil {
					current_node.Right.Next = nil
				}
				enqueue(&queue, current_node.Left)
				enqueue(&queue, current_node.Right)
			}
		}
	}
	return root
}

func main() {}
