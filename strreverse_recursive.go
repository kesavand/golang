package main

import "fmt"


var str = []byte{'a','b','c','d'};

var start int = 0;
var end int = (len(str)-1);

func main() {
	fmt.Println("The start and end is ",start,end);
	str_reverse();
	fmt.Printf("The str is %s\n",str);
}



func str_reverse() {

	var tmp byte;

	if start >= end {
		return;
	}
	//swap
	tmp,str[start]= str[start],str[end];
	str[end]= tmp;
	//call the same func	
	start +=1
        end -=1;
	//fmt.Println("The start and end is ",start,end);
	str_reverse();


}
