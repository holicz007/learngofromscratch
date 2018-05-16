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