package main

import (
	"fmt"
)

func main() {
	
	b:= make([]int, 100)
	
	for i := 1; i<=100; i++{
	    c:=100/i
	    m:=0
	    for j := 1; j<=c; j++{
	        if (m<=100){
	            m=m+i
	        }
	        if (b[m-1]==0) {
	            b[m-1]=1
	        } else {
	          b[m-1]=0
	        }
	    }
	
	}
	
	count := 0
	for _, v := range b {
		if (v == 1){
			count++
		}
	}
	fmt.Println(count)
	
}
