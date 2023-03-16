package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AppliedManifestResourceMeta = map[string]string{
	"":        "AppliedManifestResourceMeta represents the group, version, resource, name and namespace of a resource. Since these resources have been created, they must have valid group, version, resource, namespace, and name.",
	"version": "Version is the version of the Kubernetes resource.",
	"uid":     "UID is set on successful deletion of the Kubernetes resource by controller. The resource might be still visible on the managed cluster after this field is set. It is not directly settable by a client.",
}

func (AppliedManifestResourceMeta) SwaggerDoc() map[string]string {
	return map_AppliedManifestResourceMeta
}

var map_AppliedManifestWork = map[string]string{
	"":       "AppliedManifestWork represents an applied manifestwork on managed cluster that is placed on a managed cluster. An AppliedManifestWork links to a manifestwork on a hub recording resources deployed in the managed cluster. When the agent is removed from managed cluster, cluster-admin on managed cluster can delete appliedmanifestwork to remove resources deployed by the agent. The name of the appliedmanifestwork must be in the format of {hash of hub's first kube-apiserver url}-{manifestwork name}",
	"spec":   "Spec represents the desired configuration of AppliedManifestWork.",
	"status": "Status represents the current status of AppliedManifestWork.",
}

func (AppliedManifestWork) SwaggerDoc() map[string]string {
	return map_AppliedManifestWork
}

var map_AppliedManifestWorkList = map[string]string{
	"":         "AppliedManifestWorkList is a collection of appliedmanifestworks.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of appliedmanifestworks.",
}

func (AppliedManifestWorkList) SwaggerDoc() map[string]string {
	return map_AppliedManifestWorkList
}

var map_AppliedManifestWorkSpec = map[string]string{
	"":                 "AppliedManifestWorkSpec represents the desired configuration of AppliedManifestWork",
	"hubHash":          "HubHash represents the hash of the first hub kube apiserver to identify which hub this AppliedManifestWork links to.",
	"agentID":          "AgentID represents the ID of the work agent who is to handle this AppliedManifestWork.",
	"manifestWorkName": "ManifestWorkName represents the name of the related manifestwork on the hub.",
}

func (AppliedManifestWorkSpec) SwaggerDoc() map[string]string {
	return map_AppliedManifestWorkSpec
}

var map_AppliedManifestWorkStatus = map[string]string{
	"":                 "AppliedManifestWorkStatus represents the current status of AppliedManifestWork",
	"appliedResources": "AppliedResources represents a list of resources defined within the manifestwork that are applied. Only resources with valid GroupVersionResource, namespace, and name are suitable. An item in this slice is deleted when there is no mapped manifest in manifestwork.Spec or by finalizer. The resource relating to the item will also be removed from managed cluster. The deleted resource may still be present until the finalizers for that resource are finished. However, the resource will not be undeleted, so it can be removed from this list and eventual consistency is preserved.",
}

func (AppliedManifestWorkStatus) SwaggerDoc() map[string]string {
	return map_AppliedManifestWorkStatus
}

var map_DeleteOption = map[string]string{
	"propagationPolicy":  "propagationPolicy can be Foreground, Orphan or SelectivelyOrphan SelectivelyOrphan should be rarely used.  It is provided for cases where particular resources is transfering ownership from one ManifestWork to another or another management unit. Setting this value will allow a flow like 1. create manifestwork/2 to manage foo 2. update manifestwork/1 to selectively orphan foo 3. remove foo from manifestwork/1 without impacting continuity because manifestwork/2 adopts it.",
	"selectivelyOrphans": "selectivelyOrphan represents a list of resources following orphan deletion stratecy",
}

func (DeleteOption) SwaggerDoc() map[string]string {
	return map_DeleteOption
}

var map_FeedbackRule = map[string]string{
	"type":      "Type defines the option of how status can be returned. It can be jsonPaths or wellKnownStatus. If the type is JSONPaths, user should specify the jsonPaths field If the type is WellKnownStatus, certain common fields of status defined by a rule only for types in in k8s.io/api and open-cluster-management/api will be reported, If these status fields do not exist, no values will be reported.",
	"jsonPaths": "JsonPaths defines the json path under status field to be synced.",
}

func (FeedbackRule) SwaggerDoc() map[string]string {
	return map_FeedbackRule
}

var map_FeedbackValue = map[string]string{
	"name":       "Name represents the alias name for this field. It is the same as what is specified in StatuFeedbackRule in the spec.",
	"fieldValue": "Value is the value of the status field. The value of the status field can only be integer, string or boolean.",
}

func (FeedbackValue) SwaggerDoc() map[string]string {
	return map_FeedbackValue
}

var map_FieldValue = map[string]string{
	"":        "FieldValue is the value of the status field. The value of the status field can only be integer, string or boolean.",
	"type":    "Type represents the type of the value, it can be integer, string or boolean.",
	"integer": "Integer is the integer value when type is integer.",
	"string":  "String is the string value when when type is string.",
	"boolean": "Boolean is bool value when type is boolean.",
}

func (FieldValue) SwaggerDoc() map[string]string {
	return map_FieldValue
}

