package insomnium

import "time"

// Workspace represents a record of the insomnia.Workspace.db file.
type Workspace struct {
	ID          string `gedb:"_id"`
	Created     int    `gedb:"created"`
	Description string `gedb:"description"`
	Modified    int    `gedb:"modified"`
	Name        string `gedb:"name"`
	ParentID    string `gedb:"parentId"`
	Scope       string `gedb:"scope"`
	Type        string `gedb:"type"`
}

// GetID implements [Model].
func (w Workspace) GetID() string {
	return w.ID
}

// GetParentID implements [Model].
func (w Workspace) GetParentID() string {
	return w.ParentID
}

// RequestGroup represents a record of the insomnia.RequestGroup.db file.
type RequestGroup struct {
	ID                       string   `gedb:"_id"`
	Created                  int      `gedb:"created"`
	Description              string   `gedb:"description"`
	Environment              struct{} `gedb:"environment"`
	EnvironmentPropertyOrder any      `gedb:"environmentPropertyOrder"`
	MetaSortKey              int      `gedb:"metaSortKey"`
	Modified                 int      `gedb:"modified"`
	Name                     string   `gedb:"name"`
	ParentID                 string   `gedb:"parentId"`
	Type                     string   `gedb:"type"`
}

// GetID implements [Model].
func (rg RequestGroup) GetID() string {
	return rg.ID
}

// GetParentID implements [Model].
func (rg RequestGroup) GetParentID() string {
	return rg.ParentID
}

// Response represents a record of the insomnia.Response.db file.
type Response struct {
	ID                  string   `gedb:"_id"`
	BodyCompression     any      `gedb:"bodyCompression"`
	BodyPath            string   `gedb:"bodyPath"`
	BytesContent        int      `gedb:"bytesContent"`
	BytesRead           int      `gedb:"bytesRead"`
	ContentType         string   `gedb:"contentType"`
	Created             int      `gedb:"created"`
	ElapsedTime         float64  `gedb:"elapsedTime"`
	EnvironmentID       string   `gedb:"environmentId"`
	Error               string   `gedb:"error"`
	Headers             []Header `gedb:"headers"`
	HTTPVersion         string   `gedb:"httpVersion"`
	Modified            int      `gedb:"modified"`
	ParentID            string   `gedb:"parentId"`
	RequestVersionID    string   `gedb:"requestVersionId"`
	SettingSendCookies  bool     `gedb:"settingSendCookies"`
	SettingStoreCookies bool     `gedb:"settingStoreCookies"`
	StatusCode          int      `gedb:"statusCode"`
	StatusMessage       string   `gedb:"statusMessage"`
	TimelinePath        string   `gedb:"timelinePath"`
	Type                string   `gedb:"type"`
	URL                 string   `gedb:"url"`
}

// Header TODO
type Header struct {
	Name  string `gedb:"name"`
	Value string `gedb:"value"`
}

// GetID implements [Model].
func (r Response) GetID() string {
	return r.ID
}

// GetParentID implements [Model].
func (r Response) GetParentID() string {
	return r.ParentID
}

// Project represents a record of the insomnia.Project.db file.
type Project struct {
	ID       string `gedb:"_id"`
	Created  int    `gedb:"created"`
	Modified int    `gedb:"modified"`
	Name     string `gedb:"name"`
	RemoteID any    `gedb:"remoteId"`
	Type     string `gedb:"type"`
}

// GetID implements [Model].
func (p Project) GetID() string {
	return p.ID
}

// GetParentID implements [Model].
func (p Project) GetParentID() string {
	return ""
}

