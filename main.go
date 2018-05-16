package main

/*
	** package ต้องอยู่บรรทัดแรกของไฟล์เท่านั้น **

	package หมายถึง collections ของไฟล์พื้นฐานต่าง ๆ ภายในโปรเจ็ค

	package แยกได้ 2 แบบคือ

	1. Executable Package : main เวลากด build จะได้ไฟล์ exe

	2. Reusealble Packgage : lib ex. package calculator

*/

import "fmt"

/*
	import ใช้ในการเข้าถึง code ต่าง ๆ ที่อยู่นอกเหนือ package ปัจจุบัน
*/

func main() {

/*
	func : หมายถึงการประกาศ functions 
*/

	fmt.Println("Hello")

}

/*
	เราสามารถ run Go code ของเราได้อย่างไร ?
	ตอบ. เราสามารถ run Go code ของเราผ่าน Terminal โดยใช้คำสั่ง go run ตามด้วยชื่อไฟล์ (เข้าไปที่ directory ของโปรเจ็คก่อน)
*/

/*
	Go Command Lint (CLI)

		go build --> ใช้ complie อย่างเดียวแต่ยังไม่ execute จะได้ไฟล์ .exe มาหลังจากการ build

		go run -->	ใช้ run Go code ของเรา (พร้อม execute code ของเรา)

		go fmt --> ใช้จัด format ของ code ภายใน current directory

		go install --> ใช้ในการควบคุม Dependency

		go get --> ใช้ในการควบคุม Dependency

		go test --> ใช้ run test ที่เกี่ยวข้องภายในโปรเจ็ค
*/
