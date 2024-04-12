package main

import (
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()
	opt, err := redis.ParseURL("redis://default:redispw@localhost:32768")
	if err != nil {
		panic(err)
	}
	opt.DialTimeout = 10 * time.Second

	fmt.Println(unsafe.Sizeof(ctx))

	client := redis.NewClient(opt)
	fmt.Println(unsafe.Sizeof(*client))
	fmt.Println(unsafe.Sizeof(client))

	ctx = context.WithValue(ctx, "redis", *client)
	cnn := client.Conn()
	defer cnn.Close()
	fmt.Println("for")

	for i := 0; i < 10; i++ {
		ctx = getContext(ctx, *client, i)
	}

	fmt.Println("end")
	fmt.Println(unsafe.Sizeof(*client))

	fmt.Println(unsafe.Sizeof(ctx))

	fmt.Println(unsafe.Sizeof(&ctx))

	cnn.Set(ctx, "1", "1", 10)
	//cnn.SetEx(ctx, "2", "2", 10)
	//cnn.SetNX(ctx, "3", "3", 10)
	//cnn.SetXX(ctx, "4", "4", 10)

	//value := cnn.Get(ctx, "1")
	//str, err := value.Result()
	//if err != nil {
	//	client.Close()
	//	panic(err)
	//}
	//fmt.Println(str)
	cmd := cnn.Incr(ctx, "ID")
	res, err := cmd.Result()
	fmt.Println(res, err)
	cmd = cnn.IncrBy(ctx, "ID", 1)
	res, err = cmd.Result()
	fmt.Println(res, err)
	cmd2 := cnn.Get(ctx, "ID")
	res1, err := cmd2.Result()
	fmt.Println(res1, err)
	dcmd := cnn.PTTL(ctx, "ID")
	res2, err := dcmd.Result()
	fmt.Println(res2, err)
	dcmd = cnn.TTL(ctx, "ID")
	res2, err = dcmd.Result()
	fmt.Println(res2, err)
	cnn.Set(ctx, "ID", 0, 0)

	//val2, err := cnn.GetEx(ctx, "2", 10).Result()
	//if err != nil {
	//	client.Close()
	//	panic(err)
	//}
	//fmt.Println(val2)
	//mc, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/?authSource=admin"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//col := mc.Database("test").Collection("test")
	//res, err := col.InsertMany(ctx, []interface{}{
	//	bson.M{
	//		"open_id": "123",
	//	},
	//	bson.M{
	//		"open_id": "456",
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(res)
	//
	//val := col.FindOne(ctx, "open_id")
	////val.Decode()
	//fmt.Println(val)
}

func getContext(ctx context.Context, client redis.Client, number int) context.Context {
	fmt.Println(unsafe.Sizeof(ctx))

	return context.WithValue(ctx, "redis"+string(rune(number)), client)
}
