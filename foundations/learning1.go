package main
import "fmt"

// Go / Golang
// Go doesn't have generic type
// T(v) -> not in Go
// garbage collection
// emphasis on simplicity like python

func calculate(){
    
}

func main(){
    m := map[string]int{
        "A" : 1,
        "B"  : 2,
        "C"  : 3,
    }
    fmt.Println(m)
    // m["P"] -> m["P"], true/false
    value, ok := m["P"]
    if ok {
        fmt.Println("This is ok: ", ok, "This is value: ", value)
    } else {
        fmt.Println("This is not ok: ", ok, "This is value: ", value)
    }
    
    s := map[int]bool{
        
    }
    // x, y := "ABC", 1
    
    // greet := func() {
    //     y++
    // }
    // greet()
    // s := fmt.Sprint("I am ", x, " my age is ", y)
    // fmt.Println(y)
    // fmt.Println(s)
    // a := []int{}
    // a = append(a, 1,2,3,4,5)
    // // for {
        
    // // }
    // for i := 0; i < len(a); i++ {
    //     fmt.Println(i, a[i])
    // }
    // fmt.Println("---")
    // for i, v := range a {
    //     fmt.Println(i, v)
    // }
    // fmt.Println("---")
    // i := 1
    // for i < len(a) {
    //     fmt.Println(i, a[i])
    //     i++
    // }
}
// --------------------------
// class __
//     def __init__
//         attributes
    
//     methods()
// ------------------------------
// var myFunction = function () {
//     console.log("hello")
// }