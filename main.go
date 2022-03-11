package main

import (
	"fmt"
	"log"
	"math/rand"
	"test/cmd"
	"test/entity"
	"test/internal"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	s, err := cmd.NewStore(100)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer s.Free()

	_ = s.Bootstrap(10)
	if err := s.AddPeople(&entity.People{
		Id:      internal.UUID(),
		Name:    "Nguyen Van A",
		Age:     24,
		Company: "FPT Software",
		Address: "Ha Noi",
	}); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Print As Order")
	s.PrintAsOrder(internal.Asc)

}
