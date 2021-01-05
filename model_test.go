package wiseman

import (
	"testing"
)

func stringPtr(v string) *string {
	return &v
}

func TestParseBacklogItemPriority(t *testing.T) {
	tests := []struct {
		arg  string
		want BacklogItemPriority
	}{
		{"Low", BacklogItemPriorityLow},
		{"Medium", BacklogItemPriorityMedium},
		{"High", BacklogItemPriorityHigh},
		{"Urgent", BacklogItemPriorityUrgent},
		{"XXXXXX", BacklogItemPriorityUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := ParseBacklogItemPriority(tt.arg); got != tt.want {
				t.Errorf("ParseBacklogItemPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseBacklogItemType(t *testing.T) {
	tests := []struct {
		arg  string
		want BacklogItemType
	}{
		{"Unknown", BacklogItemTypeUnknown},
		{"Epic", BacklogItemTypeEpic},
		{"Feature", BacklogItemTypeFeature},
		{"Bug", BacklogItemTypeBug},
		{"Spike", BacklogItemTypeSpike},
		{"Other", BacklogItemTypeOther},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := ParseBacklogItemType(tt.arg); got != tt.want {
				t.Errorf("ParseBacklogItemType() = %v, want %v", got, tt.want)
			}
		})
	}
}
