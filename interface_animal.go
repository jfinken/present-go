package main
import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}
func (c Cat) Speak() string {
    return "Meow!"
}
func main() {
    var myPet Animal
    myPet = Dog{}
    fmt.Println(myPet.Speak())
    myPet = Cat{}
    fmt.Println(myPet.Speak())
}
