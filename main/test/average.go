package main
import "fmt"


func main(){
	ave:= average(2.3, 4.5, 3.2, 1.3)
		fmt.Println(ave)
}

func average(nums ...float64) (float64) {
	total_num := float64(0)
	for _, n :=  range nums{
		total_num += n
	}
	return total_num/float64(len(nums))
}
