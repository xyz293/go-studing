package model

type User struct{
  ID  int
  Name string
  Age int
  Eamil string
}

func (User) Table ()string{
	return "user"
}