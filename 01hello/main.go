package main
import ( 	
	"fmt" 
	"sort" 
	"strings"
	"math"
// 	"bufio"
//     "os"
// 	"strconv"
// "strings"         //string.TrimSpace(input)
)
// var tkn1 = 54
// var tkn2 = 67
// const My string ="Sushant"

func main() {

	var xx int          // Declares x as an int, default value is 0
	var yy float64      // Declares y as a float64, default value is 0.0
	var nam string     // Declares name as a string, default value is an empty string

	fmt.Println("Variables are : ", xx,yy,nam)

	// // Multiple declarations
	var aa, bb, cc int
	fmt.Println("Variables are : ", aa,bb,cc)

	// It can only be used within functions, not at the package level.
	x1 := 10           // Declares x and initializes it to 10
	namm := "Alice"   // Declares name and initializes it to "Alice"

	// Multiple declarations
	a2, b2 := 1, 2      // Declares and initializes a and b
	fmt.Println("Variables are : ", x1, a2,b2,namm)


	// fmt.Println()
	// fmt.Println("tkn1= ",tkn1)
	// fmt.Println("tkn2=",tkn2)
	// fmt.Println("My = ",My)

	// fmt.Println("Hello from lco")
	// var usrname string="sushi"
	// fmt.Println(usrname)
	// fmt.Printf("Variable is of type : %T \n", usrname)

	// var isLogged bool = false
	// fmt.Println(isLogged)
	// fmt.Printf("Variable is of type : %T \n", isLogged)

	// var smallVal int = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type : %T \n", smallVal)

	// var f1 float64 = 255.534645647
	// fmt.Println("f1 =", f1)
	// fmt.Printf("Variable is of type : %T \n", f1 )
 
	// var anotherv int 
	// fmt.Println("----\nanotherv = ",anotherv)
	// fmt.Printf("Variable is of type : %T \n", anotherv)

	// //implicit
	// var website = "abc.com"
	// fmt.Println(website)

	// //no var styled
	// noOfUser := 400003.99
	// fmt.Println(noOfUser)

	// NEW LECTURE  

	// welcome := "---- welcome to user input ----"
	// fmt.Println(welcome)

 	// reader :=  bufio.NewReader(os.Stdin)
	// fmt.Print("Enter rating for our Pizza : ")
	// input, _ := reader.ReadString('\n') //comma ok || err ok -  syntax
	// fmt.Print("Thks for rating ", input)
	// fmt.Printf("Type of rating is %T\n", input)
	// // Declare numRating and err
	// var numRating float64
	// var err error
	// numRating, err = strconv.ParseFloat(strings.TrimSpace(input),64)
	// if err!=nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added 1 to your rating:  ", numRating+1)
	// }

	// fmt.Println("----- WElcomee to our pizza app ----")
	// fmt.Print("Plz rate our pizza between 1/ Declares an array of 5 integers to 5 : ")
	// reader2 := bufio.NewReader(os.Stdin) // creates a new bufio.Reader that reads from standard input (usually the keyboard).        	take input from user and store it in refrence variable
	// input2, _ := reader2.ReadString('\n')  //reads input from standard input until a newline character (\n) is encountered, and it returns that input as a string.     => it returns 2 outputs
	// fmt.Println("Thkss for rating ", input2)
	// var numRating2 float64
	// var err2 error
	// numRating2,_ = strconv.ParseFloat(strings.TrimSpace(input2), 64) 
	// if err2!=nil {
	// 	fmt.Println(err2)
	// } else {
	// 	fmt.Println("Addded 1 to your rating : ", numRating2 + 1)
	// }

	// NEW LECTURE  

// ARRAY
	// var age [3]int = [3]int{55,66,77}
	var age = [3]int{44,55,66}
	agee := [3]int{11,22,33}
	ageee := [...]int{1, 2, 3}
	aage := [6]int{2:44,5:55}   
	strArr := [3]string{"Hello", "World", "Go"}
 	fmt.Println("age,agee,ageee,aage, strArr== ",age,agee,ageee, aage, strArr)
 
	var arr [3]int   //most common
	arr[2] = 10          
	arr[1] = 15
	arr1 := arr    //copy arr
	fmt.Println("arr, arr1  = ",arr, arr1 )

	//loop
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	} 
	
	// slice , sort
	ages := []int{45,67,2,33,9,11,55,23}
	sort.Ints(ages)
	fmt.Println("ages = ",ages)
	index := sort.SearchInts(ages,10)
	fmt.Println(index)

	fmt.Println("-----------------------")

	//strings
	greeting := "good morning frands!"
	fmt.Println(strings.Contains(greeting,"good"))
	fmt.Println(strings.ReplaceAll(greeting,"morning","afternoon"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting,"uuuuu"))
	fmt.Println(strings.Split(greeting," "))
	fmt.Println("Original value == ", greeting)

	names := []string {"yash","raj","varun","singh","harsh"}
	sort.Strings(names)
	fmt.Println("names = ",names)
	fmt.Println(sort.SearchStrings(names,"singh"))


	//while loop
	x := 0
	for x<5 {
		fmt.Println("value of x : ",x)
		x++
	}
	fmt.Println("-----")

	//for loop -----
	for i:=0; i<5; i++ {
		fmt.Println("value of i : ",i)
	}
	fmt.Println("----") 

	//loop through slice of strings - 
	// names := []string {"yash","raj"," varun","singh","harsh"}
	for i:=0; i<len(names);i++ {
		fmt.Println(names[i])
	}
	fmt.Println()

	//for-in loop -----
	//names:=[]string{"yash","raj"," varun","singh","harsh"}
	for index,value := range names {
		fmt.Printf("value at index %v is %v\n",index,value)
	}
	fmt.Println("----")
	for _,value := range names {
		fmt.Printf("value  is %v\n",value)
		value = "newstrieng"   //local to loop, doesn't change value originaly
	}
	fmt.Println("----")
	fmt.Println(names)
	fmt.Println("----")

	umar := 25
	fmt.Println(umar <=50)
	fmt.Println(umar >=50)
	fmt.Println(umar == 45)
	fmt.Println(umar != 50)

	if umar<30 {
		fmt.Println("umar is less than 30")
	} else if umar<40{
		fmt.Println("umar is less than 40")
	} else{
		fmt.Println("umar is not less than 45")
	}

	// names := []string{"yash","raj"," varun","singh","harsh"}
    for index, value := range names {
		if index==1 {
			fmt.Println("continuing at pos ", index)
			continue
		}
		fmt.Printf("the value at pos %v is %v \n", index,value)
	}
    fmt.Println()	


