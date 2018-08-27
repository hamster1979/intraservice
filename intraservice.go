package intraservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IS struct {
	Server string
	Port   string
	Auth   string
}

type Task struct {
	Priorities       interface{} `json:"Priorities"`
	Rights           interface{} `json:"Rights"`
	TaskTypeSettings interface{} `json:"TaskTypeSettings"`
	Services         interface{} `json:"Services"`
	Statuses         interface{} `json:"Statuses"`
	Task             TaskClass   `json:"Task"`
	TaskType         interface{} `json:"TaskType"`
	Users            interface{} `json:"Users"`
}

type TaskClass struct {
	// AssetIDS                     string      `json:"AssetIds"`
	// Assets                       string      `json:"Assets"`
	// Categories                   interface{} `json:"Categories"`
	// CategoryIDS                  interface{} `json:"CategoryIds"`
	// Changed                      string      `json:"Changed"`
	// Closed                       interface{} `json:"Closed"`
	// CompletionStatus             interface{} `json:"CompletionStatus"`
	// CoordinatorIDS               interface{} `json:"CoordinatorIds"`
	// Coordinators                 interface{} `json:"Coordinators"`
	// IsCoordinatedForCoordinators interface{} `json:"IsCoordinatedForCoordinators"`
	// Created                      string      `json:"Created"`
	Creator string `json:"Creator"`
	// CreatorADGUID                string      `json:"CreatorADGuid"`
	// CreatorComments              string      `json:"CreatorComments"`
	// CreatorCompanyID             int64       `json:"CreatorCompanyId"`
	// CreatorCompanyName           string      `json:"CreatorCompanyName"`
	// CreatorCompanyPath           string      `json:"CreatorCompanyPath"`
	// CreatorCurrentLanguage       interface{} `json:"CreatorCurrentLanguage"`
	CreatorEmail string `json:"CreatorEmail"`
	// CreatorIP                    interface{} `json:"CreatorIP"`
	// CreatorID                    int64       `json:"CreatorId"`
	// CreatorIsArchive             bool        `json:"CreatorIsArchive"`
	// CreatorLogin                 string      `json:"CreatorLogin"`
	// CreatorMobilePhone           interface{} `json:"CreatorMobilePhone"`
	// CreatorPhone                 string      `json:"CreatorPhone"`
	// CreatorPosition              string      `json:"CreatorPosition"`
	// CreatorRoleID                int64       `json:"CreatorRoleId"`
	// CreatorTimeZone              interface{} `json:"CreatorTimeZone"`
	Deadline    string `json:"Deadline"`
	Description string `json:"Description"`
	// EditorID                     int64       `json:"EditorId"`
	// Escalated                    interface{} `json:"Escalated"`
	// Evaluation                   interface{} `json:"Evaluation"`
	// EvaluationID                 interface{} `json:"EvaluationId"`
	// ExecutorGroup                interface{} `json:"ExecutorGroup"`
	// ExecutorGroupID              interface{} `json:"ExecutorGroupId"`
	// ExecutorIDS                  string      `json:"ExecutorIds"`
	// Executors                    string      `json:"Executors"`
	// FieryFields                  string      `json:"FieryFields"`
	Files     string `json:"Files"`
	FileNames string `json:"FileNames"`
	// FileIDS                      string      `json:"FileIds"`
	// Hours                        int64       `json:"Hours"`
	// ID                           int64       `json:"Id"`
	// IsClient                     bool        `json:"IsClient"`
	// IsMassIncident               bool        `json:"IsMassIncident"`
	Name string `json:"Name"`
	// ObserverIDS                  string      `json:"ObserverIds"`
	// Observers                    string      `json:"Observers"`
	// ParentID                     int64       `json:"ParentId"`
	// Price                        int64       `json:"Price"`
	// PriorityDescription          string      `json:"PriorityDescription"`
	// PriorityID                   int64       `json:"PriorityId"`
	// PriorityImage16URL           string      `json:"PriorityImage16Url"`
	// PriorityName                 string      `json:"PriorityName"`
	// ReactionDate                 string      `json:"ReactionDate"`
	// ReactionOverdue              bool        `json:"ReactionOverdue"`
	// ResolutionOverdue            bool        `json:"ResolutionOverdue"`
	// ServiceCode                  string      `json:"ServiceCode"`
	// ServiceDescription           string      `json:"ServiceDescription"`
	// ServiceID                    int64       `json:"ServiceId"`
	// ServiceIsArchive             bool        `json:"ServiceIsArchive"`
	// ServiceIsPublic              bool        `json:"ServiceIsPublic"`
	// ServiceName                  string      `json:"ServiceName"`
	// ServiceParentID              int64       `json:"ServiceParentId"`
	// ServicePath                  string      `json:"ServicePath"`
	// ServiceStage                 interface{} `json:"ServiceStage"`
	// ServiceStageID               interface{} `json:"ServiceStageId"`
	// StatusDescription            string      `json:"StatusDescription"`
	// StatusID                     int64       `json:"StatusId"`
	// StatusImage16URL             string      `json:"StatusImage16Url"`
	// StatusImage24URL             string      `json:"StatusImage24Url"`
	// StatusIsCommentRequired      bool        `json:"StatusIsCommentRequired"`
	// StatusIsFixed                bool        `json:"StatusIsFixed"`
	// StatusIsFinal                bool        `json:"StatusIsFinal"`
	// StatusName                   string      `json:"StatusName"`
	// TaskRepeatRuleID             interface{} `json:"TaskRepeatRuleId"`
	// Type                         string      `json:"Type"`
	// TypeID                       int64       `json:"TypeId"`
	// WorkflowID                   int64       `json:"WorkflowId"`
	// IsConcurrence                bool        `json:"IsConcurrence"`
	// IsCoordinatedByCurrentUser   bool        `json:"IsCoordinatedByCurrentUser"`
}

