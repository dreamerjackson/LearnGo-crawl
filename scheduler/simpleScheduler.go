package scheduler

import "crawl/LearnGo-crawl/engine"



type SimpleScheduler struct{
	workerChan chan  engine.Request
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
	return
}

func (s *SimpleScheduler) WorkChan() (chan engine.Request) {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func(){s.workerChan<-r}()
}
