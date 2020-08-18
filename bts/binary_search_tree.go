package bts

type Node struct {
	parent *Node
	value *int
	left *Node
	right *Node

}

func buildBST(q[] int32) *Node {

	temp := int(q[0])
	root := &Node{
		value: &temp,
	}

	for i := 1; i < len(q); i++ {
		temp2 := int(q[i])
		insertChild(root, &temp2)
	}

	return root
}

func insertChild(parent *Node, input *int) {
	if *input < *parent.value{
		if parent.left == nil {
			left := &Node{
				value: input,
				parent: parent,
			}

			parent.left = left

			left.parent = parent
			return;
		} else {
			insertChild(parent.left, input)
		}
	} else {
		if parent.right == nil {
			right := &Node{
				value: input,
				parent: parent,
			}

			parent.right = right
			return
		} else {
			insertChild(parent.right, input)
		}
	}
}
