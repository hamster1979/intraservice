package intraservice

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/tidwall/gjson"
)

// Services

// TaskService provides methods for creating and reading tasks.
type TaskService struct {
	httpClient *http.Client
	Sling      *sling.Sling
}

// NewTaskService returns a new TaskService.
func NewTaskService(httpClient *http.Client, baseURL string) *TaskService {
	return &TaskService{
		httpClient: httpClient,
		Sling:      sling.New().Client(httpClient).Base(baseURL),
	}
}

// APIError represents a IntraService API error response
type APIError struct {
	Message       string `json:"Message"`
	MessageDetail string `json:"MessageDetail"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("intraservice: %v %v", e.Message, e.MessageDetail)
}

//CustomFields - custom Fields type for parameters
type CustomFields struct {
	Field1  string `url:"Field1,omitempty"`
	Field2  string `url:"Field2,omitempty"`
	Field3  string `url:"Field3,omitempty"`
	Field4  string `url:"Field4,omitempty"`
	Field5  string `url:"Field5,omitempty"`
	Field6  string `url:"Field6,omitempty"`
	Field7  string `url:"Field7,omitempty"`
	Field8  string `url:"Field8,omitempty"`
	Field9  string `url:"Field9,omitempty"`
	Field10 string `url:"Field10,omitempty"`
	Field11 string `url:"Field11,omitempty"`
	Field12 string `url:"Field12,omitempty"`
	Field13 string `url:"Field13,omitempty"`
	Field14 string `url:"Field14,omitempty"`
	Field15 string `url:"Field15,omitempty"`
	Field16 string `url:"Field16,omitempty"`
	Field17 string `url:"Field17,omitempty"`
	Field18 string `url:"Field18,omitempty"`
	Field19 string `url:"Field19,omitempty"`
	Field20 string `url:"Field20,omitempty"`
	Field21 string `url:"Field21,omitempty"`
	Field22 string `url:"Field22,omitempty"`
	Field23 string `url:"Field23,omitempty"`
	Field24 string `url:"Field24,omitempty"`
	Field25 string `url:"Field25,omitempty"`
	Field26 string `url:"Field26,omitempty"`
	Field27 string `url:"Field27,omitempty"`
	Field28 string `url:"Field28,omitempty"`
	Field29 string `url:"Field29,omitempty"`
	Field30 string `url:"Field30,omitempty"`
	Field31 string `url:"Field31,omitempty"`
	Field32 string `url:"Field32,omitempty"`
	Field33 string `url:"Field33,omitempty"`
	Field34 string `url:"Field34,omitempty"`
	Field35 string `url:"Field35,omitempty"`
	Field36 string `url:"Field36,omitempty"`
	Field37 string `url:"Field37,omitempty"`
	Field38 string `url:"Field38,omitempty"`
	Field39 string `url:"Field39,omitempty"`
	Field40 string `url:"Field40,omitempty"`
	Field41 string `url:"Field41,omitempty"`
	Field42 string `url:"Field42,omitempty"`
	Field43 string `url:"Field43,omitempty"`
	Field44 string `url:"Field44,omitempty"`
	Field45 string `url:"Field45,omitempty"`
	Field46 string `url:"Field46,omitempty"`
	Field47 string `url:"Field47,omitempty"`
	Field48 string `url:"Field48,omitempty"`
	Field49 string `url:"Field49,omitempty"`
	Field50 string `url:"Field50,omitempty"`
	Field51 string `url:"Field51,omitempty"`
	Field52 string `url:"Field52,omitempty"`
	Field53 string `url:"Field53,omitempty"`
	Field54 string `url:"Field54,omitempty"`
	Field55 string `url:"Field55,omitempty"`
	Field56 string `url:"Field56,omitempty"`
	Field57 string `url:"Field57,omitempty"`
	Field58 string `url:"Field58,omitempty"`
	Field59 string `url:"Field59,omitempty"`
	Field60 string `url:"Field60,omitempty"`
	Field61 string `url:"Field61,omitempty"`
	Field62 string `url:"Field62,omitempty"`
	Field63 string `url:"Field63,omitempty"`
	Field64 string `url:"Field64,omitempty"`
	Field65 string `url:"Field65,omitempty"`
	Field66 string `url:"Field66,omitempty"`
	Field67 string `url:"Field67,omitempty"`
	Field68 string `url:"Field68,omitempty"`
	Field69 string `url:"Field69,omitempty"`
	Field70 string `url:"Field70,omitempty"`
	Field71 string `url:"Field71,omitempty"`
	Field72 string `url:"Field72,omitempty"`
	Field73 string `url:"Field73,omitempty"`
	Field74 string `url:"Field74,omitempty"`
	Field75 string `url:"Field75,omitempty"`
	Field76 string `url:"Field76,omitempty"`
	Field77 string `url:"Field77,omitempty"`
	Field78 string `url:"Field78,omitempty"`
	Field79 string `url:"Field79,omitempty"`
	Field80 string `url:"Field80,omitempty"`
	Field81 string `url:"Field81,omitempty"`
	Field82 string `url:"Field82,omitempty"`
	Field83 string `url:"Field83,omitempty"`
	Field84 string `url:"Field84,omitempty"`
	Field85 string `url:"Field85,omitempty"`
	Field86 string `url:"Field86,omitempty"`
	Field87 string `url:"Field87,omitempty"`
	Field88 string `url:"Field88,omitempty"`
	Field89 string `url:"Field89,omitempty"`
	Field90 string `url:"Field90,omitempty"`
	Field91 string `url:"Field91,omitempty"`
	Field92 string `url:"Field92,omitempty"`
	Field93 string `url:"Field93,omitempty"`
	Field94 string `url:"Field94,omitempty"`
	Field95 string `url:"Field95,omitempty"`
	Field96 string `url:"Field96,omitempty"`
	Field97 string `url:"Field97,omitempty"`
	Field98 string `url:"Field98,omitempty"`
	Field99 string `url:"Field99,omitempty"`
}

//FilterFields - are the parameters for the fields filtering
type FilterFields struct {
	ChangedLessThan      string `url:"ChangedLessThan,omitempty"`
	ChangedMoreThan      string `url:"ChangedMoreThan,omitempty"`
	Changed              string `url:"Changed,omitempty"`
	CreatedLessThan      string `url:"CreatedLessThan,omitempty"`
	CreatedMoreThan      string `url:"CreatedMoreThan,omitempty"`
	Created              string `url:"Created,omitempty"`
	ClosedLessThan       string `url:"ClosedLessThan,omitempty"`
	ClosedMoreThan       string `url:"ClosedMoreThan,omitempty"`
	Closed               string `url:"Closed,omitempty"`
	DeadlineLessThan     string `url:"DeadlineLessThan,omitempty"`
	DeadlineMoreThan     string `url:"DeadlineMoreThan,omitempty"`
	Deadline             string `url:"Deadline,omitempty"`
	ReactionDateLessThan string `url:"ReactionDateLessThan,omitempty"`
	ReactionDateMoreThan string `url:"ReactionDateMoreThan,omitempty"`
	ReactionDate         string `url:"ReactionDate,omitempty"`
	ResolutionOverdue    string `url:"ResolutionOverdue,omitempty"`
	ReactionOverdue      string `url:"ReactionOverdue,omitempty"`
	TypeIds              string `url:"TypeIds,omitempty"`
	StatusIds            string `url:"StatusIds,omitempty"`
	PriorityIds          string `url:"PriorityIds,omitempty"`
	ServiceIds           string `url:"ServiceIds,omitempty"`
	EditorIds            string `url:"EditorIds,omitempty"`
	CreatorIds           string `url:"CreatorIds,omitempty"`
	ExecutorIds          string `url:"ExecutorIds,omitempty"`
	ExecutorGroupIds     string `url:"ExecutorGroupIds,omitempty"`
	ObserverIds          string `url:"ObserverIds,omitempty"`
	CategoryIds          string `url:"CategoryIds,omitempty"`
	AssetIds             string `url:"AssetIds,omitempty"`
}

//TaskListParams - are the parameters for TaskService.List
type TaskListParams struct {
	Serviceid string `url:"serviceid,omitempty"`
	Fields    string `url:"fields,omitempty"`
	Search    string `url:"search,omitempty"`
	Archive   string `url:"archive,omitempty"`
	Inactive  string `url:"inactive,omitempty"`
	Filterid  string `url:"filterid,omitempty"`
	Sort      string `url:"sort,omitempty"`
	FilterFields
	Pagesize string `url:"pagesize,omitempty"`
	Page     string `url:"page,omitempty"`
	Include  string `url:"include,omitempty"`
}

//Task - is a intraservice task
type Task struct {
	ID                           int    `json:"Id,omitempty"`
	PriorityID                   int    `json:"PriorityId,omitempty"`
	StatusID                     int    `json:"StatusId,omitempty"`
	ServiceID                    int    `json:"ServiceId,omitempty"`
	ServiceStageID               int    `json:"ServiceStageId,omitempty"`
	ServiceStage                 string `json:"ServiceStage,omitempty"`
	ParentID                     int    `json:"ParentId,omitempty"`
	Name                         string `json:"Name,omitempty"`
	Description                  string `json:"Description,omitempty"`
	Deadline                     string `json:"Deadline,omitempty"`
	Created                      string `json:"Created,omitempty"`
	Changed                      string `json:"Changed,omitempty"`
	CreatorID                    int    `json:"CreatorId,omitempty"`
	Creator                      string `json:"Creator,omitempty"`
	EditorID                     int    `json:"EditorId,omitempty"`
	ExecutorIds                  string `json:"ExecutorIds,omitempty"`
	Executors                    string `json:"Executors,omitempty"`
	Files                        string `json:"Files,omitempty"`
	FileNames                    string `json:"FileNames,omitempty"`
	FileIds                      string `json:"FileIds,omitempty"`
	CategoryIds                  string `json:"CategoryIds,omitempty"`
	Categories                   string `json:"Categories,omitempty"`
	Hours                        int    `json:"Hours,omitempty"`
	Price                        int    `json:"Price,omitempty"`
	Assets                       string `json:"Assets,omitempty"`
	AssetIds                     string `json:"AssetIds,omitempty"`
	ReactionDate                 string `json:"ReactionDate,omitempty"`
	ReactionOverdue              bool   `json:"ReactionOverdue,omitempty"`
	Closed                       string `json:"Closed,omitempty"`
	TypeID                       int    `json:"TypeId,omitempty"`
	Type                         string `json:"Type,omitempty"`
	EvaluationID                 int    `json:"EvaluationId,omitempty"`
	Evaluation                   string `json:"Evaluation,omitempty"`
	IsMassIncident               bool   `json:"IsMassIncident,omitempty"`
	CrеatorIP                    string `json:"CrеatorIP,omitempty"`
	ExecutorGroupID              int    `json:"ExecutorGroupId,omitempty"`
	ExecutorGroup                string `json:"ExecutorGroup,omitempty"`
	TaskRepeatRuleID             int    `json:"TaskRepeatRuleId,omitempty"`
	ResolutionOverdue            bool   `json:"ResolutionOverdue,omitempty"`
	Coordinators                 string `json:"Coordinators,omitempty"`
	IsCoordinatedForCoordinators string `json:"IsCoordinatedForCoordinators,omitempty"`
	CoordinatorIds               string `json:"CoordinatorIds,omitempty"`
	Observers                    string `json:"Observers,omitempty"`
	ObserverIds                  string `json:"ObserverIds,omitempty"`
	Data                         string `json:"Data,omitempty"`
	CustomFields
	CompletionStatus           int    `json:"CompletionStatus,omitempty"`
	WorkflowID                 int    `json:"WorkflowId,omitempty"`
	IsClient                   bool   `json:"IsClient,omitempty"`
	FieryFields                string `json:"FieryFields,omitempty"`
	IsConcurrence              bool   `json:"IsConcurrence,omitempty"`
	IsCoordinatedByCurrentUser bool   `json:"IsCoordinatedByCurrentUser,omitempty"`
	CreatorADGuid              string `json:"CreatorADGuid,omitempty"`
	CreatorComments            string `json:"CreatorComments,omitempty"`
	CreatorCompanyID           int    `json:"CreatorCompanyId,omitempty"`
	CreatorCompanyName         string `json:"CreatorCompanyName,omitempty"`
	CreatorCompanyPath         string `json:"CreatorCompanyPath,omitempty"`
	CreatorCurrentLanguage     string `json:"CreatorCurrentLanguage,omitempty"`
	CreatorEmail               string `json:"CreatorEmail,omitempty"`
	CreatorIsArchive           bool   `json:"CreatorIsArchive,omitempty"`
	CreatorLogin               string `json:"CreatorLogin,omitempty"`
	CreatorMobilePhone         string `json:"CreatorMobilePhone,omitempty"`
	CreatorPhone               string `json:"CreatorPhone,omitempty"`
	CreatorPosition            string `json:"CreatorPosition,omitempty"`
	CreatorRoleID              int    `json:"CreatorRoleId,omitempty"`
	CreatorTimeZone            string `json:"CreatorTimeZone,omitempty"`
	PriorityDescription        string `json:"PriorityDescription,omitempty"`
	PriorityImage16Url         string `json:"PriorityImage16Url,omitempty"`
	PriorityName               string `json:"PriorityName,omitempty"`
	ServiceCode                string `json:"ServiceCode,omitempty"`
	ServiceDescription         string `json:"ServiceDescription,omitempty"`
	ServiceIsArchive           bool   `json:"ServiceIsArchive,omitempty"`
	ServiceIsPublic            bool   `json:"ServiceIsPublic,omitempty"`
	ServiceName                string `json:"ServiceName,omitempty"`
	ServiceParentID            int    `json:"ServiceParentId,omitempty"`
	ServicePath                string `json:"ServicePath,omitempty"`
	StatusName                 string `json:"StatusName,omitempty"`
	StatusDescription          string `json:"StatusDescription,omitempty"`
	StatusImage16Url           string `json:"StatusImage16Url,omitempty"`
	StatusImage24Url           string `json:"StatusImage24Url,omitempty"`
	StatusIsCommentRequired    bool   `json:"StatusIsCommentRequired,omitempty"`
	StatusIsFixed              bool   `json:"StatusIsFixed,omitempty"`
	StatusIsFinal              bool   `json:"StatusIsFinal,omitempty"`
}

//TasksResponse - json respond of resource list request
type TasksResponse struct {
	Tasks []Task `json:"Tasks,omitempty"`
}

//TaskResponse - json respond of resource request
type TaskResponse struct {
	Task Task `json:"Tasks,omitempty"`
}

//ReceiveCustom - return JSON respond decoded to nested map
func (service *TaskService) ReceiveCustom(req *http.Request, failureV interface{}) (map[string]interface{}, *http.Response, error) {
	var data map[string]interface{}
	resp, err := service.httpClient.Do(req)
	if err != nil {
		return nil, resp, err
	}
	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer resp.Body.Close()

	// Don't try to decode on 204s
	if resp.StatusCode == 204 {
		return nil, resp, nil
	}

	// Decode from json
	if failureV != nil {
		err = nil
		if code := resp.StatusCode; 200 <= code && code <= 299 {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := string(bodyBytes)
			data, _ = gjson.Parse(bodyString).Value().(map[string]interface{})
		} else {
			return nil, resp, nil
		}
	}
	return data, resp, err
}

// List  - returns the tasks, selected by parameters
func (service *TaskService) List(params *TaskListParams) (map[string]interface{}, *http.Response, error) {
	apiError := new(APIError)
	req, err := service.Sling.New().Get("task").QueryStruct(params).Request()
	if err != nil {
		return nil, nil, err
	}
	data, resp, err := service.ReceiveCustom(req, apiError)
	if err == nil {
		err = apiError
	}
	return data, resp, err
}

// Get  - returns the current tasks, selected by parameters
func (service *TaskService) Get(ID int) (TaskResponse, *http.Response, error) {
	respond := new(TaskResponse)
	apiError := new(APIError)
	path := fmt.Sprintf("task/%d", ID)
	resp, err := service.Sling.New().Get(path).Receive(respond, apiError)
	if err == nil {
		err = apiError
	}
	return *respond, resp, err
}

//AttrParams - parameters for Get Changed Attribut command
type AttrParams struct {
	Field         int `url:"field"`
	Oldworkflowid int `url:"oldworkflowid"`
	Serviced      int `url:"serviced"`
	Priorityid    int `url:"priorityid"`
	Statusid      int `url:"statusid"`
	Taskid        int `url:"taskid, omitempty"`
	Creatorid     int `url:"creatorid, omitempty"`
	Tasktypeid    int `url:"tasktypeid, omitempty"`
	Assetids      int `url:"assetids, omitempty"`
	Categoryids   int `url:"categoryids, omitempty"`
}

//AttrResponse - Attr command response
type AttrResponse struct {
	Deadline     string   `json:"Deadline"`
	ReactionDate string   `json:"ReactionDate"`
	StatusID     int      `json:"StatusId"`
	Statuses     []string `json:"Statuses"`
	WorkflowID   int      `json:"WorkflowId"`
}

// Attr  - returns the attributes of task with changes
func (service *TaskService) Attr(params *AttrParams) (AttrResponse, *http.Response, error) {
	respond := new(AttrResponse)
	apiError := new(APIError)
	resp, err := service.Sling.New().Get("task").Receive(respond, apiError)
	if err == nil {
		err = apiError
	}
	return *respond, resp, err
}

//CreateTaskParams - parameters for Create Task command
type CreateTaskParams struct {
	Name             string `url:"Name"`
	ServiceID        int    `url:"ServiceId"`
	StatusID         int    `url:"StatusId"`
	Comment          string `url:"Comment"`
	Deadline         string `url:"Deadline"`
	Description      string `url:"Description"`
	ParentID         int    `url:"ParentId"`
	PriorityID       int    `url:"PriorityId"`
	TypeID           int    `url:"TypeId"`
	CreatorID        int    `url:"CreatorId"`
	ServiceStageID   int    `url:"ServiceStageId"`
	IsPrivateComment bool   `url:"IsPrivateComment"`
	IsMassIncident   bool   `url:"IsMassIncident"`
	CompletionStatus string `url:"CompletionStatus"`
	AssetIds         string `url:"AssetIds"`
	CategoryIds      string `url:"CategoryIdsstring"`
	ExecutorIds      string `url:"ExecutorIds"`
	CoordinatorIds   string `url:"CoordinatorIds"`
	ExecutorGroupID  int    `url:"ExecutorGroupId"`
	ObserverIds      string `url:"ObserverIds"`
	FileTokens       string `url:"FileTokens"`
	DeletedFiles     string `url:"DeletedFiles"`
	CustomFields
	*UserEmailParams
}

//UserEmailParams - parameters for Create Task command by Email identification user
type UserEmailParams struct {
	UserEmail           string `url:"UserEmail"`
	UserPassword        string `url:"UserPassword"`
	UserConfirmPassword string `url:"UserConfirmPassword"`
	UserComments        string `url:"UserComments,omitempty"`
	UserCompanyID       int    `url:"UserCompanyId"`
	UserCurrentLanguage string `url:"UserCurrentLanguage, omitempty"`
	UserLogin           string `url:"UserLogin, omitempty"`
	UserMobilePhone     string `url:"UserMobilePhone, omitempty"`
	UserName            string `url:"UserName, omitempty"`
	UserPhone           string `url:"UserPhone, omitempty"`
	UserPosition        string `url:"UserPosition, omitempty"`
	UserRoleID          int    `url:"UserRoleId"`
	UserTimeZone        string `url:"UserTimeZone, omitempty"`
}

// Create  - create task with parameters
func (service *TaskService) Create(params *CreateTaskParams) (TaskResponse, *http.Response, error) {
	respond := new(TaskResponse)
	apiError := new(APIError)
	resp, err := service.Sling.New().Post("task").Receive(respond, apiError)
	if err == nil {
		err = apiError
	}
	return *respond, resp, err
}

//EditTaskParams - parameters for Edit Task command
type EditTaskParams struct {
	CreateTaskParams
	EvaluationID       int    `url:"EvaluationId, omitempty"`
	ReactionDate       string `url:"ReactionDate, omitempty"`
	ReactionDateFact   string `url:"ReactionDateFact, omitempty"`
	ResolutionDateFact string `url:"ResolutionDateFact, omitempty"`
	Coordinate         bool   `url:"Coordinate, omitempty"`
}

// Edit  - edit task by id
func (service *TaskService) Edit(ID int, params *EditTaskParams) (TaskResponse, *http.Response, error) {
	respond := new(TaskResponse)
	apiError := new(APIError)
	path := fmt.Sprintf("task/%d", ID)
	resp, err := service.Sling.New().Put(path).Receive(respond, apiError)
	if err == nil {
		err = apiError
	}
	return *respond, resp, err
}

// Client to wrap services

// Client is a tiny IntraService client
type Client struct {
	TaskService *TaskService
	// other service endpoints...
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, baseURL string) *Client {
	return &Client{
		TaskService: NewTaskService(httpClient, baseURL+"/api/"),
	}
}
