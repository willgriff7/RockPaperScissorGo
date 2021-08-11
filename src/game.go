package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	
	inTheGame()

}

func askAndGetString(i int) string{
	if(i == 0){
		fmt.Println("Thank you for starting the game. Please type P for Paper, R for Rock" +
		" and S for Scissor. Type E to end the game. This is not case sensitive")
	} else if(i == 1){
		fmt.Println("Please could you enter P, R, S or E")
	}
	var val string
    fmt.Scanln(&val)
	return val
}


func retfixedResponse(s string) string{
	var hold string = s
	for true{
		var cap string = strings.ToUpper(hold)
		if cap == "E" || cap == "R" || cap == "S" || cap == "P" {
			return string(cap[0])
		}else {
			hold = askAndGetString(0)
		}
	}
	return "E"
}

func getRandomChar() string{
	var randomNum int = rand.Intn(3)
	if randomNum == 0{
		return "P";
	}else if randomNum == 1{
		return "S";
	}else{
		return "R";
	}

}

func checkValue(c string) int{
	var randChar string = getRandomChar();
	if c == "R" && randChar == "R" {
		fmt.Println("Tied game between Rock and Rock");
		return 0;
	}else if c == "S" && randChar == "S"{
		fmt.Println("Tied game between Scissor and Scissor");
		return 0;
	}else if c == "P" && randChar == "P"{
		fmt.Println("Tied game between Paper and Paper");
		return 0;
	}else if c == "R" && randChar == "S" {
		fmt.Println("You Win. You chose Rock and it chose Scissor");
		return 1;
	}else if c == "R" && randChar == "P" {
		fmt.Println("You Lost. You chose Rock and it chose Paper");
		return -1;
	}else if c == "S" && randChar == "P" {
		fmt.Println("You Win. You chose Scissor and it chose Paper");
		return 1;
	}else if c == "S" && randChar == "R" {
		fmt.Println("You Lost. You chose Scissor and it chose Rock");
		return -1;
	}else if c == "P" && randChar == "R" {
		fmt.Println("You Win. You chose Paper and it chose Rock");
		return 1;
	}else {
		fmt.Println("You Lost. You chose Paper and it chose Scissor");
		return -1;
	}
}

func inTheGame(){

	var win int= 0
	var lose int = 0

	for true  {
		var s string = askAndGetString(1);
		var c string = retfixedResponse(s);
		if(c == "E"){
			break;
		}
		b := checkValue(c);
		if b < 0{
			lose += 1
		}else if b > 0{
			win += 1
		}
		fmt.Printf("You have %d wins and %d loses. ", win, lose);
	}
}