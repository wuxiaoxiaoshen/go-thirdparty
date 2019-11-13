package main

import "fmt"

var poolAction *PoolAction

func init() {
	poolAction = NewPoolAction("0.0.0.0:6377")
}

func main() {
	con := poolAction.Get()
	{
		// string
		var S StringAction
		S = StringAction{
			Key:   "Test-key",
			Value: "1",
			con:   &con,
		}
		con.Do("DEL", S.Key)
		fmt.Println(S.Incr(10)) // 11
		fmt.Println(S.Decr(100))
		fmt.Println(S.Run("expire", "10", 100))
	}
	{
		// keys
		KAction := NewKeyAction(con)
		fmt.Println(KAction.Exists("A")) // true
		fmt.Println(KAction.Type("A"))   // string
		fmt.Println(KAction.DBSize())    //
		fmt.Println(KAction.Exists("B"))
		fmt.Println(KAction.Type("C"))
		fmt.Println(KAction.DBSize())
		fmt.Println(KAction.Del("A"))

	}
}
