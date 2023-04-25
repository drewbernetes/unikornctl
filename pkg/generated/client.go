package generated

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

const (
	HttpBasicAuthenticationScopes = "httpBasicAuthentication.Scopes"
	Oauth2AuthenticationScopes    = "oauth2Authentication.Scopes"
)

// Defines values for Oauth2ErrorError.
const (
	AccessDenied            Oauth2ErrorError = "access_denied"
	InvalidRequest          Oauth2ErrorError = "invalid_request"
	InvalidScope            Oauth2ErrorError = "invalid_scope"
	MethodNotAllowed        Oauth2ErrorError = "method_not_allowed"
	NotFound                Oauth2ErrorError = "not_found"
	ServerError             Oauth2ErrorError = "server_error"
	TemporarilyUnavailable  Oauth2ErrorError = "temporarily_unavailable"
	UnauthorizedClient      Oauth2ErrorError = "unauthorized_client"
	UnsupportedMediaType    Oauth2ErrorError = "unsupported_media_type"
	UnsupportedResponseType Oauth2ErrorError = "unsupported_response_type"
)

// ApplicationBundle A bundle of applications.
type ApplicationBundle struct {
	// EndOfLife When the bundle is end-of-life.
	EndOfLife *time.Time `json:"endOfLife,omitempty"`

	// Name The resource name.
	Name string `json:"name"`

	// Preview Whether the bundle is in preview.
	Preview *bool `json:"preview,omitempty"`

	// Version The bundle version.
	Version string `json:"version"`
}

// ApplicationBundleAutoUpgrade When specified enables auto upgrade of application bundles.
type ApplicationBundleAutoUpgrade struct {
	// DaysOfWeek Days of the week and time windows that permit operations to be performed in.
	DaysOfWeek *AutoUpgradeDaysOfWeek `json:"daysOfWeek,omitempty"`
}

// ApplicationBundles A list of application bundles.
type ApplicationBundles = []ApplicationBundle

// ApplicationCredentialOptions Openstack application credential create options.
type ApplicationCredentialOptions struct {
	// Name Application credential name.
	Name string `json:"name"`
}

// AutoUpgradeDaysOfWeek Days of the week and time windows that permit operations to be performed in.
type AutoUpgradeDaysOfWeek struct {
	// Friday A time window that wraps into the next day if required.
	Friday *TimeWindow `json:"friday,omitempty"`

	// Monday A time window that wraps into the next day if required.
	Monday *TimeWindow `json:"monday,omitempty"`

	// Saturday A time window that wraps into the next day if required.
	Saturday *TimeWindow `json:"saturday,omitempty"`

	// Sunday A time window that wraps into the next day if required.
	Sunday *TimeWindow `json:"sunday,omitempty"`

	// Thursday A time window that wraps into the next day if required.
	Thursday *TimeWindow `json:"thursday,omitempty"`

	// Tuesday A time window that wraps into the next day if required.
	Tuesday *TimeWindow `json:"tuesday,omitempty"`

	// Wednesday A time window that wraps into the next day if required.
	Wednesday *TimeWindow `json:"wednesday,omitempty"`
}

// ControlPlane A Unikorn control plane.
type ControlPlane struct {
	// ApplicationBundle A bundle of applications.
	ApplicationBundle ApplicationBundle `json:"applicationBundle"`

	// ApplicationBundleAutoUpgrade When specified enables auto upgrade of application bundles.
	ApplicationBundleAutoUpgrade *ApplicationBundleAutoUpgrade `json:"applicationBundleAutoUpgrade,omitempty"`

	// Name The name of the resource.
	Name string `json:"name"`

	// Status A Kubernetes resource status.
	Status *KubernetesResourceStatus `json:"status,omitempty"`
}

// ControlPlanes A list of Unikorn control planes.
type ControlPlanes = []ControlPlane

// Hour An hour of the day in UTC.
type Hour = int

// KubernetesCluster Unikorn Kubernetes cluster creation parameters.
type KubernetesCluster struct {
	// Api Kubernetes API settings.
	Api *KubernetesClusterAPI `json:"api,omitempty"`

	// ApplicationBundle A bundle of applications.
	ApplicationBundle ApplicationBundle `json:"applicationBundle"`

	// ApplicationBundleAutoUpgrade When specified enables auto upgrade of application bundles.
	ApplicationBundleAutoUpgrade *ApplicationBundleAutoUpgrade `json:"applicationBundleAutoUpgrade,omitempty"`

	// ControlPlane A Kubernetes cluster machine.
	ControlPlane OpenstackMachinePool `json:"controlPlane"`

	// Features A set of optional add on features for the cluster.
	Features *KubernetesClusterFeatures `json:"features,omitempty"`

	// Name Cluster name.
	Name string `json:"name"`

	// Network A kubernetes cluster network settings.
	Network KubernetesClusterNetwork `json:"network"`

	// Openstack Unikorn Kubernetes cluster creation Openstack parameters.
	Openstack KubernetesClusterOpenstack `json:"openstack"`

	// Status A Kubernetes resource status.
	Status *KubernetesResourceStatus `json:"status,omitempty"`

	// WorkloadPools A non-empty list of Kubernetes cluster workload pools.
	WorkloadPools KubernetesClusterWorkloadPools `json:"workloadPools"`
}

// KubernetesClusterAPI Kubernetes API settings.
type KubernetesClusterAPI struct {
	// AllowedPrefixes Set of address prefixes to allow access to the Kubernetes API.
	AllowedPrefixes *[]string `json:"allowedPrefixes,omitempty"`

	// SubjectAlternativeNames Set of non-standard X.509 SANs to add to the API certificate.
	SubjectAlternativeNames *[]string `json:"subjectAlternativeNames,omitempty"`
}

// KubernetesClusterAutoscaling A Kubernetes cluster workload pool autoscaling configuration.
type KubernetesClusterAutoscaling struct {
	// MaximumReplicas The maximum number of replicas to allow.
	MaximumReplicas int `json:"maximumReplicas"`

	// MinimumReplicas The minimum number of replicas to allow.
	MinimumReplicas int `json:"minimumReplicas"`
}

// KubernetesClusterFeatures A set of optional add on features for the cluster.
type KubernetesClusterFeatures struct {
	// Autoscaling Enable auto-scaling.
	Autoscaling *bool `json:"autoscaling,omitempty"`

	// Ingress Enable an ingress controller.
	Ingress *bool `json:"ingress,omitempty"`
}

// KubernetesClusterNetwork A kubernetes cluster network settings.
type KubernetesClusterNetwork struct {
	// DnsNameservers A list of DNS name server to use.
	DnsNameservers []string `json:"dnsNameservers"`

	// NodePrefix Network prefix to provision nodes in.
	NodePrefix string `json:"nodePrefix"`

	// PodPrefix Network prefix to provision pods in.
	PodPrefix string `json:"podPrefix"`

	// ServicePrefix Network prefix to provision services in.
	ServicePrefix string `json:"servicePrefix"`
}