// Request represents a record of the insomnia.Request.db file.
type Request struct {
	ID             string `gedb:"_id"`
	Authentication struct {
		AccessTokenURL    string `gedb:"accessTokenUrl,omitempty"`
		AuthorizationURL  string `gedb:"authorizationUrl"`
		ClientID          string `gedb:"clientId,omitempty"`
		ClientSecret      string `gedb:"clientSecret,omitempty"`
		CredentialsInBody bool   `gedb:"credentialsInBody"`
		Disabled          bool   `gedb:"disabled"`
		GrantType         string `gedb:"grantType,omitempty"`
		Password          string `gedb:"password"`
		Scope             string `gedb:"scope"`
		TokenPrefix       string `gedb:"tokenPrefix,omitempty"`
		Type              string `gedb:"type,omitempty"`
		UseIso88591       bool   `gedb:"useISO88591"`
		Username          string `gedb:"username"`
	} `gedb:"authentication"`
	Body struct {
		MimeType string `gedb:"mimeType"`
		Text     string `gedb:"text"`
	} `gedb:"body"`
	Created     int    `gedb:"created"`
	Description string `gedb:"description"`
	Headers     []struct {
		Description string `gedb:"description"`
		Disabled    bool   `gedb:"disabled"`
		ID          string `gedb:"id,omitempty"`
		Name        string `gedb:"name"`
		Value       string `gedb:"value"`
	} `gedb:"headers"`
	IsPrivate   bool    `gedb:"isPrivate"`
	MetaSortKey float64 `gedb:"metaSortKey"`
	Method      string  `gedb:"method"`
	Modified    int     `gedb:"modified"`
	Name        string  `gedb:"name"`
	Parameters  []struct {
		Description string `gedb:"description"`
		Disabled    bool   `gedb:"disabled"`
		ID          string `gedb:"id"`
		Name        string `gedb:"name"`
		Value       string `gedb:"value"`
	} `gedb:"parameters"`
	ParentID                        string `gedb:"parentId"`
	PathParameters                  []any  `gedb:"pathParameters"`
	SegmentParams                   []any  `gedb:"segmentParams"`
	SettingDisableRenderRequestBody bool   `gedb:"settingDisableRenderRequestBody"`
	SettingEncodeURL                bool   `gedb:"settingEncodeUrl"`
	SettingFollowRedirects          string `gedb:"settingFollowRedirects"`
	SettingRebuildPath              bool   `gedb:"settingRebuildPath"`
	SettingSendCookies              bool   `gedb:"settingSendCookies"`
	SettingStoreCookies             bool   `gedb:"settingStoreCookies"`
	Type                            string `gedb:"type"`
	URL                             string `gedb:"url"`
}

// GetID implements [Model].
func (r Request) GetID() string {
	return r.ID
}

// GetParentID implements [Model].
func (r Request) GetParentID() string {
	return r.ParentID
}

// CookieJar TODO
type CookieJar struct {
	ID       string `gedb:"_id"`
	ParentID string `gedb:"parentId"`
	Cookies  []struct {
		ID            string    `gedb:"id"`
		Key           string    `gedb:"key"`
		Value         string    `gedb:"value"`
		Expires       any       `gedb:"expires"` // WARN
		Domain        string    `gedb:"domain"`
		Path          string    `gedb:"path"`
		Secure        bool      `gedb:"secure"`
		HTTPOnly      bool      `gedb:"httpOnly"`
		Extensions    []any     `gedb:"extensions"`
		Creation      time.Time `gedb:"creation"`
		CreationIndex int       `gedb:"creationIndex"`
		HostOnly      bool      `gedb:"hostOnly"`
		PathIsDefault bool      `gedb:"pathIsDefault"`
		LastAccessed  time.Time `gedb:"lastAccessed"`
	}
}

// GetID implements [Model].
func (cj CookieJar) GetID() string {
	return cj.ID
}

// GetParentID implements [Model].
func (cj CookieJar) GetParentID() string {
	return cj.ParentID
}

// Environment TODO
type Environment struct {
	ID       string            `gedb:"_id"`
	ParentID string            `gedb:"parentId"`
	Name     string            `gedb:"name"`
	Data     map[string]string `gedb:"data"`
	// dataPropertyOrder: Record<string, any> | null; WARN
	Color       string `gedb:"color"`
	MetaSortKey int    `gedb:"metaSortKey"`
	// For sync control
	IsPrivate bool `gedb:"isPrivate"`
}

// GetID implements [Model].
func (e Environment) GetID() string {
	return e.ID
}

// GetParentID implements [Model].
func (e Environment) GetParentID() string {
	return e.ParentID
}

// WorkspaceMeta TODO
type WorkspaceMeta struct {
	ID                        string  `gedb:"_id"`
	ActiveActivity            string  `gedb:"activeActivity"`
	ActiveEnvironmentID       string  `gedb:"activeEnvironmentId"`
	ActiveRequestID           string  `gedb:"activeRequestId"`
	ActiveUnitTestSuiteID     string  `gedb:"activeUnitTestSuiteId"`
	CachedGitLastAuthor       string  `gedb:"cachedGitLastAuthor"`
	CachedGitLastCommitTime   int     `gedb:"cachedGitLastCommitTime"`
	CachedGitRepositoryBranch string  `gedb:"cachedGitRepositoryBranch"`
	GitRepositoryID           string  `gedb:"gitRepositoryId"`
	HasSeen                   bool    `gedb:"hasSeen"`
	PaneHeight                float64 `gedb:"paneHeight"`
	PaneWidth                 float64 `gedb:"paneWidth"`
	ParentID                  string  `gedb:"parentId"`
	SidebarFilter             string  `gedb:"sidebarFilter"`
	SidebarHidden             bool    `gedb:"sidebarHidden"`
	SidebarWidth              float64 `gedb:"sidebarWidth"`
	PushSnapshotOnInitialize  bool    `gedb:"pushSnapshotOnInitialize"`
}

