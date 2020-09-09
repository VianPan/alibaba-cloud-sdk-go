package cloudcallcenter

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

// ListAgentStates invokes the cloudcallcenter.ListAgentStates API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listagentstates.html
func (client *Client) ListAgentStates(request *ListAgentStatesRequest) (response *ListAgentStatesResponse, err error) {
	response = CreateListAgentStatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAgentStatesWithChan invokes the cloudcallcenter.ListAgentStates API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listagentstates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAgentStatesWithChan(request *ListAgentStatesRequest) (<-chan *ListAgentStatesResponse, <-chan error) {
	responseChan := make(chan *ListAgentStatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAgentStates(request)
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

// ListAgentStatesWithCallback invokes the cloudcallcenter.ListAgentStates API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listagentstates.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAgentStatesWithCallback(request *ListAgentStatesRequest, callback func(response *ListAgentStatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAgentStatesResponse
		var err error
		defer close(result)
		response, err = client.ListAgentStates(request)
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

// ListAgentStatesRequest is the request struct for api ListAgentStates
type ListAgentStatesRequest struct {
	*requests.RpcRequest
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	State        string           `position:"Query" name:"State"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	Agents       string           `position:"Query" name:"Agents"`
}

// ListAgentStatesResponse is the response struct for api ListAgentStates
type ListAgentStatesResponse struct {
	*responses.BaseResponse
	RequestId      string                `json:"RequestId" xml:"RequestId"`
	Success        bool                  `json:"Success" xml:"Success"`
	Code           string                `json:"Code" xml:"Code"`
	Message        string                `json:"Message" xml:"Message"`
	HttpStatusCode int                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListAgentStates `json:"Data" xml:"Data"`
}

// CreateListAgentStatesRequest creates a request to invoke ListAgentStates API
func CreateListAgentStatesRequest() (request *ListAgentStatesRequest) {
	request = &ListAgentStatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListAgentStates", "", "")
	request.Method = requests.POST
	return
}

// CreateListAgentStatesResponse creates a response to parse from ListAgentStates response
func CreateListAgentStatesResponse() (response *ListAgentStatesResponse) {
	response = &ListAgentStatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}