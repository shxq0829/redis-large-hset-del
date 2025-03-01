package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var (
	addr        = flag.String("addr", "localhost:6379", "Redis Addr")
	clusterMode = flag.Bool("clusterMode", false, "Cluster Mode")
	key         = flag.String("key", "test", "Hset Key")
	password    = flag.String("password", "", "Password")
	pattern     = flag.String("pattern", "*", "Scan Pattern")
	batchSize   = flag.Int64("batchSize", 1, "Batch Size")
)

func main() {
	flag.Parse()

	if *key == "" {
		fmt.Println("Key flag not found")
		os.Exit(1)
	}

	var client *redis.Client
	var clusterClient *redis.ClusterClient
	if *clusterMode {
		addrs := strings.Split(*addr, ",")
		clusterClient = redis.NewClusterClient(&redis.ClusterOptions{Addrs: addrs})
		pong, err := clusterClient.Ping().Result()

		if err != nil || pong == "" {
			log.Fatal("\n\nREDIS NOT CONNECT : ", err)
		}
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:        *addr,
			Password:    *password,
			DB:          0,
			ReadTimeout: 1 * time.Minute,
		})

		pong, err := client.Ping().Result()

		if err != nil || pong == "" {
			log.Fatal("\n\nREDIS NOT CONNECT : ", err)
		}
	}

	cursor := uint64(0)
	for {
		var result []string
		var err error
		if *clusterMode {
			result, cursor, err = clusterClient.HScan(*key, cursor, *pattern, *batchSize).Result()
		} else {
			result, cursor, err = client.HScan(*key, cursor, *pattern, *batchSize).Result()
		}
		fmt.Printf("cursor:%d,size:%d,len:%d", cursor, *batchSize, len(result))
		fmt.Println()

		if err != nil {
			log.Fatalf("could not hscan: %q\n", err)
		}

		for i := 0; i < len(result); i = i + 2 {
			if *clusterMode {
				clusterClient.HDel(*key, result[i])
			} else {
				fmt.Printf("key:%s,field:%s", *key, result[i])
				fmt.Println()
				client.HDel(*key, result[i])
			}
		}

		if cursor == 0 {
			break
		}
	}
}
