package api

import (
	"testing"

	mockdb "github.com/amineadminterraform/go-app/db/mock/mockdb"
	db "github.com/amineadminterraform/go-app/db/sqlc"
	"github.com/amineadminterraform/go-app/utils"
	"github.com/golang/mock/gomock"
)
func randomProject() db.Project {
	return db.Project{
		ID: utils.RandomInt(1,10000),
		Name: utils.RandomName(),
		GitPath: utils.RandomGit(),
		Description: utils.RandomDescription(),
		}	
	}
	
func TestGetProjectAPI(t *testing.T) {
		project := randomProject()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)
		
		store.EXPECT().
			GetProject(gomock.Any(), gomock.Eq(project.ID)).
			Times(1).
			Return(project, nil)  
	}
	