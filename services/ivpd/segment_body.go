package ivpd

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SegmentBody invokes the ivpd.SegmentBody API synchronously
// api document: https://help.aliyun.com/api/ivpd/segmentbody.html
func (client *Client) SegmentBody(request *SegmentBodyRequest) (response *SegmentBodyResponse, err error) {
	response = CreateSegmentBodyResponse()
	err = client.DoAction(request, response)
	return
}

// SegmentBodyWithChan invokes the ivpd.SegmentBody API asynchronously
// api document: https://help.aliyun.com/api/ivpd/segmentbody.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SegmentBodyWithChan(request *SegmentBodyRequest) (<-chan *SegmentBodyResponse, <-chan error) {
	responseChan := make(chan *SegmentBodyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SegmentBody(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SegmentBodyWithCallback invokes the ivpd.SegmentBody API asynchronously
// api document: https://help.aliyun.com/api/ivpd/segmentbody.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SegmentBodyWithCallback(request *SegmentBodyRequest, callback func(response *SegmentBodyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SegmentBodyResponse
		var err error
		defer close(result)
		response, err = client.SegmentBody(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SegmentBodyRequest is the request struct for api SegmentBody
type SegmentBodyRequest struct {
	*requests.RpcRequest
	ImageUrl string `position:"Body" name:"ImageUrl"`
}

// SegmentBodyResponse is the response struct for api SegmentBody
type SegmentBodyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSegmentBodyRequest creates a request to invoke SegmentBody API
func CreateSegmentBodyRequest() (request *SegmentBodyRequest) {
	request = &SegmentBodyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "SegmentBody", "", "")
	request.Method = requests.POST
	return
}

// CreateSegmentBodyResponse creates a response to parse from SegmentBody response
func CreateSegmentBodyResponse() (response *SegmentBodyResponse) {
	response = &SegmentBodyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
