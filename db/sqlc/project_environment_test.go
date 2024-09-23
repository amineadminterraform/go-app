package db

import (
	"context"
	"testing"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/stretchr/testify/require"
)

//In order to test the creations of project environement we need to test it with a proper projectid because its a foreign constraint
// A test for the existence of the id will be performed later
func CreateRandomProjectEnvironment(t *testing.T) ProjectEnvironment {

	//We Need to create a random Project	
	project := createRandomProject(t)

	//Those are the args of the project env we are trying to create
	args := CreateProjectEnvironmentParams{
		GitBranch: utils.RandomName(),
		ProjectID:project.ID ,
		Description: utils.RandomDescription(),	
	}
	projectenv ,err := testQueries.CreateProjectEnvironment(context.Background(), args) 
	require.NoError(t,err)
	require.NotEmpty(t,projectenv)
	require.Equal(t,projectenv.ProjectID,args.ProjectID)
	require.Equal(t,projectenv.GitBranch,args.GitBranch)
	require.Equal(t,projectenv.Description,args.Description)
	require.NotZero(t,projectenv.ID)
	require.NotZero(t,projectenv.CreatedAt)
	require.NotZero(t,projectenv.UpdatedAt)
	return projectenv
}

//Testing the creation of a project environmenet		
func TestCreateProjectEnvironment(t *testing.T) {
	CreateRandomProjectEnvironment(t)
}

//Test the getter of a project environment
func TestGetProjectEnvironment(t *testing.T){
	createdProjectEnv  := CreateRandomProjectEnvironment(t)
	gottenProjectEnv,err := testQueries.GetProjectEnvironment(context.Background(), createdProjectEnv.ID)
	require.NoError(t,err)
	require.NotEmpty(t,gottenProjectEnv)
	require.NotZero(t,gottenProjectEnv.ID)
	require.NotZero(t,gottenProjectEnv.CreatedAt)
	require.NotZero(t,gottenProjectEnv.UpdatedAt)
	require.Equal(t,createdProjectEnv.ProjectID,gottenProjectEnv.ProjectID)
	require.Equal(t,createdProjectEnv.Description,gottenProjectEnv.Description)
	require.Equal(t,createdProjectEnv.GitBranch,gottenProjectEnv.GitBranch)
}

//Test the update of a project environment
func TestUpdateProjectEnvironment(t *testing.T){
	createdProjectEnv := CreateRandomProjectEnvironment(t)
	args := UpdateProjectEnvironmentParams{
		ID: createdProjectEnv.ID,
		ProjectID: createdProjectEnv.ProjectID,
		GitBranch: utils.RandomGitBranch(),
		Description: utils.RandomDescription(),
	}
	err:= testQueries.UpdateProjectEnvironment(context.Background(),args)
	require.NoError(t,err)
	updatedProjectEnv,err := testQueries.GetProjectEnvironment(context.Background(),args.ID)
	require.NoError(t,err)
	require.Equal(t,args.Description,updatedProjectEnv.Description)
	require.Equal(t,args.GitBranch,updatedProjectEnv.GitBranch)
	require.Equal(t,args.ProjectID,updatedProjectEnv.ProjectID)
	require.Equal(t,args.ProjectID,updatedProjectEnv.ProjectID)
	} 	

// Test delete projects 
	func TestDeleteProjectEnvironment(t *testing.T) {
		projectEnvironment := CreateRandomProjectEnvironment(t)
		err := testQueries.DeleteProjectEnvironment(context.Background(), projectEnvironment.ID)
		require.NoError(t, err)
		projectEnvironment2, err := testQueries.GetProjectEnvironment(context.Background(), projectEnvironment.ID)
		require.Error(t, err)
		//require.EqualError(t, err, sql.ErrNoRows.Error())
		require.EqualError(t, err, "no rows in result set")
		require.Empty(t, projectEnvironment2)
	}
//	Test list project environments	
	func TestListProjectEnvironments(t *testing.T) {
		for i := 0; i < 10; i++ {
			CreateRandomProjectEnvironment(t)
		}
		args := ListProjectEnvironmentsParams{
			Limit:  5,
			Offset: 5,
		}
		projectEnvs, err := testQueries.ListProjectEnvironments(context.Background(), args)
		require.NoError(t, err)
		require.Len(t, projectEnvs, 5)
	
		for _, projectEnvs := range projectEnvs {
			require.NotEmpty(t, projectEnvs)
		}
	}
