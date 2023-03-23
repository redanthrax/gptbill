package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Get[T any](ctx context.Context, url string, key string) (T, error) {
    var m T
    r, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return m, err
    }

    r.Header.Add("X-Api-Key", key)
    res, err := http.DefaultClient.Do(r)
    if err != nil {
        return m, err
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return m, err
    }

    log.Println(string(body))

    return parseJson[T](body)
}

func parseJson[T any](s []byte) (T, error) {
    var r T
    if err := json.Unmarshal(s, &r); err != nil {
        return r, err
    }

    return r, nil
}
