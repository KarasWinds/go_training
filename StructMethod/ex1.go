package main

import (
	"fmt"
	"time"
)

type email struct {
	from string
	to string
}

func (e *email) From(s string) {
  e.from = s
}

func (e *email) To(s string) {
  e.to = s
}

func (e *email) Send() {
	fmt.Printf("From: %s, To: %s\n", e.from, e.to)
}

func main() {
  for i := 0; i < 10; i++ {
		go func(i int) {
			e := &email{}
			e.From(fmt.Sprintf("example%02d@example.com", i))
			e.To(fmt.Sprintf("example%02d@example.com", i+1))
			e.Send()
		}(i)
	}

	time.Sleep(1 * time.Second)
}