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

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package gen

// Options Options specify supported methods, Operation.Status, Operation.Type, and all possible transaction submission statuses. This Options object is used by clients to validate the correctness of a Rosetta Server implementation. It is expected that these clients will error if they receive some response that contains any of the above information that is not specified here.
type Options struct {
	// All methods that this implementation supports.
	Methods []string `json:"methods"`
	// All `Operation.Status` this implementation supports. Any status that is returned during parsing that is not listed here will cause client validation to error.
	OperationStatuses []*OperationStatus `json:"operation_statuses"`
	// All `Operation.Type` this implementation supports. Any type that is returned during parsing that is not listed here will cause client validation to error.
	OperationTypes []string `json:"operation_types"`
	// All `status` that can be returned when submitting a transaction using the `/construction/submit` endpoint.
	SubmissionStatuses []*SubmissionStatus `json:"submission_statuses"`
}