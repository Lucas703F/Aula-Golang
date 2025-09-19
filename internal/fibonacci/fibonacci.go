package fibonacci

func Fibonacci(n int) int {
	if (n == 1) {// primeiro termo
        return 0;
    } else if (n == 2) { // segundo termo
    	return 1;
	}
    
	return Fibonacci(n - 1) + Fibonacci(n - 2);
}