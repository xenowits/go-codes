package linked_list

type Node struct {
    Data int
    Next *Node
}

func NewLL (arr []int) *Node {
    var head *Node
    var prev *Node

    // var head_set bool
    head_set := false

    for _, data := range arr {
        new_node := Node {Data: data}

        if head != nil {
                prev.Next = &new_node

                if head_set == false {
                    head = prev
                    head_set = true
                }

        } else {
            head = &new_node
        }

        prev = &new_node

    }

    //for the end : assign the last node to nil

    prev.Next = nil

    return head
}
