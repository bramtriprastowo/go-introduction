package main

// Interfaces are collections of method signatures

// type shape interface {
// 	area() float64
// 	perimeter() float64
// }

// func showAreaAndPerimeter(sh shape) {
// 	fmt.Printf("The area is %.2f and the perimeter is %.2f", sh.area(), sh.perimeter())
// }

// type rect struct {
// 	width, height float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }
// func (r rect) perimeter() float64 {
// 	return 2*r.width + 2*r.height
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func test1(s shape) {
// 	showAreaAndPerimeter(s)
// }

// func main() {
// 	test1(circle{
// 		radius: 7,
// 	})
// }

// Interface -> Give Arguments
// type moveFile interface {
// 	move(sourceFile string, destinationFile string) (fileName string)
// }
