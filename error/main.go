package main

import (
	"fmt"
	"errors"
)

type errUserNameExist struct {
  UserName string
}

func (e errUserNameExist) Error() string{
	return fmt.Sprintf("username %s is exist", e.UserName)
}

func isErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	return ok
}

func checkUserNameExist(username string) (bool, error){
	if username == "foo" {
		return true, errUserNameExist{UserName: username}
	}

	if username == "bar" {
		return true ,errors.New("username bar is exist")
	}

	return false, nil
}

func main(){
  if _, err := checkUserNameExist("foo"); err != nil {
		if isErrUserNameExist(err){
			fmt.Println(err)
		}
	}
	if _, err := checkUserNameExist("bar"); err != nil {
		fmt.Println(err)
	}
}