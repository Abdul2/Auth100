
# <h1> Auth100 </h1>


<h2> This is a simple Go programme that implements a file bases authentications.

I am not sure I have looked hard enough but could not find much resource 

on simple authentication and since i needed simple authentication, ii decided 

to write a module.

this work is under development but good enough to use for development 

work. </h2>


<h1> what does it do </h1>

##<h2> Auth100 does four things

   1. creates hashed passwords
   2. stores username and and associated hashed password in a client specified password file
   3. validates username and password 

</h2>


## <h1> How to use </h1>


<h2> Below is a sample client code. to use Auth100
 
 
 1. import 
 2. create a password file and point to it
 3. create some dummy data to test 
 4. validate user and password
 5. register  new users 
 
 
 
 
</h2>



package main


 
import (


	//1. import
	"github.com/Abdul2/userauthentication/Auth100"

	"fmt"
)




//2. decide where do you want to keep your hashed passwords
var PASSWRDFILE = "/location/of/passwdfile" 



func main() {

// 3 read test data into password file
	for i := 0; i < 10; i++ {

		s := fmt.Sprintf("user%v",i)
		s2 := fmt.Sprintf("password%v",i)

		fmt.Println(s,s2)
		Auth100.Encryptpassword(s,s2,PASSWRDFILE)

	}


// 4 validate 	
	Auth100.Initiatecredentialscheckingop(PASSWRDFILE)
	var a,b bool
	var s Auth100.Authenticationmessage
	a,b,s = Auth100.Validateusrandpsswr("user0","password1")
	fmt.Println(a,b,s.Message)

	//register new users and passwords
	Auth100.Encryptpassword("simon","Go further",PASSWRDFILE)
}

