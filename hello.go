// 変数
// 変数名: 1文字目に注意

// 基本的なデータ型
/*
string      "hello"
int         53
float64     10.2
bool        false true
nil

var s string // ""
var a int // 0
var f bool // false
*/

package main

import "fmt"

// func hi(name string) string {
//     // fmt.Println("hi!" + name)
//     msg := "hi!" + name
//     return msg
// }

func hi(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return
}

// func swap(a, b int) (int, int) {
//     return b, a
// }


type user struct {
	name string
	score int
}

func (u user) show() {
	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

func (u *user) hit() {
	u.score++
}

type greeter interface {
	greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
	fmt.Println("Konnnichiwa!")
}
func (a american) greet() {
	fmt.Println("Hello!")
}

func show(t interface{}) {
	// 型アサーション
	// _, ok := t.(japanese)
	// if ok {
	//     fmt.Println("i am japanese")
	// } else {
	//     fmt.Println("i am not japanese")
	// }

	// 型Switch
	switch t.(type) {
	case japanese:
		fmt.Println("i am japanese")
	default:
		fmt.Println("i am not japanese")
	}
}

func main() {
	// // var msg string
	// // msg = "hello world"
	// // var msg = "hello world"
	// msg := "hello world"
	// fmt.Println(msg)

	// // var a, b int
	// // a, b = 10, 15
	// a, b := 10, 15

	// var (
	// 	c int
	// 	d string
	// )
	// c = 20
	// d = "hoge"
	// a := 10
	// b := 12.3
	// c := "hoge"
	// var d bool
	// fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)

	// 定数
	// const name = "taguchi"
	// name = "fkoji"
	// fmt.Println(name)
	// const (
	// 	sun = iota // 0
	// 	mon        // 1
	// 	tue        // 2
	// )
	// fmt.Println(sun, mon, tue)
	// 演算

	/*
		数値 + - * / %
		文字列 +
		論理値 AND(&&) OR(||) NOT(!)
	*/
	// var x int
	// // x = 10 % 3
	// // x += 3 // x = x + 3
	// x++ // x = x++ ++x
	// fmt.Println(x)

	// var s string
	// s = "hello " + "world"
	// fmt.Println(s)

	// a := true
	// b := false
	// // fmt.Println(a && b)
	// fmt.Println(a || b)
	// fmt.Println(!a)

	// ポインタ
	// アドレスを指し示す変数
	// 演算はできない
	// a := 5
	// var pa *int
	// pa = &a // &a = aのアドレス
	// // paの領域にあるデータの値 = *pa
	// fmt.Println(pa)
	// fmt.Println(*pa)

	// 関数
	// hi("taguchi")
	// fmt.Println(hi("taguchi"))
	// fmt.Println(swap(5, 2))

	// f := func(a, b int) (int, int) {
	//     return b, a
	// }
	// fmt.Println(f(2, 3))

	// func(msg string) {
	// 	fmt.Println(msg)
	// }("taguchi")

	// 配列
	// var a [5]int // a[0] - a[4]
	// a[2] = 3
	// a[4] = 10
	// fmt.Println(a[2])
	// b := [3]int{1, 3, 5}
	// b := [...]int{1, 3, 5}
	// fmt.Println(b)
	// fmt.Println(len(b))

	// スライス
	// a := [5]int{2, 10, 8, 15, 4}
	// s := a[2:4] // [8, 15]
	// s[1] = 12
	// fmt.Println(a)
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s := make([]int, 3) // [0 0 0]
	// s := []int{1, 3, 5}
	// // append
	// s = append(s, 8, 2, 10)
	// // copy
	// t := make([]int, len(s))
	// n := copy(t, s)
	// fmt.Println(s)
	// fmt.Println(t)
	// fmt.Println(n)

	// map
	// m := make(map[string]int)
	// m["taguchi"] = 200
	// m["fkoji"] = 300
	// m := map[string]int{"taguchi": 100, "fkoji": 200}
	// fmt.Println(m)
	// fmt.Println(len(m))
	// delete(m, "taguchi")
	// fmt.Println(m)
	// v, ok := m["fkoji"]
	// fmt.Println(v)
	// fmt.Println(ok)

	// if
	// score := 43

	// if score > 80 {
	// 	fmt.Println("Great!")
	// } else if score > 60 {
	// 	fmt.Println("Good!")
	// } else {
	// 	fmt.Println("so so...")
	// }

	// fmt.Println(score)

	// switch
	// signal := "blue"
	// switch signal {
	// case "red":
	//     fmt.Println("Stop")
	// case "yellow":
	//     fmt.Println("Caution")
	// case "green", "blue":
	//     fmt.Println("Go")
	// default:
	//     fmt.Println("wrong signal")
	// }

	// score := 82
	// switch {
	// case score > 80:
	// 	fmt.Println("Great!")
	// default:
	// 	fmt.Println("so so ...")
	// }

	// for
	// for i := 0; i < 10; i++ {
	//     // if i == 3 { break }
	//     if i == 3 { continue }
	//     fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	//     fmt.Println(i)
	//     i++
	// }

	//i := 0
	//for {
	//	fmt.Println(i)
	//	i++
	//	if i == 3 {
	//		break
	//	}
	//}

	// range
	// s := []int{2, 3, 8}
	// for i, v := range s {
	//     fmt.Println(i, v)
	// }
	// for _, v := range s {
	//     fmt.Println(v)
	// }
	//m := map[string]int{"taguchi":200, "fkoji":300}
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}

	// 構造体
	// u := new(user)
	// // (*u).name = "taguchi"
	// u.name = "taguchi"
	// u.score = 20

	// u := user{"taguchi", 50}
	//u := user{name:"taguchi", score:50}
	//fmt.Println(u)

	// メソッド（データ型に紐付いた関数）
	//u := user{name:"taguchi", score:50}
	//// fmt.Println(u)
	//u.hit()
	//u.show()

	// インターフェース
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
