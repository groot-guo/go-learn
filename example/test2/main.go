package main

import "fmt"

const (
	redisTypeSigServer    = 1
	redisTypeDoubleServer = 2
	redisTypeJiQun        = 3

	choseTypeCloudServer = 1
	choseTypeRedisServer = 2
	choseTypeMysqlServer = 3
	choseTypeExit        = 4
)

type CouldServer struct {
	cpu        int
	ram        int
	memory     int
	monthPrice int
	yearPrice  int
}

type RedisServer struct {
	ram        int
	monthPrice int
	yearPrice  int
	redisType  int
}

type MysqlServer struct {
	memory     int
	monthPrice int
	yearPrice  int
}

type ProductList struct {
	CouldList []CouldServer
	RedisList []RedisServer
	MysqlList []MysqlServer
}

type User struct {
	Product ProductList
}

var (
	product        = new(ProductList)
	productListMap = map[int]string{
		1: "云服务器",
		2: "Redis",
		3: "Mysql数据库",
		4: "退出",
	}
	redisTypeMap = map[int]string{
		1: "单节点",
		2: "主父节点",
		3: "集群模式，多地部署",
	}
)

func init() {
	product.CouldList = []CouldServer{
		{
			cpu:        1,
			ram:        1,
			memory:     10,
			monthPrice: 100,
			yearPrice:  1000,
		},
		{
			cpu:        4,
			ram:        4,
			memory:     50,
			monthPrice: 200,
			yearPrice:  2000,
		},
		{
			cpu:        8,
			ram:        8,
			memory:     100,
			monthPrice: 300,
			yearPrice:  3000,
		},
	}
	product.RedisList = []RedisServer{
		{
			ram:        10,
			monthPrice: 100,
			yearPrice:  1000,
			redisType:  redisTypeSigServer,
		},
		{
			ram:        10,
			monthPrice: 200,
			yearPrice:  2000,
			redisType:  redisTypeDoubleServer,
		},
		{
			ram:        20,
			monthPrice: 300,
			yearPrice:  3000,
			redisType:  redisTypeJiQun,
		},
	}
	product.MysqlList = []MysqlServer{
		{
			memory:     100,
			monthPrice: 100,
			yearPrice:  1000,
		},
		{
			memory:     500,
			monthPrice: 300,
			yearPrice:  3000,
		},
		{
			memory:     1024,
			monthPrice: 500,
			yearPrice:  10000,
		},
	}
}

func main() {
	fmt.Println("------私有云产品预估费用--------")
	for {
		fmt.Println("根据以下选项选择对应产品查看，并回车确认:")
		fmt.Println("1.云服务器")
		fmt.Println("2.Redis")
		fmt.Println("3.Mysql数据库")
		fmt.Println("4.退出")
		chose, err := fmt.Scanf("请输入:")
		if err != nil {
			fmt.Println("请重新选择")
			continue
		}

		switch chose {
		case choseTypeCloudServer:
			for n, cloud := range product.CouldList {
				fmt.Printf("1-%d. %d核%d内存%d硬盘，包月%d, 包年%d\n",
					n, cloud.cpu, cloud.ram, cloud.memory, cloud.monthPrice,
					cloud.yearPrice)
			}
		case choseTypeRedisServer:
			for n, redis := range product.RedisList {
				fmt.Printf("2-%d. %d内存%s，包月%d, 包年%d\n",
					n, redis.ram, redisTypeMap[redis.redisType], redis.monthPrice,
					redis.yearPrice)
			}
		case choseTypeMysqlServer:
			for n, mysql := range product.MysqlList {
				fmt.Printf("3-%d. %d存储%s，包月%d, 包年%d\n",
					n, mysql.memory, mysql.monthPrice, mysql.yearPrice)
			}
		case choseTypeExit:
			return
		}
	}

}
