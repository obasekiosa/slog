package server

import (
    "fmt"
    "sync"
)

type Log struct {
    mu sync.Mutex
    records []Record
    length uint64
}

func NewLog() *Log {
    return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
    c.mu.Lock()
    defer c.mu.Unlock()
    record.Offset = c.length
    c.records = append(c.records, record)
    c.length += 1
    return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
    c.mu.Lock()
    defer c.mu.Unlock()
    if offset >= c.length {
        return Record{}, ErrOffsetNotFound
    }
    return c.records[offset], nil
}

type Record struct {
    Value []byte `json:"value"`
    Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
