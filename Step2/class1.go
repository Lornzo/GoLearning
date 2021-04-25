package Step2

import "fmt"

// Class1 ... 變數定義
func Class1() {
	// 定義單一變數：var varName {type}
	var varName string
	fmt.Println("單一變數定義，並不給予初值，變數內容 varName = ", varName)

	// 定義變數並給予初始值
	var varName2 string = "this is string"
	fmt.Println("單一變數定義，並不給予初值，變數內容 varName = ", varName2)

	// 定義多個變數：var var1,var2,var {type}
	var var1, var2, var3 string
	fmt.Printf("多變數定義，並不給予初值，變數內容 var1 = %s ,var2 = %s ,var3 = %s \n", var1, var2, var3)

	// 定義多個變數，並給予初值
	var val1, val2, val3 string = "it's val1", "it's val2", "it's val3"
	fmt.Printf("多變數定義，並給予初值，變數內容 val1 = %s ,val3 = %s ,val3 = %s \n", val1, val2, val3)

	// 簡化定義變數，編譯器會根據初始化的值自動推匯出相應的型別
	// 但是簡寫定義變數只能被用在function裡面，如果在function外定義變數還是需要使用var val {type} 的寫法
	value1, value2, value3 := "it's string", 6, true
	fmt.Printf("多變數定義，並給予初值，變數內容 value1 = %s ,value3 = %d ,value3 = %v \n", value1, value2, value3)

	// 底線定義：底變變數如果初指定值的話，會直接把指定的值丟掉，並不保存。
	// 但是底變可以用在接收多個以上return 的function時，並把不需要用到的return 丟掉
	_, str := returnTwo()
	fmt.Println("如果直接印出底線的話會報錯，但是可以得到str = ", str)

	// 常數定義：const CONST_VALUE {type}
	// 也可以不加資料型態定義常數
	const CONST_VALUE string = "123"
	const CONST_VALUE_2 = 123
	fmt.Printf("常數CONST_VALUE = %s,常數CONST_VALUE2 = %v\n", CONST_VALUE, CONST_VALUE_2)

}

//returnTwo ...
func returnTwo() (string, string) {
	return "str1", "str2"
}