//////////////////////////////////  maps //////////////////////////////////
// 1. Basic Map creation
	population := make(map[string]int)
	population["India"] = 312
	population["USA"] = 444
	population["China"] =55
	fmt.Println("Population of India ::",population["India"])

	//checking if a value exists in map or not
	fmt.Println(" ------ Checking Key exists or not ------ ")
	cntry := "FF"
	popul, exists := population[cntry]
	fmt.Printf("popul= %v, exists=%v",popul, exists) //popul = 0, exists = false
	fmt.Println(" ------>", population["aa"])
	if exists {
		fmt.Printf("Population of %s : %d\n", cntry,popul)
	} else {
		fmt.Printf("%s is not in map storage \n",cntry)
	}

	//deletion
	fmt.Println(" ------ Deletion ------ ")
	delete(population,"USA")

	//iterating
	fmt.Println(" ------ Map eleemnts are ------ ")
	for c,p := range population {
		fmt.Printf("%s: %d\n" ,c,p)
	}


///////// 2. Declaring with Initial values
	capitals := map[string]string {
		"India" : "New Delhi",
		"France": "Paris",
		"Japan" : "Tokyo",
		"Canada": "Toronto",
	}

	for country,capital := range capitals {
		fmt.Printf("the capital of %s is %s\n",country,capital)
	}

///////// Count Frequency of Numbers
    numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5} 
	
	frequency := make(map[int]int)
	for _,num := range numbers {
		frequency[num] += 1        		//requency[numbers[i]] +=1 
	}
	fmt.Println("frequency : ",frequency)

	///or

	mymap := make(map[int]int)
	for i:=0; i<len(numbers); i++ { 
		n := numbers[i]
		_, exists := mymap[n]
		if exists {
			 mymap[n]++
		} else {
			mymap[n] = 1
		} 
	}
	fmt.Println("mymap : ",mymap)


//////////// Value & Reference  ////////////

// Composite types (like maps, slices, and channels) are also passed by value, but they behave like "references" since the value is a reference to the underlying data structure
	myMapp := map[string]int{"initialKey":50}
	modifyMap(myMapp)
	fmt.Println(myMapp)

	// myslice := [3]int{45,66,7}
	// fmt.Println("myslice = ",myslice)
	// modifyArray(myslice)
	// fmt.Println("myslice = ",myslice)

// Basic types (like int, float, string) are passed by value, so a copy is made, and changes in the function don’t affect the original variable.
	// Explicitly modify a basic type, can pass a POINTER to that VALUE
	a := 10
    modifyPointer(&a)
    fmt.Println(a) 

	fmt.Println()
 


///// struct Objects in main function - 4 ways ////
 
	var p1 Person          // Value type
	p1.name = "Wilyam"
	p1.age = 56
	p1.job = "Student"
	p1.salary = 6000 
	
	p2 := new(Person)    // Refrence type
	p2.name = "Hege"
	p2.age = 45
	p2.job = "Teacher"
	p2.salary = 60000
				// Struct Literal : Value type   
	p3 := Person{name: "sushant", age: 30, job: "engineer", salary:9000000}   //3.
		//or	// Shorthand Syntax : Value type
	p4 := Person{"varun", 30, "doctor",9000000}
	fmt.Println("p3 = ",p3,"\np4 = ",p4)
	fmt.Println(p1.Greet()) // myObject.myMethod(); --  Invoked using the object or class instance followed by the dot operator (.)  --Functions that are defined with a receiver (like func (c Circle) Area()) are considered methods.   --Outputs: Hello, my name is Alice and I am 30 years old.

