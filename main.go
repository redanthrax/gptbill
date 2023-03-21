package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type LatestAction struct {
    ActionDate string `json:"actionDate"`
    Text string `json:"text"`
}

type Bill struct {
    Congress int `json:"congress"`
    LatestAction LatestAction `json:"latestAction"`
    Number string `json:"number"`
    OriginChamber string `json:"originChamber"`
    OriginChamberCode string `json:"originChamberCode"`
    Title string `json:"title"`
    Type string `json:"type"`
    UpdateDate string `json:"updateDate"`
    UpdateDateIncludingText string `json:"updateDateIncludingText"`
    Url string `json:"url"`
}

type Pagination struct {
    Count int `json:"count"`
    Next string `json:"next"`
}

type Request struct {
    ContentType string `json:"contentType"`
    Format string `json:"format"`
}

type CongressBills struct {
    Bills []Bill `json:"bills"`
    Pagination Pagination `json:"pagination"`
    Request Request `json:"request"`
}

func Get[T any](ctx context.Context, url string) (T, error) {
    var m T
    r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return m, err
    }

    res, err := http.DefaultClient.Do(r)
    if err != nil {
        return m, err
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return m, err
    }

    return parseJson[T](body)
}

func parseJson[T any](s []byte) (T, error) {
    var r T
    if err := json.Unmarshal(s, &r); err != nil {
        return r, err
    }

    return r, nil
}

func main() {
    ctx := context.Background()
    timeout := 30 * time.Second
    reqContext, _ := context.WithTimeout(ctx, timeout)
    m, err := Get[CongressBills](reqContext, "https://api.congress.gov/v3/bill?api_key=zirbBhoSdyfDIybiUwdAbxeuSAIn44DuM9VM6MTH")
    if err != nil {
        log.Fatal(err)
    }

    log.Println(len(m.Bills))
}
