package main

import "fmt"

//input example
//6
//10 10 20 20 30 40

func main() {
	var n int
	fmt.Scan(&n)

  	arr := make([]int, 101)

  	count := 0

  	for i := 0; i < n; i++ {
    		var colorCode int
    		fmt.Scan(&colorCode)

    		arr[colorCode]++

    		if (arr[colorCode]%2 == 0){
        		count++
    		}
  	}

  	fmt.Println(count)
}
