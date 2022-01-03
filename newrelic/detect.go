/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package newrelic

import (
	"fmt"
	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
	"strconv"
)

type Detect struct {
	Logger bard.Logger
}

func (d Detect) Detect(context libcnb.DetectContext) (libcnb.DetectResult, error) {

	cr, err := libpak.NewConfigurationResolver(context.Buildpack, &d.Logger)
	if err != nil {
		return libcnb.DetectResult{}, fmt.Errorf("unable to create configuration resolver\n%w", err)
	}

	if e, ok := cr.Resolve("BPL_NEW_RELIC_AGENT_ENABLED"); !ok {
		return libcnb.DetectResult{Pass: false}, nil
	} else {
		enabled, err := strconv.ParseBool(e)
		if err != nil {
			return libcnb.DetectResult{Pass: false},
				fmt.Errorf("unable to parse BPL_NEW_RELIC_AGENT_ENABLED=$BPL_NEW_RELIC_AGENT_ENABLED\n%w", err)
		} else if !enabled {
			return libcnb.DetectResult{Pass: false}, nil
		}
	}

	return libcnb.DetectResult{
		Pass: true,
		Plans: []libcnb.BuildPlan{
			{
				Provides: []libcnb.BuildPlanProvide{
					{Name: "newrelic-java-agent"},
				},
				Requires: []libcnb.BuildPlanRequire{
					{Name: "newrelic-java-agent"},
					{Name: "jvm-application"},
				},
			},
		},
	}, nil
}
