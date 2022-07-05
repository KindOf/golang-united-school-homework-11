package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func worker(jobs <-chan int64, results chan<- user) {
	for id := range jobs {
		results <- getOne(id)
	}
}

func getBatch(n int64, pool int64) (res []user) {
	num := int(n)
	jobs := make(chan int64, num)
	results := make(chan user, num)

	for j := int64(0); j < pool; j++ {
		go worker(jobs, results)
	}

	for i := int64(0); i < n; i++ {
		jobs <- i
	}
	close(jobs)

	for u := 0; u < num; u++ {
		res = append(res, <-results)
	}

	return
}
