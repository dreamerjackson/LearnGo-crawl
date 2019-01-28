package model

import "strconv"

type Profile struct {
	Name string
	Age int
	Marry string
	Constellation string
	Height int
	Weight int
	Salary string
}

func (p Profile) String() string{
	return  p.Name +" " + p.Marry + strconv.Itoa(p.Age) +"olds "+   strconv.Itoa(p.Height) + "cm " +  strconv.Itoa(p.Weight)+ "kg "
}