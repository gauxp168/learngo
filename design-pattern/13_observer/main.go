package main

import "fmt"

type Subject struct {
	observer []Observer
	context string
}

func NewSubject() *Subject {
	return &Subject{
		observer:make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer)  {
	s.observer = append(s.observer, o)
}

func (s *Subject) notify()  {
	for _, o := range s.observer {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string)  {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name:name,
	}
}

func (r *Reader) Update(o *Subject)  {
	fmt.Printf("%s receive %s\n", r.name, o.context)
}

func main() {
	subject := NewSubject()
	read1 := NewReader("read1")
	read2 := NewReader("read2")
	read3 := NewReader("read3")
	subject.Attach(read1)
	subject.Attach(read2)
	subject.Attach(read3)

	subject.UpdateContext("observer mode")
}
