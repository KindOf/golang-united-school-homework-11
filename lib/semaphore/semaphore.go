package semaphore

type Semaphore interface {
	Push()
	Pull()
}

type semaphore struct {
	semC chan int64
}

func New(pool int64) Semaphore {
	return &semaphore{
		semC: make(chan int64, pool),
	}
}

func (s *semaphore) Push() {
	s.semC <- int64(1)
}

func (s *semaphore) Pull() {
	<-s.semC
}
