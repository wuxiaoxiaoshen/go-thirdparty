package main

import "fmt"

var poolAction *PoolAction

func init() {
	poolAction = NewPoolAction("0.0.0.0:6377")
}

func Example() {
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
	{
		// list
		LAction := NewListAction(&con)
		con.Do("DEL", "list::A")
		ok := LAction.LPush("list::A", "go", "python", "java")
		fmt.Println(ok)
		v, ok := LAction.LRange("list::A", 0, -1)
		fmt.Println(v, ok)
	}
}

func main(){
	//ExampleSentinel()
	ExampleCluster()
}

func ExampleSentinel(){
	sentinel :=NewSentinelAction("mymaster", []string{":26376", ":26378", ":26377"}, "admin")
	fmt.Println(sentinel.Client.Ping().String())
	fmt.Println(sentinel.Client.DBSize())
}
func ExampleCluster(){
	cluster := NewClusterAction()
	fmt.Println(cluster.Client.Get("hello"))
	cluster.Client.Set("world", "hello",0)
	fmt.Println(cluster.Client.Get("world"))
	fmt.Println(cluster.Client.DBSize())
	fmt.Println(cluster.Client.ClusterNodes())
}