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
	"testing"

	"github.com/buildpacks/libcnb"
	"github.com/ktronkowski/new-relic-java-agent/helper"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testCredentials(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		c helper.Credentials
	)

	it("does not contribute properties if no binding exists", func() {
		Expect(c.Execute()).To(BeZero())
	})

	it("contributes credentials if NewRelic binding exists", func() {
		c.Bindings = libcnb.Bindings{
			{
				Name:   "test-binding",
				Path:   "/test/path/test-binding",
				Type:   "NewRelicAgent",
				Secret: map[string]string{"newRelicLicenseKey": "test-value"},
			},
		}

		Expect(c.Execute()).To(Equal(map[string]string{
			"NEW_RELIC_LICENSE_KEY": "/test/path/test-binding/newRelicLicenseKey",
		}))
	})
}
