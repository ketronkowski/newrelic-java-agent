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
	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/newrelic-java-agent/helper"
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testProperties(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		p helper.Properties
	)

	it("NEW_RELIC_AGENT_ENABLED=false if not specified", func() {

		it("uses configured module", func() {
			Expect(p.Execute()).To(Equal(map[string]string{
				"NEW_RELIC_AGENT_ENABLED": "false",
			}))
		})
	})

	context("NEW_RELIC_AGENT_ENABLED=true", func() {
		it.Before(func() {
			Expect(os.Setenv("BPL_NEW_RELIC_AGENT_ENABLED", "true")).To(Succeed())
		})

		it.After(func() {
			Expect(os.Unsetenv("BPL_NEW_RELIC_AGENT_ENABLED")).To(Succeed())
		})

		it("correct value for $NEW_RELIC_AGENT_ENABLED", func() {
			Expect(p.Execute()).To(Equal(map[string]string{
				"NEW_RELIC_AGENT_ENABLED": "true",
			}))
		})

		it("contributes credentials if NewRelic binding exists", func() {

			p.Bindings = libcnb.Bindings{
				{
					Name:   "test-binding",
					Path:   "/test/path/test-binding",
					Type:   "NewRelicAgent",
					Secret: map[string]string{"newRelicLicenseKey": "test-value"},
				},
			}

			Expect(p.Execute()).To(Equal(map[string]string{
				"NEW_RELIC_LICENSE_KEY":   "/test/path/test-binding/newRelicLicenseKey",
				"NEW_RELIC_AGENT_ENABLED": "true",
			}))
		})
	})

	// TODO: More tests for other variables and invalid values.
}
