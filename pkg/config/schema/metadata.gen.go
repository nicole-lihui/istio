// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// metadata.yaml
package schema

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _metadataYaml = []byte(`# Copyright 2019 Istio Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in conformance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#
# This is the main metadata file for Galley processing.
# ####  KEEP ENTRIES ALPHASORTED! ####
#

# The total set of collections, both Istio (i.e. MCP) and K8s (API Server/K8s).
collections:
  ## Istio collections
  - name: "istio/authentication/v1alpha1/meshpolicies"
    kind: "MeshPolicy"
    group: "authentication.istio.io"
    pilot: true

  - name: "istio/authentication/v1alpha1/policies"
    kind: "Policy"
    group: "authentication.istio.io"
    pilot: true

  - name: "istio/config/v1alpha2/adapters"
    kind: "adapter"
    group: "config.istio.io"

  - name: "istio/config/v1alpha2/httpapispecs"
    kind: "HTTPAPISpec"
    group: "config.istio.io"
    pilot: true

  - name: "istio/config/v1alpha2/httpapispecbindings"
    kind: "HTTPAPISpecBinding"
    group: "config.istio.io"
    pilot: true

  - name: "istio/config/v1alpha2/templates"
    kind: "template"
    group: "config.istio.io"

  - name: "istio/mesh/v1alpha1/MeshConfig"
    kind: "MeshConfig"
    group: ""

  - name: "istio/mixer/v1/config/client/quotaspecs"
    kind: "QuotaSpec"
    group: "config.istio.io"
    pilot: true

  - name: "istio/mixer/v1/config/client/quotaspecbindings"
    kind: "QuotaSpecBinding"
    group: "config.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/destinationrules"
    kind: "DestinationRule"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/envoyfilters"
    kind: "EnvoyFilter"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/gateways"
    kind: "Gateway"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/serviceentries"
    kind: "ServiceEntry"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/workloadentries"
    kind: "WorkloadEntry"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/sidecars"
    kind: "Sidecar"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/networking/v1alpha3/virtualservices"
    kind: "VirtualService"
    group: "networking.istio.io"
    pilot: true

  - name: "istio/policy/v1beta1/attributemanifests"
    kind: "attributemanifest"
    group: "config.istio.io"

  - name: "istio/policy/v1beta1/instances"
    kind: "instance"
    group: "config.istio.io"

  - name: "istio/policy/v1beta1/handlers"
    kind: "handler"
    group: "config.istio.io"

  - name: "istio/policy/v1beta1/rules"
    kind: "rule"
    group: "config.istio.io"

  - name: "istio/rbac/v1alpha1/clusterrbacconfigs"
    kind: "ClusterRbacConfig"
    group: "rbac.istio.io"
    pilot: true

  - name: "istio/rbac/v1alpha1/rbacconfigs"
    kind: "RbacConfig"
    group: "rbac.istio.io"
    pilot: true

  - name: "istio/rbac/v1alpha1/servicerolebindings"
    kind: "ServiceRoleBinding"
    group: "rbac.istio.io"
    pilot: true

  - name: "istio/rbac/v1alpha1/serviceroles"
    kind: "ServiceRole"
    group: "rbac.istio.io"
    pilot: true

  - name: "istio/security/v1beta1/authorizationpolicies"
    kind: "AuthorizationPolicy"
    group: "security.istio.io"
    pilot: true

  - name: "istio/security/v1beta1/requestauthentications"
    kind: "RequestAuthentication"
    group: "security.istio.io"
    pilot: true

  - name: "istio/security/v1beta1/peerauthentications"
    kind: "PeerAuthentication"
    group: "security.istio.io"
    pilot: true

  ### K8s collections ###

  # Built-in K8s collections
  - name: "k8s/apiextensions.k8s.io/v1beta1/customresourcedefinitions"
    kind: "CustomResourceDefinition"
    group: "apiextensions.k8s.io"

  - name: "k8s/apps/v1/deployments"
    kind: "Deployment"
    group: "apps"

  - name: "k8s/core/v1/endpoints"
    kind: "Endpoints"
    group: ""

  - name: "k8s/core/v1/namespaces"
    kind: "Namespace"
    group: ""

  - name: "k8s/core/v1/nodes"
    kind: "Node"
    group: ""

  - name: "k8s/core/v1/pods"
    kind: "Pod"
    group: ""

  - name: "k8s/core/v1/secrets"
    kind: "Secret"
    group: ""

  - name: "k8s/core/v1/services"
    kind: "Service"
    group: ""

  - name: "k8s/core/v1/configmaps"
    kind: "ConfigMap"
    group: ""

  - name: "k8s/extensions/v1beta1/ingresses"
    kind: "Ingress"
    group: "extensions"

  - kind: "GatewayClass"
    name: "k8s/service_apis/v1alpha1/gatewayclasses"
    group: "networking.x.k8s.io"

  - kind: "Gateway"
    name: "k8s/service_apis/v1alpha1/gateways"
    group: "networking.x.k8s.io"

  - kind: "HTTPRoute"
    name: "k8s/service_apis/v1alpha1/httproutes"
    group: "networking.x.k8s.io"

  - kind: "TcpRoute"
    name: "k8s/service_apis/v1alpha1/tcproutes"
    group: "networking.x.k8s.io"

  - kind: "TrafficSplit"
    name: "k8s/service_apis/v1alpha1/trafficsplits"
    group: "networking.x.k8s.io"

  # Istio CRD collections

  - name: "k8s/authentication.istio.io/v1alpha1/meshpolicies"
    kind: "MeshPolicy"
    group: "authentication.istio.io"

  - name: "k8s/authentication.istio.io/v1alpha1/policies"
    kind: "Policy"
    group: "authentication.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/adapters"
    kind: "adapter"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/attributemanifests"
    kind: "attributemanifest"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/httpapispecs"
    kind: "HTTPAPISpec"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/httpapispecbindings"
    kind: "HTTPAPISpecBinding"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/instances"
    kind: "instance"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/templates"
    kind: "template"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/quotaspecs"
    kind: "QuotaSpec"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/quotaspecbindings"
    kind: "QuotaSpecBinding"
    group: "config.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/rules"
    kind: "rule"
    group: "config.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/destinationrules"
    kind: "DestinationRule"
    group: "networking.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/envoyfilters"
    kind: "EnvoyFilter"
    group: "networking.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/gateways"
    kind: "Gateway"
    group: "networking.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/serviceentries"
    kind: "ServiceEntry"
    group: "networking.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/workloadentries"
    kind: "WorkloadEntry"
    group: "networking.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/sidecars"
    kind: "Sidecar"
    group: "networking.istio.io"

  - name: "k8s/networking.istio.io/v1alpha3/virtualservices"
    kind: "VirtualService"
    group: "networking.istio.io"

  - name: "k8s/config.istio.io/v1alpha2/handlers"
    kind: "handler"
    group: "config.istio.io"

  - name: "k8s/rbac.istio.io/v1alpha1/clusterrbacconfigs"
    kind: "ClusterRbacConfig"
    group: "rbac.istio.io"

  - name: "k8s/rbac.istio.io/v1alpha1/policy"
    kind: "ServiceRoleBinding"
    group: "rbac.istio.io"

  - name: "k8s/rbac.istio.io/v1alpha1/rbacconfigs"
    kind: "RbacConfig"
    group: "rbac.istio.io"

  - name: "k8s/rbac.istio.io/v1alpha1/serviceroles"
    kind: "ServiceRole"
    group: "rbac.istio.io"

  - name: "k8s/security.istio.io/v1beta1/authorizationpolicies"
    kind: "AuthorizationPolicy"
    group: "security.istio.io"

  - name: "k8s/security.istio.io/v1beta1/requestauthentications"
    kind: "RequestAuthentication"
    group: "security.istio.io"

  - name: "k8s/security.istio.io/v1beta1/peerauthentications"
    kind: "PeerAuthentication"
    group: "security.istio.io"

# The snapshots to generate
snapshots:
  # Used by Galley to distribute configuration.
  - name: "default"
    strategy: debounce
    collections:
      - "istio/authentication/v1alpha1/meshpolicies"
      - "istio/authentication/v1alpha1/policies"
      - "istio/config/v1alpha2/adapters"
      - "istio/config/v1alpha2/httpapispecs"
      - "istio/config/v1alpha2/httpapispecbindings"
      - "istio/config/v1alpha2/templates"
      - "istio/mesh/v1alpha1/MeshConfig"
      - "istio/mixer/v1/config/client/quotaspecbindings"
      - "istio/mixer/v1/config/client/quotaspecs"
      - "istio/networking/v1alpha3/destinationrules"
      - "istio/networking/v1alpha3/envoyfilters"
      - "istio/networking/v1alpha3/gateways"
      - "istio/networking/v1alpha3/serviceentries"
      - "istio/networking/v1alpha3/workloadentries"
      - "istio/networking/v1alpha3/sidecars"
      - "istio/networking/v1alpha3/virtualservices"
      - "istio/policy/v1beta1/attributemanifests"
      - "istio/policy/v1beta1/handlers"
      - "istio/policy/v1beta1/instances"
      - "istio/policy/v1beta1/rules"
      - "istio/rbac/v1alpha1/clusterrbacconfigs"
      - "istio/rbac/v1alpha1/rbacconfigs"
      - "istio/rbac/v1alpha1/servicerolebindings"
      - "istio/rbac/v1alpha1/serviceroles"
      - "istio/security/v1beta1/authorizationpolicies"
      - "istio/security/v1beta1/requestauthentications"
      - "istio/security/v1beta1/peerauthentications"
      - "k8s/core/v1/namespaces"
      - "k8s/core/v1/services"

    # Used by istioctl to perform analysis
  - name: "localAnalysis"
    strategy: immediate
    collections:
      - "istio/authentication/v1alpha1/meshpolicies"
      - "istio/authentication/v1alpha1/policies"
      - "istio/rbac/v1alpha1/servicerolebindings"
      - "istio/rbac/v1alpha1/serviceroles"
      - "istio/mesh/v1alpha1/MeshConfig"
      - "istio/networking/v1alpha3/envoyfilters"
      - "istio/networking/v1alpha3/destinationrules"
      - "istio/networking/v1alpha3/gateways"
      - "istio/networking/v1alpha3/serviceentries"
      - "istio/networking/v1alpha3/sidecars"
      - "istio/networking/v1alpha3/virtualservices"
      - "k8s/apiextensions.k8s.io/v1beta1/customresourcedefinitions"
      - "k8s/apps/v1/deployments"
      - "k8s/core/v1/namespaces"
      - "k8s/core/v1/pods"
      - "k8s/core/v1/secrets"
      - "k8s/core/v1/services"
      - "k8s/core/v1/configmaps"

# Configuration for resource types.
resources:
  # Kubernetes specific configuration.
  - kind: "CustomResourceDefinition"
    plural: "CustomResourceDefinitions"
    group: "apiextensions.k8s.io"
    version: "v1beta1"
    proto: "k8s.io.apiextensions_apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinition"
    protoPackage: "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

  - kind: "Deployment"
    plural: "Deployments"
    group: "apps"
    version: "v1"
    proto: "k8s.io.api.apps.v1.Deployment"
    protoPackage: "k8s.io/api/apps/v1"

  - kind: "Endpoints"
    plural: "endpoints"
    version: "v1"
    proto: "k8s.io.api.core.v1.Endpoints"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "Namespace"
    plural: "namespaces"
    version: "v1"
    clusterScoped: true
    proto: "k8s.io.api.core.v1.NamespaceSpec"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "Node"
    plural: "nodes"
    version: "v1"
    clusterScoped: true
    proto: "k8s.io.api.core.v1.NodeSpec"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "Pod"
    plural: "pods"
    version: "v1"
    proto: "k8s.io.api.core.v1.Pod"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "Secret"
    plural: "secrets"
    version: "v1"
    proto: "k8s.io.api.core.v1.Secret"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "Service"
    plural: "services"
    version: "v1"
    proto: "k8s.io.api.core.v1.ServiceSpec"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "ConfigMap"
    plural: "configmaps"
    version: "v1"
    proto: "k8s.io.api.core.v1.ConfigMap"
    protoPackage: "k8s.io/api/core/v1"

  - kind: "Ingress"
    plural: "ingresses"
    group: "extensions"
    version: "v1beta1"
    proto: "k8s.io.api.extensions.v1beta1.IngressSpec"
    protoPackage: "k8s.io/api/extensions/v1beta1"

  - Kind: "GatewayClass"
    plural: "gatewayclasses"
    group: "networking.x.k8s.io"
    version: "v1alpha1"
    clusterScoped: true
    protoPackage: "sigs.k8s.io/service-apis/api/v1alpha1"
    proto: "k8s.io.service_apis.api.v1alpha1.GatewayClassSpec"

  - Kind: "Gateway"
    plural: "gateways"
    group: "networking.x.k8s.io"
    version: "v1alpha1"
    protoPackage: "sigs.k8s.io/service-apis/api/v1alpha1"
    proto: "k8s.io.service_apis.api.v1alpha1.GatewaySpec"

  - Kind: "HTTPRoute"
    plural: "httproutes"
    group: "networking.x.k8s.io"
    version: "v1alpha1"
    protoPackage: "sigs.k8s.io/service-apis/api/v1alpha1"
    proto: "k8s.io.service_apis.api.v1alpha1.HTTPRouteSpec"

  - Kind: "TcpRoute"
    plural: "tcproutes"
    group: "networking.x.k8s.io"
    version: "v1alpha1"
    protoPackage: "sigs.k8s.io/service-apis/api/v1alpha1"
    proto: "k8s.io.service_apis.api.v1alpha1.TcpRouteSpec"

  - Kind: "TrafficSplit"
    plural: "trafficsplits"
    group: "networking.x.k8s.io"
    version: "v1alpha1"
    protoPackage: "sigs.k8s.io/service-apis/api/v1alpha1"
    proto: "k8s.io.service_apis.api.v1alpha1.TrafficSplitSpec"

  ## Istio resources
  - kind: "VirtualService"
    plural: "virtualservices"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.VirtualService"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes v1alpha3 route rules"

  - kind: "Gateway"
    plural: "gateways"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.Gateway"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes a gateway (how a proxy is exposed on the network)"

  - kind: "ServiceEntry"
    plural: "serviceentries"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.ServiceEntry"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes service entries"

  - kind: "WorkloadEntry"
    plural: "workloadentries"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.WorkloadEntry"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes workload entries"

  - kind: "DestinationRule"
    plural: "destinationrules"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.DestinationRule"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes destination rules"

  - kind: "EnvoyFilter"
    plural: "envoyfilters"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.EnvoyFilter"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes additional envoy filters to be inserted by Pilot"

  - kind: "Sidecar"
    plural: "sidecars"
    group: "networking.istio.io"
    version: "v1alpha3"
    proto: "istio.networking.v1alpha3.Sidecar"
    protoPackage: "istio.io/api/networking/v1alpha3"
    description: "describes the listeners associated with sidecars in a namespace"

  - kind: "HTTPAPISpec"
    plural: "httpapispecs"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.mixer.v1.config.client.HTTPAPISpec"
    protoPackage: "istio.io/api/mixer/v1/config/client"
    description: "describes an HTTP API specification"

  - kind: "HTTPAPISpecBinding"
    plural: "httpapispecbindings"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.mixer.v1.config.client.HTTPAPISpecBinding"
    protoPackage: "istio.io/api/mixer/v1/config/client"
    description: "describes an HTTP API specification binding"

  - kind: "QuotaSpec"
    plural: "quotaspecs"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.mixer.v1.config.client.QuotaSpec"
    protoPackage: "istio.io/api/mixer/v1/config/client"
    description: "describes an Quota specification"

  - kind: "QuotaSpecBinding"
    plural: "quotaspecbindings"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.mixer.v1.config.client.QuotaSpecBinding"
    protoPackage: "istio.io/api/mixer/v1/config/client"
    description: "describes an Quota specification binding"

  - kind: "Policy"
    plural: "policies"
    group: "authentication.istio.io"
    version: "v1alpha1"
    proto: "istio.authentication.v1alpha1.Policy"
    protoPackage: "istio.io/api/authentication/v1alpha1"
    validate: "ValidateAuthenticationPolicy"
    description: "describes an authentication policy"

  - kind: "MeshPolicy"
    plural: "meshpolicies"
    group: "authentication.istio.io"
    version: "v1alpha1"
    clusterScoped: true
    proto: "istio.authentication.v1alpha1.Policy"
    protoPackage: "istio.io/api/authentication/v1alpha1"
    validate: "ValidateAuthenticationPolicy"
    description: "describes an authentication policy at mesh level."

  - kind: "MeshConfig"
    plural: "meshconfigs"
    group: ""
    version: "v1alpha1"
    proto: "istio.mesh.v1alpha1.MeshConfig"
    protoPackage: "istio.io/api/mesh/v1alpha1"
    description: "describes the configuration for the Istio mesh."

  - kind: "ServiceRole"
    plural: "serviceroles"
    group: "rbac.istio.io"
    version: "v1alpha1"
    proto: "istio.rbac.v1alpha1.ServiceRole"
    protoPackage: "istio.io/api/rbac/v1alpha1"
    description: "describes an RBAC service role."

  - kind: "ServiceRoleBinding"
    plural: "servicerolebindings"
    group: "rbac.istio.io"
    version: "v1alpha1"
    proto: "istio.rbac.v1alpha1.ServiceRoleBinding"
    protoPackage: "istio.io/api/rbac/v1alpha1"
    description: "describes an RBAC service role."

  - kind: "RbacConfig"
    plural: "rbacconfigs"
    group: "rbac.istio.io"
    version: "v1alpha1"
    proto: "istio.rbac.v1alpha1.RbacConfig"
    protoPackage: "istio.io/api/rbac/v1alpha1"
    description: "describes the mesh level RBAC config.\n
                   Deprecated: use ClusterRbacConfig instead.\n
                   See https://github.com/istio/istio/issues/8825 for more details."

  - kind: "ClusterRbacConfig"
    plural: "clusterrbacconfigs"
    group: "rbac.istio.io"
    version: "v1alpha1"
    clusterScoped: true
    proto: "istio.rbac.v1alpha1.RbacConfig"
    protoPackage: "istio.io/api/rbac/v1alpha1"
    description: "describes the cluster level RBAC config."

  - kind: "AuthorizationPolicy"
    plural: "authorizationpolicies"
    group: "security.istio.io"
    version: "v1beta1"
    proto: "istio.security.v1beta1.AuthorizationPolicy"
    protoPackage: "istio.io/api/security/v1beta1"
    description: "describes the authorization policy."

  - kind: "RequestAuthentication"
    plural: "requestauthentications"
    group: "security.istio.io"
    version: "v1beta1"
    proto: "istio.security.v1beta1.RequestAuthentication"
    protoPackage: "istio.io/api/security/v1beta1"
    description: "describes the request authentication."

  - kind: "PeerAuthentication"
    plural: "peerauthentications"
    group: "security.istio.io"
    version: "v1beta1"
    proto: "istio.security.v1beta1.PeerAuthentication"
    protoPackage: "istio.io/api/security/v1beta1"
    validate: "ValidatePeerAuthentication"
    description: "describes the peer authentication."

  - kind: "rule"
    plural: "rules"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.policy.v1beta1.Rule"
    protoPackage: "istio.io/api/policy/v1beta1"

  - kind: "attributemanifest"
    plural: "attributemanifests"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.policy.v1beta1.AttributeManifest"
    protoPackage: "istio.io/api/policy/v1beta1"

  - kind: "instance"
    plural: "instances"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.policy.v1beta1.Instance"
    protoPackage: "istio.io/api/policy/v1beta1"

  - kind: "handler"
    plural: "handlers"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "istio.policy.v1beta1.Handler"
    protoPackage: "istio.io/api/policy/v1beta1"

  - kind: "template"
    plural: "templates"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "google.protobuf.Struct"
    protoPackage: "github.com/gogo/protobuf/types"

  - kind: "adapter"
    plural: "adapters"
    group: "config.istio.io"
    version: "v1alpha2"
    proto: "google.protobuf.Struct"
    protoPackage: "github.com/gogo/protobuf/types"

# Transform specific configurations
transforms:
  - type: direct
    mapping:
      "k8s/apiextensions.k8s.io/v1beta1/customresourcedefinitions": "k8s/apiextensions.k8s.io/v1beta1/customresourcedefinitions"
      "k8s/config.istio.io/v1alpha2/adapters": "istio/config/v1alpha2/adapters"
      "k8s/config.istio.io/v1alpha2/attributemanifests": "istio/policy/v1beta1/attributemanifests"
      "k8s/config.istio.io/v1alpha2/handlers": "istio/policy/v1beta1/handlers"
      "k8s/config.istio.io/v1alpha2/httpapispecs": "istio/config/v1alpha2/httpapispecs"
      "k8s/config.istio.io/v1alpha2/httpapispecbindings": "istio/config/v1alpha2/httpapispecbindings"
      "k8s/config.istio.io/v1alpha2/instances": "istio/policy/v1beta1/instances"
      "k8s/config.istio.io/v1alpha2/quotaspecs": "istio/mixer/v1/config/client/quotaspecs"
      "k8s/config.istio.io/v1alpha2/quotaspecbindings": "istio/mixer/v1/config/client/quotaspecbindings"
      "k8s/config.istio.io/v1alpha2/rules": "istio/policy/v1beta1/rules"
      "k8s/config.istio.io/v1alpha2/templates": "istio/config/v1alpha2/templates"
      "k8s/networking.istio.io/v1alpha3/destinationrules": "istio/networking/v1alpha3/destinationrules"
      "k8s/networking.istio.io/v1alpha3/envoyfilters": "istio/networking/v1alpha3/envoyfilters"
      "k8s/networking.istio.io/v1alpha3/gateways": "istio/networking/v1alpha3/gateways"
      "k8s/networking.istio.io/v1alpha3/serviceentries": "istio/networking/v1alpha3/serviceentries"
      "k8s/networking.istio.io/v1alpha3/workloadentries": "istio/networking/v1alpha3/workloadentries"
      "k8s/networking.istio.io/v1alpha3/sidecars": "istio/networking/v1alpha3/sidecars"
      "k8s/networking.istio.io/v1alpha3/virtualservices": "istio/networking/v1alpha3/virtualservices"
      "k8s/rbac.istio.io/v1alpha1/policy": "istio/rbac/v1alpha1/servicerolebindings"
      "k8s/rbac.istio.io/v1alpha1/rbacconfigs": "istio/rbac/v1alpha1/rbacconfigs"
      "k8s/rbac.istio.io/v1alpha1/clusterrbacconfigs": "istio/rbac/v1alpha1/clusterrbacconfigs"
      "k8s/rbac.istio.io/v1alpha1/serviceroles": "istio/rbac/v1alpha1/serviceroles"
      "k8s/security.istio.io/v1beta1/authorizationpolicies": "istio/security/v1beta1/authorizationpolicies"
      "k8s/security.istio.io/v1beta1/requestauthentications": "istio/security/v1beta1/requestauthentications"
      "k8s/security.istio.io/v1beta1/peerauthentications": "istio/security/v1beta1/peerauthentications"
      "k8s/apps/v1/deployments": "k8s/apps/v1/deployments"
      "k8s/core/v1/namespaces": "k8s/core/v1/namespaces"
      "k8s/core/v1/pods": "k8s/core/v1/pods"
      "k8s/core/v1/secrets": "k8s/core/v1/secrets"
      "k8s/core/v1/services": "k8s/core/v1/services"
      "k8s/core/v1/configmaps": "k8s/core/v1/configmaps"
      "istio/mesh/v1alpha1/MeshConfig": "istio/mesh/v1alpha1/MeshConfig"
`)

func metadataYamlBytes() ([]byte, error) {
	return _metadataYaml, nil
}

func metadataYaml() (*asset, error) {
	bytes, err := metadataYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metadata.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"metadata.yaml": metadataYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"metadata.yaml": &bintree{metadataYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
