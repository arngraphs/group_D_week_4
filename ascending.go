package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{5, 2, 9, 1, 6, 3}

    sort.Ints(numbers)

    fmt.Println("Numbers in ascending order:", numbers)
}

