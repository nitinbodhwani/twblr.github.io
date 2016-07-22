package control_structures

func collatzChainLength(n int) int {
	
	iterationCounter := 0;
	for n!= 1{		
		iterationCounter++;		
		if(n%2 == 0){
			n = n/2;
		}else{
			n= 3*n + 1;
		}
	}

	return iterationCounter;
}