//TaskLifeTime
type TaskLife struct {
	TaskLifetimes []TaskLifetime `json:"TaskLifetimes"`
	Priorities    []interface{}  `json:"Priorities"`
	Services      []interface{}  `json:"Services"`
	Statuses      []interface{}  `json:"Statuses"`
	Paginator     Paginator      `json:"Paginator"`
}

type Paginator struct {
	Count       int64       `json:"Count"`
	Page        int64       `json:"Page"`
	PageCount   int64       `json:"PageCount"`
	PageSize    int64       `json:"PageSize"`
	CountOnPage int64       `json:"CountOnPage"`
	HasNextPage interface{} `json:"HasNextPage"`
}

type TaskLifetime struct {
	Comments     *string `json:"Comments,omitempty"`
	IsPublic     *bool   `json:"IsPublic,omitempty"`
	Date         string  `json:"Date"`
	EditorID     *int64  `json:"EditorId"`
	Editor       string  `json:"Editor"`
	StatusID     *int64  `json:"StatusId,omitempty"`
	Participants *string `json:"Participants,omitempty"`
	Executors    *string `json:"Executors,omitempty"`
	Files        *string `json:"Files,omitempty"`
}

type RespErr struct {
	Message string `json:"Message"`
}

//GetTask - function to get task from Intraservice by RestAPI
func (is IS) GetTask(idTask string) (TaskClass, error) {
	var er error
	task := TaskClass{
		Creator:      "",
		CreatorEmail: "",
		Deadline:     "",
		Description:  "",
		Files:        "",
		FileNames:    "",
		Name:         "",
	}
	uri := "https://" + is.Server + ":" + is.Port + "/api/task/" + idTask
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	// добавляем заголовки
	req.Header.Add("Accept", "application/json")       // добавляем заголовок Accept
	req.Header.Add("Content-Type", "application/json") // добавляем заголовок Content-Type
	req.Header.Add("Authorization", is.Auth)           // добавляем заголовок Authorization

	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return task, err
	}
	defer resp.Body.Close()
	fmt.Printf("GetTask request header = %v, uri=%s\n", req.Header, req.URL.RawQuery)
	if resp.StatusCode != http.StatusOK {
		var result RespErr //Response code>200, then error hendler
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			errord := errors.New(fmt.Sprintf("Response code=%d, error: %s", resp.StatusCode, err.Error()))
			return task, errord

		}
		errord := errors.New(fmt.Sprintf("Response code=%d", resp.StatusCode))
		return task, errord
	}

	//All is OK< web response code is 200 OK!
	var result Task
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		errord := errors.New(fmt.Sprintf("Decode response, error: %s", err.Error()))
		return task, errord
	}

	return result.Task, er
}

//GetTaskLife - func to get task live history, order by time.
func (is IS) GetTaskLife(idTask string) ([]TaskLifetime, error) {
	var er error
	var task []TaskLifetime
	uri := "https://" + is.Server + ":" + is.Port + "/api/tasklifetime?taskid=" + idTask + "&lastcommentsontop=false"
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	// добавляем заголовки
	req.Header.Add("Accept", "application/json")       // добавляем заголовок Accept
	req.Header.Add("Content-Type", "application/json") // добавляем заголовок Content-Type
	req.Header.Add("Authorization", is.Auth)           // добавляем заголовок Authorization

	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return task, err
	}
	defer resp.Body.Close()
	// fmt.Printf("Search request header = %v, uri=%s\n", req.Header, req.URL.RawQuery)
	if resp.StatusCode != http.StatusOK {
		var result RespErr //Response code>200, then error hendler
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			errord := errors.New(fmt.Sprintf("Response code=%d, error: %s", resp.StatusCode, err.Error()))
			return task, errord

		}
		errord := errors.New(fmt.Sprintf("Response code=%d", resp.StatusCode))
		return task, errord
	}

	//All is OK< web response code is 200 OK!
	var result TaskLife
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		errord := errors.New(fmt.Sprintf("Decode response, error: %s", err.Error()))
		return task, errord
	}

	return result.TaskLifetimes, er
}
