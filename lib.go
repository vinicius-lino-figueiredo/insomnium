package insomnium

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/noirbizarre/gonja"
	"github.com/vinicius-lino-figueiredo/gedb"
	"github.com/vinicius-lino-figueiredo/gedb/domain"
)

// Insomnium loads and stores information read from the insomnia.*.db files.
type Insomnium struct {
	path           string
	responsesDir   string
	prefix         string
	suffix         string
	cl             *http.Client
	transport      *http.Transport
	workspaces     *Repository[Workspace]
	requestGroups  *Repository[RequestGroup]
	responses      *Repository[Response]
	projects       *Repository[Project]
	requests       *Repository[Request]
	cookieJars     *Repository[CookieJar]
	environments   *Repository[Environment]
	workspaceMetas *Repository[WorkspaceMeta]
	settings       *Repository[Settings]
}

// NewInsomnium creates an empty instance of *Insomnium, a type that should
// store information from the Insomnium db files. It loads a file reader
// interface that uses the package os to read the insomnia.*.db files and a json
// unmarshaler that uses the standard package encoding/json. Options can be used
// to mock the file reading and unmarshaling process.
func NewInsomnium(path, prefix, suffix string) (*Insomnium, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{TLSClientConfig: &tls.Config{}}

	i := Insomnium{
		cl:           &http.Client{Transport: transport},
		transport:    transport,
		path:         path,
		responsesDir: filepath.Join(path, "responses"),
		prefix:       prefix,
		suffix:       suffix,
	}

	err = i.createDatabases(context.Background())
	if err != nil {
		return nil, err
	}

	return &i, nil
}

// GetAllProjects lists all projects.
func (i *Insomnium) GetAllProjects(ctx context.Context) ([]Project, error) {
	return i.projects.All(ctx)
}

// GetProject returns a specific project by its id.
func (i *Insomnium) GetProject(ctx context.Context, id string) (Project, error) {
	return i.projects.Get(ctx, id)
}

// GetAllWorkspaces lists all projects.
func (i *Insomnium) GetAllWorkspaces(ctx context.Context) ([]Workspace, error) {
	return i.workspaces.All(ctx)
}

// GetChildWorkspaces lists all workspaces for a given project.
func (i *Insomnium) GetChildWorkspaces(ctx context.Context, parent string) ([]Workspace, error) {
	return i.workspaces.GetChildren(ctx, parent)
}

// GetWorkspace returns a specific workspace by its id.
func (i *Insomnium) GetWorkspace(ctx context.Context, id string) (Workspace, error) {
	return i.workspaces.Get(ctx, id)
}

// GetRequestWorkspace returns the workspace ancestor of a given request. Since
// the request could be nested within multiple request groups, the ancestor
// workspace may not be the parent of the request.
func (i *Insomnium) GetRequestWorkspace(ctx context.Context, id string) (string, error) {
	// following req groups until we find one with a non req group parent,
	// which should be a workspace instead.
	for len(id) > 4 && id[:4] == "fld_" {
		fld, err := i.requestGroups.Get(ctx, id)
		if err != nil {
			return "", err
		}
		id = fld.ParentID
	}
	return id, nil
}

// GetAllRequestGroups returns all request groups.
func (i *Insomnium) GetAllRequestGroups(ctx context.Context) ([]RequestGroup, error) {
	return i.requestGroups.All(ctx)
}

// GetChildRequestGroups lists all request groups for a given parent, which can
// be either a workspace or another request group.
func (i *Insomnium) GetChildRequestGroups(ctx context.Context, parent string) ([]RequestGroup, error) {
	return i.requestGroups.GetChildren(ctx, parent)
}

// GetRequestGroup returns a specific request group by its id.
func (i *Insomnium) GetRequestGroup(ctx context.Context, id string) (RequestGroup, error) {
	return i.requestGroups.Get(ctx, id)
}

// GetAllRequests returns all requests.
func (i *Insomnium) GetAllRequests(ctx context.Context) ([]Request, error) {
	return i.requests.All(ctx)
}

// GetChildRequests returns all requests for a given parent, which can be either
// a workspace or a request group.
func (i *Insomnium) GetChildRequests(ctx context.Context, parent string) ([]Request, error) {
	return i.requests.GetChildren(ctx, parent)
}

