# platform-go

## Go 面向对象编程：方法、函数

在 go 语言中，结构体就像是类的一种简化形式。go 的方法是作用在接收者（receiver）（接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针）上的一个函数，接收者是某种类型的变量，因此方法是一种特殊类型的变量。  
一个类型加上它的方法等价于面向对象的一个类。（类方法）类型和作用在它上面定义的方法。

```go
type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

func (p *Person) AddAge() {
    p.Age += 1
}
 
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s, I am %d years old", p.Name, p.Age)
}


func main() {
    person := NewPerson("Tom", 18)
    person.AddAge()
    person.SayHello()
}

// 通过指针定义了AddAge方法，然后定义了一个NewPerson方法，该方法返回一个指向Person对象的指针。在main函数中，我们通过NewPerson创建了一个Person对象，然后调用AddAge方法增加了Person的年龄，并输出了Person的信息。
```
