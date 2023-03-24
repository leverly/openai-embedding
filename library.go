package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

type Library struct {
	blocks          []string
	blocksEmbddings []openai.Embedding
	init            bool
}

func newLibrary() *Library {
	return &Library{init: false}
}

func (l *Library) Init(filename string) error {
	l.blocks = []string{}
	if err := l.readFile(filename); err != nil {
		return err
	}
	// exhaust api quota
	client := newClient()
	err, result := client.Embedding(l.blocks)
	if err != nil {
		return err
	}
	l.blocksEmbddings = []openai.Embedding{}
	for _, res := range result {
		l.blocksEmbddings = append(l.blocksEmbddings, res)
	}
	l.init = true
	return nil
}

// read file parses the blocks
func (l *Library) readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var block string
	for scanner.Scan() {
		line := scanner.Bytes()
		// new block start
		if bytes.HasPrefix(line, []byte("###")) {
			if len(block) > 0 {
				l.blocks = append(l.blocks, string(block))
			}
			// reset to find a new block
			block = ""
		} else {
			if len(block) != 0 {
				block += ("\n" + string(line))
			} else {
				block = string(line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (l *Library) FindSimilarBlock(query string) (error, float32, string) {
	if l.init {
		// step 1. embedding the query
		max := float32(0.0)
		temp := []string{}
		temp = append(temp, query)
		client := newClient()
		err, result := client.Embedding(temp)
		if err != nil {
			return err, max, ""
		}

		// step2. find the similarest embedding from library
		index := -1
		for i, item := range l.blocksEmbddings {
			similar := cosineSimilarity(result[0].Embedding, item.Embedding)
			if similar > max {
				max = similar
				index = i
			}
		}
		return nil, max, l.blocks[index]
	}
	return &MyError{"not init"}, float32(0.0), ""
}
