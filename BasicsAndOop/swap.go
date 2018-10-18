package main

func swapReturn(x, y int) (int, int) {
	return y,x
}

func swapPointer(x, y *int) {
	*x, *y = *y, *x
}