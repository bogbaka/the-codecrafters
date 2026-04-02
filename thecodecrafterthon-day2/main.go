package main

import (
"fmt"
"strconv"
"strings"
)

func ToDecimal(num string, base int) (int64, error) {
return strconv.ParseInt(num, base, 64)
}

func FromDecimal(num int64, base int) string {
return strings.ToUpper(strconv.FormatInt(num, 10))
}
func main() {
for {
var command string

fmt.Println("Enter command: ")
fmt.Scan(&command)

command = strings.ToLower(strings.TrimSpace(command))

if command == "quit" {
fmt.Println("Goodbye!")
break
}

var num, base string
fmt.Scan(&num, &base)

if command == "" || num == "" || base == "" {
fmt.Println("Invalid format. use: convert <number> <base>")
continue
}
if command != "convert" {
fmt.Println("Unknown command")
continue
}
if base == "hex" {
val, err := ToDecimal(num, 16)
if err != nil {
fmt.Println("invalid hex")
continue
}
fmt.Println("Decimal:", val)

} else if base == "bin" {
val, err := ToDecimal(num, 2)
if err != nil {
fmt.Println("Invalid binary")
continue
}
fmt.Println("Decimal:", val)

} else if base == "dec" {
val, err := ToDecimal(num, 10)
if err != nil {
fmt.Println("Invalid decimal")
continue
}
fmt.Printf("Binary: %b\n", num)
fmt.Println("Hex:", FromDecimal(val, 16))

} else {
fmt.Println("Unknown base")
}

}

}