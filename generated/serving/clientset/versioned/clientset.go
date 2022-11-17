/*
Copyright 2022.

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
// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	servingv1alpha2 "github.com/bentoml/yatai-deployment/generated/serving/clientset/versioned/typed/serving/v1alpha2"
	servingv1alpha3 "github.com/bentoml/yatai-deployment/generated/serving/clientset/versioned/typed/serving/v1alpha3"
	servingv2alpha1 "github.com/bentoml/yatai-deployment/generated/serving/clientset/versioned/typed/serving/v2alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ServingV1alpha2() servingv1alpha2.ServingV1alpha2Interface
	ServingV1alpha3() servingv1alpha3.ServingV1alpha3Interface
	ServingV2alpha1() servingv2alpha1.ServingV2alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	servingV1alpha2 *servingv1alpha2.ServingV1alpha2Client
	servingV1alpha3 *servingv1alpha3.ServingV1alpha3Client
	servingV2alpha1 *servingv2alpha1.ServingV2alpha1Client
}

// ServingV1alpha2 retrieves the ServingV1alpha2Client
func (c *Clientset) ServingV1alpha2() servingv1alpha2.ServingV1alpha2Interface {
	return c.servingV1alpha2
}

// ServingV1alpha3 retrieves the ServingV1alpha3Client
func (c *Clientset) ServingV1alpha3() servingv1alpha3.ServingV1alpha3Interface {
	return c.servingV1alpha3
}

// ServingV2alpha1 retrieves the ServingV2alpha1Client
func (c *Clientset) ServingV2alpha1() servingv2alpha1.ServingV2alpha1Interface {
	return c.servingV2alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.servingV1alpha2, err = servingv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.servingV1alpha3, err = servingv1alpha3.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.servingV2alpha1, err = servingv2alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.servingV1alpha2 = servingv1alpha2.New(c)
	cs.servingV1alpha3 = servingv1alpha3.New(c)
	cs.servingV2alpha1 = servingv2alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
