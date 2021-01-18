/*
Copyright 2021 The Kubernetes authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type PolicyEngineModule string

// PolicyEngineEndpointSpec defines the desired state of PolicyEngineEndpoint
type PolicyEngineEndpointSpec struct {
	Name    string    `json:"name"`
	Host    string    `json:"host"`
	Path    string    `json:"string"`
	Modules []*Module `json:"modules"`
}

type Module struct {
	Name      PolicyEngineModule      `json:"name"`
	Order     int                     `json:"order"`
	Consumers []*PolicyEngineConsumer `json:"consumers,omitempty"`
	General   *General                `json:"general,omitempty"`
	Default   *Arguments              `json:"default,omitempty"`
}

type General struct {
	Optional bool   `json:"optional,omitempty"`
	Realm    string `json:"realm,omitempty"`
	// HMAC
	Skew string `json:"skew,omitempty"`
	// oAuth2
	CallBack string `json:"callback,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	// Client secret from provider
	Secret   string `json:"secret,omitempty"`
	Provider string `json:"provider,omitempty"`
	// Custom JWT signing secret
	JWTSecret string `json:"jwt_secret,omitempty"`
	TTL       string `json:"ttl,omitempty"`
	Cookie    string `json:"cookie,omitempty"`
}

type Arguments struct {
	// Endpoint backend related
	Url          string `json:"url,omitempty"`
	PreserveHost bool   `json:"preserve_host,omitempty"`
	StripPath    bool   `json:"strip_path,omitempty"`
	// Authentication related
	Secret   string `json:"secret,omitempty"`
	Password string `json:"password,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	// IP or HTTP verb allow deny
	Order string `json:"order,omitempty"`
	Allow string `json:"allow,omitempty"`
	Deny  string `json:"deny,omitempty"`
	// Rate limit arguments
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
	Second int `json:"second,omitempty"`
}

type PolicyEngineConsumer struct {
	Name  string `json:"name"`
	EMail string `json:"email,omitempty"`
}

// PolicyEngineEndpointStatus defines the observed state of PolicyEngineEndpoint
type PolicyEngineEndpointStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PolicyEngineEndpoint is the Schema for the policyengineendpoints API
type PolicyEngineEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PolicyEngineEndpointSpec   `json:"spec,omitempty"`
	Status PolicyEngineEndpointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PolicyEngineEndpointList contains a list of PolicyEngineEndpoint
type PolicyEngineEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyEngineEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PolicyEngineEndpoint{}, &PolicyEngineEndpointList{})
}
