package insomnium

import (
	"bufio"
	"errors"
	"io"
	"path/filepath"
)

// NewInsomnium creates an empty instance of *Insomnium, a type that should
// store information from the Insomnium db files. It loads a file reader
// interface that uses the package os to read the insomnia.*.db files and a json
// unmarshaler that uses the standard package encoding/json. Options can be used
// to mock the file reading and unmarshaling process.
func NewInsomnium(path string, options ...Option) *Insomnium {
	i := Insomnium{
		path:        path,
		fileReader:  NewDefaultFileReader(),
		unmarshaler: NewDefaultUnmarshaler(),
	}
	for _, option := range options {
		option(&i)
	}
	return &i
}

// Insomnium loads and stores information read from the insomnia.*.db files.
type Insomnium struct {
	path          string
	unmarshaler   Unmarshaler
	fileReader    FileReader
	Workspaces    []Workspace
	RequestGroups []RequestGroup
	Responses     []Response
	Projects      []Project
	Requests      []Request
}

// Load reads all of the db files and stores the results. It returns the list of
// errors if it fails in reading or unmarshaling any of the files.
func (i *Insomnium) Load() (err error) {
	var errs []error
	if err = i.LoadWorkspaces(); err != nil {
		errs = append(errs, err)
	}
	if err = i.LoadRequestGroups(); err != nil {
		errs = append(errs, err)
	}
	if err = i.LoadResponses(); err != nil {
		errs = append(errs, err)
	}
	if err = i.LoadProjects(); err != nil {
		errs = append(errs, err)
	}
	if err = i.LoadRequests(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return
}

// LoadWorkspaces loads and unmarshals all of the records in the
// insomnia.Workspaces.db file or retorns an error.
func (i *Insomnium) LoadWorkspaces() (err error) {
	path := filepath.Join(i.path, "insomnia.Workspace.db")
	f, err := i.fileReader.ReadFile(path)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	var Workspaces []Workspace
	for r.Scan() {
		if err = r.Err(); err != nil {
			return
		}
		var w Workspace
		err = i.unmarshaler.Unmarshal(r.Bytes(), &w)
		if err != nil {
			return
		}
		Workspaces = append(Workspaces, w)
	}
	i.Workspaces = Workspaces
	return
}

// LoadRequestGroups loads and unmarshals all of the records in the
// insomnia.RequestGroups.db file or retorns an error.
func (i *Insomnium) LoadRequestGroups() (err error) {
	path := filepath.Join(i.path, "insomnia.RequestGroup.db")
	f, err := i.fileReader.ReadFile(path)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	var RequestGroups []RequestGroup
	for r.Scan() {
		if err = r.Err(); err != nil {
			return
		}
		var w RequestGroup
		err = i.unmarshaler.Unmarshal(r.Bytes(), &w)
		if err != nil {
			return
		}
		RequestGroups = append(RequestGroups, w)
	}
	i.RequestGroups = RequestGroups
	return
}

// LoadResponses loads and unmarshals all of the records in the
// insomnia.Responses.db file or retorns an error.
func (i *Insomnium) LoadResponses() (err error) {
	path := filepath.Join(i.path, "insomnia.Response.db")
	f, err := i.fileReader.ReadFile(path)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	var Responses []Response
	for r.Scan() {
		if err = r.Err(); err != nil {
			return
		}
		var w Response
		err = i.unmarshaler.Unmarshal(r.Bytes(), &w)
		if err != nil {
			return
		}
		Responses = append(Responses, w)
	}
	i.Responses = Responses
	return
}

// LoadProjects loads and unmarshals all of the records in the
// insomnia.Projects.db file or retorns an error.
func (i *Insomnium) LoadProjects() (err error) {
	path := filepath.Join(i.path, "insomnia.Project.db")
	f, err := i.fileReader.ReadFile(path)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	var Projects []Project
	for r.Scan() {
		if err = r.Err(); err != nil {
			return
		}
		var w Project
		err = i.unmarshaler.Unmarshal(r.Bytes(), &w)
		if err != nil {
			return
		}
		Projects = append(Projects, w)
	}
	i.Projects = Projects
	return
}

// LoadRequests loads and unmarshals all of the records in the
// insomnia.Requests.db file or retorns an error.
func (i *Insomnium) LoadRequests() (err error) {
	path := filepath.Join(i.path, "insomnia.Request.db")
	f, err := i.fileReader.ReadFile(path)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewScanner(f)

	var Requests []Request
	for r.Scan() {
		if err = r.Err(); err != nil {
			return
		}
		var w Request
		err = i.unmarshaler.Unmarshal(r.Bytes(), &w)
		if err != nil {
			return
		}
		Requests = append(Requests, w)
	}
	i.Requests = Requests
	return
}

// Unmarshaler is used to decode each line in the Insomnium db files.
type Unmarshaler interface {
	Unmarshal(data []byte, v any) error
}

// FileReader is used to access the db files content from os.
type FileReader interface {
	ReadFile(path string) (io.ReadCloser, error)
}
