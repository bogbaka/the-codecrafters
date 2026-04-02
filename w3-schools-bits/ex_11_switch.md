## The switch Statement

Use the switch statement to select one of many code blocks to be executed.

The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.


## THE SINGLE CASE SWITCH SYNTAX
The expression is evaluated once
The value of the switch expression is compared with the values of each case
If there is a match, the associated block of code is executed
The default keyword is optional. It specifies some code to run if there is no case match

## The default Keyword

The default keyword specifies some code to run if there is no case match.

All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error.

The Multi-case switch Statement

It is possible to have multiple values for each case in the switch statement.