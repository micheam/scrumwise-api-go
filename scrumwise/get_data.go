package scrumwise

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var defaultProps = []string{
	"Project.backlogItems",
	"Project.sprints",
	"Project.boards",
	// "Project.backlogs",
	"BacklogItem.tasks",
}

type GetDataParam struct {
	ProjectIDs []string
	Properties []string // TODO(micheam): Represent selectable value as an enum.
}

func NewGetDataParam(id string) *GetDataParam {
	return &GetDataParam{[]string{id}, defaultProps}
}

// joinedProjectIDs will return project ids joined with comma.
func (param *GetDataParam) joinedProjectIDs() string {
	ids := make([]string, len(param.ProjectIDs))
	for i, pid := range param.ProjectIDs {
		ids[i] = string(pid)
	}
	return strings.Join(ids, ",")
}

// joinedProperties will return properties joined with comma.
func (param *GetDataParam) joinedProperties() string {
	props := make([]string, len(param.Properties))
	for i, prop := range param.Properties {
		props[i] = string(prop)
	}
	return strings.Join(props, ",")
}

func (param *GetDataParam) asBody() io.Reader {
	prop := fmt.Sprintf(`projectIDs=%s&includeProperties=%s`,
		param.joinedProjectIDs(),
		param.joinedProperties())
	return strings.NewReader(prop)
}

type GetDataResult struct {
	DataVersion int64 `json:"dataVersion"`
	Result      Data  `json:"result"`
}

func GetData(ctx context.Context, param GetDataParam) (*GetDataResult, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", Endpoint("getData"), param.asBody())
	if err != nil {
		return nil, fmt.Errorf("failed to generate http Request: %w", err)
	}
	req.SetBasicAuth(
		// XXX(micheam): get from client
		"michito.maeda@b.so-tech.co.jp",
		"33903092E7329E0C78881C9BC10D000FC98AF2125D03817CA2A1453E3AD2D92D",
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to http.Client Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.New(resp.Status)
	}

	result := new(GetDataResult)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return result, nil
}

// GetDataVersion return current data version.
//
// https://www.scrumwise.com/api.html#getting-data
func GetDataVersion(ctx context.Context) (int64, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", Endpoint("getDataVersion"), nil)
	if err != nil {
		return -1, fmt.Errorf("failed to generate http Request: %w", err)
	}
	req.SetBasicAuth(
		// XXX(micheam): get from client
		"michito.maeda@b.so-tech.co.jp",
		"33903092E7329E0C78881C9BC10D000FC98AF2125D03817CA2A1453E3AD2D92D",
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("failed to http.Client Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return -1, errors.New(resp.Status)
	}

	result := new(struct {
		DataVersion int64 `json:"dataVersion"`
	})
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return -1, fmt.Errorf("failed to decode json: %w", err)
	}
	return result.DataVersion, nil
}
