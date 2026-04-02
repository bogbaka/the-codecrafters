## ARRAYS

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
Declare an Array

In Go, there are two ways to declare an array:
1. With the var keyword:
Syntax
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred
2. With the := sign:
Syntax
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred

Note: The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).
Array Examples
Example

This example declares two arrays (arr1 and arr2) with defined lengths.


## Access Elements of an Array

You can access a specific array element by referring to the index number.

## In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

## Change Elements of an Array

You can also change the value of a specific array element by referring to the index number.

This example shows how to change the value of the third element in the prices array.
## Array Initialization

If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

## NOTE: The default value for int is 0, and the default value for string is "".

## The len() function is used to find the length of an array.

