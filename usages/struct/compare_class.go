package main

/**
In other languages, it always has a class, but in go, it doesn't have. But it has the similar thing.
In other language, for example PHP,
class a {
	private $variable;
	public $variable;
	public function a() {
		...
	}
	private function b() {
		...
	}
}
How does it work in Go, in this file, I will show them all.
*/

import "fmt"

// This is similar as define a class, the capital variable, it means public, others mean private.

type Employee struct {
	Name        string
	age         int
	Designation string
	Salary      int
}

// This way, it is similar as define a public function in a class.
// The public or private function in a class, it depends on the first char is a capital or not.

func (staff Employee) ShowDetails() {
	fmt.Println("User Name: ", staff.Name)
}

// If you create an Employee only by value, you should define all variables in the same order.
func createStructOnlyByValue() {
	var employee1 = Employee{"Ramsey", 30, "Developer", 40}
	// var employee2 = Employee{40, "Mamba", 30, "Developer"} //If not, it will have an error.
	fmt.Println(employee1)
}

// If you create an Employee by key-value pair, you don't need to create variables using the same order.
// If you miss the variable, it will have a default value as it defines in struct.
func createStructByKeyValue() {
	employee3 := Employee{Designation: "Developer", Name: "Mamba", age: 40}
	fmt.Println(employee3)
}

// If you create a struct use the following way, you should change it using "Passing an object as a reference." way
// to change details of your struct.
// In this approach is that we have to explicitly use the address-of operator to extract the address of the object
// and send it to the function updateEmployee().

func UpdateEmployeeWay1() {
	employee1 := Employee{"Ramsey", 30, "Developer", 40}
	fmt.Println("Before update name: ", employee1.Name)
	updateEmployee(&employee1) // Passing an object as a reference.
	fmt.Println("After update: ", employee1.Name)
}

func updateEmployee(employee *Employee) { // Passing an object as a reference.
	employee.Name = "Davy"
}

// If using the keyword "new" to create a struct, you can change the details directly.
/**
When we’re creating an object with the new keyword, the address of the object is returned.
So if we’re sending the object as a parameter to the function, the reference of the object is sent.
Any changes made to the input object parameter will be reflected in the original object.

The object reference is returned from the new keyword.
Therefore, while invoking the function, we don’t need to use & to send the reference of the object.
*/

func UpdateEmployeeWay2() {
	newEmployee := new(Employee)
	fmt.Println("Use the 'new' way to create a struct, create: ", newEmployee)
	newEmployee.Name = "May"
	newEmployee.age = 30
	fmt.Println("Before update: ", newEmployee, "name is ", newEmployee.Name)
	updateEmployee(newEmployee)
	fmt.Println("After update: ", newEmployee, "name is ", newEmployee.Name)
}

func invokeFuncWayAsClass() {
	newEmployee := new(Employee)
	newEmployee.Name = "Create an employee name"

	newEmployee.ShowDetails()
}

func main() {
	createStructOnlyByValue()
	createStructByKeyValue()
	UpdateEmployeeWay1()
	UpdateEmployeeWay2()
	invokeFuncWayAsClass()
}
