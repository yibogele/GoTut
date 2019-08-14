package main

/**
 * created: 2019/8/8 9:46
 * By Will Fan
 */
func main() {
	tee := func(
		done <-chan interface{},
		in <-chan interface{},
	) (_, _ <-chan interface{}){
		out1 := make(chan interface{})
		out2 := make(chan interface{})

		go func() {
			defer close(out1)
			defer close(out2)

			//for val := range orDone
		}()
	}
}
