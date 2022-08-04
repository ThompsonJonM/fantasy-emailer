package email

import (
	"bytes"
	"fmt"
	"github.com/ThompsonJonM/fantasy-emailer/m/v2/src/data"
	"github.com/ThompsonJonM/fantasy-emailer/m/v2/src/players"
	"html/template"
	"log"
	"net/smtp"
	"strings"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func getMessageString(fromEmail, To, Subject, emailBody string) []byte {
	return []byte("From: " + fromEmail + "\r\n" + "To: " + To + "\r\n" + "Subject: " + Subject + "\r\n" + "MIME-Version: 1.0\r\n" + "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" + emailBody + "\r\n")
}

func parseTemplate(data interface{}) (string, error) {
	files := []string{
		"./templates/base.gohtml",
		"./templates/tweet.gohtml",
		"./templates/style.gohtml",
		"./templates/image.gohtml",
		"./templates/note.gohtml",
		"./templates/title.gohtml",
		"./templates/intro.gohtml",
		"./templates/tweets.gohtml",
		"./templates/images.gohtml",
		"./templates/notes.gohtml",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalln("Could not parse template", err)
	}
	buf := new(bytes.Buffer)

	if err = t.ExecuteTemplate(buf, "base", data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func SendEmail(from, pass, file, subject string) {
	ps := players.ImportPlayers(file)

	var i int

	for _, v := range ps {
		to := []string{
			v.Email,
		}
		f := strings.Split(v.Name, " ")
		first := strings.Join(f[:1], "")

		d := data.TemplateData{
			Introduction: data.Intro{
				Header:    "Week One Is in the Books",
				Subheader: "Dameon Pierce, Kyle Philips, and more.",
				Name:      first,
				Content: []string{
					"Week 1 is done, and what a sloppy weekend of football it was. Multiple teams experienced a higher number of turnovers than usual, resulting in some low scoring affairs.",
					"The game I was most invested in (aside from my Bills) was the Chargers and Raiders match. I own Adams and Waller in my League of Record, so I was interested to see what their usage would be. I came away thinking both will be great fantasy assets while Dereck Carr is not very good at this thing called football.",
					"In this issue of the Digest we're going to take a look at Javonte Williams, Dameon Pierce, Dak Prescott, and whether it is worth it to take an early round Running Back anymore.",
					"Let's get into it.",
				},
			},
			Tweets: []data.Tweet{
				{
					Header:  "Javonte Williams Impresses",
					Content: []string{},
					Link: data.TwitterLink{
						Content: []string{
							"Javonte Williams saw a majority of snaps on early downs, goalline snaps, and 100% of snaps in two minute drills🚀🚀",
						},
						Author:    "Nathan Jahnke",
						TweetLink: "https://twitter.com/PFF_NateJahnke/status/1569524633145323521",
						Date:      "Sep 12, 2022",
					},
				},
				{
					Header:  "Dak Prescott Out Less than 4 Games?",
					Content: []string{},
					Link: data.TwitterLink{
						Content: []string{
							"Jerry Jones told \n@1053thefan\n the Cowboys will not put Dak Prescott on IR with the belief he can return in the next four games.",
						},
						Author:    "Calvin Watkins",
						TweetLink: "https://twitter.com/calvinwatkins/status/1569680719680512004",
						Date:      "Sep 13, 2022",
					},
				},
				{
					Header: "Good News for Najee Owners",
					Content: []string{
						"While Najee was not a bright spot for the Steelers on Sunday, owners will be encouraged by reports that he intends to play on Sunday against the Patriots.",
						"Early speculation was that Najee would be out for at least 2-4 weeks with an ankle sprain, leading fantasy owners to check the waiver wire for backup RB Jaylen Warren.",
						"Do not roster Jaylen Warren. He had 13 attempts for 14 yards, totalling a putrid 1.07 yards per carry, well below the league average.",
					},
					Link: data.TwitterLink{
						Content: []string{
							"Steelers RB Najee Harris said on \n@MadDogRadio\n that he \"will be playing this weekend\" against the Patriots. ",
							"He's dealing with a foot injury suffered in PIT's season opener.",
						},
						Author:    "Zack Cox",
						TweetLink: "https://twitter.com/ZackCoxNESN/status/1569696091897212929",
						Date:      "Sep 13, 2022",
					},
				},
			},
			Images: []data.Image{
				{
					Header: "Kyle Philips, Target Monster",
					Image:  true,
					Content: []string{
						"With the departure of AJ Brown, the Titans needed someone to throw to. Enter Kyle Philips, a fifth round rookie out of UCLA. While most believed that fellow WR Treylon Burks would be the clear alpha, it was Philips operating out of the slot who got the most targets (9) on the day.",
						"An undrafted fantasy option, Philips is VERY likely on your waiver wire. Stash him if you're in a PPR league.",
					},
					ImageLink: "https://upload.wikimedia.org/wikipedia/commons/thumb/b/bf/Kyle_Philips_2022.jpg/640px-Kyle_Philips_2022.jpg",
				},
				{
					Header: "Samuel Shines",
					Image:  false,
					Content: []string{
						"Curtis Samuel appears to finally be healthy. He showcased his talent on Sunday against the Jaguars by catching 8 balls on 11 targets for 55 yards and a score. Samuel was on the field for the majority of the game alongside rookie Jahan Dotson and veteran Terry McLaurin.",
						"Should this usage continue, Samuel profiles to be a decent FLEX in the coming weeks.",
					},
				},
			},
			Notes: []data.Note{
				{
					Header: "Pierce Takes a Backup Role",
					Content: []string{
						"Much like Prince Daemon in Episode 3 of House of the Dragon, Houston's Dameon Pierce had a quiet day against the Indianapolis Colts.",
						"Outsnapped by veteran Rex Burkhead, Pierce saw 12 touches against the stingy Colts defense. According to HC Lovie Smith (who is sporting an impressive beard), the intent was to get Pierce more touches, but he was game scripted out.",
						"I'm not sure how a running back is game scripted out of a 20-3 lead, but it does make a bit of sense that they would ease a rookie into the starting role. Consider that Jonathan Taylor was behind Marlon Mack his rookie year before Mack tore his achilles.",
						"So, what do you do with Pierce? You hold. He will provide value, but it likely will not come until week 4 or later.",
					},
				},
				{
					Header: "Is RB Early, QB Late Dead?",
					Content: []string{
						"For years, the prevailing fantasy strategy was to take stud Running Backs early, Wide Receivers in the middle rounds, and then a QB late. Back then, early round QB's were not very valuable, so the strategy made sense.",
						"However, with week 1 in the books we saw more first and second round Running Backs bust (or be average) compared to their Wide Receiver counterparts. We also saw early round QB's (with the exception of Jalen Hurts) dominate their matches.",
						"So does this mean that the early round RB strategy is dead? Maybe.",
						"Should the trend continue, more Wide Receivers will be drafted in the first and second rounds of next year's draft than Running Backs. We already saw a first round that was fairly split between WR and RB this year, could it end up becoming 75% WR next year?",
						"At the current moment it is too early to tell, but the writing could be on the wall for the early Running Back strategy.",
					},
				},
			},
		}

		s := smtpServer{host: "smtp.gmail.com", port: "587"}
		ms, err := parseTemplate(d)
		if err != nil {
			log.Fatalln("Could not parse template", err)
		}

		b := getMessageString(from, to[0], subject, ms)

		auth := smtp.PlainAuth("", from, pass, s.host)

		err = smtp.SendMail(s.Address(), auth, from, to, b)
		if err != nil {
			log.Fatalln("Could not send e-mail", err)
		}

		i++

		fmt.Printf("%d of %d emails sent.\n", i, len(ps))
	}
}
