package main

import (
	"context"
	"fmt"
	"insect/pb/crawl"
	"insect/pb/redis"
	"log"

	rd "github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

var (
	err  error
	conn *grpc.ClientConn
	db   redis.DBServiceClient
)

func init() {
	//Redis DB
	conn, err = grpc.Dial("", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	db = redis.NewDBServiceClient(conn)
}

func main() {
	var jerr error
	defer conn.Close()
	for jerr = nil; jerr != rd.ErrNil; _, jerr = db.RangeQueue(context.Background(), nil) {
		popData, err := db.PopQueue(context.Background(), nil)
		if err != nil {
			log.Println(err)
		}
		if popData != nil && popData.Url != "" {
			//负载均衡得到一个节点IP
			//...
			//拨号那个节点
			conn, err := grpc.Dial("", grpc.WithInsecure())
			if err != nil {
				log.Println(err)
			}
			client := crawl.NewCrawlClient(conn)
			resp, err := client.Crawl(context.Background(), &crawl.CrawlRequest{Url: popData.Url})
			if err != nil {
				log.Println(err)
			}
			log.Println(resp.GetSignal())
			conn.Close()
		}
	}
}
