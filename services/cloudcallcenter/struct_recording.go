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

// Recording is a nested struct in cloudcallcenter response
type Recording struct {
	ContactId          string `json:"ContactId" xml:"ContactId"`
	FileDescription    string `json:"FileDescription" xml:"FileDescription"`
	CalledNumber       string `json:"CalledNumber" xml:"CalledNumber"`
	QualityCheckTid    string `json:"QualityCheckTid" xml:"QualityCheckTid"`
	ContactType        string `json:"ContactType" xml:"ContactType"`
	CallingNumber      string `json:"CallingNumber" xml:"CallingNumber"`
	StartTime          int64  `json:"StartTime" xml:"StartTime"`
	AgentName          string `json:"AgentName" xml:"AgentName"`
	AgentId            string `json:"AgentId" xml:"AgentId"`
	Duration           int    `json:"Duration" xml:"Duration"`
	InstanceId         string `json:"InstanceId" xml:"InstanceId"`
	Channel            string `json:"Channel" xml:"Channel"`
	QualityCheckTaskId string `json:"QualityCheckTaskId" xml:"QualityCheckTaskId"`
	FileName           string `json:"FileName" xml:"FileName"`
	FilePath           string `json:"FilePath" xml:"FilePath"`
}
