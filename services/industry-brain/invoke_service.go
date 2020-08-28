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

// InvokeService invokes the industry_brain.InvokeService API synchronously
// api document: https://help.aliyun.com/api/industry-brain/invokeservice.html
func (client *Client) InvokeService(request *InvokeServiceRequest) (response *InvokeServiceResponse, err error) {
	response = CreateInvokeServiceResponse()
	err = client.DoAction(request, response)
	return
}

// InvokeServiceWithChan invokes the industry_brain.InvokeService API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/invokeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InvokeServiceWithChan(request *InvokeServiceRequest) (<-chan *InvokeServiceResponse, <-chan error) {
	responseChan := make(chan *InvokeServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InvokeService(request)
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

// InvokeServiceWithCallback invokes the industry_brain.InvokeService API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/invokeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InvokeServiceWithCallback(request *InvokeServiceRequest, callback func(response *InvokeServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InvokeServiceResponse
		var err error
		defer close(result)
		response, err = client.InvokeService(request)
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

// InvokeServiceRequest is the request struct for api InvokeService
type InvokeServiceRequest struct {
	*requests.RpcRequest
	RequestParams            string           `position:"Query" name:"RequestParams"`
	ShowBizInfo              requests.Boolean `position:"Query" name:"ShowBizInfo"`
	ForceInvokeConfiguration requests.Boolean `position:"Query" name:"ForceInvokeConfiguration"`
	Context                  string           `position:"Query" name:"Context"`
	ServiceId                string           `position:"Query" name:"ServiceId"`
	RequestData              string           `position:"Query" name:"RequestData"`
	ShowParams               requests.Boolean `position:"Query" name:"ShowParams"`
}

// InvokeServiceResponse is the response struct for api InvokeService
type InvokeServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateInvokeServiceRequest creates a request to invoke InvokeService API
func CreateInvokeServiceRequest() (request *InvokeServiceRequest) {
	request = &InvokeServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2019-06-29", "InvokeService", "", "")
	request.Method = requests.POST
	return
}

// CreateInvokeServiceResponse creates a response to parse from InvokeService response
func CreateInvokeServiceResponse() (response *InvokeServiceResponse) {
	response = &InvokeServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