var map_JsonPath = map[string]string{
	"name":    "Name represents the alias name for this field",
	"version": "Version is the version of the Kubernetes resource. If it is not specified, the resource with the semantically latest version is used to resolve the path.",
	"path":    "Path represents the json path of the field under status. The path must point to a field with single value in the type of integer, bool or string. If the path points to a non-existing field, no value will be returned. If the path points to a structure, map or slice, no value will be returned and the status conddition of StatusFeedBackSynced will be set as false. Ref to https://kubernetes.io/docs/reference/kubectl/jsonpath/ on how to write a jsonPath.",
}

func (JsonPath) SwaggerDoc() map[string]string {
	return map_JsonPath
}

var map_Manifest = map[string]string{
	"": "Manifest represents a resource to be deployed on managed cluster.",
}

func (Manifest) SwaggerDoc() map[string]string {
	return map_Manifest
}

var map_ManifestCondition = map[string]string{
	"":               "ManifestCondition represents the conditions of the resources deployed on a managed cluster.",
	"resourceMeta":   "ResourceMeta represents the group, version, kind, name and namespace of a resoure.",
	"statusFeedback": "StatusFeedback represents the values of the feild synced back defined in statusFeedbacks",
	"conditions":     "Conditions represents the conditions of this resource on a managed cluster.",
}

func (ManifestCondition) SwaggerDoc() map[string]string {
	return map_ManifestCondition
}

var map_ManifestConfigOption = map[string]string{
	"":                   "ManifestConfigOption represents the configurations of a manifest defined in workload field.",
	"resourceIdentifier": "ResourceIdentifier represents the group, resource, name and namespace of a resoure. iff this refers to a resource not created by this manifest work, the related rules will not be executed.",
	"feedbackRules":      "FeedbackRules defines what resource status field should be returned. If it is not set or empty, no feedback rules will be honored.",
	"updateStrategy":     "UpdateStrategy defines the strategy to update this manifest. UpdateStrategy is Update if it is not set, optional",
}

func (ManifestConfigOption) SwaggerDoc() map[string]string {
	return map_ManifestConfigOption
}

var map_ManifestResourceMeta = map[string]string{
	"":          "ManifestResourceMeta represents the group, version, kind, as well as the group, version, resource, name and namespace of a resoure.",
	"ordinal":   "Ordinal represents the index of the manifest on spec.",
	"group":     "Group is the API Group of the Kubernetes resource.",
	"version":   "Version is the version of the Kubernetes resource.",
	"kind":      "Kind is the kind of the Kubernetes resource.",
	"resource":  "Resource is the resource name of the Kubernetes resource.",
	"name":      "Name is the name of the Kubernetes resource.",
	"namespace": "Name is the namespace of the Kubernetes resource.",
}

func (ManifestResourceMeta) SwaggerDoc() map[string]string {
	return map_ManifestResourceMeta
}

var map_ManifestResourceStatus = map[string]string{
	"":          "ManifestResourceStatus represents the status of each resource in manifest work deployed on managed cluster",
	"manifests": "Manifests represents the condition of manifests deployed on managed cluster. Valid condition types are: 1. Progressing represents the resource is being applied on managed cluster. 2. Applied represents the resource is applied successfully on managed cluster. 3. Available represents the resource exists on the managed cluster. 4. Degraded represents the current state of resource does not match the desired state for a certain period.",
}

func (ManifestResourceStatus) SwaggerDoc() map[string]string {
	return map_ManifestResourceStatus
}

var map_ManifestWork = map[string]string{
	"":       "ManifestWork represents a manifests workload that hub wants to deploy on the managed cluster. A manifest workload is defined as a set of Kubernetes resources. ManifestWork must be created in the cluster namespace on the hub, so that agent on the corresponding managed cluster can access this resource and deploy on the managed cluster.",
	"spec":   "Spec represents a desired configuration of work to be deployed on the managed cluster.",
	"status": "Status represents the current status of work.",
}

func (ManifestWork) SwaggerDoc() map[string]string {
	return map_ManifestWork
}

var map_ManifestWorkExecutor = map[string]string{
	"":        "ManifestWorkExecutor is the executor that applies the resources to the managed cluster. i.e. the work agent.",
	"subject": "Subject is the subject identity which the work agent uses to talk to the local cluster when applying the resources.",
}

func (ManifestWorkExecutor) SwaggerDoc() map[string]string {
	return map_ManifestWorkExecutor
}

var map_ManifestWorkExecutorSubject = map[string]string{
	"":               "ManifestWorkExecutorSubject is the subject identity used by the work agent to apply the resources. The work agent should check whether the applying resources are out-of-scope of the permission held by the executor identity.",
	"type":           "Type is the type of the subject identity. Supported types are: \"ServiceAccount\".",
	"serviceAccount": "ServiceAccount is for identifying which service account to use by the work agent. Only required if the type is \"ServiceAccount\".",
}

func (ManifestWorkExecutorSubject) SwaggerDoc() map[string]string {
	return map_ManifestWorkExecutorSubject
}

var map_ManifestWorkList = map[string]string{
	"":         "ManifestWorkList is a collection of manifestworks.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of manifestworks.",
}

