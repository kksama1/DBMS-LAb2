package driver

import (
	"context"
	"projext/internal/model"
)

func CreateProject(projectInfo model.Project) {

	client, ctx := CreateConnection()
	defer closeConn(ctx, client)

	database := client.Database("ProjectManagment").Collection("Project")

	_, err := database.InsertOne(context.TODO(), projectInfo)
	if err != nil {
		panic(err)
	}

}
