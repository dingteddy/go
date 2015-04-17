package main

import (
	"flag"
	"fmt"
	"math/rand"
	"redis"
	"time"
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	pool  *redis.Pool
	rhost = flag.String("rhost", ":6379", "")
	rpass = flag.String("rpass", "", "")
)

func main() {
	flag.Parse()
	pool = newPool(*rhost, *rpass)
	if pool == nil {
		fmt.Println("haha\n")
	}
	//conn, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	//if err != nil {
	//	panic(err)
	//}
	for i := 0; i < 5; i++ {
		go func() {
			conn := pool.Get()
			if conn == nil {
				fmt.Printf("aiai")
			}
			//fmt.Println(conn)
			for i := 0; i < 9; i++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))

				key := fmt.Sprintf("user:user%d", i)
				_, err := conn.Do("SET", key, 123)
				_, err = conn.Do("APPEND", key, 87)
				if err != nil {
					fmt.Printf("operation failed!\n")
				}

				user, err := redis.Int(conn.Do("GET", key))
				conn.Do("DEL", key)

				randnum := r.Intn(1000)
				fmt.Printf("user%d is %d, rand is %d, key is %s\n", i, user, randnum, key)
				time.Sleep(time.Millisecond * time.Duration(randnum))
			}
			conn.Close()
		}()
	}
	for {
		time.Sleep(1)
	}
	pool.Close()
}