// GetRequest returns a specific request by its id.
func (i *Insomnium) GetRequest(ctx context.Context, id string) (Request, error) {
	return i.requests.Get(ctx, id)
}

// GetChildWorkspaceMeta returns the metadata of a given workspace.
func (i *Insomnium) GetChildWorkspaceMeta(ctx context.Context, parent string) (WorkspaceMeta, error) {
	return i.workspaceMetas.GetChild(ctx, parent)
}

// GetLatestResponse TODO
func (i *Insomnium) GetLatestResponse(ctx context.Context, id string) (*Response, io.ReadCloser, error) {
	res, err := i.responses.Find(ctx,
		map[string]any{"parentId": id},
		domain.WithSort(gedb.Sort{{Key: "modified", Order: -1}}),
		domain.WithLimit(1),
	)
	if err != nil {
		return nil, nil, err
	}
	if len(res) == 0 {
		return nil, nil, nil
	}
	f, err := os.Open(res[0].BodyPath)
	if err != nil {
		return nil, nil, err
	}
	return &res[0], f, nil
}

// GetResponse returns the response and its body.
func (i *Insomnium) GetResponse(ctx context.Context, id string) (Response, io.ReadCloser, error) {
	res, err := i.responses.Get(ctx, id)
	if err != nil {
		return res, nil, err
	}
	f, err := os.Open(res.BodyPath)
	if err != nil {
		return res, nil, err
	}
	return res, f, nil
}

func (i *Insomnium) getOrCreateSettings(ctx context.Context) (s Settings, err error) {
	settings, err := i.settings.All(ctx)
	if err != nil {
		return s, err
	}

	if len(settings) == 0 {
		return s, fmt.Errorf("insomnium package does not create settings yet: %w", errors.ErrUnsupported)
	}

	return settings[0], nil
}

func (i *Insomnium) createDatabases(ctx context.Context) error {

	var err error

	file := filepath.Join(i.path, i.prefix+"Workspace"+i.suffix)
	i.workspaces, err = NewRepository[Workspace](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"RequestGroup"+i.suffix)
	i.requestGroups, err = NewRepository[RequestGroup](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"Response"+i.suffix)
	i.responses, err = NewRepository[Response](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"Project"+i.suffix)
	i.projects, err = NewRepository[Project](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"Request"+i.suffix)
	i.requests, err = NewRepository[Request](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"Environment"+i.suffix)
	i.environments, err = NewRepository[Environment](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"WorkspaceMeta"+i.suffix)
	i.workspaceMetas, err = NewRepository[WorkspaceMeta](ctx, file, false, 16)
	if err != nil {
		return err
	}

	file = filepath.Join(i.path, i.prefix+"Settings"+i.suffix)
	i.settings, err = NewRepository[Settings](ctx, file, false, 16)
	if err != nil {
		return err
	}

	return nil
}

