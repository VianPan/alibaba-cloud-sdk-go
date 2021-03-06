package dataworks_public

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

// Data is a nested struct in dataworks_public response
type Data struct {
	HourSlaDetail       string                `json:"HourSlaDetail" xml:"HourSlaDetail"`
	LastModifyTime      int64                 `json:"LastModifyTime" xml:"LastModifyTime"`
	CreatorName         string                `json:"CreatorName" xml:"CreatorName"`
	ApplicationSecret   string                `json:"ApplicationSecret" xml:"ApplicationSecret"`
	CycTime             int64                 `json:"CycTime" xml:"CycTime"`
	WhereCondition      string                `json:"WhereCondition" xml:"WhereCondition"`
	Property            string                `json:"Property" xml:"Property"`
	BaselineType        string                `json:"BaselineType" xml:"BaselineType"`
	IsDefault           bool                  `json:"IsDefault" xml:"IsDefault"`
	Content             string                `json:"Content" xml:"Content"`
	EnvType             int                   `json:"EnvType" xml:"EnvType"`
	ProjectName         string                `json:"ProjectName" xml:"ProjectName"`
	SchedulerType       string                `json:"SchedulerType" xml:"SchedulerType"`
	Creator             string                `json:"Creator" xml:"Creator"`
	FileName            string                `json:"FileName" xml:"FileName"`
	TotalCount          int64                 `json:"TotalCount" xml:"TotalCount"`
	FinishTime          int64                 `json:"FinishTime" xml:"FinishTime"`
	Id                  int64                 `json:"Id" xml:"Id"`
	RuleType            int                   `json:"RuleType" xml:"RuleType"`
	TableName           string                `json:"TableName" xml:"TableName"`
	CategoryId          int64                 `json:"CategoryId" xml:"CategoryId"`
	PredictType         int                   `json:"PredictType" xml:"PredictType"`
	Trend               string                `json:"Trend" xml:"Trend"`
	InstanceId          int64                 `json:"InstanceId" xml:"InstanceId"`
	ExpHour             int                   `json:"ExpHour" xml:"ExpHour"`
	Status              string                `json:"Status" xml:"Status"`
	RemindUnit          string                `json:"RemindUnit" xml:"RemindUnit"`
	Repeatability       string                `json:"Repeatability" xml:"Repeatability"`
	InGroupId           int                   `json:"InGroupId" xml:"InGroupId"`
	ApplicationId       int64                 `json:"ApplicationId" xml:"ApplicationId"`
	ModifyTime          int64                 `json:"ModifyTime" xml:"ModifyTime"`
	FilePropertyContent string                `json:"FilePropertyContent" xml:"FilePropertyContent"`
	ExpTime             int64                 `json:"ExpTime" xml:"ExpTime"`
	TemplateId          int                   `json:"TemplateId" xml:"TemplateId"`
	TenantId            int64                 `json:"TenantId" xml:"TenantId"`
	Detail              string                `json:"Detail" xml:"Detail"`
	PageSize            int                   `json:"PageSize" xml:"PageSize"`
	DataSize            int64                 `json:"DataSize" xml:"DataSize"`
	Comment             string                `json:"Comment" xml:"Comment"`
	Founder             string                `json:"Founder" xml:"Founder"`
	EndCast             int64                 `json:"EndCast" xml:"EndCast"`
	LastAccessTime      int64                 `json:"LastAccessTime" xml:"LastAccessTime"`
	BaselineName        string                `json:"BaselineName" xml:"BaselineName"`
	Bizdate             int64                 `json:"Bizdate" xml:"Bizdate"`
	ApplicationCode     string                `json:"ApplicationCode" xml:"ApplicationCode"`
	Priority            int                   `json:"Priority" xml:"Priority"`
	TotalColumnCount    int64                 `json:"TotalColumnCount" xml:"TotalColumnCount"`
	Buffer              float64               `json:"Buffer" xml:"Buffer"`
	NextPrimaryKey      string                `json:"NextPrimaryKey" xml:"NextPrimaryKey"`
	CommitUser          string                `json:"CommitUser" xml:"CommitUser"`
	Owner               string                `json:"Owner" xml:"Owner"`
	PageNum             int                   `json:"PageNum" xml:"PageNum"`
	EntityId            int64                 `json:"EntityId" xml:"EntityId"`
	BeginWaitTimeTime   int64                 `json:"BeginWaitTimeTime" xml:"BeginWaitTimeTime"`
	OwnerId             string                `json:"OwnerId" xml:"OwnerId"`
	UseFlag             bool                  `json:"UseFlag" xml:"UseFlag"`
	CheckerName         string                `json:"CheckerName" xml:"CheckerName"`
	RemindType          string                `json:"RemindType" xml:"RemindType"`
	UseType             string                `json:"UseType" xml:"UseType"`
	MethodId            int                   `json:"MethodId" xml:"MethodId"`
	CriticalThreshold   string                `json:"CriticalThreshold" xml:"CriticalThreshold"`
	ApplicationKey      string                `json:"ApplicationKey" xml:"ApplicationKey"`
	DagType             string                `json:"DagType" xml:"DagType"`
	DndStart            string                `json:"DndStart" xml:"DndStart"`
	BeginRunningTime    int64                 `json:"BeginRunningTime" xml:"BeginRunningTime"`
	FixCheck            bool                  `json:"FixCheck" xml:"FixCheck"`
	FileVersion         int                   `json:"FileVersion" xml:"FileVersion"`
	Operator            string                `json:"Operator" xml:"Operator"`
	HasNext             bool                  `json:"HasNext" xml:"HasNext"`
	SlaMinu             int                   `json:"SlaMinu" xml:"SlaMinu"`
	NodeName            string                `json:"NodeName" xml:"NodeName"`
	ProjectNameCn       string                `json:"ProjectNameCn" xml:"ProjectNameCn"`
	FileContent         string                `json:"FileContent" xml:"FileContent"`
	LifeCycle           int                   `json:"LifeCycle" xml:"LifeCycle"`
	HourExpDetail       string                `json:"HourExpDetail" xml:"HourExpDetail"`
	Checker             int                   `json:"Checker" xml:"Checker"`
	Version             int64                 `json:"Version" xml:"Version"`
	TopicId             int64                 `json:"TopicId" xml:"TopicId"`
	ApplicationName     string                `json:"ApplicationName" xml:"ApplicationName"`
	ProgramType         string                `json:"ProgramType" xml:"ProgramType"`
	OwnerName           string                `json:"OwnerName" xml:"OwnerName"`
	BaselineId          int64                 `json:"BaselineId" xml:"BaselineId"`
	IsCurrentProd       bool                  `json:"IsCurrentProd" xml:"IsCurrentProd"`
	FinishStatus        string                `json:"FinishStatus" xml:"FinishStatus"`
	AppGuid             string                `json:"AppGuid" xml:"AppGuid"`
	WarningThreshold    string                `json:"WarningThreshold" xml:"WarningThreshold"`
	NodeId              int64                 `json:"NodeId" xml:"NodeId"`
	RemindName          string                `json:"RemindName" xml:"RemindName"`
	FolderPath          string                `json:"FolderPath" xml:"FolderPath"`
	NodeContent         string                `json:"NodeContent" xml:"NodeContent"`
	ChangeType          string                `json:"ChangeType" xml:"ChangeType"`
	ResGroupName        string                `json:"ResGroupName" xml:"ResGroupName"`
	ParamValues         string                `json:"ParamValues" xml:"ParamValues"`
	MethodName          string                `json:"MethodName" xml:"MethodName"`
	TemplateName        string                `json:"TemplateName" xml:"TemplateName"`
	MaxAlertTimes       int                   `json:"MaxAlertTimes" xml:"MaxAlertTimes"`
	RuleName            string                `json:"RuleName" xml:"RuleName"`
	BeginWaitResTime    int64                 `json:"BeginWaitResTime" xml:"BeginWaitResTime"`
	CreateTime          int64                 `json:"CreateTime" xml:"CreateTime"`
	IsVisible           int                   `json:"IsVisible" xml:"IsVisible"`
	FolderId            string                `json:"FolderId" xml:"FolderId"`
	BlockType           int                   `json:"BlockType" xml:"BlockType"`
	Endpoint            string                `json:"Endpoint" xml:"Endpoint"`
	SlaTime             int64                 `json:"SlaTime" xml:"SlaTime"`
	PageNumber          int                   `json:"PageNumber" xml:"PageNumber"`
	RemindId            int64                 `json:"RemindId" xml:"RemindId"`
	TableGuid           string                `json:"TableGuid" xml:"TableGuid"`
	Description         string                `json:"Description" xml:"Description"`
	ExpectValue         string                `json:"ExpectValue" xml:"ExpectValue"`
	ExpMinu             int                   `json:"ExpMinu" xml:"ExpMinu"`
	OnDuty              string                `json:"OnDuty" xml:"OnDuty"`
	DndEnd              string                `json:"DndEnd" xml:"DndEnd"`
	DagId               int64                 `json:"DagId" xml:"DagId"`
	CronExpress         string                `json:"CronExpress" xml:"CronExpress"`
	AlertUnit           string                `json:"AlertUnit" xml:"AlertUnit"`
	LastDdlTime         int64                 `json:"LastDdlTime" xml:"LastDdlTime"`
	ModifiedTime        int64                 `json:"ModifiedTime" xml:"ModifiedTime"`
	Useflag             bool                  `json:"Useflag" xml:"Useflag"`
	AlertInterval       int                   `json:"AlertInterval" xml:"AlertInterval"`
	CommitTime          int64                 `json:"CommitTime" xml:"CommitTime"`
	ProjectId           int64                 `json:"ProjectId" xml:"ProjectId"`
	SlaHour             int                   `json:"SlaHour" xml:"SlaHour"`
	AlertTargets        []string              `json:"AlertTargets" xml:"AlertTargets"`
	TableGuidList       []string              `json:"TableGuidList" xml:"TableGuidList"`
	AlertMethods        []string              `json:"AlertMethods" xml:"AlertMethods"`
	Deployment          Deployment            `json:"Deployment" xml:"Deployment"`
	NodeConfiguration   NodeConfiguration     `json:"NodeConfiguration" xml:"NodeConfiguration"`
	File                File                  `json:"File" xml:"File"`
	BlockInstance       BlockInstance         `json:"BlockInstance" xml:"BlockInstance"`
	LastInstance        LastInstance          `json:"LastInstance" xml:"LastInstance"`
	TableEntityList     []TableEntityListItem `json:"TableEntityList" xml:"TableEntityList"`
	Influences          []InfluencesItem      `json:"Influences" xml:"Influences"`
	Baselines           []BaselinesItem       `json:"Baselines" xml:"Baselines"`
	DataEntityList      []DataEntityListItem  `json:"DataEntityList" xml:"DataEntityList"`
	Rules               []RulesItem           `json:"Rules" xml:"Rules"`
	ColumnList          []ColumnListItem      `json:"ColumnList" xml:"ColumnList"`
	Robots              []RobotsItem          `json:"Robots" xml:"Robots"`
	BizProcesses        []BizProcessesItem    `json:"BizProcesses" xml:"BizProcesses"`
	Projects            []ProjectsItem        `json:"Projects" xml:"Projects"`
	Nodes               []NodesItem           `json:"Nodes" xml:"Nodes"`
}
