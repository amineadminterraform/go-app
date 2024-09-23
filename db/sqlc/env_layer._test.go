package db

import (
	"context"
	"testing"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

//In order to test the creations of project environement we need to test it with a proper EnvironmentID because its a foreign constraint
// A test for the existence of the id will be performed later
func CreateRandomEnvLayer(t *testing.T) EnvLayer {

	//We Need to create a random Environment ( in a random project)
	environment := CreateRandomProjectEnvironment(t)
	// we Need to create a random Process with a random argoworkflow and a random template	
	Process := CreateRandomProcess(t)
	//Those are the args of the project env we are trying to create
	args := CreateEnvLayerParams{
		EnvironmentID: environment.ID,
		S3Path: utils.RandomName(),
		ProcessID: Process.ID ,	
		CurrentRequestID: pgtype.Int8{},
	}
	EnvLayer ,err := testQueries.CreateEnvLayer(context.Background(), args) 
	require.NoError(t,err)
	require.NotEmpty(t,EnvLayer)
	require.Equal(t,EnvLayer.EnvironmentID,args.EnvironmentID)
	require.Equal(t,EnvLayer.S3Path,args.S3Path)
	require.Equal(t,EnvLayer.ProcessID,args.ProcessID)
	require.NotZero(t,EnvLayer.ID)
	require.NotZero(t,EnvLayer.CreatedAt)
	require.NotZero(t,EnvLayer.UpdatedAt)
	return EnvLayer
}

//Testing the creation of a project environmenet		
func TestCreateEnvLayer(t *testing.T) {
	CreateRandomEnvLayer(t)
}

//Test the getter of a project environment
func TestGetEnvLayer(t *testing.T){
	createdEnvLayer  := CreateRandomEnvLayer(t)
	gottenEnvLayer,err := testQueries.GetEnvLayer(context.Background(), createdEnvLayer.ID)
	require.NoError(t,err)
	require.NotEmpty(t,gottenEnvLayer)
	require.NotZero(t,gottenEnvLayer.ID)
	require.NotZero(t,gottenEnvLayer.CreatedAt)
	require.NotZero(t,gottenEnvLayer.UpdatedAt)
	require.Equal(t,createdEnvLayer.EnvironmentID,gottenEnvLayer.EnvironmentID)
	require.Equal(t,createdEnvLayer.ProcessID,gottenEnvLayer.ProcessID)
	require.Equal(t,createdEnvLayer.S3Path,gottenEnvLayer.S3Path)
}

//Test the update of a project environment
func TestUpdateEnvLayer(t *testing.T){
	createdEnvlayer:= CreateRandomEnvLayer(t)
	createdProcess := CreateRandomProcess(t)
	args := UpdateEnvLayerParams{
		ID: createdEnvlayer.ID,
	
		S3Path: utils.RandomS3Path(),
		ProcessID: createdProcess.ID,
	}
	err:= testQueries.UpdateEnvLayer(context.Background(),args)
	require.NoError(t,err)
	updatedEnvLayer,err := testQueries.GetEnvLayer(context.Background(),args.ID)
	require.NoError(t,err)
	require.Equal(t,args.ProcessID,updatedEnvLayer.ProcessID)
	require.Equal(t,args.S3Path,updatedEnvLayer.S3Path)
	require.Equal(t,createdEnvlayer.EnvironmentID,updatedEnvLayer.EnvironmentID)
	} 	

// Test delete projects 
	func TestDeleteEnvLayer(t *testing.T) {
		EnvLayer := CreateRandomEnvLayer(t)
		err := testQueries.DeleteEnvLayer(context.Background(), EnvLayer.ID)
		require.NoError(t, err)
		EnvLayer2, err := testQueries.GetEnvLayer(context.Background(), EnvLayer.ID)
		require.Error(t, err)
		//require.EqualError(t, err, sql.ErrNoRows.Error())
		require.EqualError(t, err, "no rows in result set")
		require.Empty(t, EnvLayer2)
	}
//	Test list project environments	
	func TestListEnvLayers(t *testing.T) {
		for i := 0; i < 10; i++ {
			CreateRandomEnvLayer(t)
		}
		args := ListEnvLayersParams{
			Limit:  5,
			Offset: 5,
		}
		envLayers, err := testQueries.ListEnvLayers(context.Background(), args)
		require.NoError(t, err)
		require.Len(t, envLayers, 5)
	
		for _, envLayer := range envLayers {
			require.NotEmpty(t, envLayer)
		}
	}
