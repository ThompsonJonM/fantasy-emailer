package main

import "github.com/ThompsonJonM/fantasy-emailer/m/v2/src/email"

func main() {
	email.SendEmail("", "", "./csvFiles/player.csv", "Fantasy Digest")
}
