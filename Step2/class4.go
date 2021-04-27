package Step2

import "fmt"

func Class4() {

	// 正常的if else
	x := 10
	if x >= 10 {
		fmt.Println("x >= 10")
	} else {
		fmt.Println("x < 10")
	}

	// 也可以在if裡面宣告一個變數，但是只能宣告一個
	// 這一個變數只能作用在這一個if else的範圍
	if a := x; a >= 10 {
		fmt.Println("a >= 10")
	} else {
		fmt.Println("a < 10")
	}

	// 多個條件的巢狀結構
	if x > 10 {
		fmt.Println("x > 10")
	} else if x == 10 {
		fmt.Println("x = 10")
	} else {
		fmt.Println("x < 10")
	}

	// goto 跳躍執行

	goto Here
	fmt.Println("理論上這裡不會執行")
Here:
	fmt.Println("這裡是HERE區塊")
	goto Then
Then:
	// for
	sum := 0
	for i := 0; i < 10; i += 1 {
		sum += i
		fmt.Printf("我是第%d圈，目前total為：%d\n", i, sum)
	}

	// 也可以把for裡面的初始條件跟前進條件去掉，變成只有執行條件
	for sum < 100 {
		sum += 5

		if sum >= 90 && sum < 95 {
			fmt.Println("sum >= 90 => continue")
			continue
		} else if sum >= 95 {
			fmt.Println("sum >= 95 => break")
			break
		}
		fmt.Printf("目前total為：%d\n", sum)
	}

	// 用slice或map跑loop的時候，可以使用range
	mapping := make(map[string]int)
	mapping["one"], mapping["two"], mapping["three"] = 1, 2, 3
	for key, value := range mapping {
		fmt.Printf("key %s = %d\n", key, value)
	}

	// 如果在需要接收多回傳值，但又用不到的時候可以用底線來接
	for _, value := range mapping {
		fmt.Printf("因為沒有key，所以值接輸出value => %d\n", value)
	}

	// switch的使用
	// Golang裡面switch case區塊裡面預設有break，所以就不用再加上去
	// 但是我們可以使用「falltrouph」來強制它往下做下一個case裡面的內容(無論有沒有合規則都會執行)，最後default也會被執行進去
	var str string = "a"
	switch str {
	case "a":
		fmt.Println("a")
		fallthrough
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("not in (a,b,c)")
	}

}
