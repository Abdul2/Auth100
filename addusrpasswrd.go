package Auth100
import (
	"os"
"errors"
)



//function adds a pair of usrnaem:hashed password to passwords file

func appendtopasswrdfile(username string, hashedpasswrd []byte, passwordsfile string) (bool, error) {


	file, err := os.OpenFile(passwordsfile, os.O_APPEND|os.O_WRONLY,0600)
	if err != nil {
		//panic(err)
		err = errors.New("couldnt open fiel"+passwordsfile + "... make sure that the file exists")
		return false, err
	}
	defer file.Close()

	if _, err = file.WriteString("\n"+username+":"); err != nil {
		//panic(err)
		err = errors.New("couldnt open fiel"+passwordsfile)
		return false, err
	}

	if _, err = file.Write(hashedpasswrd); err != nil {
		//panic(err)
		err = errors.New("couldnt open fiel"+passwordsfile)
		return false, err
	}


	return true, nil
}

