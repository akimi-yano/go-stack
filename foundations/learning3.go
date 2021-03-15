package main
import "fmt"

type User struct {
    FirstName string
    LastName string
    Email string
}

type DataBase struct {
    rows map[string]User
    counter int 
}

type DataBaseMethods interface {
    Get()
    GetOne()
    Post()
    Put()
    Delete()
}

func (d DataBase) Get() []User {
    users := []User{}
    for _, user := range d.rows {
        users = append(users, user)
    }
    return users
}

func (d DataBase) GetOne(id int) User {
    return d.rows[string(id)]
}

func (d DataBase) Post(payload User) DataBase {
    d.rows[string(d.counter)] = payload
    d.counter++
    return d
}

func (d DataBase) Put(id int, payload User) DataBase {
    d.rows[string(id)] = payload 
    return d
}

func (d DataBase) Delete(id int) DataBase {
    delete(d.rows, string(id))
    return d
}

func main(){
	db := DataBase{map[string]User{}, 0} // Important: its {} to initialize a data structure
    db.Get()
    payload := User{"P","J","email@email.com"}
    newpayload := User{"P", "A", "email@email.com"}
    fmt.Println(payload)
    db = db.Post(payload)
    fmt.Println(db.GetOne(0))
    fmt.Println(db)
    fmt.Println(db.counter)
    db.Get()
    db = db.Put(0,newpayload)
    fmt.Println(db)
    db = db.Delete(0)
    fmt.Println(db)
    db = db.Post(User{"a","b","email2@email.com"})
    db = db.Post(User{"c","d","emai32@email.com"})
    db = db.Post(User{"e","f","234@email.com"})
    db = db.Post(User{"g","h","12312312@email.com"})
    fmt.Println(db)
}
