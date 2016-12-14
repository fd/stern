//   Copyright 2016 Wercker Holding BV
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package kubernetes

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/tools/clientcmd"

	// auth providers
	_ "k8s.io/client-go/1.5/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/1.5/plugin/pkg/client/auth/oidc"
)

// NewClientSet returns a new Kubernetes client set for a context
func NewClientSet(configPath string, contextName string) (*kubernetes.Clientset, error) {
	c, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: configPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: contextName,
		},
	).ClientConfig()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get client config")
	}

	clientset, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create clientset")
	}

	return clientset, nil
}
