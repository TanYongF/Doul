package svc

import (
	"bufio"
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"go_code/Doul/app/security/internal/config"
	"log"
	"os"
)

type ServiceContext struct {
	Config config.Config
	Filter stringx.Trie
}

func NewServiceContext(c config.Config) *ServiceContext {
	//构建敏感词字典树
	sensitiveWordPath := "C:/Projects/Go/src/Doul/app/security/doc/words.txt"
	return &ServiceContext{
		Config: c,
		Filter: loadSensitiveTrie(sensitiveWordPath),
	}
}

func loadSensitiveTrie(wordsPath string) stringx.Trie {
	//auth.LoadFromFile(words)
	words := make([]string, 0)
	fd, err := os.Open(wordsPath)
	if err != nil {
		return nil
	}
	defer fd.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stringx.NewTrie(words, stringx.WithMask('?'))
}
