// CodeCrafters — Hackathon 002
// Squad: [Channels]
/*
Members: [
   All names:
1.  Okopi Ebo Joy
2.  Adikwu Regina
3.  Antony Agbo
4.  Daniel Akpa
5.  Jonathan Ahubi
6.  Okoh Daniel Oche
7.  Raymond Nicholas
8.  Blessing Ogbaka
9.  Moses Ochife
]
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
back:
	fmt.Println("════════════════════════════════════════════════")
	fmt.Println(" SENTINEL — COMMAND & CONTROL CONSOLE")
	fmt.Println("All systems nominal. Type 'help' to begin.")
	fmt.Println("════════════════════════════════════════════════")

	fmt.Print("C&C> ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	input := scanner.Text()

	cmd := strings.Fields(input)

	inputed := scanner.Text()

	







	

	switch inputed {
		
	case "base help" :
		fmt.Println("1. C&C> base dec 255\n✦ Binary : 11111111\n✦ Hex    : FF")
 		fmt.Println("2. C&C> base hex 1F\n✦ Decimal: 31")
  		fmt.Println("3. C&C> base bin 1010\n ✦ Decimal: 10")
		goto back

		case "cal help":
		fmt.Println("calc add <a> <b> - addition")
		fmt.Println("calc sub <a> <b> substraction")
		fmt.Println("calc mul <a> <b> multiplication")
		fmt.Println("calc div <a> <b> division")
		fmt.Println("calc mod <a> <b> remainder")
		fmt.Println("calc pow <a> <b> a to the power of b")
		fmt.Println("calc last → show the last result")
		fmt.Println("calc last  → show last 5 calc results")
		goto back

		case "quit":
			fmt.Println("THANK YOU AND GOODBYE")
			return
	}

	

	if cmd[0] != "calc <a> <b>" && cmd[1] != "add" && cmd[1] != "sub" && cmd[1] != "mul" && cmd[1] != "div" && cmd[1] != "mod" && cmd[1] != "pow" && cmd[1] != "hex" && cmd[1] != "bin" && cmd[1] != "dec"  && cmd[1] != "upper" && cmd[1] != "lower" && cmd[1] != "cap" && cmd[1] != "title" && cmd[1] != "snake" && cmd[1] != "reverse"{
		fmt.Println("Invalid command")
		goto back
	}

	switch cmd[1] {
		case "upper":
		new := cmd[2:]

		result := strings.Join(new, " ")

		fmt.Println(strings.ToUpper(result))
		goto back

    case "lower":
		new := cmd[2:]

		result := strings.Join(new, " ")

		fmt.Println(strings.ToLower(result))
		goto back

	case "cap":
		new2 := cmd[2:]

		result := strings.Join(new2, " ")

		fmt.Println(strings.Title(result))
		goto back
	case "reverse" :
		reverse := cmd[2:]
		var result []string
		for i := len(reverse)-1; i >= 0; i--{
			result = append( result, reverse[i])
		}
		fmt.Println(strings.Join(result, ""))
		goto back

	case "hex":
		if len(cmd) != 3 {
			fmt.Println("Invalid command")
			goto back
		} else {
		add1,err := strconv.ParseInt(cmd[2], 16, 64)
		if err != nil{
			fmt.Println("NOT A VALID HEX CHARACTER")
		} else {
			add2 :=  strconv.FormatInt(add1, 10)
		fmt.Println("result", add2)
		
		}
		}
		goto back
		
	case "bin":
		if len(cmd) != 3 {
			fmt.Println("Invalid command")
			goto back
		} else{
		add1,err := strconv.ParseInt(cmd[2], 2, 64)
				if err != nil{
			fmt.Println("NOT A VALID BIN CHARACTER")
		} else {
			add2 :=  strconv.FormatInt(add1, 10)
		    fmt.Println("result", add2)
		}
		}
		goto back
		
	case "dec":
		if len(cmd) != 3 {
			fmt.Println("Invalid command")
			goto back
		} else{
		add1,err := strconv.ParseInt(cmd[2], 10, 64)
			if err != nil{
			fmt.Println("NOT A VALID DEC CHARACTER")
		} else {
	
		add2 :=  strconv.FormatInt(add1, 2)
		add3 :=  strings.ToUpper(strconv.FormatInt(add1, 16)) + "\n" 
		fmt.Println("result:", add2)
		fmt.Println("result:",add3)
		
		}
	}
		
		goto back
		
	case "add":
		if len(cmd) != 4 {
			fmt.Println("Invalid command")
			goto back
		} else{
		add1, _ := strconv.ParseInt(cmd[2], 10, 64)
		add2, _ := strconv.ParseInt(cmd[3], 10, 64)

		fmt.Println("✦ Result: ", add1+add2)
		}
		goto back

	case "sub":
		if len(cmd) != 4 {
			fmt.Println("Invalid command")
			goto back
		} else{
		sub1, _ := strconv.ParseInt(cmd[2], 10, 64)
		sub2, _ := strconv.ParseInt(cmd[3], 10, 64)
		fmt.Println("✦ Result: ", sub1-sub2)
		}
		goto back

	case "mul":
		if len(cmd) != 4 {
			fmt.Println("Invalid command")
			goto back
		} else{
		mul1, _ := strconv.ParseInt(cmd[2], 10, 64)
		mul2, _ := strconv.ParseInt(cmd[3], 10, 64)
		fmt.Println("✦ Result: ", mul1*mul2)
		}
		goto back

	case "div":
		if len(cmd) != 4 {
			fmt.Println("Invalid command")
			goto back
		} else{
		div1, _ := strconv.ParseInt(cmd[2], 10, 64)
		div2, _ := strconv.ParseInt(cmd[3], 10, 64)
		if div2 == 0 {
			fmt.Println("can't be divided by ZERO")
			return
		}
		fmt.Println("✦ Result: ", float64(div1)/float64(div2))
	}
		goto back


	case "mod":
		if len(cmd) != 4 {
			fmt.Println("Invalid command")
			goto back
		} else{
		mod1, _ := strconv.ParseInt(cmd[2], 10, 64)
		mod2, _ := strconv.ParseInt(cmd[3], 10, 64)
		if mod2 == 0 {
			fmt.Println("error")
			return
		}
		fmt.Println("✦ Result: ", mod1%mod2)
	}
		goto back

	case "pow":
		if len(cmd) != 4 {
			fmt.Println("Invalid command")
			goto back
		} else{
		pow1, _ := strconv.ParseInt(cmd[2], 10, 64)
		pow2, _ := strconv.ParseInt(cmd[3], 10, 64)
		fmt.Println("✦ Result: ", math.Pow(float64(pow1), float64(pow2)))
		}
		goto back

	}

}
