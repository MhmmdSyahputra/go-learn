package main
import "fmt"

func main(){
	pola1_cara1(5)
	fmt.Println()
	pola1_cara2(5)
	fmt.Println()
	pola2(5)
	fmt.Println()
	pola3(5)
	fmt.Println()
	pola4(5)
	fmt.Println()
	pola5(5)
	fmt.Println()
	pola6(6)
}

func pola1_cara1(val int) {
	for i := 0; i < val; i++{
		for j := 0; j < i; j++{
			fmt.Print("*")
		} 
		fmt.Println()
	} 
}
func pola1_cara2(val int) {
	bentuk := "*"
	for i := 0; i < val; i++ {
        fmt.Println(bentuk)
        bentuk += "*" 
    }
}

func pola2(val int) {
	for i := 0; i < val; i++ {
       for j := 0; j < val - i - 1; j++{
		fmt.Print(" ")
	   } 
	   for k := 0; k <= i; k++{
		fmt.Print("*")
	   }
	   fmt.Println("")
    }
}

func pola3(val int){
	for i := 0; i < val; i++{
		for j := 0; j <= i; j++{
			fmt.Print("*")
		}
		for k := 0; k < val - i; k++{
			fmt.Print(" ")
		}

		for l := 0; l < val - i; l++{
			fmt.Print(" ")
		}
		for m := 0; m <= i; m++{
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func pola4(val int){
	for i := 0; i < val; i++{
		for j := 0; j < val - i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pola5(val int){
	for i := 0; i < val; i++{
		for j := 0; j < val - i; j++ {
			fmt.Print("*")
		}
		for k := 0; k <= val - val + i; k++ {
			fmt.Print(" ")
		}
		for l := 0; l <= val - val + i; l++{
			fmt.Print(" ")
		}
		for m := 0; m < val - i; m++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pola6(val int){
	//pola X 
	for i := 0; i < val; i++ {
		for a := 0; a < val - val + i ; a++ {
			fmt.Print("i")
		}
		for j := 0; j < val - val + 1; j++ {
			fmt.Print("*")
		}
		for u := 0; u < val - i - i; u++ {
			fmt.Print(" ")
		}
		for t := 0; t < val - val + 1; t++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}