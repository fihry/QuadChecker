package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// hadi bach tkon mteakad bli argument raha mn l pip, o ila kano chi arguments return
	// ila makanch arguments mn l pip return
	// ay haja makhdamalnach return

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return
	}
	var arguments []string
	for scanner.Scan() {
		line := scanner.Text()
		arguments = append(arguments, line)
	}

	if len(arguments) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	x := len(arguments[0])
	y := len(arguments)

	if x == 0 || y == 0 {
		return
	}

	var sl []string
	for a := 'A'; a <= 'E'; a++ {
		cmd := exec.Command("./quad"+string([]rune{a}), strconv.Itoa(x), strconv.Itoa(y))
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue
		}
		outputStr := string(output)
		sl = append(sl, outputStr)
	}

	var match []bool
	joinedArgs := strings.Join(arguments, "\n")
	// fmt.Println(joinedArgs)
	for _, shape := range sl {
		// fmt.Println(joinedArgs)
		// fmt.Println(shape)

		if strings.TrimSpace(shape) == strings.TrimSpace(joinedArgs) {
			match = append(match, true)
		} else {
			match = append(match, false)
		}
	}

	// for _, word := range match{
	// 	if word {
	// 		fmt.Print("true")
	// 	}else{
	// 		fmt.Println("false")
	// 	}
	// }

	X := strconv.Itoa(x)
	Y := strconv.Itoa(y)

	output := ""
	for i, isMatch := range match {
		if isMatch {
			if output != "" {
				output += " || "
			}
			output += fmt.Sprintf("[quad%c] [%s] [%s]", 'A'+i, X, Y)
		}
	}

	if output == "" {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(output)
	}
}
