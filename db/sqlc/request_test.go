package db

import (
	"context"
	"testing"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/stretchr/testify/require"
)

//In order to test the creations of project environement we need to test it with a proper projectid because its a foreign constraint
// A test for the existence of the id will be performed later
func CreateRandomRequest(t *testing.T) Request {

	//We Need to create a random Project	
	RandomEnvLayer := CreateRandomEnvLayer(t)

	//Those are the args of the project env we are trying to create
	args := CreateRequestParams{
		LayerID: RandomEnvLayer.ID,
		Source: utils.RandomName(),
		Status: utils.RandomName(),
		Payload: utils.RandomName(),

	}
	Request ,err := testQueries.CreateRequest(context.Background(), args) 
	require.NoError(t,err)
	require.NotEmpty(t,Request)
	require.Equal(t,Request.LayerID,args.LayerID)
	require.Equal(t,Request.Status,args.Status)
	require.Equal(t,Request.Source,args.Source)
	require.Equal(t,Request.Payload,args.Payload)
	require.NotZero(t,Request.ID)
	require.NotZero(t,Request.UpdatedAt)

	require.NotZero(t,Request.CreatedAt)
	return Request
}

//Testing the creation of a project environmenet		
func TestCreateRequest(t *testing.T) {
	CreateRandomRequest(t)
}

//Test the getter of a project environment
func TestGetRequest(t *testing.T){
	createdRequest  := CreateRandomRequest(t)
	gottenRequest,err := testQueries.GetRequest(context.Background(), createdRequest.ID)
	require.NoError(t,err)
	require.NotEmpty(t,gottenRequest)
	require.NotZero(t,gottenRequest.ID)
	require.NotZero(t,gottenRequest.CreatedAt)
	require.NotZero(t,gottenRequest.UpdatedAt)
	require.Equal(t,createdRequest.LayerID,gottenRequest.LayerID)
	require.Equal(t,createdRequest.Source,gottenRequest.Source)
	require.Equal(t,createdRequest.Payload,gottenRequest.Payload)
	require.Equal(t,createdRequest.UpdatedAt,gottenRequest.UpdatedAt)
	require.Equal(t,createdRequest.CreatedAt,gottenRequest.CreatedAt)

	require.Equal(t,createdRequest.Status,gottenRequest.Status)
}

//Test the update of a project environment
func TestUpdateRequest(t *testing.T){
	createdRequest := CreateRandomRequest(t)
	createEnvLayer := CreateRandomEnvLayer(t)
	args := UpdateRequestParams{
		ID: createdRequest.ID,
		Source: utils.RandomName(),
		Payload: utils.RandomName(),		
		Status: utils.RandomName(),
		LayerID: createEnvLayer.ID,
	}
	err:= testQueries.UpdateRequest(context.Background(),args)
	require.NoError(t,err)
	updatedRequest,err := testQueries.GetRequest(context.Background(),args.ID)
	require.NoError(t,err)
	require.Equal(t,args.LayerID,updatedRequest.LayerID)
	require.Equal(t,args.ID,updatedRequest.ID)
	require.Equal(t,args.Payload,updatedRequest.Payload)
	require.Equal(t,args.Status,updatedRequest.Status)
	require.Equal(t,args.Source,updatedRequest.Source)
	require.NotZero(t,updatedRequest.CreatedAt)
	require.NotZero(t,updatedRequest.UpdatedAt)
	
	} 	

// Test delete projects 
	func TestDeleteRequest(t *testing.T) {
		Request := CreateRandomRequest(t)
		err := testQueries.DeleteRequest(context.Background(), Request.ID)
		require.NoError(t, err)
		Request2, err := testQueries.GetRequest(context.Background(), Request.ID)
		require.Error(t, err)
		//require.EqualError(t, err, sql.ErrNoRows.Error())
		require.EqualError(t, err, "no rows in result set")
		require.Empty(t, Request2)
	}
//	Test list project environments	
	func TestListRequests(t *testing.T) {
		for i := 0; i < 10; i++ {
			CreateRandomRequest(t)
		}
		args := ListRequestsParams{
			Limit:  5,
			Offset: 5,
		}
		Requests, err := testQueries.ListRequests(context.Background(), args)
		require.NoError(t, err)
		require.Len(t, Requests, 5)
	
		for _, Request := range Requests {
			require.NotEmpty(t, Request)
		}
	}
