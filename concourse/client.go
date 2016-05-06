package concourse

import (
	"io"
	"net/http"

	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse/internal"
)

//go:generate counterfeiter . Client

type Client interface {
	URL() string
	HTTPClient() *http.Client

	Builds(Page) ([]atc.Build, Pagination, error)
	Build(buildID string) (atc.Build, bool, error)
	BuildEvents(buildID string) (Events, error)
	BuildResources(buildID int) (atc.BuildInputsOutputs, bool, error)
	AbortBuild(buildID string) error
	BuildInputsForJob(pipelineName string, jobName string) ([]atc.BuildInput, bool, error)
	CreateBuild(plan atc.Plan) (atc.Build, error)
	CreateJobBuild(pipelineName string, jobName string) (atc.Build, error)
	BuildPlan(buildID int) (atc.PublicBuildPlan, bool, error)
	CreateOrUpdatePipelineConfig(pipelineName string, configVersion string, passedConfig atc.Config) (bool, bool, []ConfigWarning, error)
	CreatePipe() (atc.Pipe, error)
	DeletePipeline(pipelineName string) (bool, error)
	PausePipeline(pipelineName string) (bool, error)
	UnpausePipeline(pipelineName string) (bool, error)
	RenamePipeline(pipelineName, name string) (bool, error)
	Job(pipelineName, jobName string) (atc.Job, bool, error)
	JobBuild(pipelineName, jobName, buildName string) (atc.Build, bool, error)
	JobBuilds(pipelineName string, jobName string, page Page) ([]atc.Build, Pagination, bool, error)
	PauseJob(pipelineName string, jobName string) (bool, error)
	UnpauseJob(pipelineName string, jobName string) (bool, error)
	ListContainers(queryList map[string]string) ([]atc.Container, error)
	ListPipelines() ([]atc.Pipeline, error)
	ListVolumes() ([]atc.Volume, error)
	ListWorkers() ([]atc.Worker, error)
	GetInfo() (atc.Info, error)
	PipelineConfig(pipelineName string) (atc.Config, atc.RawConfig, string, bool, error)
	GetCLIReader(arch, platform string) (io.ReadCloser, error)
	ListAuthMethods() ([]atc.AuthMethod, error)
	AuthToken() (atc.AuthToken, error)
	Pipeline(name string) (atc.Pipeline, bool, error)
	Resource(pipelineName string, resourceName string) (atc.Resource, bool, error)
	ResourceVersions(pipelineName string, resourceName string, page Page) ([]atc.VersionedResource, Pagination, bool, error)
	CheckResource(pipelineName string, resourceName string, version atc.Version) (bool, error)
	DeleteResourceVersion(pipelineName string, resourceName string, id int) (bool, error)

	BuildsWithVersionAsInput(pipelineName string, resourceName string, resourceVersionID int) ([]atc.Build, bool, error)
	BuildsWithVersionAsOutput(pipelineName string, resourceName string, resourceVersionID int) ([]atc.Build, bool, error)

	SetTeam(teamName string, team atc.Team) (atc.Team, bool, bool, error)
}

type client struct {
	connection internal.Connection
}

func NewClient(apiURL string, httpClient *http.Client) Client {
	return &client{connection: internal.NewConnection(apiURL, httpClient)}
}

func (client *client) URL() string {
	return client.connection.URL()
}

func (client *client) HTTPClient() *http.Client {
	return client.connection.HTTPClient()
}
