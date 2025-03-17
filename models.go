package insomnium

// Workspace represents a record of the insomnia.Workspace.db file.
type Workspace struct {
	ID          string `json:"_id"`
	Created     int    `json:"created"`
	Description string `json:"description"`
	Modified    int    `json:"modified"`
	Name        string `json:"name"`
	ParentID    string `json:"parentId"`
	Scope       string `json:"scope"`
	Type        string `json:"type"`
}

// RequestGroup represents a record of the insomnia.RequestGroup.db file.
type RequestGroup struct {
	ID                       string   `json:"_id"`
	Created                  int      `json:"created"`
	Description              string   `json:"description"`
	Environment              struct{} `json:"environment"`
	EnvironmentPropertyOrder any      `json:"environmentPropertyOrder"`
	MetaSortKey              int      `json:"metaSortKey"`
	Modified                 int      `json:"modified"`
	Name                     string   `json:"name"`
	ParentID                 string   `json:"parentId"`
	Type                     string   `json:"type"`
}

// Response represents a record of the insomnia.Response.db file.
type Response struct {
	ID              string  `json:"_id"`
	BodyCompression any     `json:"bodyCompression"`
	BodyPath        string  `json:"bodyPath"`
	BytesContent    int     `json:"bytesContent"`
	BytesRead       int     `json:"bytesRead"`
	ContentType     string  `json:"contentType"`
	Created         int     `json:"created"`
	ElapsedTime     float64 `json:"elapsedTime"`
	EnvironmentID   string  `json:"environmentId"`
	Error           string  `json:"error"`
	Headers         []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	HTTPVersion         string `json:"httpVersion"`
	Modified            int    `json:"modified"`
	ParentID            string `json:"parentId"`
	RequestVersionID    string `json:"requestVersionId"`
	SettingSendCookies  bool   `json:"settingSendCookies"`
	SettingStoreCookies bool   `json:"settingStoreCookies"`
	StatusCode          int    `json:"statusCode"`
	StatusMessage       string `json:"statusMessage"`
	TimelinePath        string `json:"timelinePath"`
	Type                string `json:"type"`
	URL                 string `json:"url"`
}

// Project represents a record of the insomnia.Project.db file.
type Project struct {
	ID       string `json:"_id"`
	Created  int    `json:"created"`
	Modified int    `json:"modified"`
	Name     string `json:"name"`
	ParentID any    `json:"parentId"`
	RemoteID any    `json:"remoteId"`
	Type     string `json:"type"`
}

// Request represents a record of the insomnia.Request.db file.
type Request struct {
	ID             string `json:"_id"`
	Authentication struct {
		Disabled    bool   `json:"disabled"`
		Password    string `json:"password"`
		Type        string `json:"type,omitempty"`
		UseIso88591 bool   `json:"useISO88591"`
		Username    string `json:"username"`
	} `json:"authentication"`
	Body struct {
		MimeType string `json:"mimeType,omitempty"`
		Text     string `json:"text,omitempty"`
	} `json:"body"`
	Created     int    `json:"created"`
	Description string `json:"description"`
	Headers     []struct {
		ID    string `json:"id,omitempty"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	IsPrivate   bool    `json:"isPrivate"`
	MetaSortKey float64 `json:"metaSortKey"`
	Method      string  `json:"method"`
	Modified    int     `json:"modified"`
	Name        string  `json:"name"`
	Parameters  []struct {
		Description string `json:"description"`
		Disabled    bool   `json:"disabled"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Value       string `json:"value"`
	} `json:"parameters"`
	ParentID                        string `json:"parentId"`
	SegmentParams                   []any  `json:"segmentParams"`
	SettingDisableRenderRequestBody bool   `json:"settingDisableRenderRequestBody"`
	SettingEncodeURL                bool   `json:"settingEncodeUrl"`
	SettingFollowRedirects          string `json:"settingFollowRedirects"`
	SettingRebuildPath              bool   `json:"settingRebuildPath"`
	SettingSendCookies              bool   `json:"settingSendCookies"`
	SettingStoreCookies             bool   `json:"settingStoreCookies"`
	Type                            string `json:"type"`
	URL                             string `json:"url"`
}
