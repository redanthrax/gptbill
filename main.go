package main

import (
	"context"
	"log"
)


func main() {
    log.Println("Starting gobill...")
    congress, err := CongressApi()
    if err != nil {
        log.Fatalln(err)
    }
    
    bills := BillService{}
    ctx := context()
    bills.ListAll()
}
