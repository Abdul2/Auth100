package main


import (


	//1. import
	"github.com/Abdul2/userauthentication/Auth100"

	"fmt"
)




//1. decide where do you want to keep your hashed passwords - one time operation
var PASSWRDFILE = "/home/aabdi/goprojects/src/github.com/Abdul2/userauthentication/htpasswd"



func main() {

//2. read test data into password file
	for i := 0; i < 10; i++ {

		s := fmt.Sprintf("user%v",i)
		s2 := fmt.Sprintf("password%v",i)

		fmt.Println(s,s2)
		Auth100.Encryptpassword(s,s2,PASSWRDFILE)

	}

	// 3. validate users
	var a,b bool
	var s Auth100.Authenticationmessage
	a,b,s = Auth100.Validateusrandpsswr("user0","password1")
	fmt.Println(a,b,s.Message)

	//register new users and passwords

	Auth100.Encryptpassword("simon","Go further",PASSWRDFILE)
}


func init(){

	// 2. load the password file into memory to check access
	Auth100.Initiatecredentialscheckingop(PASSWRDFILE)


}