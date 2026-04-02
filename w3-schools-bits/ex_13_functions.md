A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.
Create a Function

To create (often referred to as declare) a function, do the following:

    Use the func keyword.
    Specify a name for the function, followed by parentheses ().
    Finally, add code that defines what the function should do, inside curly braces {}.

Call a Function

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

A function can be called multiple times.

 ## Naming Rules for Go Functions

    A function name must start with a letter
    A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
    Function names are case-sensitive
    A function name cannot contain spaces
    If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

## NOTE: Give the function a name that reflects what the function does!

Parameters and Arguments

Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma.

## Note: When a parameter is passed to the function, it is called an argument.

Multiple Parameters

Inside the function, you can add as many parameters as you want.

## Note: When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.

## Return Values

If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function.

## Recursion Functions

Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.
