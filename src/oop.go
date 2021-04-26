package main
import (
	animal2 "../src/animal"
	"fmt"
)

func main()  {
	ani := animal2.NewAnimal("狗")
	dog := animal2.Dog{ani}
	fmt.Println(dog.GetName(), dog.Call(), dog.FavorFood())
}
