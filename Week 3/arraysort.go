package main
import (
	"fmt"
	"sync"
)
func main(){
	//finalarr:=make([]int,12)
	fmt.Printf("Enter 12 integers : ")
	arr:=make([]int,12)
	arr1:=make([]int,3)
	arr2:=make([]int,3)
	arr3:=make([]int,3)
	arr4:=make([]int,3)
	for i:=0;i<len(arr);i++ {
		fmt.Scan(&arr[i])
	}
	c1:=make(chan []int)
	c2:=make(chan []int)
	c3:=make(chan []int)
	c4:=make(chan []int)
	var wg sync.WaitGroup
	go sort(arr,0,3,c1)
	arr1 = <-c1
	go sort(arr,3,6,c2)
	arr2 = <-c2
	wg.Wait()
	go sort(arr,6,9,c3)
	arr3= <-c3
	go sort(arr,9,12,c4)
	arr4 = <-c4
	wg.Wait()	
	fmt.Printf("\nArr 1 %d \n",arr1)
	fmt.Printf("Arr 2 %d \n",arr2)
	fmt.Printf("Arr 3 %d \n",arr3)
	fmt.Printf("Arr 4 %d \n",arr4)
	final := merge(merge(arr1, arr2), merge(arr3, arr4))
	fmt.Println("Sorted Array : ",final)
}
func sort(arr []int,start int,end int,c chan []int){
	temparr:=make([]int,3)
	for j:=start;j<end;j++{
		for k:=j+1;k<end;k++{
			if(arr[j]>arr[k]){
				temp:=arr[j]
				arr[j]=arr[k]
				arr[k]=temp
			}
		}
	}
	temparr[0] = arr[start]
	temparr[1] = arr[start+1]
	temparr[2] = arr[end-1]
	//fmt.Printf("[%d, %d, %d]",temparr[0],temparr[1],temparr[2])
	c <- temparr
}
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}