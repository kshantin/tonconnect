package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cameo-engineering/tonconnect"
)

var pathfile string = "./examples/exist-connection/config.json"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	f, err := os.ReadFile(pathfile)
	if err != nil {
		log.Fatalf("Can not read file %v", err)
	}
	log.Print("Log")

	s := tonconnect.Session{}
	json.Unmarshal([]byte(f), &s)
	log.Println(s)

	msg, err := tonconnect.NewMessage("0QBZ_35Wy144n2GBM93YpcV4KOKcIjDJk8DdX4kyXEEHcbLZ", "100000000")
	if err != nil {
		log.Fatal(err)
	}
	tx, err := tonconnect.NewTransaction(
		tonconnect.WithTimeout(10*time.Minute),
		tonconnect.WithTestnet(),
		tonconnect.WithMessage(*msg),
	)
	if err != nil {
		log.Fatal(err)
	}
	boc, err := s.SendTransaction(ctx, *tx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Bag of Cells: %x", boc)
	}
	if err := s.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
