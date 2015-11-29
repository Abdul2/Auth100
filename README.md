
# Auth100 


This is a simple Go programme that implements a file based authentications.I am not sure I have looked hard enough but I couldn't find much on file based
authentication and since i needed an  authentication module, i decided to write pkg. 

this work is under development and i intend to improve on it but for now is good enough for testing and development work


## what does it do?

Auth100 does four things

1. creates hashed passwords
2. stores username and and associated hashed password in a client specified password file
3. validates username and password 


## How to use ?


Below is a sample client code. to use Auth100
 
 
1. import 
2. create a password file and point to it
3. create some dummy data to test 
4. validate user and password
5. register  new users 
 
 

## Methods 

1. Encryptpassword(username string, password string, passwordfile string) 

  this method creates a new line in the password file and writes the
  username  and password to it in the format of ```username:$2a$10$S9T.BTkpi11iAszdKpqEUeToC4eiGoe.spQWH5TmOQlhpT6Kg.wdq```
  .Colon (:) separets the user name and the hashed password.  
  
2. Validateusrandpsswr(providedusername string, providedpassword string) (bool,bool,Authenticationmessage)

   given a username and password, Validateusrandpsswr checks if username exists and if the does validates 
   the supplied password. it returns two booleans and a message confirming the outcome of the validation 
   processes.  the first boolean confirms or otherwise the username. the second boolean confirms or otherwise
   the validity of the provided password.  note that password validation is always false if username search returns
   a false value.
   

## Improvements (time permitting)

1. more testing 
2. error handling 
3. improving the code 

# known issues
 
1. slow file IO (on my machine)

 
## feedback:
 
  abdulrashid2@gmail.com
  
  

## sample client code  

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

