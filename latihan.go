package main
import "fmt"

func main(){
	//ganjil genap
	fmt.Println("Ganjil Genap")
	for i := 0; i < 6; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "adalah genap")
			}else{
				fmt.Println(i, "adalah ganjil")
			}
		}

	//fizz buzz
	fmt.Println("\nFizz Buzz")
	var maxFib int
	fmt.Print("Max Bilangan : ")
	fmt.Scanln(&maxFib)
	for i := 0; i < maxFib; i++{
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("Fizz buzz")
		}else if i % 3 == 0 {
			fmt.Println("fizz")
		}else if i % 5 == 0 {
			fmt.Println("buzz")
		}else{
			fmt.Println(i)
		}
	} 

	//reverse string
	fmt.Println("\nReverse String")
	var word string
	var resWord string
	fmt.Print("masukan kata : ")
	fmt.Scanln(&word)
	
	for i := 0; i < len(word); i++ {
		resWord = string(word[i]) + resWord
	}
	fmt.Println(resWord)
}