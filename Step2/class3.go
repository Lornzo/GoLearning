package Step2

import "fmt"

// Class3 ... array,slice,map
func Class3() {

	// 陣列定義：var arr [n]type
	// 定義一個長度為3的int陣列，並給予值
	// 長度也是陣列型別的一部份，∴[3]int != [4]int
	// 陣列一但被定義完成，是無法改變長度的
	var arr [3]int
	arr[1] = 1
	arr[2] = 2
	fmt.Printf("arr => ")
	fmt.Println(arr)

	// 也可以這樣定義陣列，這樣就不用寫死陣列大小，但是還是沒有辦法在定義完陣列後改變陣列大小
	arr2 := [...]int{1, 2, 3}
	fmt.Printf("arr2 => ")
	fmt.Println(arr2)

	// 定義多維陣列
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{2, 2, 3, 4}}

	// 也可以簡寫
	dArray := [2][4]int{{1, 2, 3, 4}, {2, 2, 3, 4}}
	_ = dArray

	fmt.Printf("多維陣列長這樣 => ")
	fmt.Println(doubleArray)

	// 把陣列帶入Function的時候，傳入的其實是該Array的副本，而不是它的指標(Pointer)。

	// 動態陣列(Slice)
	// Slice並不是真正意義上的動態陣列，而是一個參考型別。Slice總是指向一個底層Array
	// slice像一個結構體，這個結構體包含三個元素
	// 指標：指向陣列指定的開始位置
	// 長度
	// 最大長度

	var sliceArray = []byte{'a', 'b', 'c', 'd'}
	fmt.Println("sliceArray => ", sliceArray)

	// 可以直接存取Slice，使用 slice[start:end]
	// 存取slice 3 到 4
	fmt.Println("sliceArray[2:4] => " + string(sliceArray[2:4]))

	// slice的基本操作
	//  slice預設開始位置是0，在讀取時可以這樣
	//  slice預設最後一個值就是整個長度
	fmt.Println("sliceArray[:4] => " + string(sliceArray[:4]))
	fmt.Println("sliceArray[1:] => " + string(sliceArray[1:]))

	// slice是參考型別，當參考改變其中的元素值的時候，其它所有的參考都會改變
	sliceArray2 := sliceArray
	fmt.Printf("定義一個sliceArray2，使之等於sliceArray：sliceArray=>%s,sliceArray2=>%s\n", string(sliceArray), string(sliceArray2))

	// 只更動sliceArray
	sliceArray[0] = 'A'
	fmt.Printf("改變sliceArray[0]='A'：sliceArray=>%s,sliceArray2=>%s\n", string(sliceArray), string(sliceArray2))

	// 動態改變長度
	// 如果參考的其它陣列長度不足的話，就會參考原大小的位置，所以其它的slice不會被改變，除非在增加原slice長度之前，先增加其它slice長度的位置
	sliceArray = append(sliceArray, 'e')
	fmt.Printf("為sliceArray增加一個元素：sliceArray = append(sliceArray,'e') ： sliceArray=>%s,sliceArray2=>%s\n", string(sliceArray), string(sliceArray2))

	// map操作
	// map可以用在key/value的操作

	// 基本定義法中，要在開始要給map值之前，用make一下，預先在記憶體內「分配」一個空間給map使用，不然會報錯
	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	fmt.Printf("然後map就變成這樣了")
	fmt.Println(numbers)

	// map的特性：
	// 1. map是無序的，所以每次顯示出來的map都有可能不一樣
	// 2. map只能透過key來存取
	// 3. 跟slice一樣，map也是一個參考型別，也就是說，如果map改變元素的時候，其它參考它的「map」也會跟著改變
	// 4. map的長度不是固定的，所以想加就加
	// 5. 最重要的是：map不是thread-safe，所以在多執行緒的時候，需要使用資料鎖來確表存取安全

	number := numbers
	fmt.Printf("定義一個number變數，並參考numbers[one]： number=>%d,numbers[one]=>%d\n", number["one"], numbers["one"])

	numbers["one"] = 999
	fmt.Printf("「只」改變numbers[one]之後：number=>%d,number[one]=>%d\n", number["one"], numbers["one"])

	delete(numbers, "one")
	fmt.Printf("刪除一個map元素 => numbers ")
	fmt.Println(numbers)
	fmt.Printf("參考numbers的map => number")
	fmt.Println(number)

	// new就可以用於各種型別的記憶體分配，並return 一個用預設初直填充的指標
	// make 跟 new的區別
	str := new(string)
	fmt.Printf("str => %v , *str => %s\n", str, *str)

	// make 只能用在建立slice、map還有channel的記憶體分配，並return 一個有初始值的type型別，而不是指標
	// 這三個型別必需使用make的原因，是因為指向資料結構的指標在使用前必需要被初始化
	// 如果slice、map、還有channel沒有使用make來初始化，它們的初值就會是nil
	var mapexample map[string]int
	fmt.Println(mapexample)

	// Golang的零值
	// 「變數未填充前」的預設值
	var integer int
	fmt.Printf("%T => %v\n", integer, integer)

	var integer8 int8
	fmt.Printf("%T => %v\n", integer8, integer8)

	var integer32 int32
	fmt.Printf("%T => %v\n", integer32, integer32)

	var integer64 int64
	fmt.Printf("%T => %v\n", integer64, integer64)

	var uinteger uint
	fmt.Printf("%T => %v\n", uinteger, uinteger)

	var runeVal rune
	fmt.Printf("%T => %v\n", runeVal, runeVal)

	var byteVal byte
	fmt.Printf("%T => %v\n", byteVal, byteVal)

	var floatNumber32 float32
	fmt.Printf("%T => %v\n", floatNumber32, floatNumber32)

	var floatNumber64 float64
	fmt.Printf("%T => %v\n", floatNumber64, floatNumber64)

	var boolean bool
	fmt.Printf("%T => %v\n", boolean, boolean)

	var strVal string
	fmt.Printf("%T => %v\n", strVal, strVal)

}
