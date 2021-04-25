package Step2

import (
	"fmt"
)

// Class2 ... Golang內建的資料型態
func Class2() {
	// Boolean 布林型態，預設為false
	var valBoolean bool
	fmt.Printf("%T 型態變數，如果沒有初值，就會預設false , %v\n", valBoolean, valBoolean)

	// 數值型別，依長度來區分，一個大重點：不用型別之前不可以相互進行操作，除非先轉換
	// 數值定義的初值為0
	var intVal int
	var intVal2 int32 = 32
	fmt.Printf("intVal(%T)=%d,intVal2(%T)=%d,兩者相加int32(intVal)+intVal2=%d\n", intVal, intVal, intVal2, intVal2, (int32(intVal) + intVal2))

	// 浮點數只有float32跟float64，如果使用簡寫定義的話，預設會是float64
	floatVal := 3.14
	fmt.Printf("floatVal(%T)=%f\n", floatVal, floatVal)

	// 字串定義上，預設的定元變編是為UTF-8，可使用雙引號 「"」或是大寫單引號來「`」定義，若無定義初值，為設為空""

	var str1 string
	str2 := "this is string"
	fmt.Printf("str1=%s,str2=%s\n", str1, str2)

	// 如果要更動字串內的單一內容，必需要先轉成[]byte，再根據[]byte陣列的索引來改變，最後再轉回string輸出
	// 因為[]byte會被視為字元，所以要用單引號「'」來給值
	str3 := []byte(str2)
	str3[0] = 't'
	str3[2] = 'a'
	str3[3] = 't'
	str2 = string(str3)
	fmt.Printf("改變後就會變成這樣：str2=%s\n", str2)

	// 也可以使用切片(slice)的方式進行更動
	str := "hello"
	str = "b" + str[1:]
	fmt.Printf("改變後字串=>%s\n", str)

	// 可以透過「`」來定義一個多行的字串
	str = `這是第一行
	這是第二行`
	fmt.Printf("%s\n", str)

	// 連結兩個字串可以用「+」來做連結
	str1 = "this is str1"
	str2 = " this is str2"
	fmt.Printf("str1 + str2 = %s\n", str1+str2)

	//  分組宣告寫法
	const (
		CONST1 string  = "CONST1 string"
		CONST2 int64   = 12
		CONST3 float64 = 3.14159
	)
	fmt.Printf("分組宣告：CONST1(%T)=%v , CONST2(%T)=%d , CONST3(%T)=%f\n", CONST1, CONST1, CONST2, CONST2, CONST3, CONST3)

	// iota 列舉
	// 在分組宣告裡，iota會自動從0往上加
	// iota每碰到一次const就會變0
	const (
		a = iota
		b = iota
		c = iota
	)
	const d = iota
	fmt.Printf("a=%v,b=%v,c=%v,d=%v\n", a, b, c, d)
}
