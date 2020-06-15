package scheduler

import "simqo.com/mygospace/learngo/example/crawl/engine"



type SimpleScheduler struct {
	workerchan chan engine.Requset
}

func (s *SimpleScheduler) Run() {
	s.workerchan = make(chan engine.Requset)
}

func (s *SimpleScheduler) WorkReady(chan engine.Requset) {
	return
}

func (s *SimpleScheduler) WorkChan() (chan engine.Requset) {
	return  s.workerchan
}

func (s *SimpleScheduler) Submit(r engine.Requset) {
	go func() {s.workerchan <- r}()
}
