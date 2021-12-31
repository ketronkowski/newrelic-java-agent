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

package newrelic_test

import (
	"github.com/paketo-buildpacks/newrelic-java-agent/newrelic"
	"os"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx    libcnb.DetectContext
		detect newrelic.Detect
	)

	it("Not detected without $BPL_NEW_RELIC_AGENT_ENABLED", func() {
		Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{Pass: false}))
	})

	context("Not detected with $BPL_NEW_RELIC_AGENT_ENABLED=false", func() {
		it.Before(func() {
			Expect(os.Setenv("BPL_NEW_RELIC_AGENT_ENABLED", "false")).To(Succeed())
		})

		it.After(func() {
			Expect(os.Unsetenv("BPL_NEW_RELIC_AGENT_ENABLED")).To(Succeed())
		})

		it("not detected", func() {
			Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{Pass: false}))
		})
	})

	context("Detected with $BPL_NEW_RELIC_AGENT_ENABLED=true", func() {
		it.Before(func() {
			Expect(os.Setenv("BPL_NEW_RELIC_AGENT_ENABLED", "true")).To(Succeed())
		})

		it.After(func() {
			Expect(os.Unsetenv("BPL_NEW_RELIC_AGENT_ENABLED")).To(Succeed())
		})

		it("detected", func() {
			Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{
				Pass: true,
				Plans: []libcnb.BuildPlan{
					{
						Provides: []libcnb.BuildPlanProvide{
							{Name: "new-relic"},
						},
						Requires: []libcnb.BuildPlanRequire{
							{Name: "new-relic"},
							{Name: "jvm-application"},
						},
					},
				},
			}))
		})
	})

	context("Fail with invalid $BPL_NEW_RELIC_AGENT_ENABLED value", func() {
		it.Before(func() {
			Expect(os.Setenv("BPL_NEW_RELIC_AGENT_ENABLED", "invalid")).To(Succeed())
		})

		it.After(func() {
			Expect(os.Unsetenv("BPL_NEW_RELIC_AGENT_ENABLED")).To(Succeed())
		})

		it("failure", func() {
			Expect(detect.Detect(ctx)).Error()
		})
	})

}
