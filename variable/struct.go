package main

import "fmt"

// Here define a type named Books. When use books in this file, it can use Books type directly.
// Each defined type uses a struct as its underlying type. Define a type, it is defined at the top is better.
// In struct, it does not have any restrict with data type. In array and slice, both of them have only one type data.

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	// Here Books is the type of Book1
	var Book1 Books
	var Book2 Books

	/* book 1 desc */
	Book1.title = "Go lang"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go lang"
	Book1.bookId = 6495407

	/* book 2 desc */
	Book2.title = "Python"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python"
	Book2.bookId = 6495700

	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookId : %d\n", Book1.bookId)

	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookId : %d\n", Book2.bookId)

	fmt.Printf("Book 2 all info is %v", Book2)
}
