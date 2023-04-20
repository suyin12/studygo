package main

import "fmt"

type Zhifubao struct {
}

func (z *Zhifubao) Pay(amount int64) {
	fmt.Printf("使用支付宝支付了%.2f元\n", float64(amount))
	// fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

func CheckoutWithZhifubao(obj *Zhifubao) {
	obj.Pay(100)
}

type WeChatPay struct {
}

func (w *WeChatPay) Pay(amount int64) {
	fmt.Printf("使用微信支付了%.2f元\n", float64(amount))
}

func checkoutWithWeChatPay(obj *WeChatPay) {
	obj.Pay(200)
}

type Payer interface {
	Pay(int64)
}

func Checkout(obj Payer, x int64) {
	obj.Pay(x)
}

type Sayer interface {
	Say()
}

func main() {
	// CheckoutWithZhifubao(new(Zhifubao))
	// checkoutWithWeChatPay(&WeChatPay{})

	// Checkout(new(Zhifubao), 200)
	// Checkout(&WeChatPay{}, 1000)

	// var x Sayer
	var m Mover
	c := &Cat{}
	d := Dog{}
	d2 := &Dog{}

	// x = c
	// x.Say()
	// x = d
	// x.Say()

	m = c
	m.Move()
	m = d
	m.Move()
	m = d2
	m.Move()
}

type Cat struct {
}

func (c Cat) Say() {
	fmt.Println("cat say")
}

type Dog struct {
}

func (c Dog) Say() {
	fmt.Println("dog say")
}

type Mover interface {
	Move()
}

func (d Dog) Move() {
	fmt.Println("Dog Move")
}

func (c *Cat) Move() {
	fmt.Println("Cat Move")
}
