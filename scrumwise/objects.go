package scrumwise

// The response object that you will receive when a request succeeds.
//
// https://www.scrumwise.com/api.html#result-object
// type Result struct {
// 	DataVersion DataVersion `json:"dataVersion"`
// 	Result      interface{} `json:"result"`
// }

// Data is an object that is returned as the result of the getData method.
//
// https://www.scrumwise.com/api.html#data-object
type Data struct {
	Persons        *[]Person `json:"persons"`
	DeletedPersons *[]Person `json:"deletedPersons"`
	Projects       []Project `json:"projects"`
}

// Person is a person.
//
// https://www.scrumwise.com/api.html#person-object
type Person struct {
	ID              string  `json:"id"`
	ExternalID      *string `json:"externalID"`
	FirstName       string  `json:"firstName"`
	LastName        *string `json:"lastName"`
	Nickname        *string `json:"nickname"`
	EmailAddress    string  `json:"emailAddress"`
	Phone           *string `json:"phone"`
	Im              *string `json:"im"`
	Description     *string `json:"description"`
	PhotoURL        *string `json:"photoURL"`
	IsActivated     bool    `json:"isActivated"`
	IsAdministrator bool    `json:"isAdministrator"`
}

// Project is a project.
//
// https://www.scrumwise.com/api.html#project-object
type Project struct {
	ID                   string           `json:"id"`                   // "729-11230-1",
	ExternalID           *string          `json:"externalID"`           // null,
	Name                 string           `json:"name"`                 // "Project 1",
	Description          *string          `json:"description"`          // null,
	Link                 *string          `json:"link"`                 // "http://www.example.com",
	Status               string           `json:"status"`               // "Active" or "Archived"
	RoughEstimateUnit    string           `json:"roughEstimateUnit"`    // "Points", "Days", or "Hours".
	DetailedEstimateUnit string           `json:"detailedEstimateUnit"` // "Points", "Days", or "Hours".
	TimeTrackingUnit     string           `json:"timeTrackingUnit"`     // "Points", "Days", or "Hours".
	Checklists           *[]Checklist     `json:"checklists"`           // null,
	Comments             *[]Comment       `json:"comments"`             // null,
	Attachments          *[]Attachment    `json:"attachments"`          // null,
	CustomFields         *[]CustomField   `json:"customFields"`         // null,
	Tags                 *[]Tag           `json:"tags"`                 // null,
	ProductOwnerIDs      []string         `json:"productOwnerIDs"`      // [ "729-334-1" ],
	StakeholderIDs       []string         `json:"stakeholderIDs"`       // [ "729-4151-1", "729-8107-6" ],
	Teams                *[]Team          `json:"teams"`                // null,
	Messages             *[]Message       `json:"messages"`             // null,
	Backlogs             *[]Backlog       `json:"backlogs"`             // null,
	BacklogItems         *[]BacklogItem   `json:"backlogItems"`         // null,
	Releases             *[]Release       `json:"releases"`             // null,
	Sprints              *[]Sprint        `json:"sprints"`              // null,
	Boards               *[]Board         `json:"boards"`               // null,
	Retrospectives       *[]Retrospective `json:"retrospectives"`       // null,
	Files                *[]File          `json:"files"`                // null,
	Relationships        *[]Relationship  `json:"relationships"`        // null
}

// Checklist is a checklist.
//
// https://www.scrumwise.com/api.html#checklist-object
type Checklist interface{}

// Comment is a comment on another object.
//
// https://www.scrumwise.com/api.html#comment-object
type Comment interface{}

// Attachment is a file attachment on another object.
//
//https://www.scrumwise.com/api.html#attachment-object
type Attachment interface{}

// CustomField is a custom field.
//
// https://www.scrumwise.com/api.html#customfield-object
type CustomField interface{}

// CustomField is a value of custom field on an object.
//
// https://www.scrumwise.com/api.html#customfieldvalue-object
type CustomFieldValue interface{}

// Tag is a tag.
//
// https://www.scrumwise.com/api.html#tag-object
type Tag interface{}

// Team is a team.
//
// https://www.scrumwise.com/api.html#team-object
type Team interface{}

// Message is a message in the message board of a project.
//
// https://www.scrumwise.com/api.html#message-object
type Message interface{}

