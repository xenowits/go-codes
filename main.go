package main

import (
    "fmt"
    "github.com/xenowits/go-codes/ds/linked_list"
)

func main () {
    arr := []int {2, 3, 5, 7, 11, 13, 15}

    head := linked_list.NewLL(arr)

    //traverse the linked list
    traversal_node := head

    for traversal_node != nil {
        fmt.Print(traversal_node.Data, " ---> ")
        traversal_node = traversal_node.Next
    }
    fmt.Println("nil")

}
