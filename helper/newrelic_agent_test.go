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

package helper_test

import (
	"os"
	"testing"

	"github.com/ktronkowski/new-relic-java-agent/helper"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testNewRelic(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		n helper.Newrelic
	)

	it("Error that BPL_NEW_RELIC_AGENT_ENABLED is not set", func() {
		Expect(n.Execute()).Error()
	})

	context("$BPL_NEW_RELIC_AGENT_ENABLED=true", func() {
		it.Before(func() {
			Expect(os.Setenv("BPL_NEW_RELIC_AGENT_ENABLED", "true")).To(Succeed())
		})

		it.After(func() {
			Expect(os.Unsetenv("BPL_NEW_RELIC_AGENT_ENABLED")).To(Succeed())
		})

		it("sets enabled value", func() {
			Expect(n.Execute()).To(Equal(map[string]string{
				"NEW_RELIC_AGENT_ENABLED": "true",
			}))
		})
	})

}
