package utils

import (
	"fmt"
    "gorm.io/gorm"
)

type RR struct {
    Index int
    Addrs []string
}

func (rr *RR) Add(url string) {
    if url == "" {
        fmt.Println("Add error: empty url")
        return
    }
    rr.Addrs = append(rr.Addrs, url)
}

func (rr *RR) RoundRobin() string {
    if rr.Index >= len(rr.Addrs) {
        rr.Index = 0
    }
    res := rr.Addrs[rr.Index]
    rr.Index++
    return res
}

type DBRR struct {
    Index int
    Addrs []*gorm.DB
}

func (rr *DBRR) Add(db *gorm.DB) {
    if &db == nil {
        fmt.Println("Add error: empty url")
        return
    }
    rr.Addrs = append(rr.Addrs, db)
}

func (rr *DBRR) RoundRobin() *gorm.DB {
    if rr.Index >= len(rr.Addrs) {
        rr.Index = 0
    }
    res := rr.Addrs[rr.Index]
    rr.Index++
    return res
}