package control_structures

const add string = "+"
const subtract string = "-"
const multiply string = "*"
const divide = "/"

func calculate(op string, arg1, arg2 float32) float32 {
	
	switch op{
	case "+":
		return arg1 + arg2;
	case "-":
		return arg1 - arg2;
	case "*":
		return arg1 * arg2;
	case "/":
		if(arg2 != 0){
		return arg1/arg2;
	}else{
		return -1;
	}
	}
		return -1;
}
