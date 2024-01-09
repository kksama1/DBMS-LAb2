package driver

import (
	"context"
	"projext/internal/model"
	"time"
)

func InsertTask() {
	newTasks := []interface{}{
		model.Task{
			RelatedProject: model.Project{
				Name:     "DBMS",
				Code:     1,
				Sp:       144,
				Deadline: time.Date(2024, 1, 10, 14, 00, 00, 00, time.Local),
			},
			Executor: model.Executor{
				Id:         1,
				FirstName:  "Dmitry",
				LastName:   "Golubyatnikov",
				FatherName: "Sergeevich",
				Position:   "Student",
			},
			IssueDate:    time.Date(2023, 9, 1, 00, 00, 00, 00, time.Local),
			Sp:           14,
			ExpDeadline:  time.Date(2024, 1, 10, 14, 00, 00, 00, time.Local),
			RealDeadline: time.Date(2024, 2, 6, 23, 59, 59, 59, time.Local),
		},

		model.Task{
			RelatedProject: model.Project{
				Name:     "Pugachev",
				Code:     2,
				Sp:       144,
				Deadline: time.Date(2024, 1, 15, 14, 00, 00, 00, time.Local),
			},
			Executor: model.Executor{
				Id:         0,
				FirstName:  "Dmitry",
				LastName:   "Golubyatnikov",
				FatherName: "Sergeevich",
				Position:   "Student",
			},
			IssueDate:    time.Date(2023, 9, 1, 00, 00, 00, 00, time.Local),
			Sp:           14,
			ExpDeadline:  time.Date(2024, 1, 10, 14, 00, 00, 00, time.Local),
			RealDeadline: time.Date(2024, 2, 6, 23, 59, 59, 59, time.Local),
		},

		model.Task{
			RelatedProject: model.Project{
				Name:     "MakeTasks",
				Code:     3,
				Sp:       10,
				Deadline: time.Date(2024, 1, 15, 20, 00, 00, 00, time.Local),
			},
			Executor: model.Executor{
				Id:         0,
				FirstName:  "NotDmitry",
				LastName:   "NotGolubyatnikov",
				FatherName: "NotSergeevich",
				Position:   "Intern",
			},
			IssueDate:    time.Date(2023, 12, 25, 19, 00, 00, 00, time.Local),
			Sp:           14,
			ExpDeadline:  time.Date(2024, 1, 10, 14, 00, 00, 00, time.Local),
			RealDeadline: time.Date(2024, 2, 6, 23, 59, 59, 59, time.Local),
		},
	}

	client, ctx := CreateConnection()
	defer closeConn(ctx, client)

	database := client.Database("ProjectManagment").Collection("Task")

	_, err := database.InsertMany(context.TODO(), newTasks)
	if err != nil {
		panic(err)
	}

}
