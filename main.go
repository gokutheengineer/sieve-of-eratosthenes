package main

import "fmt"

var primesUntil = 1_000_000

func main() {
	primes := sieve()

	for p := range primes {
		fmt.Println(p)
	}

}

func sieve() chan int {
	out_ch := make(chan int)

	go func() {
		in_ch := generate()
		for {
			prime, ok := <-in_ch
			if !ok {
				close(out_ch)
				break
			}
			in_ch = filter(in_ch, prime)
			out_ch <- prime
		}
	}()
	return out_ch
}

func generate() chan int {
	in_ch := make(chan int)

	go func() {
		for i := 2; i < primesUntil; i++ {
			in_ch <- i
		}
		close(in_ch)
	}()
	return in_ch
}

func filter(in_ch chan int, prime int) chan int {
	out_ch := make(chan int)

	go func() {
		for {
			i, ok := <-in_ch
			if !ok {
				close(out_ch)
				break
			}
			if i%prime != 0 {
				out_ch <- i
			}
		}
	}()

	return out_ch
}
