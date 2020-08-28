package industry_brain

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

// GetServiceResultAsync invokes the industry_brain.GetServiceResultAsync API synchronously
// api document: https://help.aliyun.com/api/industry-brain/getserviceresultasync.html
func (client *Client) GetServiceResultAsync(request *GetServiceResultAsyncRequest) (response *GetServiceResultAsyncResponse, err error) {
	response = CreateGetServiceResultAsyncResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceResultAsyncWithChan invokes the industry_brain.GetServiceResultAsync API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getserviceresultasync.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceResultAsyncWithChan(request *GetServiceResultAsyncRequest) (<-chan *GetServiceResultAsyncResponse, <-chan error) {
	responseChan := make(chan *GetServiceResultAsyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceResultAsync(request)
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

// GetServiceResultAsyncWithCallback invokes the industry_brain.GetServiceResultAsync API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getserviceresultasync.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetServiceResultAsyncWithCallback(request *GetServiceResultAsyncRequest, callback func(response *GetServiceResultAsyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceResultAsyncResponse
		var err error
		defer close(result)
		response, err = client.GetServiceResultAsync(request)
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

// GetServiceResultAsyncRequest is the request struct for api GetServiceResultAsync
type GetServiceResultAsyncRequest struct {
	*requests.RpcRequest
	ServiceId string `position:"Query" name:"ServiceId"`
	TaskId    string `position:"Query" name:"TaskId"`
}

// GetServiceResultAsyncResponse is the response struct for api GetServiceResultAsync
type GetServiceResultAsyncResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateGetServiceResultAsyncRequest creates a request to invoke GetServiceResultAsync API
func CreateGetServiceResultAsyncRequest() (request *GetServiceResultAsyncRequest) {
	request = &GetServiceResultAsyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2019-06-29", "GetServiceResultAsync", "", "")
	request.Method = requests.POST
	return
}

// CreateGetServiceResultAsyncResponse creates a response to parse from GetServiceResultAsync response
func CreateGetServiceResultAsyncResponse() (response *GetServiceResultAsyncResponse) {
	response = &GetServiceResultAsyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
