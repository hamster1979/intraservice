package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/hamster1979/intraservice"
)

func main() {
	user := "tester"
	password := "watcom"

	client := intraservice.NewClient(nil, "http://sd-test.watcom.ru")

	///api/task?fields=Id,Name,StatusId
	params := &intraservice.TaskListParams{Fields: "Id,Name,StatusId"}
	client.TaskService.Sling.SetBasicAuth(user, password) //Basic auth

	tasks, resp, err := client.TaskService.List(params)
	//fmt.Printf("Task list:\n%+v\n", tasks)
	fmt.Printf("Task list:\n")
	spew.Dump(tasks)
	fmt.Printf("Respond:\n%+v", resp)
	//spew.Dump(resp)
	fmt.Printf("Error\n")
	spew.Dump(err)

}
