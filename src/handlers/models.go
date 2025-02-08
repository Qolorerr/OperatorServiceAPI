package handlers

type requestCreateTag struct {
	Name string `json:"name"`
}

type requestCreateAppeal struct {
	UserId string   `json:"userId"`
	TagIds []string `json:"tagIds"`
}

type requestChangeTagsInAppeal struct {
	AppealId string   `json:"appealId"`
	TagIds   []string `json:"tagIds"`
}

type requestCreateGroup struct {
	TagIds []string `json:"tagIds"`
}

type requestChangeOperatorsInGroup struct {
	GroupId     string   `json:"groupId"`
	OperatorIds []string `json:"operatorIds"`
}

type requestChangeTagsInGroup struct {
	GroupId string   `json:"groupId"`
	TagIds  []string `json:"tagIds"`
}
