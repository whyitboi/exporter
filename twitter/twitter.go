package twitter

//Twitter struct does
type Twitter struct {
	URL       string
	Userame   string
	Followers int
}

//Feed does
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"coding is basically copying",
	}
}

//Fame does
func (t *Twitter) Fame() int {
	return t.Followers
}
