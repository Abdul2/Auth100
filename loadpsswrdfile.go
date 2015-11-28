package Auth100

import (
	"strings"
	"bufio"
	"os"
)


//internal construct
type hashedpasswords struct {

	username string
	hash     string
}


func (h *hashedpasswords) addhashedpassword(password string, hash string) {

	h.username = password
	h.hash = hash

}

func (h hashedpasswords) gethashedpasswords() (string, string) {

	return h.username,h.hash

}


//slice to hold pairs of passwords and their hash

var infilehashedpasswords = make([]hashedpasswords,0)
var passwordandhashpair hashedpasswords



//this function reads passwords and hash pair
func Initiatecredentialscheckingop(path string) []hashedpasswords{

	//todo - handle errors
	inFile, _ := os.Open(path)

	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)


	for scanner.Scan() {

		//fmt.Println(scanner.Text())

		tokens := strings.Split(scanner.Text(), ":")

		//fmt.Println(tokens[0])
		//fmt.Println(tokens[len(tokens)-1])

		passwordandhashpair.addhashedpassword(tokens[0],tokens[len(tokens)-1])

		infilehashedpasswords = append(infilehashedpasswords,passwordandhashpair)
	}

	return infilehashedpasswords
}


