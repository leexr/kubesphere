/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha2

import (
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"

	kubesphereconfig "kubesphere.io/kubesphere/pkg/apiserver/config"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
)

const (
	GroupName = "config.kubesphere.io"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

func AddToContainer(c *restful.Container, config *kubesphereconfig.Config) error {
	webservice := runtime.NewWebService(GroupVersion)

	webservice.Route(webservice.GET("/configs/oauth").
		Doc("Information about the authorization server are published.").
		To(func(request *restful.Request, response *restful.Response) {
			response.WriteEntity(config.AuthenticationOptions.OAuthOptions)
		}))

	webservice.Route(webservice.GET("/configs/configz").
		Doc("Information about the server configuration").
		To(func(request *restful.Request, response *restful.Response) {
			response.WriteAsJson(config.ToMap())
		}))

	webservice.Route(webservice.GET("/configs/gpu/kinds").
		Doc("Get all supported GPU kinds.").
		To(func(request *restful.Request, response *restful.Response) {
			response.WriteAsJson(config.GPUOptions.Kinds)
		}))

	c.Add(webservice)
	return nil
}
