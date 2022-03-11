package entity

import "fmt"

type People struct {
	Id      string
	Name    string
	Age     int
	Company string
	Address string
}

func (p People) SayHello() {
	fmt.Println("Hello World")
}

func (p *People) UpdateAddress(in string) {
	p.Address = in
}

func (p *People) UpdateCompany(in string) {
	p.Company = in
}

func (p People) ToString() string {
	result := fmt.Sprintf("Id: %v, Name: %v, Age: %v, Company: %v, Address: %v", p.Id, p.Name, p.Age, p.Company, p.Address)
	return result
}
