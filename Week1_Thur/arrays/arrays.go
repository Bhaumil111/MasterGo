package arrays

import "fmt"

var a [3]int

func change(arr [3]int) {
	arr[0] = 99
}

func main() {
	b := [3]int{1, 2, 3}
	//compiler infered length on own
	c := [...]int{10, 20, 30, 40}

	d := [5]int{1, 2}

	e := [5]int{
		0: 32,
		3: 323,
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(d)
	fmt.Println(e)

	f := [3]int{1, 2, 3}
	change(f)
	fmt.Println(f)

	fmt.Println("................")

	for i := 0; i < len(f); i++ {
		fmt.Println(f[i])

	}

	for index, value := range a {
		a[index] =199
		fmt.Println(index, value)
	}
	fmt.Println(a)



	// fmt.Println("Hi i am from array")


	


	matrix := [2][3]int{
		{1,2,3},
		{3,4,5},
	}
	fmt.Println(matrix)



	//Array  of Structs



	type User struct{
		id int 
	}


	users:=[2]User{
		{id:1}, 
		{id:2},

	}
	fmt.Println(users)


	



}

// 7 3 5 9 3 8 7 4 1 4
