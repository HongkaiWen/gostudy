package concurrency

func put(ch chan) {
	ch <- 1
	a <- ch
}