package main



import "fmt"


//declaration and assignemnt outside of funcs with var

var g string="kostas"

func main(){
	//shorthand declaration and assignemnt

	a:=10
	b:="golang"
	c:=1.45
	d:= true || false

    var r bool

	fmt.Print("Hello World\n")
	fmt.Print(" \n",a)
	fmt.Print("%T \n",b)
	fmt.Print("%v \n",c)
	fmt.Print("%v \n",d)
	fmt.Print("%v \n",r)

}