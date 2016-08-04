package output

type UserInfo struct {
	userId              float64 `json:"userId"`
	nickname            string  `json:"nickname"`
	loginToken          string  `json:"loginToken"`
	avatar              string  `json:"avatar"`
	registerCountryCode string  `json:"registerCountryCode"`
	registerLocation    string  `json:"registerLocation"`
	userType            float64 `json:"userType"`
	registerTime        float64 `json:"registerTime"`
}
