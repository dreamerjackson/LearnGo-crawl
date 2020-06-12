package scheduler

import "github.com/dreamerjackson/fast-crawl/engine"

type QueueScheduler struct {
	requestChan chan *engine.Request
	workerChan chan chan *engine.Request
}

func (s *QueueScheduler) WorkChan() (chan *engine.Request) {
	return make(chan *engine.Request)
}

func (s *QueueScheduler) Submit(r *engine.Request) {
	s.requestChan <-r
}

func (s *QueueScheduler) GetChanLen() int{
	return len(s.requestChan)
}


func (s * QueueScheduler) WorkReady(w chan *engine.Request){
	s.workerChan<-w
}

func (s * QueueScheduler)  Run(){
	s.workerChan = make(chan chan *engine.Request)
	s.requestChan = make(chan *engine.Request)
	go func() {
		var requestQ []*engine.Request
		var workQ [] chan *engine.Request
		for{
			var activeRequest *engine.Request
			var activework chan *engine.Request

			if len(requestQ)>0 && len(workQ)>0{
				activeRequest = requestQ[0]
				activework = workQ[0]
			}

			select {
				case r:=<-s.requestChan:
					requestQ = append(requestQ,r)
				case w:=<-s.workerChan:
					workQ = append(workQ,w)
				case activework<-activeRequest:
					workQ =workQ[1:]
					requestQ = requestQ[1:]
			}
		}
	}()
}