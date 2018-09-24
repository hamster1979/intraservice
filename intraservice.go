package intraservice

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// Services

// TaskService provides methods for creating and reading tasks.
type TaskService struct {
	sling *sling.Sling
}

// NewTaskService returns a new TaskService.
func NewTaskService(httpClient *http.Client, baseURL string, baseAUTH string) *TaskService {
	return &TaskService{
		sling: sling.New().Client(httpClient).Base(baseURL).Set("Authorization", "Basic "+baseAUTH),
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
	CompletionStatus             int    `json:"CompletionStatus,omitempty"`
	WorkflowID                   int    `json:"WorkflowId,omitempty"`
	IsClient                     bool   `json:"IsClient,omitempty"`
	FieryFields                  string `json:"FieryFields,omitempty"`
	IsConcurrence                bool   `json:"IsConcurrence,omitempty"`
	IsCoordinatedByCurrentUser   bool   `json:"IsCoordinatedByCurrentUser,omitempty"`
	CreatorADGuid                string `json:"CreatorADGuid,omitempty"`
	CreatorComments              string `json:"CreatorComments,omitempty"`
	CreatorCompanyID             int    `json:"CreatorCompanyId,omitempty"`
	CreatorCompanyName           string `json:"CreatorCompanyName,omitempty"`
	CreatorCompanyPath           string `json:"CreatorCompanyPath,omitempty"`
	CreatorCurrentLanguage       string `json:"CreatorCurrentLanguage,omitempty"`
	CreatorEmail                 string `json:"CreatorEmail,omitempty"`
	CreatorIsArchive             bool   `json:"CreatorIsArchive,omitempty"`
	CreatorLogin                 string `json:"CreatorLogin,omitempty"`
	CreatorMobilePhone           string `json:"CreatorMobilePhone,omitempty"`
	CreatorPhone                 string `json:"CreatorPhone,omitempty"`
	CreatorPosition              string `json:"CreatorPosition,omitempty"`
	CreatorRoleID                int    `json:"CreatorRoleId,omitempty"`
	CreatorTimeZone              string `json:"CreatorTimeZone,omitempty"`
	PriorityDescription          string `json:"PriorityDescription,omitempty"`
	PriorityImage16Url           string `json:"PriorityImage16Url,omitempty"`
	PriorityName                 string `json:"PriorityName,omitempty"`
	ServiceCode                  string `json:"ServiceCode,omitempty"`
	ServiceDescription           string `json:"ServiceDescription,omitempty"`
	ServiceIsArchive             bool   `json:"ServiceIsArchive,omitempty"`
	ServiceIsPublic              bool   `json:"ServiceIsPublic,omitempty"`
	ServiceName                  string `json:"ServiceName,omitempty"`
	ServiceParentID              int    `json:"ServiceParentId,omitempty"`
	ServicePath                  string `json:"ServicePath,omitempty"`
	StatusName                   string `json:"StatusName,omitempty"`
	StatusDescription            string `json:"StatusDescription,omitempty"`
	StatusImage16Url             string `json:"StatusImage16Url,omitempty"`
	StatusImage24Url             string `json:"StatusImage24Url,omitempty"`
	StatusIsCommentRequired      bool   `json:"StatusIsCommentRequired,omitempty"`
	StatusIsFixed                bool   `json:"StatusIsFixed,omitempty"`
	StatusIsFinal                bool   `json:"StatusIsFinal,omitempty"`
}

// List returns the tasks, selected by parameters
func (service *TaskService) List(params *TaskListParams) ([]Task, *http.Response, error) {
	tasks := new([]Task)
	apiError := new(APIError)
	resp, err := service.sling.New().Path("issues").QueryStruct(params).Receive(tasks, apiError)
	if err == nil {
		err = apiError
	}
	return *tasks, resp, err
}

// Client to wrap services

// Client is a tiny IntraService client
type Client struct {
	TaskService *TaskService
	// other service endpoints...
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, baseURL string, baseAUTH string) *Client {
	return &Client{
		TaskService: NewTaskService(httpClient, baseURL, baseAUTH),
	}
}