// Backlog is a backlog in a project.
// The backlog items in a project are organized into backlogs and backlog lists.
// A project has one or more backlogs, which each contain one or more backlog lists,
// which in turn each contain zero or more backlog items.
//
// https://www.scrumwise.com/api.html#backlog-object
type Backlog interface{}

// A backlog list in a backlog.
// A backlog list contains zero or more backlog items.
//
// https://www.scrumwise.com/api.html#backloglist-object
type BacklogList struct {
	ID          string `json:"id"`          //  "729-43612-1229",
	ExternalID  string `json:"externalID"`  //  null,
	ProjectID   string `json:"projectID"`   //  "729-1321-1244",
	BacklogID   string `json:"backlogID"`   //  "729-5931-3524",
	Name        string `json:"name"`        //  "Example backlog list",
	Description string `json:"description"` //  "An example description.",
	Color       string `json:"color"`       //  "Green",
	IsArchived  bool   `json:"isArchived"`  //  true,

	// BacklogItems []string `json:"backlogItems"`
}

//  BacklogItem is a backlog item.
//
//  https://www.scrumwise.com/api.html#backlogitem-object
type BacklogItem struct {
	ID                string        `json:"id"`                //  "729-33220-627",
	ExternalID        *string       `json:"externalID"`        //  null,
	ItemNumber        int64         `json:"itemNumber"`        //  2149,
	ProjectID         string        `json:"projectID"`         //  "729-11230-132",
	ParentEpicID      *string       `json:"parentEpicID"`      //  null,
	BacklogListID     string        `json:"backlogListID"`     //  "729-15220-532",
	Name              string        `json:"name"`              //  "Example backlog item 1",
	Description       string        `json:"description"`       //  "An example description.",
	Link              string        `json:"link"`              //  "http://www.example.com",
	Priority          string        `json:"priority"`          //  "High",
	Typ               string        `json:"type"`              //  "Feature",
	Color             string        `json:"color"`             //  "Blue",
	CreatorID         string        `json:"creatorID"`         //  "729-32144-2210",
	CreationDate      int64         `json:"creationDate"`      //  1547836680000,
	BugState          *string       `json:"bugState"`          //  null,
	Reproducible      *string       `json:"reproducible"`      //  null,
	Resolution        *string       `json:"resolution"`        //  null,
	ReleaseID         string        `json:"releaseID"`         //  "729-3490-4392",
	Status            string        `json:"status"`            //  "In progress",
	ToDoDate          int64         `json:"toDoDate"`          //  1547984280000,
	InProgressDate    int64         `json:"inProgressDate"`    //  1548081480000,
	ToTestDate        *string       `json:"toTestDate"`        //  null,
	DoneDate          *int64        `json:"doneDate"`          //  null,
	RoughEstimate     float64       `json:"roughEstimate"`     //  15.0,
	Estimate          float64       `json:"estimate"`          //  21.25,
	UsedTime          float64       `json:"usedTime"`          //  9.5,
	RemainingWork     float64       `json:"remainingWork"`     //  12.5,
	SprintID          string        `json:"sprintID"`          //  "729-5445-4334",
	TeamID            string        `json:"teamID"`            //  "729-6764-8437",
	BoardID           string        `json:"boardID"`           //  "729-2345-322",
	BoardColumnID     string        `json:"boardColumnID"`     //  "729-74532-4342",
	IsArchivedInBoard bool          `json:"isArchivedInBoard"` //  false,
	AssignedPersonIDs []string      `json:"assignedPersonIDs"` //  [ "729-5444-343" ],
	DueDate           *string       `json:"dueDate"`           //  null,
	CustomFieldValues *string       `json:"customFieldValues"` //  null,
	TagIDs            *[]Tag        `json:"tagIDs"`            //  [ "729-14820-466", "729-6893-337" ],
	Checklists        *[]Checklist  `json:"checklists"`        //  null,
	Comments          *[]Comment    `json:"comments"`          //  null,
	Attachments       *[]Attachment `json:"attachments"`       //  null,
	TimeEntries       interface{}   `json:"timeEntries"`       //  null,
	Commits           interface{}   `json:"commits"`           //  null,
	Tasks             *[]Task       `json:"tasks"`             //  null,
	ChildBacklogItems []BacklogItem `json:"childBacklogItems"` //  [ ]
}

