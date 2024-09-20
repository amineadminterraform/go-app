package db

import (
	"context"
	"testing"
	"time"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

func createRandomProject(t *testing.T) Project{
	arg := CreateProjectParams{
		Name: utils.RandomName(),
		GitPath: utils.RandomGit(),
		Description: utils.RandomDescription(),
	}
	project, err := testQueries.CreateProject(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,project)
	require.Equal(t, arg.Name, project.Name)
	require.Equal(t, arg.GitPath, project.GitPath)
	require.Equal(t, arg.Description, project.Description)
	require.NotZero(t, project.ID)
	require.NotZero(t, project.CreatedAt)
	//require.NotZero(t, created_at)
	return project
}
	

func TestCreateProject(t *testing.T){
	createRandomProject(t)
}
func TestGetProject(t *testing.T){
	project1 := createRandomProject(t)
	project2, err := testQueries.GetProject(context.Background(),project1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, project2)
	require.Equal(t, project1.ID,project2.ID)
	require.Equal(t, project1.Description,project2.Description)
	require.Equal(t, project1.GitPath,project2.GitPath)
	require.Equal(t, project1.Name,project2.Name)
    require.WithinDuration(t, project1.CreatedAt.Time, project2.CreatedAt.Time, time.Second)
}
func TestUpdateProject(t *testing.T){
	project1 := createRandomProject(t)
    arg := UpdateProjectParams{
		ID: project1.ID,
		Name: utils.RandomName(),
		GitPath: utils.RandomGit(),
		Description: utils.RandomDescription(),
	}
	err := testQueries.UpdateProject(context.Background(), arg)
	require.NoError(t, err)
	project2, err2 := testQueries.GetProject(context.Background(),project1.ID)
	
	require.NoError(t, err2)
	require.NotEmpty(t, project2)
	require.Equal(t, project1.ID,project2.ID)
	require.Equal(t, arg.GitPath,project2.GitPath)
	require.Equal(t, arg.Name,project2.Name)
	require.Equal(t, arg.Description,project2.Description)
    require.WithinDuration(t, project1.CreatedAt.Time, project2.CreatedAt.Time, time.Second)
}
func TestDeleteProject(t* testing.T){
	project := createRandomProject(t)
	err := testQueries.DeleteProject(context.Background(),project.ID)
	require.NoError(t, err)
	project2, err := testQueries.GetProject(context.Background(),project.ID) 
	require.Error(t, err)
	//require.EqualError(t, err, sql.ErrNoRows.Error())
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, project2)
}

func TestListProjects(t *testing.T) {
	for i:= 0; i< 10; i++ {
		createRandomProject(t)
	}
	args := ListProjectsParams{
		Limit: 5, 
        Offset: 5,
	}
	projects, err := testQueries.ListProjects(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, projects, 5)

	for _, project := range projects {
		require.NotEmpty(t, project)
	} 
}
