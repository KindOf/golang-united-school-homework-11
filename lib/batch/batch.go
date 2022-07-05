package batch

import (
	"gitlab.com/wshaman/hw-concurrency/lib/semaphore"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	sem := semaphore.New(pool)
	done := make(chan bool, 1)

	for i := int64(0); i < n; i++ {
		sem.Push()
		go func(j int64) {
			defer sem.Pull()

			res = append(res, getOne(j))

			if j == n {
				done <- true
			}
		}(i)
	}
	<-done

	return
}
