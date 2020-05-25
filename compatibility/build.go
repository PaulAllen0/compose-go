/*
   Copyright 2020 The Compose Specification Authors.

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

package compatibility

import "github.com/compose-spec/compose-go/types"

func (c *WhiteList) CheckBuild(service *types.ServiceConfig) bool {
	if !c.supported("services.build") && service.Build != nil {
		service.Build = nil
		c.error("services.build")
		return false
	}
	return true
}

func (c *WhiteList) CheckBuildArgs(build *types.BuildConfig) {
	if !c.supported("services.build.args") && len(build.Args) != 0 {
		build.Args = nil
		c.error("services.build.args")
	}
}

func (c *WhiteList) CheckBuildLabels(build *types.BuildConfig) {
	if !c.supported("services.build.labels") && len(build.Labels) != 0 {
		build.Labels = nil
		c.error("services.build.labels")
	}
}

func (c *WhiteList) CheckBuildCacheFrom(build *types.BuildConfig) {
	if !c.supported("services.build.cache_from") && len(build.CacheFrom) != 0 {
		build.CacheFrom = nil
		c.error("services.build.cache_from")
	}
}

func (c *WhiteList) CheckBuildNetwork(build *types.BuildConfig) {
	if !c.supported("services.build.network") && build.Network != "" {
		build.Network = ""
		c.error("services.build.network")
	}
}

func (c *WhiteList) CheckBuildTarget(build *types.BuildConfig) {
	if !c.supported("services.build.target") && build.Target != "" {
		build.Target = ""
		c.error("services.build.target")
	}
}