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

package helper

import (
	"fmt"
	"github.com/paketo-buildpacks/libpak/bard"
	"os"
	"strconv"
)

type Newrelic struct {
	Logger bard.Logger
}

func (n Newrelic) Execute() (map[string]string, error) {

	n.Logger.Info("Configuring New Relic environment variables")

	environment := make(map[string]string)

	if e, ok := os.LookupEnv("BPL_NEW_RELIC_AGENT_ENABLED"); ok {
		enabled, err := strconv.ParseBool(e)
		if err != nil {
			return nil, fmt.Errorf("unable to parse BPL_NEW_RELIC_AGENT_ENABLED=$BPL_NEW_RELIC_AGENT_ENABLED\n%w", err)
		}
		environment["NEW_RELIC_AGENT_ENABLED"] = fmt.Sprintf("%t", enabled)
	} else {
		return nil, fmt.Errorf("unable to find required environment variable BPL_NEW_RELIC_AGENT_ENABLED")
	}

	if e, ok := os.LookupEnv("BPL_NEW_RELIC_DISTRIBUTED_TRACING_ENABLED"); ok {
		enabled, err := strconv.ParseBool(e)
		if err != nil {
			return nil, fmt.Errorf("unable to parse BPL_NEW_RELIC_DISTRIBUTED_TRACING_ENABLED"+
				"=$BPL_NEW_RELIC_DISTRIBUTED_TRACING_ENABLED\n%w", err)
		}
		environment["NEW_RELIC_DISTRIBUTED_TRACING_ENABLED"] = fmt.Sprintf("%t", enabled)
	}

	if name, ok := os.LookupEnv("BPL_NEW_RELIC_APP_NAME"); ok {
		environment["NEW_RELIC_APP_NAME"] = name
	}

	if key, ok := os.LookupEnv("BPL_NEW_RELIC_LICENSE_KEY"); ok {
		environment["NEW_RELIC_LICENSE_KEY"] = key
	}

	return environment, nil

}
