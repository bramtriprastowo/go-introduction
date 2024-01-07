package main

import (
	"fmt"
)

type messageToSend struct {
	message     string
	phoneNumber int
}

func tests(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %d\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

// func main() {
// 	tests(messageToSend{
// 		message:     "Halo",
// 		phoneNumber: 123123,
// 	})
// }

// Nested Structs
type car struct {
	model  string
	height int
	wheel  wheel
}

type wheel struct {
	radius   int
	material string
}

// Embedded Structs
// type car struct {
// 	model  string
// 	height int
// 	wheel
// }

// type wheel struct {
// 	radius   int
// 	material string
// }

// func main() {
// 	car1 := car{}
// 	car1.wheel.radius = 14
// }

// func main() {
// 	supra := car{
// 		model:  "sedan",
// 		height: 26,
// 		wheel: wheel{
// 			radius:   14,
// 			material: "aluminium",
// 		},
// 	}
// 	fmt.Printf(supra.model)
// }

// Anonymouse struct
// myCar := struct {
// 	Make string
// 	Model string
//   } {
// 	Make: "tesla",
// 	Model: "model 3",
//   }

// type car struct {
// 	Make string
// 	Model string
// 	Height int
// 	Width int
// 	// Wheel is a field containing an anonymous struct
// 	Wheel struct {
// 	  Radius int
// 	  Material string
// 	}
//   }

// Method for struct
// type rect struct {
// 	width  int
// 	height int
// }
// func (r rect) area() int {
// 	return r.width * r.height
// }
// var rect1 = rect{
// 	width:  5,
// 	height: 3,
// }
// func main() {
// 	fmt.Println(rect1.area())
// }
