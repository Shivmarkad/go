package main
import "fmt"

func main(){
	welcome := "taking the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the details :")

	// comma ok || err ok

	input , _ := reader.ReadString('\n') // if you don't care about the err use _
	fmt.Println("Thanks for providing the details: ", input)
	fmt.Printf("Type of the details: %T", input)
}