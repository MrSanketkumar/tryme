package main

import (
	"fmt"
)

// "encoding/json"
// "fmt"

// type Car struct{
// 	name string
// 	age int
// 	owner []string
// }

// func main(){
// 	car:=Car{name: "BMW",age: 1,owner: []string{"shivam","pal"}}

// 	car1:=car
// 	fmt.Println(car)
// 	car1.owner[0]="Sanket"
// 	fmt.Println(car1,car)

// }

// func deep(src,dest interface{})error{
// 	bytes,err:=json.Marshal(src)
// 	if err!=nil{
// 		return err
// 	}
// 	return json.Unmarshal(bytes,dest)
// }

// type Engine struct {
//    Capacity int
//    Power    int
// }

// type Car struct {
//    Brand  string
//    Model  string
//    Owners []string
//    Engine Engine
// }

// func deepCopy(src, dst interface{}) error {
//      bytes, err := json.Marshal(src)
//      if err != nil {
//       return err
//      }
//      return json.Unmarshal(bytes, dst)
// }

// func main() {
//      car1 := Car{
//       Brand:  "Tesla",
//       Model:  "Model X",
//       Owners: []string{"Alice", "Bob"},
//       Engine: Engine{Capacity: 100, Power: 300},
//      }

//      var car2 Car
//      err := deepCopy(&car1, &car2)
//      fmt.Println(err)

//      car2.Owners[0] = "John"
//      car2.Engine.Capacity = 200

//      fmt.Println(car1)
// 	 fmt.Println(car2)
//      fmt.Printf("Car engine: %v\nCar2: %v\n", car1.Engine, car2.Engine)
// }

// type Animal interface{
// 	sound()string
// }

// type Dog struct{
// 	name string
// }

// type Cat struct{
// 	name string
// }

// func (c *Cat) sound()string{
// 	return "meooow !! by "+c.name
// }

// func (d *Dog)sound()string{
// 	return "weof by"+d.name
// }

// func pass(a Animal){
// 	a.sound()
// }

// func main(){
// 	var animal Animal
// 	animal =&Cat{name: " sweety"}
// 	fmt.Println(animal.sound())

// }

func main() {

	fmt.Print(plaindrome("abb"))

}


func plaindrome(n string)bool{
	i,j:=0,len(n)-1

	for i<j{
		if n[i]!=n[j]{
			return false
		}
		i++
		j--
	}
	return true
}

// func binarySearch(array []int, toFind int) int {
// 	var left, right, m int
// 	left = 0
// 	right = len(array)

// 	for left < right {
// 		m = (right + left) / 2
// 		if array[m] == toFind {
// 			return m
// 		} else if array[m] > toFind {
// 			right = m - 1
// 		} else {
// 			left = m
// 		}
// 	}

// 	return -1
// }
