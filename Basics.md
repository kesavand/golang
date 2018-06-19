Golang:


Basic Types:
===========

Numbers
---------
  Integers 
   uint8, uint16, uint32, uint64,
   int8, int16, int32 and int64
  Float
   float32, float64
 
NOTE : In addition there two alias types: byte which is the same as uint8 and rune which is the same as
int32.  

Operations allowed on Numbers:

+ addition
- subtraction
* multiplication
/ division
% remainder

Strings
--------
String literals can be created using double quotes
"Hello World" or back ticks `Hello World`. 

Operations allowed on Strings:
length of a string: len("Hello World")
accessing an individual character in the string: "Hello World"[1], 
concatenating two strings together: "Hello " + "World". 


Booleans
---------
A boolean value (named after George Boole) is a special
1 bit integer type used to represent true and false
(or on and off). Three logical operators are used with
boolean values:

operations allowed with boolean values:
&& and
|| or
! not



Variables:
==========
A variable is a storage location, with a specific type and an associated name. 

example:
var x string = "hello" // Valid in inside an doutside function.
x := "Hello World" // valid only in the inside of the function.

Variable Scope
--------------
The variable exists within the nearest curly braces { } (a block) including any nested curly braces (blocks), but
not outside of them. 

Defining Multiple Variables
---------------------------
var (
 a = 5
 b = 10
 c = 15
)
Use the keyword var (or const) followed by parentheses
with each variable on its own line.




Constants
==========
Constants are basically variables whose values cannot be changed later.
const x string = "Hello World"




Control Structures
==================

For
----
for i := 1; i <= 10; i++ {
 fmt.Println(i)
 }
 
IF
---
if i % 2 == 0 {
 // even
} else {
 // odd
}

Switch
------

switch i {
case 0: fmt.Println("Zero")
case 1: fmt.Println("One")
case 2: fmt.Println("Two")
case 3: fmt.Println("Three")
case 4: fmt.Println("Four")
case 5: fmt.Println("Five")
default: fmt.Println("Unknown Number")
}


Arrays, Slices and Maps
=======================
Arrays:
-------
Declaration and Initialization:

1) A basic declaration of an array goes like this
[length]element_type
For example:
var intArray [5]int

2) Declare an array with some or all values to it at the time of declaration
The literal value for an array is composed of the array type definition followed by a set of comma-separated values, enclosed in curly brackets.

For example:

var intArray = [5]int{10,20,30}


You can provide values for specific elements as shown here:

var intArray = [5]int{0:10,2:30,4:50}


3) Declare an array inside function using the := shortcut

intArray := [5]int{10, 20, 30, 40, 50}


4) Declaring an array with ellipses ...
The length of an array may be omitted and replaced by ellipses during initialization. Capacity is determined based on the number of values initialized.
The following will assign type [5]int to variable intArray:

intArray := [...]int{10, 20, 30, 40, 50}

for loop to iterate over the Array

Create integer array named intArray, with a length of 5. Using a for loop to iterate over the array, display the contents of the array.

package main
 
import "fmt"
 
func main() {
    intArray := [5]int{10, 20, 30, 40, 50}
     
    for i := 0; i < len(intArray); i++ {
        fmt.Println(intArray[i])
    }
     
}


Slices
------

A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length
is allowed to change. Here's an example of a slice:

 var x []float64
 
If you want to create a slice you should use the built-in make function:
x := make([]float64, 5)
This creates a slice that is associated with an underlying
float64 array of length 5.

x := make([]float64, 5, 10)
The make function also allows a 3rd parameter:
x := make([]float64, 5, 10) 10 represents the capacity of the underlying array
which the slice points to:


Another way to create slices is to use the [low : high]
expression:
arr := []float64{1,2,3,4,5}
x := arr[0:5]

Slice Functions
---------------
Go includes two built-in functions to assist with slices:
append and copy. 

append
-------
func main() {
slice1 := []int{1,2,3}
slice2 := append(slice1, 4, 5)
fmt.Println(slice1, slice2)
}

copy
----
func main() {
slice1 := []int{1,2,3}
slice2 := make([]int, 2)
copy(slice2, slice1)
fmt.Println(slice1, slice2)
}


Maps
----
A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a
dictionary, maps are used to look up a value by its associated key. 
Here's an example of a map in Go:

var x map[string]int

Maps should be initialized before using 
Initialization of a map 
x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])

We can also delete items from a map using the built-in
delete function:
delete(x, 1)


Accessing an element of a map can return two values
instead of just one. The first value is the result of the
lookup, the second tells us whether or not the lookup

