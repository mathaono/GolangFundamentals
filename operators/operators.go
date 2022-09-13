package operators

type Numbers struct {
	num1 int64
	num2 int64
}

func Operators(number1, number2 int64) int64 {
	n1 := Numbers{}
	n1.num1 = number1
	n1.num2 = number2
	soma := n1.num1 + n1.num2
	return soma
}
