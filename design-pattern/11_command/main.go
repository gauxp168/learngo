package main

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
 	mb *MotherBoard
}

func NewStarCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb:mb,
	}
}

func (s *StartCommand) Execute()  {
	s.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb:mb,
	}
}

func (r *RebootCommand) Execute()  {
	r.mb.Reboot()
}

type MotherBoard struct {

}

func (m *MotherBoard) Start()  {
	fmt.Println("system starting")
}

func (m *MotherBoard) Reboot()  {
	fmt.Println("system rebooting")
}

type Box struct {
	buttion1 Command
	buttion2 Command
}

func NewBox(buttion1,buttion2 Command) *Box {
	return &Box{
		buttion1:buttion1,
		buttion2:buttion2,
	}
}

func (b *Box) PressButtion1()  {
	b.buttion1.Execute()
}

func (b *Box) pressButtion2()  {
	b.buttion2.Execute()
}

func main() {
	mb := &MotherBoard{}
	StartCommand := NewStarCommand(mb)
	rebootCommand := NewRebootCommand(mb)
	box := NewBox(StartCommand, rebootCommand)
	box.PressButtion1()
	box.pressButtion2()
}
