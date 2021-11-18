package main

import ("fmt")

func add( x, y int ) (out1,out2 int){
  out1 = x + y;
  out2 = x - y;
  return out1,out2
}


func main(){
	
	fmt.Println("Enter two numbers\n");
	var x,y int
	fmt.Scan(&x,&y);

	result1 , result2 := add(x, y);
	if x>y{
		fmt.Println(result2);		
	}else{
		fmt.Println(result1);
	}
	
}