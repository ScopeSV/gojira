package issues

type IssueSearch struct {
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
	Issues     []struct {
		Id     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Created           string `json:"created"`
			Updated           string `json:"updated"`
			IssueName         string `json:"summary"`
			CustomField_10020 []struct {
				SprintName string `json:"name"`
			} `json:"customfield_10020"`
		} `json:"fields"`
	} `json:"issues"`
}

type Issue struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Assignee    struct {
			DisplayName string `json:"displayName"`
		} `json:"assignee"`
		Creator struct {
			DisplayName string `json:"displayName"`
		} `json:"creator"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		SubTask []struct {
			Id  string `json:"id"`
			Key string `json:"key"`
		} `json:"subtasks"`
		CustomField_10020 []struct {
			SprintName string `json:"name"`
		} `json:"customfield_10020"`
		Comment struct {
			Comments []struct {
				Id string `json:"id"`
			}
			Total      int `json:"total"`
			MaxResults int `json:"maxResults"`
			StartAt    int `json:"startAt"`
		} `json:"comment"`
	} `json:"fields"`
}
