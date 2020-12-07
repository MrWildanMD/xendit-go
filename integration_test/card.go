package main

import (
	"fmt"
	"github.com/xendit/xendit-go/card"
	"log"
)

func cardTest() {
	_, err := card.GetCharge(&card.GetChargeParams{
		ChargeID: "5e046a736113354249aab8bd",
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Card integration tests done!")
}
