package stack

const (
    default_max_size = 100
    EMPTY_STACK = "empty"
)

func Size(arr *[]int) int {
    return len(*arr)
}

func Empty(arr *[]int) bool {
    if len(*arr) > 0 {
        return false
    } else {
        return true
    }
}

func Push (arr *[]int, ele int) {
    *arr = append(*arr, ele)
}

func Top (arr *[]int) (top int) {
    actual_slice := *arr
    if !Empty(arr) {
        top = actual_slice[len(actual_slice)-1]
    }
    return
}

func Pop (arr *[]int) {
    if !Empty(arr) {
        *arr = (*arr)[:len(*arr)-1]
    }
    return
}

func NewStack() *[]int {
    arr := make([]int, 0, default_max_size)
    return &arr
}