// GetID implements [Model].
func (wm WorkspaceMeta) GetID() string {
	return wm.ID
}

// GetParentID implements [Model].
func (wm WorkspaceMeta) GetParentID() string {
	return wm.ParentID
}

// Settings TODO
type Settings struct {
	ID                          string `gedb:"_id"`
	ParentID                    string `gedb:"parentId"` // always null
	AutoDetectColorScheme       bool   `gedb:"autoDetectColorScheme"`
	AutoHideMenuBar             bool   `gedb:"autoHideMenuBar"`
	AutocompleteDelay           int    `gedb:"autocompleteDelay"`
	ClearOAuth2SessionOnRestart bool   `gedb:"clearOAuth2SessionOnRestart"`
	DarkTheme                   string `gedb:"darkTheme"`
	DeviceID                    any    `gedb:"deviceId"` // WARN null
	DisableHTMLPreviewJs        bool   `gedb:"disableHtmlPreviewJs"`
	DisableResponsePreviewLinks bool   `gedb:"disableResponsePreviewLinks"`
	DisableUpdateNotification   bool   `gedb:"disableUpdateNotification"`
	EditorFontSize              int    `gedb:"editorFontSize"`
	EditorIndentSize            int    `gedb:"editorIndentSize"`
	EditorIndentWithTabs        bool   `gedb:"editorIndentWithTabs"`
	EditorKeyMap                string `gedb:"editorKeyMap"`
	EditorLineWrapping          bool   `gedb:"editorLineWrapping"`
	EnableAnalytics             bool   `gedb:"enableAnalytics"`
	ShowVariableSourceAndValue  bool   `gedb:"showVariableSourceAndValue"`
	FilterResponsesByEnv        bool   `gedb:"filterResponsesByEnv"`
	FollowRedirects             bool   `gedb:"followRedirects"`
	FontInterface               any    `gedb:"fontInterface"` // WARN null
	FontMonospace               any    `gedb:"fontMonospace"` // WARN null
	FontSize                    int    `gedb:"fontSize"`
	FontVariantLigatures        bool   `gedb:"fontVariantLigatures"`
	ForceVerticalLayout         bool   `gedb:"forceVerticalLayout"`
	HotKeyRegistry              any    `gedb:"hotKeyRegistry"` // WARN
	HTTPProxy                   string `gedb:"httpProxy"`
	HTTPSProxy                  string `gedb:"httpsProxy"`
	LightTheme                  string `gedb:"lightTheme"`
	MaxHistoryResponses         int    `gedb:"maxHistoryResponses"`
	MaxRedirects                int    `gedb:"maxRedirects"`
	MaxTimelineDataSizeKB       int    `gedb:"maxTimelineDataSizeKB"`
	NoProxy                     string `gedb:"noProxy"`
	NunjucksPowerUserMode       bool   `gedb:"nunjucksPowerUserMode"`
	PluginConfig                any    `gedb:"pluginConfig"` // WARN {}
	PluginPath                  string `gedb:"pluginPath"`
	PreferredHTTPVersion        string `gedb:"preferredHttpVersion"`
	ProxyEnabled                bool   `gedb:"proxyEnabled"`
	ShowPasswords               bool   `gedb:"showPasswords"`
	Theme                       string `gedb:"theme"`
	// milliseconds
	Timeout                 int    `gedb:"timeout"`
	UpdateAutomatically     bool   `gedb:"updateAutomatically"` //do not auto
	UpdateChannel           string `gedb:"updateChannel"`
	UseBulkHeaderEditor     bool   `gedb:"useBulkHeaderEditor"`
	UseBulkParametersEditor bool   `gedb:"useBulkParametersEditor"`
	ValidateAuthSSL         bool   `gedb:"validateAuthSSL"`
	ValidateSSL             bool   `gedb:"validateSSL"`
}

// GetID implements [Model].
func (s Settings) GetID() string {
	return s.ID
}

// GetParentID implements [Model].
func (s Settings) GetParentID() string {
	return ""
}
