package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

func main() {
	lambda.Start(runLambda)
}

func runLambda() {
	start := time.Now()

	runHttpTests(1000)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func runHttpTests(loopAmount int) {
	var wg sync.WaitGroup
	wg.Add(loopAmount)

	for i := 0; i < loopAmount; i++ {
		go func(i int) {
			defer wg.Done()
			_, err := runHttpRequest(i)

			if err != nil {
				fmt.Println("Http Error!!!")
			}
		}(i)
	}

	wg.Wait()
}

func runHttpRequest(iteration int) ([]byte, error) {
	const url = "https://h3km0z2853.execute-api.eu-central-1.amazonaws.com/daniele-node"

	var jsonStr = fmt.Sprintf(`{"Id":"%v-%v", "Fullname":"%v"}`, iteration, time.Now().UTC().Format(time.RFC3339Nano), uuid.New())

	fmt.Println("Doing request:", jsonStr)

	var jsonData = []byte(jsonStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
