package main
import "fmt"
func main(){
	numbers:=[]int{1,2,3,4}
	num:=doubledNumber(&numbers)
	fmt.Println(num)

}
func doubledNumber(numbers *[]int)[]int{
	dnumbers:=[]int{}
	for _,val:=range *numbers{
		dnumbers=append(dnumbers, val*2)
	}
	return dnumbers
}