// Eg2
	s1:=Student{"mark",22,"2020-01-01"}
	s1.ReceiverFunc() //oop pattern i.e calling object public members using DOT Operator
	s2:=Student{"ee",242,"2025-01-01"}
	nname := s2.ReceiverFunc() 
	fmt.Println("nname = ",nname)

// ---- main block code ----  // Factory Design Pattern 
//sc := SedanCar{Name:"Honda City"}
// sc := getCar()  
// fmt.Println(sc)          	//CLIENT 
	getCarFactory(1)    //CLIENT
	getCarFactory(2)    // doesn't know what actual objects are instentiated behind the scene - just know i need to pass this paramater actual logic

   

}
 // -------------------Outside main-----------
// type SedanCar struct {    // 1.Struct  // APPP CODE
// 	Name string
//  } 
// func getCar() SedanCar {         // FACTORY
// 	return SedanCar{Name:"Honda City"}
// }

// func getNewSedan() *SedanType{      // FACTORY
// 	return &SedanType
// }
//---------
// APP (actual logic)
type Car interface {     //car Interface
	getCar() string
}

type SedanType struct {   //1. Struct
	Name string  
}
type HatchBackType struct {  //2. Struct
	Name string  
} 
func getNewSedan() *SedanType {     
	return &SedanType{} // return SedanType OR object of Struct
}
func getNewHatchBack() *HatchBackType {    //  NoFACTORY
	return &HatchBackType{} // return object of Struct
}
// getCar() Method - getter for name
func (s *SedanType) getCar() string {  //1.Implemented Car Interface on SedanType struct 
	return "Honda City V2"
}
func (h *HatchBackType) getCar() string {  //2.Implemented Car Interface on HatchBackType struct 
	return "Swift VDI"
}
// FACTORY CLASS/OBJECT
func getCarFactory(cartype int) {   
 	var car Car          //both struct implements Car interface
	if cartype == 1 {
		car = getNewHatchBack()
	} else {
		car = getNewSedan()
	}
	carInfo := car.getCar()
	fmt.Println(carInfo)
}

/////////  Structs in GOlang    /////////
// DEFINATION: Structs are used to create complex data types that model real-world entities, allowing you to encapsulate related data. Collection of Other-datatypes together. 
func (s Student)  ReceiverFunc() string  {
	// fmt.Println(s)   --without return value string
	return "Name in struct object is "+ s.name
} 
type Student struct {
	name string
	age  int
	dob  string
}

type Person struct {
	name string
	age int
	job string
	salary int
  }
// Methods on Structs
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.name, p.age)
}
// NOTE: Methods are functions that are associated with a specific type (usually a struct). They have a special receiver argument that specifies the type they are associated with.

/*
Example 2 :
----------
type Circle struct {          //circle struct
    Radius float64
}
func (c Circle) Area() float64 {   //area method
    return 3.14 * c.Radius * c.Radius
}

c1 := Circle(5)
fmt.Println(c1.Area())
*/

//---------- OUTSIDE MAIN FUNCTION ----------

func modifyPointer(p *int){
	*p = 100
}
// func modifyArray(var a [3]int ){
// 	a[0]=99
// }
// see array , slice methods  ????????    https://go.dev/tour/moretypes/10    SEE BEHAVIOUR

func modifyMap(m map[string]int) {
	m["newKey"] = 100
}
 

// 	// //////////////  interface ////////////////
// 	shapes := []shape{
// 		square{length: 15.3},
// 		circle{radius:3}
// 	}
// 		// rectangle{length:4, breadth:5}
	
// 	for _,v := range shapes {
// 		printShapeInfo(v)
// 		fmt.Println("---")
// 	}
// 	// a1 := square(3)
// 	// fmt.Println(printShapeInfo(a1))	
// }

// //shape interface -groups type together based on their methods
// a1 := square(3)
// fmt.Println(a1.printShapeInfo())


//// structs (define here variables names & their datatypes)
type square struct {
	length float64
}
type circle struct {
	radius float64
}
// type rectangle struct {
// 	length float64
// 	breadth float64
// }

//// square methods/functions  - use struct variables and return output from function   
func (s square) area() float64 {
	return s.length * s.length      
}
func (s square) circumf() float64 {
	return s.length * 4
}

//// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// //rectangle methods
// func (r rectangle) area() float64 {
// 	return r.length * r.breadth
// }
// func (r rectangle) circumf() float64 {
// 	return 2 * (r.length + r.breadth)
// }

//func print
func printShapeInfo(s shape) {
	fmt.Printf("area of %T is : %0.2f \n",s,s.area())
	fmt.Printf("circumference of %T is : %0.2f \n",s,s.circumf)
}

type shape interface {
	area() float64
	circumf() float64
}