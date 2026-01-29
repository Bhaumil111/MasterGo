package maps 

import "fmt"

func main() {

	websites := map[string]string{

		"Google.com":  "https://google.com",
		"Amazon.com ": "https://amazon.com",
	}

	fmt.Println(websites)


	fmt.Println(websites["Amazon.com "])


}
