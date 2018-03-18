package main


import "fmt"


var x=0



func increment () int {
	x++
	return x

}


func main(){

	x:=42
	fmt.Println(x)
	fmt.Println(increment()) // x=1 
	fmt.Println(increment())  // x=2
	{

		fmt.Println(x)
		y:="Hello im getting bored thats too simple"
		fmt.Println(y)

	}
	//fmt.Println(y) !!!!can do that actually


}