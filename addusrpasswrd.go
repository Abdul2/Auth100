package Auth100
import (
	"os"
	"fmt"
)



//function adds a pair of usrnaem:hashed password to passwords file

func appendtopasswrdfile(username string, hashedpasswrd []byte, passwordsfile string) {


	file, err := os.OpenFile(passwordsfile, os.O_APPEND|os.O_WRONLY,0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString("\n"+username+":"); err != nil {
		panic(err)
	}

	if _, err = file.Write(hashedpasswrd); err != nil {
		panic(err)
	}


	fmt.Printf("Appended into file\n")
}

