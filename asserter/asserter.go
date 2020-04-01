// Copyright 2020 Coinbase, Inc.
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

package asserter

import (
	"context"

	rosetta "github.com/coinbase/rosetta-sdk-go/gen"
)

// Asserter contains all logic to perform static
// validation on Rosetta Server responses.
type Asserter struct {
	operationTypes      []string
	operationStatusMap  map[string]bool
	submissionStatusMap map[string]bool
	genesisIndex        int64
}

// New constructs a new Asserter.
func New(
	ctx context.Context,
	networkResponse *rosetta.NetworkStatusResponse,
) *Asserter {
	asserter := &Asserter{
		operationTypes: networkResponse.Options.OperationTypes,
		genesisIndex:   networkResponse.NetworkStatus.NetworkInformation.GenesisBlockIdentifier.Index,
	}

	asserter.operationStatusMap = map[string]bool{}
	for _, status := range networkResponse.Options.OperationStatuses {
		asserter.operationStatusMap[status.Status] = status.Successful
	}

	asserter.submissionStatusMap = map[string]bool{}
	for _, status := range networkResponse.Options.SubmissionStatuses {
		asserter.submissionStatusMap[status.Status] = status.Successful
	}

	return asserter
}

func (a *Asserter) operationStatuses() []string {
	statuses := []string{}
	for k := range a.operationStatusMap {
		statuses = append(statuses, k)
	}

	return statuses
}

func (a *Asserter) submissionStatuses() []string {
	statuses := []string{}
	for k := range a.submissionStatusMap {
		statuses = append(statuses, k)
	}

	return statuses
}