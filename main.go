package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"syscall/js"
)

//go:export add
func Add2(x, y int) int {
	return x + y
}

// 给字符串生成md5
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @return str 返回md5码

//go:export Md5Crypt
func Md5Crypt(str string, salt ...interface{}) string {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

const salt = "ftmsabcd@1234!"

func MD5_SALT(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

//export sign2
func MD5_SALT2(str string) string {
	fmt.Println("============== begin")
	fmt.Println(str)
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	result := hex.EncodeToString(h.Sum(nil))
	fmt.Println(result)
	fmt.Println("============== end")
	return result
}

//export multiply
func multiply(x, y int) int {
	fmt.Println(x, y)
	return x * y
}

func onInput(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(map[string]interface{}{
		"key":    60,
		"remove": 1,
	})
}

func Hello(this js.Value, args []js.Value) interface{} {
	message := args[0].String() // get the parameters
	return "Hello " + message
}

func sign(this js.Value, args []js.Value) interface{} {
	fmt.Println("============== begin")
	fmt.Println(args[0].String())
	b := []byte(args[0].String())
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	result := hex.EncodeToString(h.Sum(nil))
	fmt.Println(result)
	fmt.Println("============== end")
	return result
}

func main() {
	fmt.Println("hello wasm123 ...")

	js.Global().Set("onInput", js.FuncOf(onInput))
	js.Global().Set("Hello", js.FuncOf(Hello))
	js.Global().Set("sign", js.FuncOf(sign))

	message := "👋 Hello World 🌍"

	document := js.Global().Get("document")
	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", message)
	document.Get("body").Call("appendChild", h2)

	<-make(chan bool)

	// h := md5.New()
	// h.Write([]byte("sharejs.com"))                     // 需要加密的字符串为 sharejs.com
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果

	// str := MD5_SALT2("sharejs.com")
	// fmt.Println(str)
}
