package main

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"insect/pb/crawl"
	"insect/pb/redis"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"

	"google.golang.org/grpc"

	"golang.org/x/text/encoding/simplifiedchinese"

	"golang.org/x/text/transform"

	"github.com/PuerkitoBio/goquery"
)

type server struct{}

var (
	ArticleReg = regexp.MustCompile(`https://new.qq.com/omn/20200429/20200429A0QK1100.html`)
	LinkReg    = regexp.MustCompile(``)
)

func (s *server) Crawl(ctx context.Context, in *crawl.CrawlRequest) (*crawl.CrawlResponse, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr}
	if in.Timeout > 0 {
		client.Timeout = time.Millisecond * time.Duration(in.Timeout)
	}
	var req *http.Request
	var err error
	req, err = http.NewRequest("", in.Url, nil)
	if err != nil {
		fmt.Println(err)
	}
	for _, header := range in.Headers {
		req.Header.Add(header.Key, header.Value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	r := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {

	}
	//DB conn
	conn, err := grpc.Dial("", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	db := redis.NewDBServiceClient(conn)
	links := make([]string, 0, 100)
	_, err = db.SadD(ctx, &redis.SadDRequest{CrawledURL: []string{in.Url}})
	if err != nil {
		fmt.Println(err)
	}
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		url, isExist := selection.Attr("href")
		if !isExist {
			return
		}
		//估计需要用一下正则表达式过滤一下
		if LinkReg.MatchString(url) || ArticleReg.MatchString(url) {
			judge, err := db.SisMember(ctx, &redis.SisMemberRequest{Url: fmt.Sprintf("%x", md5.Sum([]byte(url)))})
			if err != nil {
				fmt.Println(err)
				return
			}
			if judge != nil && !judge.IsExist {
				links = append(links, url)
			}
		}
	})
	if len(links) != 0 {
		_, err = db.PushQueue(ctx, &redis.PushQueueRequest{Urls: links})
		if err != nil {
			fmt.Println(err)
		}
	}
	content := make([]string, 0, 100)
	doc.Find("#content #clearfix #content-article p").Each(func(i int, selection *goquery.Selection) {
		content = append(content, selection.Text())
	})
	//下面就开始将文章内容存储到MySQL数据库
	//...
	return &crawl.CrawlResponse{Signal: in.GetUrl() + " is completed."}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	ser := grpc.NewServer()
	//将WaitingService服务注册到ser,也就是注册到gRpc服务里面
	crawl.RegisterCrawlServer(ser, &server{})
	err = ser.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}