func (ManifestWorkList) SwaggerDoc() map[string]string {
	return map_ManifestWorkList
}

var map_ManifestWorkSpec = map[string]string{
	"":                "ManifestWorkSpec represents a desired configuration of manifests to be deployed on the managed cluster.",
	"workload":        "Workload represents the manifest workload to be deployed on a managed cluster.",
	"deleteOption":    "DeleteOption represents deletion strategy when the manifestwork is deleted. Foreground deletion strategy is applied to all the resource in this manifestwork if it is not set.",
	"manifestConfigs": "ManifestConfigs represents the configurations of manifests defined in workload field.",
	"executor":        "Executor is the configuration that makes the work agent to perform some pre-request processing/checking. e.g. the executor identity tells the work agent to check the executor has sufficient permission to write the workloads to the local managed cluster. Note that nil executor is still supported for backward-compatibility which indicates that the work agent will not perform any additional actions before applying resources.",
}

func (ManifestWorkSpec) SwaggerDoc() map[string]string {
	return map_ManifestWorkSpec
}

var map_ManifestWorkStatus = map[string]string{
	"":               "ManifestWorkStatus represents the current status of managed cluster ManifestWork.",
	"conditions":     "Conditions contains the different condition statuses for this work. Valid condition types are: 1. Applied represents workload in ManifestWork is applied successfully on managed cluster. 2. Progressing represents workload in ManifestWork is being applied on managed cluster. 3. Available represents workload in ManifestWork exists on the managed cluster. 4. Degraded represents the current state of workload does not match the desired state for a certain period.",
	"resourceStatus": "ResourceStatus represents the status of each resource in manifestwork deployed on a managed cluster. The Klusterlet agent on managed cluster syncs the condition from the managed cluster to the hub.",
}

func (ManifestWorkStatus) SwaggerDoc() map[string]string {
	return map_ManifestWorkStatus
}

var map_ManifestWorkSubjectServiceAccount = map[string]string{
	"":          "ManifestWorkSubjectServiceAccount references service account in the managed clusters.",
	"namespace": "Namespace is the namespace of the service account.",
	"name":      "Name is the name of the service account.",
}

func (ManifestWorkSubjectServiceAccount) SwaggerDoc() map[string]string {
	return map_ManifestWorkSubjectServiceAccount
}

var map_ManifestsTemplate = map[string]string{
	"":          "ManifestsTemplate represents the manifest workload to be deployed on a managed cluster.",
	"manifests": "Manifests represents a list of kuberenetes resources to be deployed on a managed cluster.",
}

func (ManifestsTemplate) SwaggerDoc() map[string]string {
	return map_ManifestsTemplate
}

var map_ResourceIdentifier = map[string]string{
	"":          "ResourceIdentifier identifies a single resource included in this manifestwork",
	"group":     "Group is the API Group of the Kubernetes resource, empty string indicates it is in core group.",
	"resource":  "Resource is the resource name of the Kubernetes resource.",
	"name":      "Name is the name of the Kubernetes resource.",
	"namespace": "Name is the namespace of the Kubernetes resource, empty string indicates it is a cluster scoped resource.",
}

func (ResourceIdentifier) SwaggerDoc() map[string]string {
	return map_ResourceIdentifier
}

var map_SelectivelyOrphan = map[string]string{
	"":               "SelectivelyOrphan represents a list of resources following orphan deletion stratecy",
	"orphaningRules": "orphaningRules defines a slice of orphaningrule. Each orphaningrule identifies a single resource included in this manifestwork",
}

func (SelectivelyOrphan) SwaggerDoc() map[string]string {
	return map_SelectivelyOrphan
}

var map_ServerSideApplyConfig = map[string]string{
	"force":        "Force represents to force apply the manifest.",
	"fieldManager": "FieldManager is the manager to apply the resource. It is work-agent by default, but can be other name with work-agent as the prefix.",
}

func (ServerSideApplyConfig) SwaggerDoc() map[string]string {
	return map_ServerSideApplyConfig
}

var map_StatusFeedbackResult = map[string]string{
	"":       "StatusFeedbackResult represents the values of the feild synced back defined in statusFeedbacks",
	"values": "Values represents the synced value of the interested field.",
}

func (StatusFeedbackResult) SwaggerDoc() map[string]string {
	return map_StatusFeedbackResult
}

var map_UpdateStrategy = map[string]string{
	"":                "UpdateStrategy defines the strategy to update this manifest",
	"type":            "type defines the strategy to update this manifest, default value is Update. Update type means to update resource by an update call. CreateOnly type means do not update resource based on current manifest. ServerSideApply type means to update resource using server side apply with work-controller as the field manager. If there is conflict, the related Applied condition of manifest will be in the status of False with the reason of ApplyConflict.",
	"serverSideApply": "serverSideApply defines the configuration for server side apply. It is honored only when type of updateStrategy is ServerSideApply",
}

func (UpdateStrategy) SwaggerDoc() map[string]string {
	return map_UpdateStrategy
}

// AUTO-GENERATED FUNCTIONS END HERE