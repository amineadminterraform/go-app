package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/amineadminterraform/go-app/db/mock"
	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/amineadminterraform/go-app/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomProject() db.Project {
	return db.Project{
		ID:          utils.RandomInt(1, 10000),
		Name:        utils.RandomName(),
		GitPath:     utils.RandomGit(),
		Description: utils.RandomDescription(),
	}
}

func TestGetProjectAPI(t *testing.T) {
	project := randomProject()
	testCases := []struct {
		name string
		ID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
		}{
			{
				name:"OK", 
				ID: project.ID,
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
					GetProject(gomock.Any(), gomock.Eq(project.ID)).
					Times(1).
					Return(project, nil)
				},
				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusOK, recorder.Code)
					requiredBodyMatchAccount(t, recorder.Body, project)
				},
			},
			{
				name:"NotFound", 
				ID: project.ID,
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
					GetProject(gomock.Any(), gomock.Eq(project.ID)).
					Times(1).
					Return(db.Project{}, sql.ErrNoRows)
				},
				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusNotFound, recorder.Code)
				},
			},
			{
				name:"InternalError", 
				ID: project.ID,
				buildStubs: func(store *mockdb.MockStore) {
					store.EXPECT().
					GetProject(gomock.Any(), gomock.Eq(project.ID)).
					Times(1).
					Return(project, sql.ErrConnDone)
				},
				checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
					require.Equal(t, http.StatusInternalServerError, recorder.Code)
				},
			},
		
		}
		for i := range testCases{
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
			
				store := mockdb.NewMockStore(ctrl)
				tc.buildStubs(store)		
				server := NewServer(store)
				recorder := httptest.NewRecorder()
				url := fmt.Sprintf("/project/%d", tc.ID)
				request,err :=http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)
			
				server.router.ServeHTTP(recorder, request)
				// check response
				tc.checkResponse(t, recorder)
			})
		}
	}
	
		
		
			
			



//Check if body of the request is similar to the project
func requiredBodyMatchAccount(t* testing.T, body *bytes.Buffer, project db.Project) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotProject db.Project
	err = json.Unmarshal(data, &gotProject)
	require.NoError(t, err)
	require.Equal(t, project, gotProject)
}