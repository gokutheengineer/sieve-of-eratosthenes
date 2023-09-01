package main

import "fmt"

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
			prime := <-in_ch
			in_ch = filter(in_ch, prime)
			out_ch <- prime
		}
	}()
	return out_ch
}

func generate() chan int {
	in_ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			in_ch <- i
		}
	}()

	return in_ch
}

func filter(in_ch chan int, prime int) chan int {
	out_ch := make(chan int)

	go func() {
		for {
			if i := <-in_ch; i%prime != 0 {
				out_ch <- i
			}
		}
	}()

	return out_ch
}
