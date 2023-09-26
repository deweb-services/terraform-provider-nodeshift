package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider("jx3n67f4a5zspnhqdae54j2sdtga",
			"jyoiszqe7zjsaupi26b6kjdese4nn3z7nvuwdajz5eoecubsyn3xe", "")))
	if err != nil {
		panic(err)
	}

	cfg.Region = "us-west-1"
	//cfg.Credentials = endpointcreds.New("https://s3.dws.so")
	cli := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String("https://s3.dws.local/")
	})

	op, err := cli.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		log.Println(err)
		log.Println(op)
		panic(err)
	}
	log.Println(op)
}
