package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

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
	ExampleSentinel()
	//ExampleCluster()
}

func ExampleSentinel(){
	sentinel :=NewSentinelAction("mymaster", []string{":26376", ":26378", ":26377"}, "admin")
	fmt.Println(sentinel.Client.Ping().String())
	fmt.Println(sentinel.Client.DBSize())
	//for i:=0;i<1000000000;i++{
	//	sentinel.Client.Do("sadd", "sentinel:test", i)
	//}
	keys := "tv:anchor:rank"
	results, _ := sentinel.Client.ZRangeWithScores(keys, 0, -1).Result()
	for _,i := range results{
		fmt.Println(i)
	}
	re, _ := sentinel.Client.ZRevRangeByScoreWithScores(
		keys, &redis.ZRangeBy{
			Min: "0",
			Max: "+inf",
			Count: 10,
		}).Result()
	for _, i :=range re{
		fmt.Println(i)
	}


}
func ExampleCluster(){
	cluster := NewClusterAction()
	fmt.Println(cluster.Client.Get("hello"))
	cluster.Client.Set("world", "hello",0)
	fmt.Println(cluster.Client.Get("world"))
	fmt.Println(cluster.Client.DBSize())
	fmt.Println(cluster.Client.ClusterNodes())
}