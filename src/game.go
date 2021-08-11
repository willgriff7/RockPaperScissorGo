package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Playtime p := new Playtime()

}

func askAndGetString(int i) string{
	if(i == 0){
		fmt.println(PLEASE_ENTER_CORRECT)
	}
	else if(i == 1){
		fmt.printlnn(INTROLINE)
	}
	var val string
    fmt.Scanln(&val)
	return val
}


func retfixedResponse(string s) string{
	var hold string := s
	for true{
		var cap string = hold.toUpperCase();
		if cap == "E" || cap == "R" || cap == "S" || cap == "P" {
			return cap.charAt(0)
		}else {
			hold = askAndGetString(0)
		}
	}
}

func getRandomChar() string{
	var int randomNum := rand.Intn(3)
	if randomNum == 0{
		return "P";
	}else if randomNum == 1{
		return "S";
	}else{
		return "R";
	}

}

func checkValue(string c) int{
	var string randChar := getRandomChar();
	if c == "R" && randChar == "R") {
		System.out.println("Tied game between Rock and Rock");
		return 0;
	}else if c == "S" && randChar == "S"{
		System.out.println("Tied game between Scissor and Scissor");
		return 0;
	}else if c == "P" && randChar == "P"{
		System.out.println("Tied game between Paper and Paper");
		return 0;
	}else if c == "R" && randChar == "S" {
		System.out.println("You Win. You chose Rock and it chose Scissor");
		return 1;
	}else if c == "R" && randChar == "P" {
		System.out.println("You Lost. You chose Rock and it chose Paper");
		return -1;
	}else if c == "S" && randChar == "P" {
		System.out.println("You Win. You chose Scissor and it chose Paper");
		return 1;
	}else if c == "S" && randChar == "R" {
		System.out.println("You Lost. You chose Scissor and it chose Rock");
		return -1;
	}else if c == "P" && randChar == "R" {
		System.out.println("You Win. You chose Paper and it chose Rock");
		return 1;
	}else {
		System.out.println("You Lost. You chose Paper and it chose Scissor");
		return -1;
	}
}

func inTheGame(){
	var int win_tracker = 0;
	var int lose_tracker = 0;
	while true{
		String s := askAndGetString(1);
		char c := retfixedResponse(s);
		if(c == 'E'){
			break;
		}
		var int b := checkValue(c);
		if b < 0{
			lose_tracker++;
		}else if b > 0{
			win_tracker++;
		}
		System.out.printf("You have %d wins and %d loses. ", win_tracker, lose_tracker);
	}
}