if name, ok := elements["Un"]; ok {
fmt.Println(name, ok)
}

Map of Maps
-----------

Maps are also often used to store general information

func main() {
elements := map[string]map[string]string{
"H": map[string]string{
"name":"Hydrogen",
"state":"gas",
},
"He": map[string]string{
"name":"Helium",
"state":"gas",
},
"Li": map[string]string{
"name":"Lithium",
"state":"solid",
},



Functions
==========
func main() {}

func average(xs []float64) float64 {
panic("Not Implemented")
}
The names of the parameters don't have to match in
the calling function.

Functions don't have access to anything in the calling
function.

 We can also name the return type:
 func f2() (r int) {
 r = 1
 return
}

Returning Multiple Values
--------------------------
Go is also capable of returning multiple values from a
function:
func f() (int, int) {
return 5, 6
}
func main() {
x, y := f()
}



Defer, Panic & Recover
======================


Defer
-----
Go has a special statement called defer which schedules a function call to be run after the function completes

package main
import "fmt"
func first() {
fmt.Println("1st")
}
func second() {
fmt.Println("2nd")
}
func main() {
defer second()
first()
}

f, _ := os.Open(filename)
defer f.Close()

This has 3 advantages: 
(1) it keeps our Close call near our Open call so its easier to understand,
2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before
both of them and 
(3) deferred functions are run even if a run-time panic occurs



Panic & Recover
---------------
panic function causes a run time error.

We can handle a run-time panic with the built-in recover function.
recover stops the panic and returns the value that was passed to the call to panic.
 We might be tempted to use it like this:
 
 
 package main
import "fmt"
func main() {
panic("PANIC")
str := recover()
fmt.Println(str)
}
But the call to recover will never happen in this case
because the call to panic immediately stops execution
of the function. Instead we have to pair it with defer:
package main
import "fmt"
func main() {
defer func() {
str := recover()
fmt.Println(str)
}()
panic("PANIC")
}
A panic generally indicates a programmer error (for
example attempting to access an index of an array
that's out of bounds, forgetting to initialize a map, etc.)
or an exceptional condition that there's no easy way to
recover from. (Hence the name “panic”)




Pointers
=========

In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value. 
* is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the value the
pointer points to.

Finally we use the & operator to find the address of a Pointers variable

NOTE:

There is only one way to pass parameters in Go - by value. Every time a variable is passed as parameter, a new copy of the variable is created and passed to called function or method. The copy is allocated at a different memory address.

In case a variable is passed by pointer, a new copy of pointer to the same memory address is created. 

new
----
Another way to get a pointer is to use the built-in new function:
func one(xPtr *int) {
*xPtr = 1
}
func main() {
xPtr := new(int)
one(xPtr)
fmt.Println(*xPtr) // x is 1
}


Pointers are rarely used with Go's built-in types, but as we will see in the next chapter, they are extremely
useful when paired with structs


Structs and Interfaces
======================

Structs
--------
type Circle struct {
x float64
y float64
r float64
}

type Circle struct {
x, y, r float64
}


Initialization of structure
----------------------------

var c Circle

c := new(Circle)
This allocates memory for all the fields, sets each of
them to their zero value and returns a pointer.
(*Circle) 


We can initialize the structure variable with initial values. this can be done in in two ways:

c := Circle{x: 0, y: 0, r: 5}
Or we can leave off the field names if we know the order
they were defined:
c := Circle{0, 0, 5}


We can access fields using the . operator:


Methods
-------

Method is a special type of function.In between the keyword func and the name of the
function we've added a “receiver”. The receiver is like a parameter – it has a name and a type – but by creating
the function in this way it allows us to call the function using the "." operator



Embedded Types
--------------

A struct's fields usually represent the has-a relationship.
For example a Circle has a radius. Suppose we
had a person struct:

type Person struct {
 Name string
}
func (p *Person) Talk() {
 fmt.Println("Hi, my name is", p.Name)
}
And we wanted to create a new Android struct. We
could do this:
type Android struct {
 Person Person
 Model string
}
This would work, but we would rather say an Android
is a Person, rather than an Android has a Person. Go
supports relationships like this by using an embedded
type. Also known as anonymous fields, embedded
types look like this:
type Android struct {
 Person
 Model string
}
We use the type (Person) and don't give it a name.
When defined this way the Person struct can be accessed
using the type name:
a := new(Android)
a.Person.Talk()
But we can also call any Person methods directly on
the Android:
a := new(Android)
a.Talk()
The is-a relationship works this way intuitively: People
can talk, an android is a person, therefore an android
can talk.


Interfaces
----------

Interfaces are named collections of method signatures.
