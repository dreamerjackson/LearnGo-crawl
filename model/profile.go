package model

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
	return  p.Name +" " + p.Constellation
}