package main
import fmt

func eval(a,b int,op string) (int error){
	switch opo {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		q , _ := dev(a , b)
		return q , nil
	default:
		return  0 , fmt.Errorf(
		"unsupported operation: %s" , op)		
	}
}

func dev(a, b int) (q, r int){
	return a / b , a % b
}

func main() {
	if result , err := eval(3,4,"x"); err !=nil{
		fmt.Println("Error:",err)
	}else{
		fmt.Println(result)
	}
	q,r := div(13,3)
	fmt.Println(q,r)
}