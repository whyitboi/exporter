package linkedin

type Linkedin struct {
	URL         string
	Name        string
	Connections int
}

//Feed does
func (l *Linkedin) Feed() []string {
	return []string{
		"Linkedin feeds",
		"Hey, Just started mynew position at hotels.ng",
	}
}

//Fame does
func (l *Linkedin) Fame() int {
	return l.Connections
}
