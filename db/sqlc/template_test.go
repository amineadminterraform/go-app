package db

import (
	"context"
	"testing"

	"github.com/amineadminterraform/go-app/utils"
	"github.com/stretchr/testify/require"
)

//A function to create a RandomTemplate
func createRandomTemplate(t *testing.T) Template{
	args := CreateTemplateParams{
		Name: utils.RandomName(),
		Path: utils.RandomName(),
		Type: utils.RandomName(),
	}
	Template , err := testQueries.CreateTemplate(context.Background(),args)
	require.NoError(t,err)
	require.NotEmpty(t,Template) 
	require.Equal(t,Template.Name,args.Name)
	require.Equal(t,Template.Type,args.Type)
	require.Equal(t,Template.Path,args.Path)
	require.NotZero(t,Template.ID)

	return Template
}
func TestCreateTemplate(t * testing.T) {
	createRandomTemplate(t)
}
//	Test get an  Template
func TestGetTemplate(t *testing.T) {
	template1 := createRandomTemplate(t)
	template2, err := testQueries.GetTemplate(context.Background(), template1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, template2)
	require.Equal(t, template1.ID, template2.ID)
	require.Equal(t, template1.Type, template2.Type)
	require.Equal(t, template1.Path, template2.Path)
	require.Equal(t, template1.Name, template2.Name)
}

//	Test update Templates
func TestUpdateTemplate(t *testing.T){
	createdTemplate := createRandomTemplate(t)

	args := UpdateTemplateParams{
		ID: createdTemplate.ID ,
		Name: utils.RandomName() ,
		Type: utils.RandomName() ,
		Path: utils.RandomName(),
	}
	err := testQueries.UpdateTemplate(context.Background(),args)
	require.NoError(t,err)
	updatedTemplate, err := testQueries.GetTemplate(context.Background(),args.ID)
	require.NoError(t,err)
	require.NotEmpty(t,updatedTemplate)
	require.Equal(t,args.ID,updatedTemplate.ID)
	require.Equal(t,args.Type,updatedTemplate.Type)
	require.Equal(t,args.Name,updatedTemplate.Name)
	require.Equal(t,args.Path,updatedTemplate.Path)
}

//	Test delete Templates
func TestDeleteTemplate(t *testing.T){
	newTemplate := createRandomTemplate(t)
	err := testQueries.DeleteTemplate(context.Background(),newTemplate.ID)
	require.NoError(t,err)
	testedTemplate , err := testQueries.GetTemplate(context.Background(),newTemplate.ID)
	require.Error(t,err)
	require.Empty(t,testedTemplate)
}

//	Test list Templates	
func TestListTemplates(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTemplate(t)
	}
	args := ListTemplatesParams{
		Limit:  5,
		Offset: 5,
	}
	Templates, err := testQueries.ListTemplates(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, Templates, 5)

	for _, Template := range Templates{
		require.NotEmpty(t, Template)
	}
}
