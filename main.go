package main

import (
    "fmt"
    "github.com/xenowits/go-codes/ds/linked_list"
    "github.com/xenowits/go-codes/ds/stack"
)

func psh(a *[]int) {
    *a = append(*a, 10)
}

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

    //STACK section starts

    stk := stack.NewStack()
    stack.Push(stk, 1)
    stack.Push(stk, 2)

    fmt.Println(stack.Top(stk))
    stack.Pop(stk)

    fmt.Println(stack.Top(stk))
    stack.Pop(stk)

    if !stack.Empty(stk) {
        fmt.Println(stack.Top(stk))
        stack.Pop(stk)
    }

    //STACK section ends

}
