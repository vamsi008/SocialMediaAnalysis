package communicator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpCommunicator struct {
}

func (httpCommunication *HttpCommunicator) Communicate(encodedUrl string) (string, error) {
	var responsebody string
	var err error
	fmt.Println("Request = ", encodedUrl)
	time.Sleep(3000 * time.Millisecond)
	resp, err := http.Get(encodedUrl)

	if err != nil {
		fmt.Println("The error is ::", err)

	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("The error is::", err)

		} else {
			responsebody = string(body)
			//fmt.Println("The response is ::\n", responsebody)
		}
	}
	fmt.Println("Response = ", responsebody)
	return responsebody, err
}