// A task in a backlog item.
//
// https://www.scrumwise.com/api.html#task-object
type Task struct {
	Id                string      `json:"id"`                //  "729-3322-69378",
	ExternalID        *string     `json:"externalID"`        //  null,
	TaskNumber        int64       `json:"taskNumber"`        //  1325,
	ProjectID         string      `json:"projectID"`         //  "729-11230-1",
	BacklogItemID     string      `json:"backlogItemID"`     //  "729-6743-353",
	Name              string      `json:"name"`              //  "Example task 1",
	Description       *string     `json:"description"`       //  "An example description.",
	Link              *string     `json:"link"`              //  "http://www.example.com",
	Color             *string     `json:"color"`             //  "Blue",
	CreationDate      int64       `json:"creationDate"`      //  1547836680000,
	Status            string      `json:"status"`            //  "In progress",
	ToDoDate          int64       `json:"toDoDate"`          //  1547836680000,
	InProgressDate    *int64      `json:"inProgressDate"`    //  1548081480000,
	ToTestDate        *int64      `json:"toTestDate"`        //  null,
	DoneDate          *int64      `json:"doneDate"`          //  null,
	BoardColumnID     *string     `json:"boardColumnID"`     //  "729-4765-83",
	Estimate          float64     `json:"estimate"`          //  10.25,
	UsedTime          float64     `json:"usedTime"`          //  4.5,
	RemainingWork     float64     `json:"remainingWork"`     //  5.0,
	AssignedPersonIDs []string    `json:"assignedPersonIDs"` //  [ "729-7654-243" ],
	DueDate           *string     `json:"dueDate"`           //  null,
	CustomFieldValues interface{} `json:"customFieldValues"` //  null,
	TagIDs            interface{} `json:"tagIDs"`            //  [ "729-14820-466", "729-12820-336" ],
	Checklists        interface{} `json:"checklists"`        //  null,
	Comments          interface{} `json:"comments"`          //  null,
	Attachments       interface{} `json:"attachments"`       //  null,
	TimeEntries       interface{} `json:"timeEntries"`       //  null,
	Commits           interface{} `json:"commits"`           //  null
}

// Release is a release.
//
// https://www.scrumwise.com/api.html#release-object
type Release interface{}

// Sprint is a sprint.
//
// https://www.scrumwise.com/api.html#sprint-object
type Sprint struct {
	ID                       string                    `json:"id"` // "729-8333-622".
	ExternalID               *string                   `json:"externalID"`
	ProjectID                string                    `json:"projectID"`   // "729-1321-1244".
	Name                     string                    `json:"name"`        // "Example sprint".
	Description              *string                   `json:"description"` // "An example description.".
	Link                     *string                   `json:"link"`        // "http://www.example.com".
	StartDate                *string                   `json:"startDate"`   // "2019-05-06" or null.
	EndDate                  *string                   `json:"endDate"`     // "2019-05-17" or null.
	BoardID                  *string                   `json:"boardID"`     // "729-2351-553" or null.
	Status                   string                    `json:"status"`      // "In planning", "In progress", "Completed", or "Aborted".
	IsArchived               bool                      `json:"isArchived"`
	CustomFieldValues        *[]CustomFieldValue       `json:"customFieldValues"`
	TagIDs                   []string                  `json:"tagIDs"` // [ "729-10820-446", "729-10820-226" ].
	Checklists               *[]Checklist              `json:"checklists"`
	Comments                 *[]Comment                `json:"comments"`
	TeamSprintParticipations []TeamSprintParticipation `json:"teamSprintParticipations"` //
}

// Board is a Scrum task board or Kanban board.
//
// https://www.scrumwise.com/api.html#board-object
type Board interface{}

// Retrospective is a retrospective.
//
// https://www.scrumwise.com/api.html#retrospective-object
type Retrospective interface{}

// File is a file or a folder in a project.
//
// https://www.scrumwise.com/api.html#file-object
type File interface{}

// Relationship is a relationship between two objects.
//
// https://www.scrumwise.com/api.html#relationship-object
type Relationship interface{}

// TeamSprintParticipation is an object that contains information about a team’s participation in a sprint.
//
// https://www.scrumwise.com/api.html#teamsprintparticipation-object
type TeamSprintParticipation struct{}
