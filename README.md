# intraservice
Golang package for working with Intraservice API

# Documentation
[IntraService API](https://intraservice.ru/upload/iblock/7bc/IntraService_API_v4_46_for_clients.pdf)

# Usage

```
	client := intraservice.NewClient(nil, "http://helpdesk.test.com:8888", "aXZhbm92Om15aXZhbm92MjM=")

	///api/task?fields=Id,Name,StatusId
	params := &intraservice.TaskListParams{Fields: "Id,Name,StatusId"}

	tasks, _, _ := client.TaskService.List(params)
	fmt.Printf("Task list:\n%v\n", tasks)
```
