package main

import "github.com/ThompsonJonM/fantasy-emailer/src/email"

func main() {
	email.SendEmail("jonathan.thompson@pendo.io", "kohocdwztrxxwfau", "../../csvFiles/player.csv", "Pendo Fantasy Digest", "../../templates/20210921.gohtml")
}