// KubernetesClusterOpenstack Unikorn Kubernetes cluster creation Openstack parameters.
type KubernetesClusterOpenstack struct {
	// ApplicationCredentialID Application credential ID.
	ApplicationCredentialID string `json:"applicationCredentialID"`

	// ApplicationCredentialSecret Application credential secret.
	ApplicationCredentialSecret string `json:"applicationCredentialSecret"`

	// ComputeAvailabilityZone Compute availability zone for control plane, and workload pool default.
	ComputeAvailabilityZone string `json:"computeAvailabilityZone"`

	// ExternalNetworkID Openstack external network ID.
	ExternalNetworkID string `json:"externalNetworkID"`

	// SshKeyName Openstack SSH Key to install on all machines.
	SshKeyName *string `json:"sshKeyName,omitempty"`

	// VolumeAvailabilityZone Volume availability zone for control plane, and workload pool default.
	VolumeAvailabilityZone string `json:"volumeAvailabilityZone"`
}

// KubernetesClusterWorkloadPool A Kuberntes cluster workload pool.
type KubernetesClusterWorkloadPool struct {
	// Autoscaling A Kubernetes cluster workload pool autoscaling configuration.
	Autoscaling *KubernetesClusterAutoscaling `json:"autoscaling,omitempty"`

	// AvailabilityZone Workload pool availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Labels Workload pool labels to apply on node creation.
	Labels *map[string]string `json:"labels,omitempty"`

	// Machine A Kubernetes cluster machine.
	Machine OpenstackMachinePool `json:"machine"`

	// Name Workload pool name.
	Name string `json:"name"`
}

// KubernetesClusterWorkloadPools A non-empty list of Kubernetes cluster workload pools.
type KubernetesClusterWorkloadPools = []KubernetesClusterWorkloadPool

// KubernetesClusters A list of Unikorn Kubernetes clusters.
type KubernetesClusters = []KubernetesCluster

// KubernetesResourceStatus A Kubernetes resource status.
type KubernetesResourceStatus struct {
	// CreationTime The time the resource was created.
	CreationTime time.Time `json:"creationTime"`

	// DeletionTime The time a control plane was deleted.
	DeletionTime *time.Time `json:"deletionTime,omitempty"`

	// Name The name of the resource.
	Name string `json:"name"`

	// Status The current status of the resource.
	Status string `json:"status"`
}

// Oauth2Error Generic error message.
type Oauth2Error struct {
	// Error A terse error string expaning on the HTTP error code.
	Error Oauth2ErrorError `json:"error"`

	// ErrorDescription Verbose message describing the error.
	ErrorDescription string `json:"error_description"`
}

// Oauth2ErrorError A terse error string expaning on the HTTP error code.
type Oauth2ErrorError string

// OpenstackApplicationCredential An Openstack application credential.
type OpenstackApplicationCredential struct {
	// Id Application credential ID.
	Id string `json:"id"`

	// Name Application credential name.
	Name string `json:"name"`

	// Secret Application credential secret, this is only present on creation.
	Secret *string `json:"secret,omitempty"`
}

// OpenstackAvailabilityZone An Openstack availability zone.
type OpenstackAvailabilityZone struct {
	// Name The availability zone name.
	Name string `json:"name"`
}

// OpenstackAvailabilityZones A list of Openstack availability zones.
type OpenstackAvailabilityZones = []OpenstackAvailabilityZone

// OpenstackExternalNetwork An Openstack external network.
type OpenstackExternalNetwork struct {
	// Id Openstack external network ID.
	Id string `json:"id"`

	// Name Opestack external network name.
	Name string `json:"name"`
}

// OpenstackExternalNetworks A list of Openstack external networks.
type OpenstackExternalNetworks = []OpenstackExternalNetwork

// OpenstackFlavor An Openstack flavor.
type OpenstackFlavor struct {
	// Cpus The number of CPUs.
	Cpus int `json:"cpus"`

	// Gpus The number of GPUs, if not set there are none.
	Gpus *int `json:"gpus,omitempty"`

	// Id The unique flavor ID.
	Id string `json:"id"`

	// Memory The amount of memory in GiB.
	Memory int `json:"memory"`

	// Name The flavor name.
	Name string `json:"name"`
}

// OpenstackFlavors A list of Openstack flavors.
type OpenstackFlavors = []OpenstackFlavor

// OpenstackImage And Openstack image.
type OpenstackImage struct {
	// Created Time when the image was created.
	Created time.Time `json:"created"`

	// Id The unique image ID.
	Id string `json:"id"`

	// Modified Time when the image was last modified.
	Modified time.Time `json:"modified"`

	// Name The image name.
	Name string `json:"name"`

	// Versions Image version metadata.
	Versions struct {
		// Kubernetes The kubernetes semantic version.
		Kubernetes string `json:"kubernetes"`

		// NvidiaDriver The nvidia driver version.
		NvidiaDriver string `json:"nvidiaDriver"`
	} `json:"versions"`
}

// OpenstackImages A list of Openstack images that are compatible with this platform.
type OpenstackImages = []OpenstackImage

// OpenstackKeyPair An Openstack key pair.
type OpenstackKeyPair struct {
	// Name The key pair name.
	Name string `json:"name"`
}

// OpenstackKeyPairs A list of Openstack key pairs.
type OpenstackKeyPairs = []OpenstackKeyPair

// OpenstackMachinePool A Kubernetes cluster machine.
type OpenstackMachinePool struct {
	// Disk An Openstack volume.
	Disk *OpenstackVolume `json:"disk,omitempty"`

	// FlavorName Openstack flavor name.
	FlavorName string `json:"flavorName"`

	// ImageName Openstack image name.
	ImageName string `json:"imageName"`

	// Replicas Number of machines.
	Replicas int `json:"replicas"`

	// Version Kubernetes version.
	Version string `json:"version"`
}

// OpenstackProject An Openstack project.
type OpenstackProject struct {
	// Description A verbose description of the project.
	Description *string `json:"description,omitempty"`

	// Id Globally unique project ID.
	Id string `json:"id"`

	// Name The name of the project within the scope of the domain.
	Name string `json:"name"`
}

// OpenstackProjects A list of Openstack projects.
type OpenstackProjects = []OpenstackProject

// OpenstackVolume An Openstack volume.
type OpenstackVolume struct {
	// AvailabilityZone Volume availability zone.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Size Disk size in GiB.
	Size int `json:"size"`
}

// Project A Unikorn project.
type Project struct {
	// Status A Kubernetes resource status.
	Status *KubernetesResourceStatus `json:"status,omitempty"`
}

// StringParameter A basic string parameter.
type StringParameter = string

// TimeWindow A time window that wraps into the next day if required.
type TimeWindow struct {
	// End An hour of the day in UTC.
	End Hour `json:"end"`

	// Start An hour of the day in UTC.
	Start Hour `json:"start"`
}

// Token Authentication token result.
type Token struct {
	// Email The user's email address.
	Email string `json:"email"`

	// Token Authentication token.
	Token string `json:"token"`
}

// TokenScope Password authentication scope.
type TokenScope struct {
	// Project Openstack project scoping.
	Project struct {
		// Id Openstack project ID.
		Id string `json:"id"`
	} `json:"project"`
}

