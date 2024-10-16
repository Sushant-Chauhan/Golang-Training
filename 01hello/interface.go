// package customer

// //define structs
// type Car struct {
// 	Brand string
// 	Model string
//     Year int
// }

// type Truck struct {
// 	Brand string
// 	Model string
// 	Year int
// }

// type Bike struct {
// 	Brand string
// 	Model string
// 	Year int
// }

// // factory functions
// func NewCar(brand,model string, year int) *Car {
// 	return &Car {
// 		Brand: brand
// 		Model: model
// 		Year: year
// 	}
// }
// func NewTruck(brand,model string, year int) *Truck {
// 	return &Truck{
// 		Brand: brand
// 		Model: model
// 		Year: year
// 	}
// }
// func NewBike(brand, model string, year int) *Bike {
// 	return &Bike{
// 		Brand: brand
// 		Model: model
// 		Year: year
// 	}
// }
// func main()



/////// Eg: interface  /////

package main
import "fmt"

//interface
type geometry interface {
	area() float64
	perimeter() float64
}

//struct
type square struct {
	side float64
}
type rectangle struct {
	length, breadth float64
}

//method to implement interface square functions()
func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4*s.side
}

//method to implement interface perimeter functions()
func (r rectangle) area() float64 {
	return r.length*r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2*(r.length + r.breadth)
}

//Measure
func measure(g geometry){
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perimeter())
}

//---------  main function --------

func main() {
	fmt.Println("Running code from interface.go!")
	s:= square{side:10.0}
	r:= rectangle{length:10, breadth:5}
	measure(s)
	measure(r)
}