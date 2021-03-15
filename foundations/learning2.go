package main
import "fmt"

func main(){
	
    // num := 2
    // fmt.Println(num)
    
    // 1) to get pointer/ memory address
    // pointer := &num
    // fmt.Println(pointer)
    
    // 2) to get value in this memory address
    // val := *pointer
    // fmt.Println(val)
    
    // 3) as type
    // []*User
    // its slice of pointer type for user objects
    
    // const example = 3 // not defined the type yet
    
    // fmt.Println(example + 1)
    const example = 3
    // fmt.Println(example + 0.5)
    
    const first = iota
    const second = iota
    
    const (
        third = iota
        fourth
    )
    // fmt.Println(first)
    // fmt.Println(second)
    // fmt.Println(third)
    // fmt.Println(fourth)
    m := map[string]int{}
    m["A"] = 1
    fmt.Println(m)
    
    delete(m, "A")
    fmt.Println(m)
}
