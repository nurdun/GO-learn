package main 

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	// fmt.Println("hello world!")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	input,err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Printf("found an error : %s\n",err)
	}else{
		input = input[:len(input)-1] //为了避免感叹号出现在下一行
		fmt.Printf("Hello,%s!\n",input)
	}
}