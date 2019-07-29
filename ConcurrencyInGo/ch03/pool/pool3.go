package pool

import "time"

/**
 * created: 2019/7/29 15:39
 * By Will Fan
 */
func main() {
	connectToService := func () interface{} {
		time.Sleep(1 * time.Second)
		return struct {}{}
	}


}
