Variables are containers for storing data values.
# Go Variable Types

# In Go, there are different types of variables, for example:

    int- stores integers (whole numbers), such as 123 or -123
    float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
    string - stores text, such as "Hello World". String values are surrounded by double quotes
    bool- stores values with two states: true or false

More about different variable types, will be explained in the Go Data Types chapter.
Declaring (Creating) Variables

# In Go, there are two ways to declare a variable:
1. With the var keyword:

Use the var keyword, followed by variable name and type:
Syntax
var variablename type = value

# Note: You always have to specify either type or value (or both).
2. With the := sign:

Use the := sign, followed by the variable value:
Syntax
variablename := value

# Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

Note: It is not possible to declare a variable using :=, without assigning a value to it.
Variable Declaration With Initial Value

If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:


# Variable Declaration Without Initial Value

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type.

# Value Assignment After Declaration

It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.

#  Difference Between var and :=

There are some small differences between the var 
# var 	:=                                       var    :=:
Can be used inside and outside of functions	while, Can only be used inside functions
Variable declaration and value assignment while, can be done separately while, Variable declaration and value assignment cannot be done separately (must be done in the same line)


In Go, it is possible to declare multiple variables on the same line.

If you use the type keyword, it is only possible to declare one type of variable per line.

If the type keyword is not specified, you can declare different types of variables on the same line.

Go Variable Naming Rules

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

# Go variable naming rules:

    A variable name must start with a letter or an underscore character (_).
    A variable name cannot start with a digit.
    A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ ).
    Variable names are case-sensitive (age, Age and AGE are three different variables).
    There is no limit on the length of the variable name.
    A variable name cannot contain spaces.
    The variable name cannot be any Go keywords.
