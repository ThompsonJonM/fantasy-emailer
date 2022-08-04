package data

type Intro struct {
	Header    string
	Subheader string
	Name      string
	Content   []string
}

type Image struct {
	Header    string
	Image     bool
	ImageLink string
	Content   []string
}

type Note struct {
	Header  string
	Content []string
}

type Tweet struct {
	Header  string
	Content []string
	Link    TwitterLink
}

type TwitterLink struct {
	Content   []string
	Author    string
	TweetLink string
	Date      string
}

type TemplateData struct {
	Introduction Intro
	Images       []Image
	Tweets       []Tweet
	Notes        []Note
}
