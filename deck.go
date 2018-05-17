package main

import (
	"fmt"
)

/*
	ถ้าหากเราต้องการสร้างอะไรให้เรามองว่ามันเป็นประเภทตัวแปรอะไรก่อนโดยดูจากภาพรวม
	หลังจากนั้นให้ทำการ custom type ด้วยการ

	type ชื่อไทป์ ประเภทของไทป์
*/

type deck []string /* เสมือนประกาศคาสอ่ะ */

/*
	methods receiver : หมายถึงให้ตัวแปรอะไรก็ได้ที่มีไทป์ตรงกับ receiver สามารถเรียกใช้งานได้
	จากตัวอย่างด้านล่าง สร้าง methods ชื่อ print โดยมี receiver ของ deck (ชื่อไทป์)
	ต่อไปนี้ถ้ามีตัวแปรอะไรมีไทป์เป็น deck จะสามารถใช้ methods print ได้
*/

func (d deck) print() {
	/*
		d ใน go นิยมตั้งตัวแปรของ receiver เป็นตัวแปรเดี่ยว ๆ
		d คืออะไร d ก็คือ deck deck คือไทป์ที่ทำการสืบทอดมาจาก slice string
		ก็เลยทำให้มันสามารถวนอ่านข้อมูลจาก d ได้ กล่าวคือมันก็เหมือนกับ this ในจาวาอะ
	*/
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
	functions นี้ไม่จำเป็นต้องใช้ Receiver Fucntions เพราะ เป็นขั้นตอนการสร้าง
	ต้องสรา้งก่อนสิถึงจะเรียกใช้ methods ของตัวมันเองได้ถูกมั้ย ?
*/

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	/*
		สังเกตุที่ _ ใน loop ปกติมันจะต้องเป็นค่าของ index แต่ถ้าเราไม่ได้ใข้ตัวแปรนั้น ๆ
		Go จะบอกว่ามัน Error เราสามารถไม่ใช้ค่านั้น ๆ ได้ด้วยการใช้ _ ที่ตัวแปรที่มีการคืนค่า
		หรือส่งค่าได้เพื่อบอก Go ว่าเราไม่ต้องการที่จะใช้ค่านั้น ๆ
	*/
	for _, suit := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suit)
		}
	}

	return cards

}

/*
	การประกาศ Parameters ของ functions ใน Go
	(ชื่อตัวแปน ประเภทของตัวแปร , ชื่อตัวแปร ประเภทของตัวแปร)
*/

/*
	Go functions สามารถคืนค่าได้มากกว่า 1 ค่าโดยสามารถทำได้ดังนี้

	func functionName(some parameter) (return values 1 , return values 2) {
		// do somethings stuff

		return values1, values2
	}

*/

/*
	การรับค่าจาก functions ที่มีการ return หลายค่า
	ตัวแปร1 , ตัวแปร2 := function() (v1 , v2)

		ตัวแปร1 รับค่าจาก v1
		ตัวแปร2 รับค่าจาก v2

*/

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
