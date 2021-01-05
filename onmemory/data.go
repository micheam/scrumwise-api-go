package onmemory

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/micheam/wiseman"
	"github.com/micheam/wiseman/scrumwise"
)

var (
	lock sync.RWMutex

	dataVersion  int64 = -1
	projectID    wiseman.ProjectID
	project      *wiseman.Project
	backlogItems map[wiseman.BacklogItemID]*wiseman.BacklogItem
)

func Current() int64 {
	lock.RLock()
	defer lock.RUnlock()
	return dataVersion
}

func init() {
	id := wiseman.ProjectID(os.Getenv("SCRUMWISE_PROJECT"))
	if id == "" {
		log.Println("skip to init project-id: env var `SCRUMWISE_PROJECT` not found")
		return
	}
	SetProjectID(id)
}

func SetProjectID(id wiseman.ProjectID) {
	lock.Lock()
	defer lock.Unlock()
	projectID = id
}

func LoadData(ctx context.Context) error {
	lock.Lock()
	defer lock.Unlock()

	// TODO: Skip to load data if current-data is latest

	param := scrumwise.GetDataParam{}
	param.ProjectIDs = []string{projectID.String()}
	param.Properties = []string{
		"Project.backlogItems", "Project.sprints", "Project.boards",
		"BacklogItem.tasks",
	}

	data, err := scrumwise.GetData(ctx, param)
	if err != nil {
		return fmt.Errorf("failed to get data from scrumwise: %w", err)
	}

	if len(data.Result.Projects) == 0 {
		return fmt.Errorf("project %q not found", projectID)
	}

	dataVersion = data.DataVersion

	// TODO: extranct as a function


	//
	// WIP: 
	//
	// scrumwise.Project から、 wiseman.Project への変換処理を実装途中
	prj := wiseman.Project{
	ID          projectID,
	Name        
	Description NullString
	Link        NullURL
	Archived    bool
	}

	&data.Result.Projects[0]

	return nil
}
