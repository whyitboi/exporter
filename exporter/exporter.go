package exporter

import (
	"encoding/json"
	"io/ioutil"
	"week3"
)

type Message struct {
	Name string
	Body string
}

//Export writes to the specifid file
func Export(u week3.SocialMedia, filename string) error {

	/* 	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	   	if err != nil {
	   		return errors.New("An error occured opeming the file:" + err.Error())
	   	} */
	//for _, fd := range u.Feed() {
	//n, err := f.Write([]byte(fd + "\n"))
	file, _ := json.MarshalIndent(u.Feed(), "", " ")

	_ = ioutil.WriteFile(filename+".json", file, 0644)
	/* 	if file == nil {
		return errors.New("An error occured while writing to the file:" + file.Error())
	} */
	//fmt.Printf("wrote %d bytes\n", n)
	//}
	return nil
}
