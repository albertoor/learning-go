package main

// import "fmt"

import (
  "math"
  "strings"
)

func main()  {
	// fmt.Println("Hello World")
  math.Floor(2.75)
  strings.Title("Head first go")

  // Strings
  fmt.Println("Hello, GO!")
  fmt.Println("Hello, \tGO!")
  fmt.Println("Quotes: \"\"")
  fmt.Println("Backslash: \\")

  // Booleans
  var isLoggin = true
  var isLogout = false
  fmt.Println(isLoggin)
  fmt.Println(isLogout)

  // Numbers
  var weight = 134.14
  var steps = 340

  // Math operations and comparisions
  fmt.Println("Add: ", 10 +  2)
  fmt.Println("Sub: ", 10 -  2)
  fmt.Println("Mult", 10 *  2)
  fmt.Println("Div: ", 10 /  2)
  fmt.Println(4 < 6)
  fmt.Println(4 > 6)
  fmt.Println(2 + 2 == 5)
  fmt.Println(2 + 2 != 5)
  fmt.Println(4 <= 6)
  fmt.Println(4 >= 6)

  // how to know type of variable do you using
  
  // int (holds whole numbers)
  fmt.Println(reflect.TypeOf(42))
  // float64 (Holds numbers with a fractional part) 
  fmt.Println(reflect.TypeOf(3.1415))
  // bool (Can only be true or false)
  fmt.Println(reflect.TypeOf(true))
  // string (A series of data usually represents text characters)
  fmt.Println(reflect.TypeOf("Hello, Go!"))

  // Declaring variables
  var quantity int
  var length, width float64
  var customerName string 
  quantity = 2
  customerName = "John Zapk"
  length, width = 1.2, 2.5

  fmt.Println(customerName)
  fmt.Println("has ordered", quantity, "sheets")
  fmt.Println("each with an area of")
  fmt.Println(length * width, "square meters")

  // Zero values
  var myInt int // is 0
  // for bool variables the value is false
  var myBool bool

  // Short variable declarations
  instrument := "Piano"
  fmt.Println(instrument)

  // Conversions
  var note int = 2
  fmt.Println(reflect.TypeOf(note))
  fmt.Println(reflect.TypeOf(float64(note)))
}