// ApplicationCredentialParameter A basic string parameter.
type ApplicationCredentialParameter = StringParameter

// ClusterNameParameter A basic string parameter.
type ClusterNameParameter = StringParameter

// ControlPlaneNameParameter A basic string parameter.
type ControlPlaneNameParameter = StringParameter

// ApplicationBundleResponse A list of application bundles.
type ApplicationBundleResponse = ApplicationBundles

// BadRequestResponse Generic error message.
type BadRequestResponse = Oauth2Error

// ControlPlaneResponse A Unikorn control plane.
type ControlPlaneResponse = ControlPlane

// ControlPlanesResponse A list of Unikorn control planes.
type ControlPlanesResponse = ControlPlanes

// InternalServerErrorResponse Generic error message.
type InternalServerErrorResponse = Oauth2Error

// KubernetesClusterResponse Unikorn Kubernetes cluster creation parameters.
type KubernetesClusterResponse = KubernetesCluster

// KubernetesClustersResponse A list of Unikorn Kubernetes clusters.
type KubernetesClustersResponse = KubernetesClusters

// NullResponse defines model for nullResponse.
type NullResponse = map[string]interface{}

// OpenstackApplicationCredentialResponse An Openstack application credential.
type OpenstackApplicationCredentialResponse = OpenstackApplicationCredential

// OpenstackBlockStorageAvailabilityZonesResponse A list of Openstack availability zones.
type OpenstackBlockStorageAvailabilityZonesResponse = OpenstackAvailabilityZones

// OpenstackComputeAvailabilityZonesResponse A list of Openstack availability zones.
type OpenstackComputeAvailabilityZonesResponse = OpenstackAvailabilityZones

// OpenstackExternalNetworksResponse A list of Openstack external networks.
type OpenstackExternalNetworksResponse = OpenstackExternalNetworks

// OpenstackFlavorsResponse A list of Openstack flavors.
type OpenstackFlavorsResponse = OpenstackFlavors

// OpenstackImagesResponse A list of Openstack images that are compatible with this platform.
type OpenstackImagesResponse = OpenstackImages

// OpenstackKeyPairsResponse A list of Openstack key pairs.
type OpenstackKeyPairsResponse = OpenstackKeyPairs

// OpenstackProjectsResponse A list of Openstack projects.
type OpenstackProjectsResponse = OpenstackProjects

// ProjectResponse A Unikorn project.
type ProjectResponse = Project

// TokenResponse Authentication token result.
type TokenResponse = Token

// UnauthorizedResponse Generic error message.
type UnauthorizedResponse = Oauth2Error

// ApplicationCredentialRequest Openstack application credential create options.
type ApplicationCredentialRequest = ApplicationCredentialOptions

// CreateControlPlaneRequest A Unikorn control plane.
type CreateControlPlaneRequest = ControlPlane

// CreateKubernetesClusterRequest Unikorn Kubernetes cluster creation parameters.
type CreateKubernetesClusterRequest = KubernetesCluster

// TokenScopeRequest Password authentication scope.
type TokenScopeRequest = TokenScope

// PostApiV1AuthTokensTokenJSONRequestBody defines body for PostApiV1AuthTokensToken for application/json ContentType.
type PostApiV1AuthTokensTokenJSONRequestBody = TokenScope

// PostApiV1ControlplanesJSONRequestBody defines body for PostApiV1Controlplanes for application/json ContentType.
type PostApiV1ControlplanesJSONRequestBody = ControlPlane

// PutApiV1ControlplanesControlPlaneNameJSONRequestBody defines body for PutApiV1ControlplanesControlPlaneName for application/json ContentType.
type PutApiV1ControlplanesControlPlaneNameJSONRequestBody = ControlPlane

// PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody defines body for PostApiV1ControlplanesControlPlaneNameClusters for application/json ContentType.
type PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody = KubernetesCluster

// PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody defines body for PutApiV1ControlplanesControlPlaneNameClustersClusterName for application/json ContentType.
type PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody = KubernetesCluster

// PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody defines body for PostApiV1ProvidersOpenstackApplicationCredentials for application/json ContentType.
type PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody = ApplicationCredentialOptions

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetApiV1ApplicationBundlesCluster request
	GetApiV1ApplicationBundlesCluster(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ApplicationBundlesControlPlane request
	GetApiV1ApplicationBundlesControlPlane(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1AuthTokensPassword request
	PostApiV1AuthTokensPassword(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1AuthTokensToken request with any body
	PostApiV1AuthTokensTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV1AuthTokensToken(ctx context.Context, body PostApiV1AuthTokensTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1Controlplanes request
	GetApiV1Controlplanes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1Controlplanes request with any body
	PostApiV1ControlplanesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV1Controlplanes(ctx context.Context, body PostApiV1ControlplanesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV1ControlplanesControlPlaneName request
	DeleteApiV1ControlplanesControlPlaneName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ControlplanesControlPlaneName request
	GetApiV1ControlplanesControlPlaneName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV1ControlplanesControlPlaneName request with any body
	PutApiV1ControlplanesControlPlaneNameWithBody(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV1ControlplanesControlPlaneName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PutApiV1ControlplanesControlPlaneNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ControlplanesControlPlaneNameClusters request
	GetApiV1ControlplanesControlPlaneNameClusters(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1ControlplanesControlPlaneNameClusters request with any body
	PostApiV1ControlplanesControlPlaneNameClustersWithBody(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV1ControlplanesControlPlaneNameClusters(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV1ControlplanesControlPlaneNameClustersClusterName request
	DeleteApiV1ControlplanesControlPlaneNameClustersClusterName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ControlplanesControlPlaneNameClustersClusterName request
	GetApiV1ControlplanesControlPlaneNameClustersClusterName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutApiV1ControlplanesControlPlaneNameClustersClusterName request with any body
	PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithBody(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutApiV1ControlplanesControlPlaneNameClustersClusterName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, body PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig request
	GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV1Project request
	DeleteApiV1Project(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1Project request
	GetApiV1Project(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1Project request
	PostApiV1Project(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostApiV1ProvidersOpenstackApplicationCredentials request with any body
	PostApiV1ProvidersOpenstackApplicationCredentialsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostApiV1ProvidersOpenstackApplicationCredentials(ctx context.Context, body PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential request
	DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential request
	GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorage request
	GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorage(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackAvailabilityZonesCompute request
	GetApiV1ProvidersOpenstackAvailabilityZonesCompute(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackExternalNetworks request
	GetApiV1ProvidersOpenstackExternalNetworks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackFlavors request
	GetApiV1ProvidersOpenstackFlavors(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackImages request
	GetApiV1ProvidersOpenstackImages(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackKeyPairs request
	GetApiV1ProvidersOpenstackKeyPairs(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1ProvidersOpenstackProjects request
	GetApiV1ProvidersOpenstackProjects(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetApiV1ApplicationBundlesCluster(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ApplicationBundlesClusterRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ApplicationBundlesControlPlane(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ApplicationBundlesControlPlaneRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1AuthTokensPassword(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1AuthTokensPasswordRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1AuthTokensTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1AuthTokensTokenRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1AuthTokensToken(ctx context.Context, body PostApiV1AuthTokensTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1AuthTokensTokenRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1Controlplanes(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ControlplanesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1ControlplanesWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ControlplanesRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1Controlplanes(ctx context.Context, body PostApiV1ControlplanesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ControlplanesRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteApiV1ControlplanesControlPlaneName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV1ControlplanesControlPlaneNameRequest(c.Server, controlPlaneName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ControlplanesControlPlaneName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ControlplanesControlPlaneNameRequest(c.Server, controlPlaneName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutApiV1ControlplanesControlPlaneNameWithBody(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV1ControlplanesControlPlaneNameRequestWithBody(c.Server, controlPlaneName, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutApiV1ControlplanesControlPlaneName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PutApiV1ControlplanesControlPlaneNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV1ControlplanesControlPlaneNameRequest(c.Server, controlPlaneName, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ControlplanesControlPlaneNameClusters(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ControlplanesControlPlaneNameClustersRequest(c.Server, controlPlaneName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1ControlplanesControlPlaneNameClustersWithBody(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ControlplanesControlPlaneNameClustersRequestWithBody(c.Server, controlPlaneName, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1ControlplanesControlPlaneNameClusters(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ControlplanesControlPlaneNameClustersRequest(c.Server, controlPlaneName, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteApiV1ControlplanesControlPlaneNameClustersClusterName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV1ControlplanesControlPlaneNameClustersClusterNameRequest(c.Server, controlPlaneName, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ControlplanesControlPlaneNameClustersClusterName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ControlplanesControlPlaneNameClustersClusterNameRequest(c.Server, controlPlaneName, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithBody(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequestWithBody(c.Server, controlPlaneName, clusterName, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutApiV1ControlplanesControlPlaneNameClustersClusterName(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, body PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequest(c.Server, controlPlaneName, clusterName, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigRequest(c.Server, controlPlaneName, clusterName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteApiV1Project(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV1ProjectRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1Project(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProjectRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1Project(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ProjectRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1ProvidersOpenstackApplicationCredentialsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ProvidersOpenstackApplicationCredentialsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostApiV1ProvidersOpenstackApplicationCredentials(ctx context.Context, body PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostApiV1ProvidersOpenstackApplicationCredentialsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialRequest(c.Server, applicationCredential)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialRequest(c.Server, applicationCredential)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorage(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackAvailabilityZonesCompute(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackAvailabilityZonesComputeRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackExternalNetworks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackExternalNetworksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackFlavors(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackFlavorsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackImages(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackImagesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackKeyPairs(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackKeyPairsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1ProvidersOpenstackProjects(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1ProvidersOpenstackProjectsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetApiV1ApplicationBundlesClusterRequest generates requests for GetApiV1ApplicationBundlesCluster
func NewGetApiV1ApplicationBundlesClusterRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/applicationBundles/cluster")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ApplicationBundlesControlPlaneRequest generates requests for GetApiV1ApplicationBundlesControlPlane
func NewGetApiV1ApplicationBundlesControlPlaneRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/applicationBundles/controlPlane")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1AuthTokensPasswordRequest generates requests for PostApiV1AuthTokensPassword
func NewPostApiV1AuthTokensPasswordRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/auth/tokens/password")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1AuthTokensTokenRequest calls the generic PostApiV1AuthTokensToken builder with application/json body
func NewPostApiV1AuthTokensTokenRequest(server string, body PostApiV1AuthTokensTokenJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV1AuthTokensTokenRequestWithBody(server, "application/json", bodyReader)
}

// NewPostApiV1AuthTokensTokenRequestWithBody generates requests for PostApiV1AuthTokensToken with any type of body
func NewPostApiV1AuthTokensTokenRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/auth/tokens/token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1ControlplanesRequest generates requests for GetApiV1Controlplanes
func NewGetApiV1ControlplanesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1ControlplanesRequest calls the generic PostApiV1Controlplanes builder with application/json body
func NewPostApiV1ControlplanesRequest(server string, body PostApiV1ControlplanesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV1ControlplanesRequestWithBody(server, "application/json", bodyReader)
}

// NewPostApiV1ControlplanesRequestWithBody generates requests for PostApiV1Controlplanes with any type of body
func NewPostApiV1ControlplanesRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteApiV1ControlplanesControlPlaneNameRequest generates requests for DeleteApiV1ControlplanesControlPlaneName
func NewDeleteApiV1ControlplanesControlPlaneNameRequest(server string, controlPlaneName ControlPlaneNameParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ControlplanesControlPlaneNameRequest generates requests for GetApiV1ControlplanesControlPlaneName
func NewGetApiV1ControlplanesControlPlaneNameRequest(server string, controlPlaneName ControlPlaneNameParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutApiV1ControlplanesControlPlaneNameRequest calls the generic PutApiV1ControlplanesControlPlaneName builder with application/json body
func NewPutApiV1ControlplanesControlPlaneNameRequest(server string, controlPlaneName ControlPlaneNameParameter, body PutApiV1ControlplanesControlPlaneNameJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV1ControlplanesControlPlaneNameRequestWithBody(server, controlPlaneName, "application/json", bodyReader)
}

// NewPutApiV1ControlplanesControlPlaneNameRequestWithBody generates requests for PutApiV1ControlplanesControlPlaneName with any type of body
func NewPutApiV1ControlplanesControlPlaneNameRequestWithBody(server string, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1ControlplanesControlPlaneNameClustersRequest generates requests for GetApiV1ControlplanesControlPlaneNameClusters
func NewGetApiV1ControlplanesControlPlaneNameClustersRequest(server string, controlPlaneName ControlPlaneNameParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s/clusters", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1ControlplanesControlPlaneNameClustersRequest calls the generic PostApiV1ControlplanesControlPlaneNameClusters builder with application/json body
func NewPostApiV1ControlplanesControlPlaneNameClustersRequest(server string, controlPlaneName ControlPlaneNameParameter, body PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV1ControlplanesControlPlaneNameClustersRequestWithBody(server, controlPlaneName, "application/json", bodyReader)
}

// NewPostApiV1ControlplanesControlPlaneNameClustersRequestWithBody generates requests for PostApiV1ControlplanesControlPlaneNameClusters with any type of body
func NewPostApiV1ControlplanesControlPlaneNameClustersRequestWithBody(server string, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s/clusters", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteApiV1ControlplanesControlPlaneNameClustersClusterNameRequest generates requests for DeleteApiV1ControlplanesControlPlaneNameClustersClusterName
func NewDeleteApiV1ControlplanesControlPlaneNameClustersClusterNameRequest(server string, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s/clusters/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ControlplanesControlPlaneNameClustersClusterNameRequest generates requests for GetApiV1ControlplanesControlPlaneNameClustersClusterName
func NewGetApiV1ControlplanesControlPlaneNameClustersClusterNameRequest(server string, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s/clusters/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequest calls the generic PutApiV1ControlplanesControlPlaneNameClustersClusterName builder with application/json body
func NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequest(server string, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, body PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequestWithBody(server, controlPlaneName, clusterName, "application/json", bodyReader)
}

// NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequestWithBody generates requests for PutApiV1ControlplanesControlPlaneNameClustersClusterName with any type of body
func NewPutApiV1ControlplanesControlPlaneNameClustersClusterNameRequestWithBody(server string, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s/clusters/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigRequest generates requests for GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig
func NewGetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigRequest(server string, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "controlPlaneName", runtime.ParamLocationPath, controlPlaneName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "clusterName", runtime.ParamLocationPath, clusterName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/controlplanes/%s/clusters/%s/kubeconfig", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteApiV1ProjectRequest generates requests for DeleteApiV1Project
func NewDeleteApiV1ProjectRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/project")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProjectRequest generates requests for GetApiV1Project
func NewGetApiV1ProjectRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/project")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1ProjectRequest generates requests for PostApiV1Project
func NewPostApiV1ProjectRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/project")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostApiV1ProvidersOpenstackApplicationCredentialsRequest calls the generic PostApiV1ProvidersOpenstackApplicationCredentials builder with application/json body
func NewPostApiV1ProvidersOpenstackApplicationCredentialsRequest(server string, body PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostApiV1ProvidersOpenstackApplicationCredentialsRequestWithBody(server, "application/json", bodyReader)
}

// NewPostApiV1ProvidersOpenstackApplicationCredentialsRequestWithBody generates requests for PostApiV1ProvidersOpenstackApplicationCredentials with any type of body
func NewPostApiV1ProvidersOpenstackApplicationCredentialsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/application-credentials")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialRequest generates requests for DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential
func NewDeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialRequest(server string, applicationCredential ApplicationCredentialParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "applicationCredential", runtime.ParamLocationPath, applicationCredential)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/application-credentials/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialRequest generates requests for GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential
func NewGetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialRequest(server string, applicationCredential ApplicationCredentialParameter) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "applicationCredential", runtime.ParamLocationPath, applicationCredential)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/application-credentials/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageRequest generates requests for GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorage
func NewGetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/availability-zones/block-storage")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackAvailabilityZonesComputeRequest generates requests for GetApiV1ProvidersOpenstackAvailabilityZonesCompute
func NewGetApiV1ProvidersOpenstackAvailabilityZonesComputeRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/availability-zones/compute")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackExternalNetworksRequest generates requests for GetApiV1ProvidersOpenstackExternalNetworks
func NewGetApiV1ProvidersOpenstackExternalNetworksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/external-networks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackFlavorsRequest generates requests for GetApiV1ProvidersOpenstackFlavors
func NewGetApiV1ProvidersOpenstackFlavorsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/flavors")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackImagesRequest generates requests for GetApiV1ProvidersOpenstackImages
func NewGetApiV1ProvidersOpenstackImagesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/images")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackKeyPairsRequest generates requests for GetApiV1ProvidersOpenstackKeyPairs
func NewGetApiV1ProvidersOpenstackKeyPairsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/key-pairs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1ProvidersOpenstackProjectsRequest generates requests for GetApiV1ProvidersOpenstackProjects
func NewGetApiV1ProvidersOpenstackProjectsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/providers/openstack/projects")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetApiV1ApplicationBundlesCluster request
	GetApiV1ApplicationBundlesClusterWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ApplicationBundlesClusterResponse, error)

	// GetApiV1ApplicationBundlesControlPlane request
	GetApiV1ApplicationBundlesControlPlaneWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ApplicationBundlesControlPlaneResponse, error)

	// PostApiV1AuthTokensPassword request
	PostApiV1AuthTokensPasswordWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV1AuthTokensPasswordResponse, error)

	// PostApiV1AuthTokensToken request with any body
	PostApiV1AuthTokensTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1AuthTokensTokenResponse, error)

	PostApiV1AuthTokensTokenWithResponse(ctx context.Context, body PostApiV1AuthTokensTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1AuthTokensTokenResponse, error)

	// GetApiV1Controlplanes request
	GetApiV1ControlplanesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesResponse, error)

	// PostApiV1Controlplanes request with any body
	PostApiV1ControlplanesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesResponse, error)

	PostApiV1ControlplanesWithResponse(ctx context.Context, body PostApiV1ControlplanesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesResponse, error)

	// DeleteApiV1ControlplanesControlPlaneName request
	DeleteApiV1ControlplanesControlPlaneNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*DeleteApiV1ControlplanesControlPlaneNameResponse, error)

	// GetApiV1ControlplanesControlPlaneName request
	GetApiV1ControlplanesControlPlaneNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameResponse, error)

	// PutApiV1ControlplanesControlPlaneName request with any body
	PutApiV1ControlplanesControlPlaneNameWithBodyWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameResponse, error)

	PutApiV1ControlplanesControlPlaneNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PutApiV1ControlplanesControlPlaneNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameResponse, error)

	// GetApiV1ControlplanesControlPlaneNameClusters request
	GetApiV1ControlplanesControlPlaneNameClustersWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameClustersResponse, error)

	// PostApiV1ControlplanesControlPlaneNameClusters request with any body
	PostApiV1ControlplanesControlPlaneNameClustersWithBodyWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesControlPlaneNameClustersResponse, error)

	PostApiV1ControlplanesControlPlaneNameClustersWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesControlPlaneNameClustersResponse, error)

	// DeleteApiV1ControlplanesControlPlaneNameClustersClusterName request
	DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error)

	// GetApiV1ControlplanesControlPlaneNameClustersClusterName request
	GetApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error)

	// PutApiV1ControlplanesControlPlaneNameClustersClusterName request with any body
	PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithBodyWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error)

	PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, body PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error)

	// GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig request
	GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse, error)

	// DeleteApiV1Project request
	DeleteApiV1ProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DeleteApiV1ProjectResponse, error)

	// GetApiV1Project request
	GetApiV1ProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProjectResponse, error)

	// PostApiV1Project request
	PostApiV1ProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV1ProjectResponse, error)

	// PostApiV1ProvidersOpenstackApplicationCredentials request with any body
	PostApiV1ProvidersOpenstackApplicationCredentialsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1ProvidersOpenstackApplicationCredentialsResponse, error)

	PostApiV1ProvidersOpenstackApplicationCredentialsWithResponse(ctx context.Context, body PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1ProvidersOpenstackApplicationCredentialsResponse, error)

	// DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential request
	DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse, error)

	// GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential request
	GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse, error)

	// GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorage request
	GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse, error)

	// GetApiV1ProvidersOpenstackAvailabilityZonesCompute request
	GetApiV1ProvidersOpenstackAvailabilityZonesComputeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse, error)

	// GetApiV1ProvidersOpenstackExternalNetworks request
	GetApiV1ProvidersOpenstackExternalNetworksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackExternalNetworksResponse, error)

	// GetApiV1ProvidersOpenstackFlavors request
	GetApiV1ProvidersOpenstackFlavorsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackFlavorsResponse, error)

	// GetApiV1ProvidersOpenstackImages request
	GetApiV1ProvidersOpenstackImagesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackImagesResponse, error)

	// GetApiV1ProvidersOpenstackKeyPairs request
	GetApiV1ProvidersOpenstackKeyPairsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackKeyPairsResponse, error)

	// GetApiV1ProvidersOpenstackProjects request
	GetApiV1ProvidersOpenstackProjectsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackProjectsResponse, error)
}

type GetApiV1ApplicationBundlesClusterResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ApplicationBundles
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ApplicationBundlesClusterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ApplicationBundlesClusterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ApplicationBundlesControlPlaneResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ApplicationBundles
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ApplicationBundlesControlPlaneResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ApplicationBundlesControlPlaneResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1AuthTokensPasswordResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Token
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostApiV1AuthTokensPasswordResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1AuthTokensPasswordResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1AuthTokensTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Token
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostApiV1AuthTokensTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1AuthTokensTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ControlplanesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ControlPlanes
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ControlplanesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ControlplanesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1ControlplanesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostApiV1ControlplanesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1ControlplanesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteApiV1ControlplanesControlPlaneNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r DeleteApiV1ControlplanesControlPlaneNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteApiV1ControlplanesControlPlaneNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ControlplanesControlPlaneNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ControlPlane
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ControlplanesControlPlaneNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ControlplanesControlPlaneNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutApiV1ControlplanesControlPlaneNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PutApiV1ControlplanesControlPlaneNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutApiV1ControlplanesControlPlaneNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ControlplanesControlPlaneNameClustersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *KubernetesClusters
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ControlplanesControlPlaneNameClustersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ControlplanesControlPlaneNameClustersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1ControlplanesControlPlaneNameClustersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostApiV1ControlplanesControlPlaneNameClustersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1ControlplanesControlPlaneNameClustersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *KubernetesCluster
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteApiV1ProjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON202      *map[string]interface{}
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r DeleteApiV1ProjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteApiV1ProjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Project
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1ProjectResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostApiV1ProjectResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1ProjectResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostApiV1ProvidersOpenstackApplicationCredentialsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackApplicationCredential
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r PostApiV1ProvidersOpenstackApplicationCredentialsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostApiV1ProvidersOpenstackApplicationCredentialsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackApplicationCredential
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackAvailabilityZones
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackAvailabilityZones
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackExternalNetworksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackExternalNetworks
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackExternalNetworksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackExternalNetworksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackFlavorsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackFlavors
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackFlavorsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackFlavorsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackImagesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackImages
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackImagesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackImagesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackKeyPairsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackKeyPairs
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackKeyPairsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackKeyPairsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1ProvidersOpenstackProjectsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OpenstackProjects
	JSON400      *Oauth2Error
	JSON401      *Oauth2Error
	JSON500      *Oauth2Error
}

// Status returns HTTPResponse.Status
func (r GetApiV1ProvidersOpenstackProjectsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1ProvidersOpenstackProjectsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetApiV1ApplicationBundlesClusterWithResponse request returning *GetApiV1ApplicationBundlesClusterResponse
func (c *ClientWithResponses) GetApiV1ApplicationBundlesClusterWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ApplicationBundlesClusterResponse, error) {
	rsp, err := c.GetApiV1ApplicationBundlesCluster(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ApplicationBundlesClusterResponse(rsp)
}

// GetApiV1ApplicationBundlesControlPlaneWithResponse request returning *GetApiV1ApplicationBundlesControlPlaneResponse
func (c *ClientWithResponses) GetApiV1ApplicationBundlesControlPlaneWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ApplicationBundlesControlPlaneResponse, error) {
	rsp, err := c.GetApiV1ApplicationBundlesControlPlane(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ApplicationBundlesControlPlaneResponse(rsp)
}

// PostApiV1AuthTokensPasswordWithResponse request returning *PostApiV1AuthTokensPasswordResponse
func (c *ClientWithResponses) PostApiV1AuthTokensPasswordWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV1AuthTokensPasswordResponse, error) {
	rsp, err := c.PostApiV1AuthTokensPassword(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1AuthTokensPasswordResponse(rsp)
}

// PostApiV1AuthTokensTokenWithBodyWithResponse request with arbitrary body returning *PostApiV1AuthTokensTokenResponse
func (c *ClientWithResponses) PostApiV1AuthTokensTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1AuthTokensTokenResponse, error) {
	rsp, err := c.PostApiV1AuthTokensTokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1AuthTokensTokenResponse(rsp)
}

func (c *ClientWithResponses) PostApiV1AuthTokensTokenWithResponse(ctx context.Context, body PostApiV1AuthTokensTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1AuthTokensTokenResponse, error) {
	rsp, err := c.PostApiV1AuthTokensToken(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1AuthTokensTokenResponse(rsp)
}

// GetApiV1ControlplanesWithResponse request returning *GetApiV1ControlplanesResponse
func (c *ClientWithResponses) GetApiV1ControlplanesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesResponse, error) {
	rsp, err := c.GetApiV1Controlplanes(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ControlplanesResponse(rsp)
}

// PostApiV1ControlplanesWithBodyWithResponse request with arbitrary body returning *PostApiV1ControlplanesResponse
func (c *ClientWithResponses) PostApiV1ControlplanesWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesResponse, error) {
	rsp, err := c.PostApiV1ControlplanesWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ControlplanesResponse(rsp)
}

func (c *ClientWithResponses) PostApiV1ControlplanesWithResponse(ctx context.Context, body PostApiV1ControlplanesJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesResponse, error) {
	rsp, err := c.PostApiV1Controlplanes(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ControlplanesResponse(rsp)
}

// DeleteApiV1ControlplanesControlPlaneNameWithResponse request returning *DeleteApiV1ControlplanesControlPlaneNameResponse
func (c *ClientWithResponses) DeleteApiV1ControlplanesControlPlaneNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*DeleteApiV1ControlplanesControlPlaneNameResponse, error) {
	rsp, err := c.DeleteApiV1ControlplanesControlPlaneName(ctx, controlPlaneName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV1ControlplanesControlPlaneNameResponse(rsp)
}

// GetApiV1ControlplanesControlPlaneNameWithResponse request returning *GetApiV1ControlplanesControlPlaneNameResponse
func (c *ClientWithResponses) GetApiV1ControlplanesControlPlaneNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameResponse, error) {
	rsp, err := c.GetApiV1ControlplanesControlPlaneName(ctx, controlPlaneName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ControlplanesControlPlaneNameResponse(rsp)
}

// PutApiV1ControlplanesControlPlaneNameWithBodyWithResponse request with arbitrary body returning *PutApiV1ControlplanesControlPlaneNameResponse
func (c *ClientWithResponses) PutApiV1ControlplanesControlPlaneNameWithBodyWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameResponse, error) {
	rsp, err := c.PutApiV1ControlplanesControlPlaneNameWithBody(ctx, controlPlaneName, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV1ControlplanesControlPlaneNameResponse(rsp)
}

func (c *ClientWithResponses) PutApiV1ControlplanesControlPlaneNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PutApiV1ControlplanesControlPlaneNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameResponse, error) {
	rsp, err := c.PutApiV1ControlplanesControlPlaneName(ctx, controlPlaneName, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV1ControlplanesControlPlaneNameResponse(rsp)
}

// GetApiV1ControlplanesControlPlaneNameClustersWithResponse request returning *GetApiV1ControlplanesControlPlaneNameClustersResponse
func (c *ClientWithResponses) GetApiV1ControlplanesControlPlaneNameClustersWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameClustersResponse, error) {
	rsp, err := c.GetApiV1ControlplanesControlPlaneNameClusters(ctx, controlPlaneName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ControlplanesControlPlaneNameClustersResponse(rsp)
}

// PostApiV1ControlplanesControlPlaneNameClustersWithBodyWithResponse request with arbitrary body returning *PostApiV1ControlplanesControlPlaneNameClustersResponse
func (c *ClientWithResponses) PostApiV1ControlplanesControlPlaneNameClustersWithBodyWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesControlPlaneNameClustersResponse, error) {
	rsp, err := c.PostApiV1ControlplanesControlPlaneNameClustersWithBody(ctx, controlPlaneName, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ControlplanesControlPlaneNameClustersResponse(rsp)
}

func (c *ClientWithResponses) PostApiV1ControlplanesControlPlaneNameClustersWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, body PostApiV1ControlplanesControlPlaneNameClustersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1ControlplanesControlPlaneNameClustersResponse, error) {
	rsp, err := c.PostApiV1ControlplanesControlPlaneNameClusters(ctx, controlPlaneName, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ControlplanesControlPlaneNameClustersResponse(rsp)
}

// DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse request returning *DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse
func (c *ClientWithResponses) DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	rsp, err := c.DeleteApiV1ControlplanesControlPlaneNameClustersClusterName(ctx, controlPlaneName, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp)
}

// GetApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse request returning *GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse
func (c *ClientWithResponses) GetApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	rsp, err := c.GetApiV1ControlplanesControlPlaneNameClustersClusterName(ctx, controlPlaneName, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp)
}

// PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithBodyWithResponse request with arbitrary body returning *PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse
func (c *ClientWithResponses) PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithBodyWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	rsp, err := c.PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithBody(ctx, controlPlaneName, clusterName, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp)
}

func (c *ClientWithResponses) PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, body PutApiV1ControlplanesControlPlaneNameClustersClusterNameJSONRequestBody, reqEditors ...RequestEditorFn) (*PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	rsp, err := c.PutApiV1ControlplanesControlPlaneNameClustersClusterName(ctx, controlPlaneName, clusterName, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp)
}

// GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigWithResponse request returning *GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse
func (c *ClientWithResponses) GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigWithResponse(ctx context.Context, controlPlaneName ControlPlaneNameParameter, clusterName ClusterNameParameter, reqEditors ...RequestEditorFn) (*GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse, error) {
	rsp, err := c.GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfig(ctx, controlPlaneName, clusterName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse(rsp)
}

// DeleteApiV1ProjectWithResponse request returning *DeleteApiV1ProjectResponse
func (c *ClientWithResponses) DeleteApiV1ProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DeleteApiV1ProjectResponse, error) {
	rsp, err := c.DeleteApiV1Project(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV1ProjectResponse(rsp)
}

// GetApiV1ProjectWithResponse request returning *GetApiV1ProjectResponse
func (c *ClientWithResponses) GetApiV1ProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProjectResponse, error) {
	rsp, err := c.GetApiV1Project(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProjectResponse(rsp)
}

// PostApiV1ProjectWithResponse request returning *PostApiV1ProjectResponse
func (c *ClientWithResponses) PostApiV1ProjectWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostApiV1ProjectResponse, error) {
	rsp, err := c.PostApiV1Project(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ProjectResponse(rsp)
}

// PostApiV1ProvidersOpenstackApplicationCredentialsWithBodyWithResponse request with arbitrary body returning *PostApiV1ProvidersOpenstackApplicationCredentialsResponse
func (c *ClientWithResponses) PostApiV1ProvidersOpenstackApplicationCredentialsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostApiV1ProvidersOpenstackApplicationCredentialsResponse, error) {
	rsp, err := c.PostApiV1ProvidersOpenstackApplicationCredentialsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ProvidersOpenstackApplicationCredentialsResponse(rsp)
}

func (c *ClientWithResponses) PostApiV1ProvidersOpenstackApplicationCredentialsWithResponse(ctx context.Context, body PostApiV1ProvidersOpenstackApplicationCredentialsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostApiV1ProvidersOpenstackApplicationCredentialsResponse, error) {
	rsp, err := c.PostApiV1ProvidersOpenstackApplicationCredentials(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostApiV1ProvidersOpenstackApplicationCredentialsResponse(rsp)
}

// DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse request returning *DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse
func (c *ClientWithResponses) DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse, error) {
	rsp, err := c.DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential(ctx, applicationCredential, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse(rsp)
}

// GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse request returning *GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse(ctx context.Context, applicationCredential ApplicationCredentialParameter, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredential(ctx, applicationCredential, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse(rsp)
}

// GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageWithResponse request returning *GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorage(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse(rsp)
}

// GetApiV1ProvidersOpenstackAvailabilityZonesComputeWithResponse request returning *GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackAvailabilityZonesComputeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackAvailabilityZonesCompute(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse(rsp)
}

// GetApiV1ProvidersOpenstackExternalNetworksWithResponse request returning *GetApiV1ProvidersOpenstackExternalNetworksResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackExternalNetworksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackExternalNetworksResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackExternalNetworks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackExternalNetworksResponse(rsp)
}

// GetApiV1ProvidersOpenstackFlavorsWithResponse request returning *GetApiV1ProvidersOpenstackFlavorsResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackFlavorsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackFlavorsResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackFlavors(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackFlavorsResponse(rsp)
}

// GetApiV1ProvidersOpenstackImagesWithResponse request returning *GetApiV1ProvidersOpenstackImagesResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackImagesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackImagesResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackImages(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackImagesResponse(rsp)
}

// GetApiV1ProvidersOpenstackKeyPairsWithResponse request returning *GetApiV1ProvidersOpenstackKeyPairsResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackKeyPairsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackKeyPairsResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackKeyPairs(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackKeyPairsResponse(rsp)
}

// GetApiV1ProvidersOpenstackProjectsWithResponse request returning *GetApiV1ProvidersOpenstackProjectsResponse
func (c *ClientWithResponses) GetApiV1ProvidersOpenstackProjectsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1ProvidersOpenstackProjectsResponse, error) {
	rsp, err := c.GetApiV1ProvidersOpenstackProjects(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1ProvidersOpenstackProjectsResponse(rsp)
}

// ParseGetApiV1ApplicationBundlesClusterResponse parses an HTTP response from a GetApiV1ApplicationBundlesClusterWithResponse call
func ParseGetApiV1ApplicationBundlesClusterResponse(rsp *http.Response) (*GetApiV1ApplicationBundlesClusterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ApplicationBundlesClusterResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ApplicationBundles
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ApplicationBundlesControlPlaneResponse parses an HTTP response from a GetApiV1ApplicationBundlesControlPlaneWithResponse call
func ParseGetApiV1ApplicationBundlesControlPlaneResponse(rsp *http.Response) (*GetApiV1ApplicationBundlesControlPlaneResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ApplicationBundlesControlPlaneResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ApplicationBundles
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1AuthTokensPasswordResponse parses an HTTP response from a PostApiV1AuthTokensPasswordWithResponse call
func ParsePostApiV1AuthTokensPasswordResponse(rsp *http.Response) (*PostApiV1AuthTokensPasswordResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1AuthTokensPasswordResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Token
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1AuthTokensTokenResponse parses an HTTP response from a PostApiV1AuthTokensTokenWithResponse call
func ParsePostApiV1AuthTokensTokenResponse(rsp *http.Response) (*PostApiV1AuthTokensTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1AuthTokensTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Token
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ControlplanesResponse parses an HTTP response from a GetApiV1ControlplanesWithResponse call
func ParseGetApiV1ControlplanesResponse(rsp *http.Response) (*GetApiV1ControlplanesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ControlplanesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ControlPlanes
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1ControlplanesResponse parses an HTTP response from a PostApiV1ControlplanesWithResponse call
func ParsePostApiV1ControlplanesResponse(rsp *http.Response) (*PostApiV1ControlplanesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1ControlplanesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteApiV1ControlplanesControlPlaneNameResponse parses an HTTP response from a DeleteApiV1ControlplanesControlPlaneNameWithResponse call
func ParseDeleteApiV1ControlplanesControlPlaneNameResponse(rsp *http.Response) (*DeleteApiV1ControlplanesControlPlaneNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV1ControlplanesControlPlaneNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ControlplanesControlPlaneNameResponse parses an HTTP response from a GetApiV1ControlplanesControlPlaneNameWithResponse call
func ParseGetApiV1ControlplanesControlPlaneNameResponse(rsp *http.Response) (*GetApiV1ControlplanesControlPlaneNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ControlplanesControlPlaneNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ControlPlane
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePutApiV1ControlplanesControlPlaneNameResponse parses an HTTP response from a PutApiV1ControlplanesControlPlaneNameWithResponse call
func ParsePutApiV1ControlplanesControlPlaneNameResponse(rsp *http.Response) (*PutApiV1ControlplanesControlPlaneNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV1ControlplanesControlPlaneNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ControlplanesControlPlaneNameClustersResponse parses an HTTP response from a GetApiV1ControlplanesControlPlaneNameClustersWithResponse call
func ParseGetApiV1ControlplanesControlPlaneNameClustersResponse(rsp *http.Response) (*GetApiV1ControlplanesControlPlaneNameClustersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ControlplanesControlPlaneNameClustersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest KubernetesClusters
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1ControlplanesControlPlaneNameClustersResponse parses an HTTP response from a PostApiV1ControlplanesControlPlaneNameClustersWithResponse call
func ParsePostApiV1ControlplanesControlPlaneNameClustersResponse(rsp *http.Response) (*PostApiV1ControlplanesControlPlaneNameClustersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1ControlplanesControlPlaneNameClustersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse parses an HTTP response from a DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse call
func ParseDeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp *http.Response) (*DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV1ControlplanesControlPlaneNameClustersClusterNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse parses an HTTP response from a GetApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse call
func ParseGetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp *http.Response) (*GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ControlplanesControlPlaneNameClustersClusterNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest KubernetesCluster
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse parses an HTTP response from a PutApiV1ControlplanesControlPlaneNameClustersClusterNameWithResponse call
func ParsePutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse(rsp *http.Response) (*PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutApiV1ControlplanesControlPlaneNameClustersClusterNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse parses an HTTP response from a GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigWithResponse call
func ParseGetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse(rsp *http.Response) (*GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ControlplanesControlPlaneNameClustersClusterNameKubeconfigResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteApiV1ProjectResponse parses an HTTP response from a DeleteApiV1ProjectWithResponse call
func ParseDeleteApiV1ProjectResponse(rsp *http.Response) (*DeleteApiV1ProjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV1ProjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 202:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON202 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProjectResponse parses an HTTP response from a GetApiV1ProjectWithResponse call
func ParseGetApiV1ProjectResponse(rsp *http.Response) (*GetApiV1ProjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Project
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1ProjectResponse parses an HTTP response from a PostApiV1ProjectWithResponse call
func ParsePostApiV1ProjectResponse(rsp *http.Response) (*PostApiV1ProjectResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1ProjectResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParsePostApiV1ProvidersOpenstackApplicationCredentialsResponse parses an HTTP response from a PostApiV1ProvidersOpenstackApplicationCredentialsWithResponse call
func ParsePostApiV1ProvidersOpenstackApplicationCredentialsResponse(rsp *http.Response) (*PostApiV1ProvidersOpenstackApplicationCredentialsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostApiV1ProvidersOpenstackApplicationCredentialsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackApplicationCredential
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse parses an HTTP response from a DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse call
func ParseDeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse(rsp *http.Response) (*DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse parses an HTTP response from a GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialWithResponse call
func ParseGetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackApplicationCredentialsApplicationCredentialResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackApplicationCredential
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse parses an HTTP response from a GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageWithResponse call
func ParseGetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackAvailabilityZonesBlockStorageResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackAvailabilityZones
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse parses an HTTP response from a GetApiV1ProvidersOpenstackAvailabilityZonesComputeWithResponse call
func ParseGetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackAvailabilityZonesComputeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackAvailabilityZones
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackExternalNetworksResponse parses an HTTP response from a GetApiV1ProvidersOpenstackExternalNetworksWithResponse call
func ParseGetApiV1ProvidersOpenstackExternalNetworksResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackExternalNetworksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackExternalNetworksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackExternalNetworks
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackFlavorsResponse parses an HTTP response from a GetApiV1ProvidersOpenstackFlavorsWithResponse call
func ParseGetApiV1ProvidersOpenstackFlavorsResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackFlavorsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackFlavorsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackFlavors
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackImagesResponse parses an HTTP response from a GetApiV1ProvidersOpenstackImagesWithResponse call
func ParseGetApiV1ProvidersOpenstackImagesResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackImagesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackImagesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackImages
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackKeyPairsResponse parses an HTTP response from a GetApiV1ProvidersOpenstackKeyPairsWithResponse call
func ParseGetApiV1ProvidersOpenstackKeyPairsResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackKeyPairsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackKeyPairsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackKeyPairs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiV1ProvidersOpenstackProjectsResponse parses an HTTP response from a GetApiV1ProvidersOpenstackProjectsWithResponse call
func ParseGetApiV1ProvidersOpenstackProjectsResponse(rsp *http.Response) (*GetApiV1ProvidersOpenstackProjectsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1ProvidersOpenstackProjectsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OpenstackProjects
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Oauth2Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
