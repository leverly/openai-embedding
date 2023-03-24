package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"math/rand"
	"time"
)

var apiList = [3]string{
	"sk-ziIjWUX7ZmfiVA86pSPH",
	"sk-Fsf8MmcrssrKy5B7Pqvx",
	"sk-t3wJTgGCLSv6nwNrAmBL",
}

type OpenaiClient struct {
	client *openai.Client
}

func newClient() *OpenaiClient {
	rand.Seed(time.Now().Unix())
	return &OpenaiClient{client: openai.NewClient(apiList[rand.Int()%len(apiList)])}
}

func (c *OpenaiClient) Embedding(blocks []string) (error, []openai.Embedding) {
	resp, err := c.client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: blocks,
			Model: openai.AdaEmbeddingV2,
		})
	if err != nil {
		fmt.Printf("Create Embeddings error: %v\n", err)
		return err, nil
	}
	return nil, resp.Data
}
