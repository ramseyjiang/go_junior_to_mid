Memory in a program is divided into two areas, one for the stack and one for the heap.
The stack area has a specific structure and addressing method, and is very fast to address with little overhead.
The heap, on the other hand, is an area of memory that has no specific structure and no fixed size and can be adjusted as needed.

Global variables, local variables with a large memory footprint,
and local variables that cannot be reclaimed immediately after a function call are all stored inside the heap.
Variables are allocated and reclaimed on the heap with much more overhead than on the stack.

When possible, the Go compilers will allocate variables that are local to a function in that function’s stack frame.
However, if the compiler cannot prove that the variable is not referenced after the function returns,
then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors.
Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.

In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap.
However, a basic escape analysis recognizes some cases when such variables will not live past the return
from the function and can reside on the stack.

Dynamic size variables are in heap, static(not always change size) variables are in stack.

go tool compile -m example.go
The above command is used to check Golang’s memory escape behavior.

When using reflection in Golang, it involves introspection and manipulation of values at runtime, rather than at compile time.
The Go runtime must perform additional work to determine the type and structure of the reflected value, 
which can add overhead and slow down the program.
