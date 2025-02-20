// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checks

import (
	"github.com/ossf/scorecard/v3/checker"
	"github.com/ossf/scorecard/v3/checks/evaluation"
	"github.com/ossf/scorecard/v3/checks/raw"
	sce "github.com/ossf/scorecard/v3/errors"
)

// CheckBinaryArtifacts is the exported name for Binary-Artifacts check.
const CheckBinaryArtifacts string = "Binary-Artifacts"

//nolint
func init() {
	registerCheck(CheckBinaryArtifacts, BinaryArtifacts)
}

// BinaryArtifacts  will check the repository contains binary artifacts.
func BinaryArtifacts(c *checker.CheckRequest) checker.CheckResult {
	rawData, err := raw.BinaryArtifacts(c)
	if err != nil {
		e := sce.WithMessage(sce.ErrScorecardInternal, err.Error())
		return checker.CreateRuntimeErrorResult(CheckBinaryArtifacts, e)
	}

	return evaluation.BinaryArtifacts(CheckBinaryArtifacts, c.Dlogger, &rawData)
}
