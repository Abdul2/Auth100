package Auth100



import  (

	"golang.org/x/crypto/bcrypt"

)

//function to hash password
//calls appendtopasswrdfile to write out generated hashed password to passwords file
func Encryptpassword(username string, password string, passwordfile string) {

	passwordbytes := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordbytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}


	appendtopasswrdfile(username,hashedPassword,passwordfile)

}


