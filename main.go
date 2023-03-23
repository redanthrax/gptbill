package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const (
    interval int = 30
    conBaseUrl string = "https://api.congress.gov/v3"
)

func main() {
    conKey := os.Getenv("CONKEY")
    if len(conKey) == 0 {
        log.Fatalln("Congress API Key \"CONKEY\" not set.")
    }

    lastBill := Bill{}

    ctx := context.Background()
    timeout := 30 * time.Second
    reqContext, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    billsUrl := fmt.Sprintf("%s/bill?limit=1", conBaseUrl)
    m, err := Get[CongressBills](
        reqContext, 
        billsUrl, 
        conKey,
    )

    if err != nil {
        log.Fatal(err)
    }

    firstBill := m.Bills[0]
    if firstBill.Number != lastBill.Number {
        //get bill details
        billUrl := fmt.Sprintf(
            "%s/bill/%d/%s/%s",
            conBaseUrl,
            firstBill.Congress,
            firstBill.Type,
            firstBill.Number,
        )

        m, err := Get[BillRequest](
            reqContext,
            billUrl,
            conKey,
        )

        if err != nil {
            log.Fatalln(err)
        }

        lastBill = m.Bill
        textUrl := fmt.Sprintf(
            "%s/bill/%d/%s/%s/text",
            conBaseUrl,
            firstBill.Congress,
            firstBill.Type,
            firstBill.Number,
        )

        var t TextRequest
        t, err = Get[TextRequest](
            reqContext,
            textUrl,
            conKey,
        )

        if err != nil {
            log.Fatalln(err)
        }

        log.Printf("%#v\n", t)
    }
}
