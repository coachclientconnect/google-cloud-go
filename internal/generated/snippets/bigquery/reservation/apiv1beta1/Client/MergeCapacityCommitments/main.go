// Copyright 2021 Google LLC
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

// [START bigqueryreservation_generated_bigquery_reservation_apiv1beta1_Client_MergeCapacityCommitments]

package main

import (
	"context"

	reservation "cloud.google.com/go/bigquery/reservation/apiv1beta1"
	reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1beta1"
)

func main() {
	// import reservationpb "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1beta1"

	ctx := context.Background()
	c, err := reservation.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &reservationpb.MergeCapacityCommitmentsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MergeCapacityCommitments(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END bigqueryreservation_generated_bigquery_reservation_apiv1beta1_Client_MergeCapacityCommitments]