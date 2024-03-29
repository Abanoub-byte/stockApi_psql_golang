package main

import "fmt"

func main() {

	fmt.Println("Welcome to quiz.Bob !! \n Let us know your name!")
	var name string
	var age int
	result := 3

	fmt.Scan(&name)
	fmt.Printf("Hello %v! \n", name)

	fmt.Println("How old are you ?")

	fmt.Scan(&age)
	if age <= 10 {
		fmt.Println("sorry you cant play sir, you are too young!!")
	} else {
		var answer, answer2 string
		
		fmt.Println("Loading...")
		fmt.Println("-First question: \n who is better Messi or Ronaldo ?")

		fmt.Scan(&answer)
		if answer == "ronaldo" || answer == "Ronaldo" {
			fmt.Println("Wrong !")
			result--

		} else if answer == "messi" || answer == "Messi" {
			fmt.Println("correct answer!")
		} else {
			fmt.Println("Wrong !")
			result--
		}

		fmt.Println("-Second question: \n Who won the FIFA Men's World Cup in 2018, and in which country was the tournament held?")
		fmt.Scan(&answer)
		if answer == "france"|| answer == "France"{
			fmt.Println("correct answer!")
			
		} else {
			fmt.Println("wrong")
			result--
		}

		fmt.Println("-Third question: \n Who holds the record for the most goals scored in a single season of the English Premier League? ")
		fmt.Scan(&answer, &answer2)
		if (answer + answer2 == "MohamedSalah" || answer + answer2 == "mohamedsalah"){
			fmt.Println(answer + answer2)
			fmt.Println("Correct answer!")
			fmt.Println(result, "/3")
			return
			
		}else {
			fmt.Println("Wrong!")
			result--
			fmt.Println("You final result is ", result, "/3")
			return
		}
	}
	
}
