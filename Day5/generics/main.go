package main
import "fmt" 

type ordered interface{
	int | float64
}




func bigFunc[T ordered] ( a , b T ) T{
	if( a>b ){
		return  a
	}
	return b
}
func main(){



	fmt.Println(bigFunc(1,3))
	fmt.Println(bigFunc(2.3,4.2))
	fmt.Println(bigFunc(4.0,4))

	// bigFunc("342","32")
	fmt.Println("Learning generics for the first time")
}
