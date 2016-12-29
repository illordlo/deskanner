package main

import (
	"fmt"
	"sync"
)

type executionLoad struct {
	Payload string
	Used    bool
}

type processingStack struct {
	Loads []*executionLoad
	Mtx   sync.Mutex
}

func newProcessingStack(payloads []string) *processingStack {
	ps := &processingStack{
		Loads: make([]*executionLoad, len(payloads)),
	}
	for idx, loads := range payloads {
		ps.Loads[idx] = &executionLoad{
			Payload: loads,
			Used:    false,
		}
	}
	return ps
}

func (p *processingStack) GetNext() (string, error) {
	p.Mtx.Lock()
	defer p.Mtx.Unlock()

	for _, payload := range p.Loads {
		if payload.Used != true {
			load := payload.Payload
			payload.Used = true
			return load, nil
		}
	}
	return "", fmt.Errorf("no more available payloads")
}
