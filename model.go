package wiseman

import (
	"time"
)

type (
	ID            string
	ProjectID     ID
	SprintID      ID
	BacklogItemID ID
)

func (pid ProjectID) String() string {
	return string(pid)
}

type ExternalID NullString

type Project struct {
	ID          ProjectID
	Name        string
	Description NullString
	Link        NullURL
	Archived    bool
}

type Sprint struct {
	ID          ID
	ExternalID  ExternalID
	Project     *Project
	Name        string
	Description NullString
	Link        NullURL
	StartDate   NullTime
	EndDate     NullTime
	Status      SprintStatus
	Archived    bool
}

type SprintStatus int

const (
	SprintStatusUnknown = iota
	SprintStatusInPlanning
	SprintStatusInProgress
	SprintStatusCompleted
	SprintStatusAborted
)

func (s SprintStatus) String() string {
	return [...]string{
		"Unknown", "In planning", "In progress", "Completed", "Aborted",
	}[s]
}

func ParseSprintStatus(s string) SprintStatus {
	switch s {
	case "In planning":
		return SprintStatusInPlanning
	case "In progress":
		return SprintStatusInProgress
	case "Completed":
		return SprintStatusCompleted
	case "Aborted":
		return SprintStatusAborted
	}
	return SprintStatusUnknown
}

type BacklogItem struct {
	ID         BacklogItemID
	ExternalID ExternalID
	Project    *Project

	Name        string
	Description NullString
	Link        NullURL
	Priority    BacklogItemPriority
	Typ         BacklogItemType
	Creator     string
	CreatedAt   time.Time

	Status          BacklogItemStatus
	InProgressSince NullTime
	ToTestSince     NullTime
	DoneAt          NullTime
	RoughEstimate   TimeUnit
	Estimate        TimeUnit
	UsedTime        TimeUnit
	RemainingWork   TimeUnit

	Parent   *BacklogItem
	Children []BacklogItem
}

type BacklogItemStatus int

const (
	BacklogItemStatusUnknown = iota
	BacklogItemStatusNew
	BacklogItemStatusReadyForEstimation
	BacklogItemStatusReadyForSprint
	BacklogItemStatusAssignedToSprint
	BacklogItemStatusToDo
	BacklogItemStatusInProgress
	BacklogItemStatusToTest
	BacklogItemStatusDone
	BacklogItemStatusSprintCompleted
	BacklogItemStatusReleased
)

func (b BacklogItemStatus) String() string {
	return [...]string{
		"Unknown",
		"New",
		"Ready for estimation",
		"Ready for sprint",
		"Assigned to sprint",
		"To do",
		"In progress",
		"To test",
		"Done",
		"Sprint completed",
		"Released",
	}[b]
}

func ParseBacklogItemStatus(s string) BacklogItemStatus {
	switch s {
	case "New":
		return BacklogItemStatusNew
	case "Ready for estimation":
		return BacklogItemStatusReadyForEstimation
	case "Ready for sprint":
		return BacklogItemStatusReadyForSprint
	case "Assigned to sprint":
		return BacklogItemStatusAssignedToSprint
	case "To do":
		return BacklogItemStatusToDo
	case "In progress":
		return BacklogItemStatusInProgress
	case "To test":
		return BacklogItemStatusToTest
	case "Done":
		return BacklogItemStatusDone
	case "Sprint completed":
		return BacklogItemStatusSprintCompleted
	case "Released":
		return BacklogItemStatusReleased
	}
	return BacklogItemStatusUnknown
}

type BacklogItemPriority int

const (
	BacklogItemPriorityUnknown = iota
	BacklogItemPriorityLow
	BacklogItemPriorityMedium
	BacklogItemPriorityHigh
	BacklogItemPriorityUrgent
)

func (b BacklogItemPriority) String() string {
	return [...]string{
		"Unknown", "Low", "Medium", "High", "Urgent",
	}[b]
}

func ParseBacklogItemPriority(s string) BacklogItemPriority {
	switch s {
	case "Low":
		return BacklogItemPriorityLow
	case "Medium":
		return BacklogItemPriorityMedium
	case "High":
		return BacklogItemPriorityHigh
	case "Urgent":
		return BacklogItemPriorityUrgent
	}
	return BacklogItemPriorityUnknown
}

type BacklogItemType int

const (
	BacklogItemTypeUnknown = iota
	BacklogItemTypeEpic
	BacklogItemTypeFeature
	BacklogItemTypeBug
	BacklogItemTypeSpike
	BacklogItemTypeOther
)

func (b BacklogItemType) String() string {
	return [...]string{
		"Unknown", "Epic", "Feature", "Bug", "Spike", "Other",
	}[b]
}

func ParseBacklogItemType(s string) BacklogItemType {
	switch s {
	case "Epic":
		return BacklogItemTypeEpic
	case "Feature":
		return BacklogItemTypeFeature
	case "Bug":
		return BacklogItemTypeBug
	case "Spike":
		return BacklogItemTypeSpike
	case "Other":
		return BacklogItemTypeOther
	}
	return BacklogItemTypeUnknown
}
