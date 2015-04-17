package main

import (
	"errors"
	"fmt"
	"reflect"
)

/*var MAX_POOL_SIZE int
var MySQLPool chan *mysql.MySQL

func getMySQL() *mysql.MySQL {
	if MySQLPool == nil {
		MySQLPool = make(chan *mysql.MySQL, MAX_POOL_SIZE)
	}
	if len(MySQLPool) == 0 {
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				mysql := mysql.New()
				err := mysql.Connect("127.0.0.1", "root", "", "teamtop3", 3306)
				if err != nil {
					panic(err.String())
				}
				putMySQL(mysql)
			}
		}()
	}
	return <-MySQLPool
}
func putMySQL(conn *mysql.MySQL) {
	if MySQLPool == nil {
		MySQLPool = make(chan *mysql.MySQL, MAX_POOL_SIZE)
	}
	if len(MySQLPool) == MAX_POOL_SIZE {
		conn.Close()
		return
	}
	MySQLPool <- conn
}*/

func foobar() {
	fmt.Printf("i am foobar\n")
}

func foo() {
	fmt.Printf("i am foo\n")
}

func bar(a int, b int, c int) {
	fmt.Printf("i am bar, params: %d %d %d, sum: %d\n", a, b, c, a+b+c)
}

type myfunc map[string]func()

type funcs map[string]interface{}

func test() {
	tmpmyfunc := make(myfunc)
	tmpmyfunc["foobar"] = foobar
	tmpmyfunc["foobar"]()
	tmpfuncs := make(funcs)
	tmpfuncs["foo"] = foo
	tmpfuncs["bar"] = bar
	Call(tmpfuncs, "foo")
	Call(tmpfuncs, "bar", 1, 2, 3)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		fmt.Printf("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {
	test()
}
