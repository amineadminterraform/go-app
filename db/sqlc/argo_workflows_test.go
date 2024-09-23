package db

import (
	"context"
	"testing"
	"time"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/stretchr/testify/require"
)

//A function to create a RandomArgoworkflow
func createRandomArgoworkflow(t *testing.T) ArgoWorkflow{
	args := CreateArgoWorkflowParams{
		Name: utils.RandomName(),
		Path: utils.RandomName(),
		Description: utils.RandomDescription(),
	}
	argoworkflow , err := testQueries.CreateArgoWorkflow(context.Background(),args)
	require.NoError(t,err)
	require.NotEmpty(t,argoworkflow) 
	require.Equal(t,argoworkflow.Name,args.Name)
	require.Equal(t,argoworkflow.Description,args.Description)
	require.Equal(t,argoworkflow.Path,args.Path)
	require.NotZero(t,argoworkflow.ID)
	require.NotZero(t,argoworkflow.CreatedAt)

	return argoworkflow
}
func TestCreateArgoworkflow(t * testing.T) {
	createRandomArgoworkflow(t)
}
//	Test get an  argoworkflow
func TestGetArgoworkflow(t *testing.T) {
	workflow1 := createRandomArgoworkflow(t)
	workflow2, err := testQueries.GetArgoWorkflow(context.Background(), workflow1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, workflow2)
	require.Equal(t, workflow1.ID, workflow2.ID)
	require.Equal(t, workflow1.Description, workflow2.Description)
	require.Equal(t, workflow1.Path, workflow2.Path)
	require.Equal(t, workflow1.Name, workflow2.Name)
	require.Equal(t, workflow1.CreatedAt,workflow2.CreatedAt)
	require.WithinDuration(t, workflow1.CreatedAt.Time,workflow2.CreatedAt.Time, time.Second)
}

//	Test update argoworkflows
func TestUpdateArgoworkflow(t *testing.T){
	createdArgoWorkflow := createRandomArgoworkflow(t)

	args := UpdateArgoWorkflowParams{
		ID: createdArgoWorkflow.ID ,
		Name: utils.RandomName() ,
		Description: utils.RandomDescription(),
		Path: utils.RandomName(),
	}
	err := testQueries.UpdateArgoWorkflow(context.Background(),args)
	require.NoError(t,err)
	updatedArgoWorkflow, err := testQueries.GetArgoWorkflow(context.Background(),args.ID)
	require.NoError(t,err)
	require.NotEmpty(t,updatedArgoWorkflow)
	require.Equal(t,args.ID,updatedArgoWorkflow.ID)
	require.Equal(t,args.Description,updatedArgoWorkflow.Description)
	require.Equal(t,args.Name,updatedArgoWorkflow.Name)
	require.Equal(t,args.Path,updatedArgoWorkflow.Path)
}

//	Test delete argoworkflows
func TestDeleteArgoworkflow(t *testing.T){
	newargoworkflow := createRandomArgoworkflow(t)
	err := testQueries.DeleteArgoWorkflow(context.Background(),newargoworkflow.ID)
	require.NoError(t,err)
	testedargoworkflow , err := testQueries.GetArgoWorkflow(context.Background(),newargoworkflow.ID)
	require.Error(t,err)
	require.Empty(t,testedargoworkflow)
}

//	Test list argoworkflows	
func TestListArgoworkflows(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomArgoworkflow(t)
	}
	args := ListArgoWorkflowsParams{
		Limit:  5,
		Offset: 5,
	}
	argoworkflows, err := testQueries.ListArgoWorkflows(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, argoworkflows, 5)

	for _, argoworkflow := range argoworkflows{
		require.NotEmpty(t, argoworkflow)
	}
}
