package main 

import "fmt"


var str = []byte {'a','b','c','d'};

func main () {
	var tmp byte;

	for i, j:= 0,(len(str) -1);i<len(str);i++ {
		tmp, str[i] = str[i],str[j];
		str[j] = tmp;

	}

	fmt.Printf("str is %s %d\n",str,len(str));

}
