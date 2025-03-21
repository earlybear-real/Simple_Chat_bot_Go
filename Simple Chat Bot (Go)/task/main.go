package main

import "fmt"

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me of your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Printf("Your age is %d; that's a good time to start programming!\n", age)
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d !\n", i)
	}
}

func showOptions() {
	fmt.Println("1. Because this makes fun.")
	fmt.Println("2. To make it harder to understand.")
	fmt.Println("3. To make the code more readable and make it easier to make changes to single parts.")
	fmt.Println("4. Be other developers so can't copy my code.")
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	// write the question here followed by the options
	fmt.Println("Why should we decompose big function into smaller ones?")
	showOptions()

	// use an infinite for loop until the user enters the correct answer
	for {
		var answer int
		fmt.Scan(&answer)
		if answer == 3 { // replace ? with the correct answer number
			return
		}
		fmt.Println("Please, try again.")
	}
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Aid", "2023") // change it as you need
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}
