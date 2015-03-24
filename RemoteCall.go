package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//var client = &http.Client{}
type CallWork struct {
	url    string
	result chan string
}

func Call(urls []string) map[string]string {
	workerCount := len(urls)
	fmt.Println(workerCount)
	m := make(map[string]string, workerCount)
	jobs, addJobDone := make([]CallWork, workerCount), make(chan bool)

	go addJob(jobs, urls, addJobDone)

	if addDone := <-addJobDone; addDone {
		for i := 0; i < workerCount; i++ {
			go DoJobs(jobs)
		}
	}

	for i := 0; i < workerCount; i++ {
		select {
		case callResult := <-jobs[i].result:
			m[jobs[i].url] = callResult
		case <-time.After(3 * time.Second):
			panic(ERR_TIMEOUT)
		}
	}

	return m
}

func (this CallWork) Get(url string) string {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")

	response, _ := client.Do(request)

	if response.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()

		return string(body)
	} else {
		panic(ERR_REQUEST_TIMEOUT)
	}

}

func addJob(jobs []CallWork, urls []string, addJobDone chan bool) {
	defer func() {
		close(addJobDone)
	}()
	for k, url := range urls {
		jobs[k] = CallWork{url: url, result: make(chan string)}
	}

	addJobDone <- true
}

func DoJobs(jobs []CallWork) {
	for _, job := range jobs {
		callResult := job.Get(job.url)
		job.result <- callResult
	}
}
