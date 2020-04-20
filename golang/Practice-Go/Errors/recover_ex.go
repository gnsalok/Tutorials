package main 

import(
	"fmt"
)

func main(){
	
	//after executing this line program still continue the executing becasue of recover 
	//recover allows us to continue the execution 
	fmt.Println(safeDiv(5,0))
	fmt.Println(safeDiv(54,8))

}

func safeDiv(n1, n2 int) int{
		defer func(){
		fmt.Println(recover())
	}()

	solution := n1/n2
	return solution


}
