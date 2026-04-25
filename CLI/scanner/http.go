package scanner

import (
    "net/http"
)

type Result struct {
    StatusCode     int
    Server         string
    MissingHeaders []string
}

func ScanURL(target string) Result {
    resp, err := http.Get(target)
    if err != nil {
        return Result{}
    }
    defer resp.Body.Close()

    missing := checkHeaders(resp.Header)

    return Result{
        StatusCode:     resp.StatusCode,
        Server:         resp.Header.Get("Server"),
        MissingHeaders: missing,
    }
}
