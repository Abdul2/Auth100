package Auth100


import (
	"golang.org/x/crypto/bcrypt"
)

//to hold failure / success messages
type Authenticationmessage struct {

	Message string

}



func Validateusrandpsswr(providedusername string, providedpassword string) (bool,bool,Authenticationmessage) {

	//assume user and password do not exist
	checkuser := false
	checkpassword := false

	//set message
	m := Authenticationmessage{""}

	//walk through the objects in the slice
	for _, userhashedpasswrd1 := range infilehashedpasswords {


		//does user exits ?
		if providedusername == userhashedpasswrd1.username {

			m.Message = "user exists"
			checkuser = true

			//user exists so lets check the password

			for _, userhashedpasswrd2 := range infilehashedpasswords{


				s,_ := userhashedpasswrd2.gethashedpasswords()

				// check associated hash only
				if providedusername == s {

					_, s1 := userhashedpasswrd2.gethashedpasswords()

					p1 := []byte(providedpassword)
					p2 := []byte(s1)


					if (bcrypt.CompareHashAndPassword(p2,p1)) == nil {

						checkpassword = true
						m.Message = "user exists and password is Ok"

					}

				}

			m.Message = "user exists but password is incorrect"
			}

		}


	}

	return  checkuser, checkpassword,m

}


