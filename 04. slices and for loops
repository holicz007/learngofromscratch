package main

import (
	"fmt"
)

/*
	arrays : กำหนดขนาดชัดเจน

	slice : ไม่่จำเป็นต้องระบุขนาด (List) และ ไทป์ของสิ่งที่เก็บต้องเหมือนกันทั้งหมด
*/

func main() {

	/*
		การประกาศ slice
		[] ประเภทตัวแปร{values,values}
	*/

	cards := []string{newCard(), newCard(), newCard()}

	/*
		การเพิ่มค่าให้กับ slice
		ตัวแปรประเภท slice = append(ตัวแปรประเภท slice นั้น ๆ , values)

		** Note : เหมือนกับว่ามันไม่ได้แก้ไข slice ตัวก่อนแต่มันสร้างใหม่แล้วเอาไปทับของเดิม
	*/

	cards = append(cards, "Six of Spades")

	/*
		การนำค่าใน slice มาใช้งานด้วย for loops

		for อินเด็กซ์ , ชื่อตัวแปร (ค่าของช่องนั้น ๆ จะอยู่ในตัวแปรนี้) := range ตัวแปร slice ที่เราจะวน {
			
		}
	*/

	for i, card := range cards {
		fmt.Println(i, card)

	}

}

func newCard() string {
	return "Five of Diamonds"
}
