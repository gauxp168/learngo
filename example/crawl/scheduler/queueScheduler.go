package scheduler

import "simqo.com/mygospace/learngo/example/crawl/engine"

type QueueScheduler struct {
	requestChan chan engine.Requset
	workChan chan chan engine.Requset
}

func (q *QueueScheduler) WorkChan() (chan engine.Requset) {
	return make(chan engine.Requset)
}

func (q *QueueScheduler) Submit(r engine.Requset) {
	q.requestChan <-r
}

func (q *QueueScheduler) WorkReady (w chan engine.Requset)  {
	q.workChan <- w
}

func (q *QueueScheduler) Run()  {
	q.workChan = make(chan  chan engine.Requset)
	q.requestChan = make(chan engine.Requset)
	go func() {
		var requsetQ []engine.Requset
		var workQ [] chan engine.Requset
		for  {
			var activeRequset  engine.Requset
			var activework chan engine.Requset
			if len(requsetQ) > 0 && len(workQ) > 0 {
				activeRequset = requsetQ[0]
				activework = workQ[0]
			}
			select {
			case r:= <-q.requestChan:
				requsetQ = append(requsetQ, r)
			case w := <-q.workChan:
				workQ = append(workQ, w)
			case activework <- activeRequset:
				workQ = workQ[1:]
				requsetQ = requsetQ[1:]
			}
		}
	}()
}