// Load reads all of the db files and stores the results. It returns the list of
// errors if it fails in reading or unmarshaling any of the files.
func (i *Insomnium) Load() (err error) {
	var errs []error
	if err = i.workspaces.Load(); err != nil {
		errs = append(errs, err)
	}
	if err = i.requestGroups.Load(); err != nil {
		errs = append(errs, err)
	}
	if err = i.responses.Load(); err != nil {
		errs = append(errs, err)
	}
	if err = i.projects.Load(); err != nil {
		errs = append(errs, err)
	}
	if err = i.requests.Load(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return
}

// DoRequest does a request, stores the response and returns the response ID.
func (i *Insomnium) DoRequest(ctx context.Context, id string) (string, error) {

	ref, err := i.requests.Get(ctx, id)
	if err != nil {
		return "", err
	}

	wrk, err := i.GetRequestWorkspace(ctx, ref.ParentID)
	if err != nil {
		return "", err
	}

	meta, err := i.GetChildWorkspaceMeta(ctx, wrk)
	if err != nil {
		return "", err
	}

	req, err := i.createRequest(ctx, ref, wrk, meta.ActiveEnvironmentID)
	if err != nil {
		return "", err
	}

	now := time.Now()

	res, err := i.doRequest(ctx, req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	then := time.Now()

	random, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	file := filepath.Join(i.responsesDir, random.String()+".response")

	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	written, err := io.Copy(f, res.Body)
	if err != nil {
		return "", err
	}

	resID, err := i.generateID("res")
	if err != nil {
		return "", err
	}

	obj := Response{
		ID:            resID,
		ParentID:      id,
		URL:           res.Request.URL.String(),
		Modified:      int(now.Unix()),
		StatusCode:    res.StatusCode,
		ElapsedTime:   float64(then.Sub(now)) / float64(time.Second),
		ContentType:   res.Header.Get("Content-Type"),
		Created:       int(then.Unix()),
		BytesContent:  int(written),
		Headers:       make([]Header, 0, len(res.Header)),
		EnvironmentID: meta.ActiveEnvironmentID,
		BodyPath:      file,
	}

	for k, values := range res.Header {
		for _, v := range values {
			obj.Headers = append(obj.Headers, Header{Name: k, Value: v})
		}
	}

	err = i.responses.Create(ctx, obj)
	if err != nil {
		return "", err
	}

	return resID, nil
}

func (i *Insomnium) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {

	settings, err := i.getOrCreateSettings(ctx)
	if err != nil {
		return nil, err
	}

	i.transport.TLSClientConfig.InsecureSkipVerify = !settings.ValidateSSL

	return i.cl.Do(req)
}

func (i *Insomnium) createRequest(ctx context.Context, ref Request, wrk string, env string) (*http.Request, error) {

	// creating a template for the URL
	tplEnvs, err := i.getEnvs(ctx, wrk, env)
	if err != nil {
		return nil, err
	}

	urlTpl, err := gonja.FromString(ref.URL)
	if err != nil {
		return nil, err
	}

	// applying the template
	u, err := urlTpl.Execute(gonja.Context{"_": tplEnvs})
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(ref.Body.Text)

	req, err := http.NewRequestWithContext(ctx, ref.Method, u, body)
	if err != nil {
		return nil, err
	}

	if err = i.setAuth(req, ref); err != nil {
		return nil, err
	}

	return req, nil

}

func (i *Insomnium) getEnvs(ctx context.Context, key, activeEnv string) (gonja.Context, error) {
	envs, err := i.environments.GetChildren(ctx, key)
	if err != nil {
		return nil, err
	}
	var tplEnvs gonja.Context
	if len(envs) == 1 {
		curr, err := i.environments.Get(ctx, activeEnv)
		if err != nil {
			return nil, err
		}

		tplEnvs = make(gonja.Context, len(envs[0].Data))
		for k, v := range envs[0].Data {
			tplEnvs[k] = v
		}

		for k, v := range curr.Data {
			tplEnvs[k] = v
		}
	}
	return tplEnvs, nil
}

func (i *Insomnium) setAuth(req *http.Request, ref Request) error {
	// TODO: add these:
	// export const AUTH_API_KEY = 'apikey';
	// export const AUTH_OAUTH_2 = 'oauth2';
	// export const AUTH_OAUTH_1 = 'oauth1';
	// export const AUTH_DIGEST = 'digest';
	// export const AUTH_BEARER = 'bearer';
	// export const AUTH_NTLM = 'ntlm';
	// export const AUTH_HAWK = 'hawk';
	// export const AUTH_AWS_IAM = 'iam';
	// export const AUTH_NETRC = 'netrc';
	// export const AUTH_ASAP = 'asap';
	// export const HAWK_ALGORITHM_SHA256 = 'sha256';
	// export const HAWK_ALGORITHM_SHA1 = 'sha1';

	auth := ref.Authentication
	switch ref.Authentication.Type {
	case "basic":
		req.SetBasicAuth(auth.Username, auth.Password)
	case "none":
		return nil
	default:
		return fmt.Errorf("for now, we only accept basic or nothing: %w", errors.ErrUnsupported)
	}

	return nil
}

func (i *Insomnium) generateID(prefix string) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	b := strings.ReplaceAll(id.String(), "-", "")
	if prefix != "" {
		return prefix + "_" + string(b), nil
	}
	return string(b), nil
}
