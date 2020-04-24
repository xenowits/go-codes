package main

import (
    "fmt"
    temp "cp/ds/linked_list"
)

func main () {
    arr := []int {2, 3, 5, 7, 11, 13, 15}

    head := temp.NewLL(arr)

    //traverse the linked list
    traversal_node := head

    for traversal_node != nil {
        fmt.Print(traversal_node.Data, " ---> ")
        traversal_node = traversal_node.Next
    }
    fmt.Println("nil")

}
