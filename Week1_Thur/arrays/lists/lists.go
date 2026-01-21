package lists

//struct used to group data together 
// used for organize data
import "fmt"
// type Product struct{
// 	title string 
// 	id string 
// 	price float64


// }
type Product struct {
	title string 
	id string 
	price float64
}







func  main()  {
	
	// var prices [4]int64 = [4]int64 { 3,3,2,1}

	// fmt.Println(prices[3])
	// fmt.Println(len (prices))
	// fmt.Println(prices)
	// //infer data types 
	// infer_prices := [4] float64{32.32, 32.21,4.23, 2.1}

	// slice_prices := infer_prices[0:4]

	// fmt.Println(slice_prices)
	// fmt.Println(infer_prices)





	// fmt.Println("Hello from me")

	hobbies := [3]string {"Sports", "Cooking", "Reading "}
	fmt.Println(hobbies)
	

	fmt.Println(hobbies[0])

	fmt.Println(hobbies[1:])



	// 3)
	mainHobbiess := hobbies[:2]

	fmt.Println(mainHobbiess)

	// 4)
	fmt.Println(cap(mainHobbiess))


	//capacity means in slicing we use first 2 element of 5 ele size array so cap should 5 because we can use later values for that we reassign slice
	mainHobbiess = mainHobbiess[0:3]

	fmt.Println(mainHobbiess)
	

	// 5) 


	courseGoals := []string {"learn go ", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6) 
	courseGoals[1] = "learn c++ nw "
	courseGoals = append(courseGoals, "learn cpp basic now ")
	fmt.Println ( courseGoals)
	



	products := []Product{
		{"first=product", "hello", 32.32} , 
		{"second-product" ,"name " , 32.32}, 
	}
	fmt.Println(products)
	fmt.Println("..................................")


	super_prices := hobbies[0:2]

	super_prices = append(super_prices, "hello", "map" , " good-morning")

	fmt.Println(super_prices)


	discountPrices := []string{"hewe" , "bye", "google "}

	super_prices = append(super_prices, discountPrices...)
	fmt.Println(super_prices)

}


