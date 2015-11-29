package Auth100

import (

	"fmt"
	"testing"

)

func TestAddusrpasswd(t *testing.T)  {


	f := []byte("jjjjkl")

	fmt.Println("testing")

	result, err := appendtopasswrdfile("hhkjk",f,"passwrd")

	if err != nil && result != true {

		t.Errorf(" FAILURE.... didnt create an entry ")
	}

}