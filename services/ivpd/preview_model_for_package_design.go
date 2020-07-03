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

// PreviewModelForPackageDesign invokes the ivpd.PreviewModelForPackageDesign API synchronously
// api document: https://help.aliyun.com/api/ivpd/previewmodelforpackagedesign.html
func (client *Client) PreviewModelForPackageDesign(request *PreviewModelForPackageDesignRequest) (response *PreviewModelForPackageDesignResponse, err error) {
	response = CreatePreviewModelForPackageDesignResponse()
	err = client.DoAction(request, response)
	return
}

// PreviewModelForPackageDesignWithChan invokes the ivpd.PreviewModelForPackageDesign API asynchronously
// api document: https://help.aliyun.com/api/ivpd/previewmodelforpackagedesign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreviewModelForPackageDesignWithChan(request *PreviewModelForPackageDesignRequest) (<-chan *PreviewModelForPackageDesignResponse, <-chan error) {
	responseChan := make(chan *PreviewModelForPackageDesignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreviewModelForPackageDesign(request)
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

// PreviewModelForPackageDesignWithCallback invokes the ivpd.PreviewModelForPackageDesign API asynchronously
// api document: https://help.aliyun.com/api/ivpd/previewmodelforpackagedesign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreviewModelForPackageDesignWithCallback(request *PreviewModelForPackageDesignRequest, callback func(response *PreviewModelForPackageDesignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreviewModelForPackageDesignResponse
		var err error
		defer close(result)
		response, err = client.PreviewModelForPackageDesign(request)
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

// PreviewModelForPackageDesignRequest is the request struct for api PreviewModelForPackageDesign
type PreviewModelForPackageDesignRequest struct {
	*requests.RpcRequest
	MaterialName string                                     `position:"Body" name:"MaterialName"`
	ElementList  *[]PreviewModelForPackageDesignElementList `position:"Body" name:"ElementList"  type:"Repeated"`
	DataId       string                                     `position:"Body" name:"DataId"`
	MaterialType string                                     `position:"Body" name:"MaterialType"`
	ModelType    string                                     `position:"Body" name:"ModelType"`
	Category     string                                     `position:"Body" name:"Category"`
}

// PreviewModelForPackageDesignElementList is a repeated param struct in PreviewModelForPackageDesignRequest
type PreviewModelForPackageDesignElementList struct {
	ImageUrl string `name:"ImageUrl"`
	SideName string `name:"SideName"`
}

// PreviewModelForPackageDesignResponse is the response struct for api PreviewModelForPackageDesign
type PreviewModelForPackageDesignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreatePreviewModelForPackageDesignRequest creates a request to invoke PreviewModelForPackageDesign API
func CreatePreviewModelForPackageDesignRequest() (request *PreviewModelForPackageDesignRequest) {
	request = &PreviewModelForPackageDesignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "PreviewModelForPackageDesign", "", "")
	request.Method = requests.POST
	return
}

// CreatePreviewModelForPackageDesignResponse creates a response to parse from PreviewModelForPackageDesign response
func CreatePreviewModelForPackageDesignResponse() (response *PreviewModelForPackageDesignResponse) {
	response = &PreviewModelForPackageDesignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
