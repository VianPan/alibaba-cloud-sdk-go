package vs

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

// ModifyTemplate invokes the vs.ModifyTemplate API synchronously
func (client *Client) ModifyTemplate(request *ModifyTemplateRequest) (response *ModifyTemplateResponse, err error) {
	response = CreateModifyTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTemplateWithChan invokes the vs.ModifyTemplate API asynchronously
func (client *Client) ModifyTemplateWithChan(request *ModifyTemplateRequest) (<-chan *ModifyTemplateResponse, <-chan error) {
	responseChan := make(chan *ModifyTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTemplate(request)
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

// ModifyTemplateWithCallback invokes the vs.ModifyTemplate API asynchronously
func (client *Client) ModifyTemplateWithCallback(request *ModifyTemplateRequest, callback func(response *ModifyTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTemplateResponse
		var err error
		defer close(result)
		response, err = client.ModifyTemplate(request)
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

// ModifyTemplateRequest is the request struct for api ModifyTemplate
type ModifyTemplateRequest struct {
	*requests.RpcRequest
	HlsTs            string           `position:"Query" name:"HlsTs"`
	OssEndpoint      string           `position:"Query" name:"OssEndpoint"`
	Description      string           `position:"Query" name:"Description"`
	OssFilePrefix    string           `position:"Query" name:"OssFilePrefix"`
	JpgOverwrite     string           `position:"Query" name:"JpgOverwrite"`
	StartTime        string           `position:"Query" name:"StartTime"`
	JpgOnDemand      string           `position:"Query" name:"JpgOnDemand"`
	Id               string           `position:"Query" name:"Id"`
	Retention        requests.Integer `position:"Query" name:"Retention"`
	ShowLog          string           `position:"Query" name:"ShowLog"`
	HlsM3u8          string           `position:"Query" name:"HlsM3u8"`
	OssBucket        string           `position:"Query" name:"OssBucket"`
	TransConfigsJSON string           `position:"Query" name:"TransConfigsJSON"`
	EndTime          string           `position:"Query" name:"EndTime"`
	Trigger          string           `position:"Query" name:"Trigger"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	JpgSequence      string           `position:"Query" name:"JpgSequence"`
	Mp4              string           `position:"Query" name:"Mp4"`
	Flv              string           `position:"Query" name:"Flv"`
	Name             string           `position:"Query" name:"Name"`
	Callback         string           `position:"Query" name:"Callback"`
	Interval         requests.Integer `position:"Query" name:"Interval"`
	FileFormat       string           `position:"Query" name:"FileFormat"`
	Region           string           `position:"Query" name:"Region"`
}

// ModifyTemplateResponse is the response struct for api ModifyTemplate
type ModifyTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateModifyTemplateRequest creates a request to invoke ModifyTemplate API
func CreateModifyTemplateRequest() (request *ModifyTemplateRequest) {
	request = &ModifyTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ModifyTemplate", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyTemplateResponse creates a response to parse from ModifyTemplate response
func CreateModifyTemplateResponse() (response *ModifyTemplateResponse) {
	response = &ModifyTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
