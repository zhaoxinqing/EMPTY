package lib

import "fmt"

// 不定参数
func foo(args ...int) { // 接受不定数量的参数，这些参数都是int类型
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func TestRun() {
	foo(2, 3, 4)
	foo(1, 3, 7, 13)
}

// 不定参数传任意类型，可以指定类型为interface{}，如标准库中的fmt.Printf()的函数原型：
func Printf(format string, args ...interface{}) {
	fmt.Println("arg")
}

func foo2(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case float32:
			fmt.Println(arg, "is a float32 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
