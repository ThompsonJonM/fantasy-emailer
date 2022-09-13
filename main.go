package main

import "github.com/ThompsonJonM/fantasy-emailer/m/v2/src/email"

func main() {
	email.SendEmail("jonathan.thompson@pendo.io", "kohocdwztrxxwfau", "./csvFiles/players.csv", "Pendo Fantasy Digest")
}
