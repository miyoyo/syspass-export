package main

import "encoding/json"

func UnmarshalAccountSearch(data []byte) (AccountSearch, error) {
	var r AccountSearch
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccountSearch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccountSearch struct {
	Jsonrpc string              `json:"jsonrpc"`
	Result  AccountSearchResult `json:"result"`
	ID      int64               `json:"id"`
}

type AccountSearchResult struct {
	ItemID        int64           `json:"itemId"`
	Result        []ResultElement `json:"result"`
	ResultCode    int64           `json:"resultCode"`
	ResultMessage interface{}     `json:"resultMessage"`
	Count         int64           `json:"count"`
}

type ResultElement struct {
	ID                        int64       `json:"id"`
	UserID                    int64       `json:"userId"`
	UserGroupID               int64       `json:"userGroupId"`
	UserEditID                int64       `json:"userEditId"`
	Name                      string      `json:"name"`
	ClientID                  int64       `json:"clientId"`
	CategoryID                int64       `json:"categoryId"`
	Login                     string      `json:"login"`
	URL                       string      `json:"url"`
	Notes                     string      `json:"notes"`
	OtherUserEdit             int64       `json:"otherUserEdit"`
	OtherUserGroupEdit        int64       `json:"otherUserGroupEdit"`
	IsPrivate                 int64       `json:"isPrivate"`
	IsPrivateGroup            int64       `json:"isPrivateGroup"`
	DateEdit                  *string     `json:"dateEdit"`
	PassDate                  int64       `json:"passDate"`
	PassDateChange            interface{} `json:"passDateChange"`
	ParentID                  int64       `json:"parentId"`
	CategoryName              string      `json:"categoryName"`
	ClientName                string      `json:"clientName"`
	UserGroupName             string      `json:"userGroupName"`
	UserName                  string      `json:"userName"`
	UserLogin                 string      `json:"userLogin"`
	UserEditName              string      `json:"userEditName"`
	UserEditLogin             string      `json:"userEditLogin"`
	NumFiles                  int64       `json:"num_files"`
	PublicLinkHash            interface{} `json:"publicLinkHash"`
	PublicLinkDateExpire      interface{} `json:"publicLinkDateExpire"`
	PublicLinkTotalCountViews interface{} `json:"publicLinkTotalCountViews"`
	CountView                 int64       `json:"countView"`
}

func UnmarshalJSONRPCSearch(data []byte) (JSONRPCSearch, error) {
	var r JSONRPCSearch
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JSONRPCSearch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func MakeJSONRPCSearch(apiKey string) *JSONRPCSearch {
	return &JSONRPCSearch{
		Jsonrpc: "2.0",
		Method:  "account/search",
		ID:      1,
		Params: JSONRPCSearchParams{
			AuthToken: apiKey,
		},
	}
}

type JSONRPCSearch struct {
	Jsonrpc string              `json:"jsonrpc"`
	Method  string              `json:"method"`
	Params  JSONRPCSearchParams `json:"params"`
	ID      int64               `json:"id"`
}

type JSONRPCSearchParams struct {
	AuthToken string `json:"authToken"`
}

func UnmarshalJSONRPCViewPass(data []byte) (JSONRPCViewPass, error) {
	var r JSONRPCViewPass
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JSONRPCViewPass) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func MakeJSONRPCViewPass(apiKey string, password string, id int64) *JSONRPCViewPass {
	return &JSONRPCViewPass{
		Jsonrpc: "2.0",
		Method:  "account/viewPass",
		ID:      1,
		Params: JSONRPCViewPassParams{
			AuthToken: apiKey,
			TokenPass: password,
			ID:        id,
		},
	}
}

type JSONRPCViewPass struct {
	Jsonrpc string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  JSONRPCViewPassParams `json:"params"`
	ID      int64                 `json:"id"`
}

type JSONRPCViewPassParams struct {
	AuthToken string `json:"authToken"`
	TokenPass string `json:"tokenPass"`
	ID        int64  `json:"id"`
}

func UnmarshalPasswordSearch(data []byte) (PasswordSearch, error) {
	var r PasswordSearch
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PasswordSearch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PasswordSearch struct {
	Jsonrpc string               `json:"jsonrpc"`
	Result  PasswordSearchResult `json:"result"`
	ID      int64                `json:"id"`
}

type PasswordSearchResult struct {
	ItemID        int64        `json:"itemId"`
	Result        ResultResult `json:"result"`
	ResultCode    int64        `json:"resultCode"`
	ResultMessage interface{}  `json:"resultMessage"`
	Count         int64        `json:"count"`
}

type ResultResult struct {
	Password string `json:"password"`
}
