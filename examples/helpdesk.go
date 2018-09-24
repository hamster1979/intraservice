package helpdesk

import (
	"fmt"

	"github.com/hamster1979/intraservice"
)

func main() {
	client := intraservice.NewClient(nil, "http://helpdesk.test.com:8888", "aXZhbm92Om15aXZhbm92MjM=")

	///api/task?fields=Id,Name,StatusId
	params := &intraservice.TaskListParams{Fields: "Id,Name,StatusId"}

	tasks, _, _ := client.TaskService.List(params)
	fmt.Printf("Task list:\n%v\n", tasks)

}
