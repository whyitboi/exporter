package facebook

//Facebook does
type Facebook struct {
	URL     string
	Name    string
	Friends int
}

//Feed does
func (f *Facebook) Feed() []string {
	return []string{
		"FaceBook feeds",
		"Hey, here's my cool new selfie",
		"What is code",
	}
}

//Fame does it
func (f *Facebook) Fame() int {
	return f.Friends
}
