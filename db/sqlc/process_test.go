package db

import (
	"context"
	"testing"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/stretchr/testify/require"
)

//In order to test the creations of project environement we need to test it with a proper projectid because its a foreign constraint
// A test for the existence of the id will be performed later
func CreateRandomProcess(t *testing.T) Process {

	//We Need to create a random Project	
	RandomArgoworkflow := createRandomArgoworkflow(t)
	randomTemplate := createRandomTemplate(t)

	//Those are the args of the project env we are trying to create
	args := CreateProcessParams{
		ArgoID: RandomArgoworkflow.ID,
		TemplateID: randomTemplate.ID,
		Name: utils.RandomName(),
	}
	process ,err := testQueries.CreateProcess(context.Background(), args) 
	require.NoError(t,err)
	require.NotEmpty(t,process)
	require.Equal(t,process.ArgoID,args.ArgoID)
	require.Equal(t,process.TemplateID,args.TemplateID)
	require.Equal(t,process.Name,args.Name)
	require.NotZero(t,process.ID)
	require.NotZero(t,process.CreatedAt)
	return process
}

//Testing the creation of a project environmenet		
func TestCreateProcess(t *testing.T) {
	CreateRandomProcess(t)
}

//Test the getter of a project environment
func TestGetProcess(t *testing.T){
	createdprocess  := CreateRandomProcess(t)
	gottenprocess,err := testQueries.GetProcess(context.Background(), createdprocess.ID)
	require.NoError(t,err)
	require.NotEmpty(t,gottenprocess)
	require.NotZero(t,gottenprocess.ID)
	require.NotZero(t,gottenprocess.CreatedAt)
	require.NotZero(t,gottenprocess.UpdatedAt)
	require.Equal(t,createdprocess.ArgoID,gottenprocess.ArgoID)
	require.Equal(t,createdprocess.TemplateID,gottenprocess.TemplateID)
	require.Equal(t,createdprocess.Name,gottenprocess.Name)
}

//Test the update of a project environment
func TestUpdateProcess(t *testing.T){
	createdprocess := CreateRandomProcess(t)
	createargo := createRandomArgoworkflow(t)
	createtemplate := createRandomTemplate(t)
	args := UpdateProcessParams{
		ID: createdprocess.ID,
		ArgoID: createargo.ID,
		TemplateID: createtemplate.ID,
		Name: utils.RandomName(),
	}
	err:= testQueries.UpdateProcess(context.Background(),args)
	require.NoError(t,err)
	updatedprocess,err := testQueries.GetProcess(context.Background(),args.ID)
	require.NoError(t,err)
	require.Equal(t,args.ArgoID,updatedprocess.ArgoID)
	require.Equal(t,args.ID,updatedprocess.ID)
	require.Equal(t,args.Name,updatedprocess.Name)
	require.Equal(t,args.TemplateID,updatedprocess.TemplateID)
	} 	

// Test delete projects 
	func TestDeleteProcess(t *testing.T) {
		Process := CreateRandomProcess(t)
		err := testQueries.DeleteProcess(context.Background(), Process.ID)
		require.NoError(t, err)
		Process2, err := testQueries.GetProcess(context.Background(), Process.ID)
		require.Error(t, err)
		//require.EqualError(t, err, sql.ErrNoRows.Error())
		require.EqualError(t, err, "no rows in result set")
		require.Empty(t, Process2)
	}
//	Test list project environments	
	func TestListProcesss(t *testing.T) {
		for i := 0; i < 10; i++ {
			CreateRandomProcess(t)
		}
		args := ListProcesssParams{
			Limit:  5,
			Offset: 5,
		}
		processs, err := testQueries.ListProcesss(context.Background(), args)
		require.NoError(t, err)
		require.Len(t, processs, 5)
	
		for _, process := range processs {
			require.NotEmpty(t, process)
		}
	}
