package main

import (
    "fmt"  
    "math"
    "os"
)

//input example
//10
//DUDUDUDUDU

func main() {
	var n int
        fmt.Scanln(&n)    
 	
	if (float64(n) > math.Pow(10,6)){
        	fmt.Println("please input n no more than 1.000.000")    
        	os.Exit(1)    
    	}

    	var arr string
    	fmt.Scanln(&arr)    

    	arrLen := len(arr)    
    	if (arrLen!=n){
        	fmt.Println("string input not match with n")    
        	os.Exit(1)    
    	}
    
    
    	count := 0
    	valley := 0
    	valid := true
    
    	hiking := make([]string, arrLen)
    	for i:=1;i<=n;i++{    
        	var compare string = string(arr[i-1])
        	if (compare=="D"){
            		hiking[i-1]=compare
        	} else if(compare=="U"){
            		hiking[i-1]=compare
        	} else{
            		valid = false
            		break
        	}
    	}
    
      	if (valid){
        	for i:=1;i<=n;i++{    

                	if(hiking[i-1]=="D"){
                        	count--
                	} else {
                        	if((count+1)==0){
                
                        	valley++
                	}
                    		count++
                }
        }
    
        fmt.Println(valley)
        } else {
        fmt.Println("wrong input")
    }
    
}
