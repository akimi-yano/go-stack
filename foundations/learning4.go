package main
import "fmt"

    
type Bucket struct {
    values []string
}

type Bucketer interface {
    size() int
}

type HashTable struct{
  // array of pointer that holds bucket type
    buckets []*Bucket
}

type HashTabler interface {
    get(key string) []string
    put(key string, value string)
    remove(key string, value string)
}

func main(){
    h := initHashTable()
    h.put("pete", "sensei")
    h.put("epte", "asdf")
    fmt.Println(h)
    fmt.Println(h.buckets[0].values)
    // fmt.Println(hash("petefdfsfdsfdsfdsfds"))
    
    fmt.Println(h.get("pete"))
    fmt.Println(h.get("epte"))
    h.remove("pete", "sensei")
    fmt.Println(h.buckets[0].values)

    fmt.Println(h.buckets[0].size())
}

func initHashTable() *HashTable {
    h := &HashTable{}
    // define size here (10)
    for i := 0; i < 10; i++ {
        h.buckets = append(h.buckets, &Bucket{})
    }
    return h
}

func hash(key string) int {
    sum := 0
    for _, c := range key{
         sum += int(c)
    }
    return sum % 10 
}

func (h *HashTable) put(key string, value string) {
    idx := hash(key)
    found := false
    for _, word := range h.buckets[idx].values{
        if word == value{
            found = true
            break
        }
    }
    if found == false{
        h.buckets[idx].values = append(h.buckets[idx].values, value)
    }
}

func (h *HashTable) get(key string) []string {
    idx := hash(key)
    return h.buckets[idx].values
}

func (h *HashTable) remove(key string, value string) {
    idx := hash(key)
    for i, word := range h.buckets[idx].values{
        if word == value{
            h.buckets[idx].values = append(h.buckets[idx].values[:i], h.buckets[idx].values[i+1:]...)
            break
        }
    }
}

func (b *Bucket) size() int {
    return len(b.values)
}

// HashTable = {
//     buckets = [
//         B{
//             values:[]
//         }
//     ]
// }
// init() 
// hash(key) -> sum += int(char) % 10
// put(key, value)
// get(key)
// remove(key, value)

// put("Pete", "Jung")
//   hash("Pete") -> 4

// get("Pete")
//   hash("Pete") -> 4
  
// remove("Pete", "Jung")
//   hash("Pete") -> 4
//   if it finds "Jung" then slice out
  
// 0 -> []
// 1 -> []
// 2 -> []
// 3 -> []
// 4 -> ["Jung", "Chris", ...]

