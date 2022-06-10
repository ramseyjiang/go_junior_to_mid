A running program will store objects in two memory locations, the heap and the stack.

Garbage collection operates on the heap, not the stack. The stack is a LIFO data structure that stores function values.

Calling another function from within a function pushes a new frame onto the stack, 
which will contain the values of that function and so on. When the called function returns, its stack frame is popped off the stack.

In contrast, the heap contains values that are referenced outside a function.
For example, statically defined constants at the start of a program, or more complex objects, like Go structs.
The heap is a graph where objects are represented as nodes which are referred to in code or by other objects in the heap. 
As a program runs, the heap will continue to grow as objects are added unless the heap is cleaned up.

Go prefers to allocate memory on the stack, so most memory allocations will end up there. 
This means that Go has a stack per goroutine and when possible Go will allocate variables to this stack. 
The Go compiler attempts to prove that a variable is not needed outside the function by performing escape analysis to see 
if an object “escapes” the function. 

If the compiler can determine a variable lifetime, it will be allocated to a stack. 
However, if the variable’s lifetime is unclear it will be allocated on the heap.
Generally if a Go program has a pointer to an object then that object is stored on the heap.

GOGC controls the aggressiveness of the garbage collector.

Go compiler tries to determine a variable’s lifetime and size. 
If it succeeds, then it will be allocated to a stack frame. 
Stack frames are individual memory spaces for each function call, 
this means that Go has a stack per function, and when possible, Go will allocate variables to this stack.

Determining the variable lifetime is done by escape analysis, 
it chases across functions and packages to check if the memory should be allocated to a heap or stack.

Go GC means Go Garbage collectors. GC has two key parts, a mutator, and a collector.
The collector executes garbage collection logic and finds objects that should have their memory freed.
The collector is the section or functionality where collecting and detecting of the garbage take place according to algorithm.

The mutator executes application code and allocates new objects to the heap. 
It also updates existing objects on the heap as the program runs, 
which includes making some objects unreachable when they’re no longer needed.
The mutator is the initial phase before any garbage collecting where the heap allocation happens.

Go’s garbage collector is a non-generational concurrent, tri-color mark and sweep garbage collector.

1.non-generational concurrent
A generational garbage collector focuses on recently allocated objects.
Compiler optimisations allow the Go compiler to allocate objects with a known lifetime to the stack.
This means fewer objects will be on the heap, so fewer objects will be garbage collected.
This means that a generational garbage collector is not necessary in Go.
So, Go uses a non-generational garbage collector.

Concurrent means that the collector runs at the same time as mutator threads.

That's why Go uses a non-generational concurrent garbage collector.

2. Mark and sweep is the type of garbage collector
The algorithm has two phases, marking and sweeping.
In the mark phase, the garbage will be detected in the heap and marks objects that are no longer needed.
In the sweep phase, the garbage will be removed.
Mark and sweep is an indirect algorithm, as it marks live objects, and removes everything else.
They are not direct, it means that the mark does not mark the garbage, it will mark the non-garbage. 
The sweep will remove everything, but not non-garbage.

Implementation steps :
Stop the world (STW) is the crucial phase to check memory status, 
it actually means stopping the running goroutines and turn write barrier on so that data on heap can be maintained.
Once all goroutines has writing barriers on, the world will be started, and the workers will perform the garbage collection.

3.  tri-color algorithm
Marking is implemented by using a tri-color algorithm.
The algorithm divides objects in the heap into three sets according to the color assigned, black, white, and gray.

When marking begins, all objects are white except for the root objects which are grey. 
Roots are an object that all other heap objects come from, and are instantiated as part of running the program.
The garbage collector begins marking by scanning stacks, globals and heap pointers to understand what is in use.
When scanning a stack, the worker stops the goroutine and marks all found objects grey by traversing downwards from the roots. 
It resumes the goroutine after the scan finishes.

The grey objects are then enqueued to be turned black, which indicates that they’re still in use.
Once all grey objects have been turned black, 
the collector will stop the world again and clean up all the white nodes that are no longer needed.
The program can now continue running until it needs to clean up more memory again.

White objects are garbage. 
Grey ones need to be checked and figure out if they are still in use and have reference to the white ones. 
If they have, they will turn white. 
If not, then we have our black set which is guaranteed to not have any reference to the white set, and they will stay safe.

The collector finds grey objects by traversing downwards from the roots objects,once that happened, 
the goroutine will be resumed and then new greyed object will be in a queue to turn black, 
then the world will stop again and the collector cleans up all the white nodes that are no longer needed.