package main
import "fmt"


type Box[T any ] struct {
	Value T
}

func ( b Box[T]) Get( )T{
	return b.Value
}


func main(){

	b1 := Box[int] {
		Value :20,
	}

	b2:=Box[float64]{
		Value: 40.32,
	}

	fmt.Println(b1.Get())
	fmt.Println(b2.Get())
	
}
