// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package depository

import (
	"encoding/json"
	"fmt"
	"net/http"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/paygate/internal/route"

	"github.com/go-kit/kit/log"
)

// getMicroDeposits is an http.HandlerFunc for paygate's admin server to return micro-deposits for a given Depository
//
// This endpoint should not be exposed on the business http port as it would allow anyone to automatically verify a Depository
// without micro-deposits.
func getMicroDeposits(logger log.Logger, depositoryRepo Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = route.Wrap(logger, w, r)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Method != "GET" {
			moovhttp.Problem(w, fmt.Errorf("unsupported HTTP verb: %s", r.Method))
			return
		}

		id, userID := GetDepositoryID(r), route.HeaderUserID(r)
		requestID := moovhttp.GetRequestID(r)
		if id == "" {
			// 404 - A depository with the specified ID was not found.
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "depository not found"}`))
			return
		}

		microDeposits, err := depositoryRepo.GetMicroDeposits(id)
		if err != nil {
			logger.Log("microDeposits", fmt.Sprintf("admin: problem reading micro-deposits for depository=%s: %v", id, err), "requestID", requestID, "userID", userID)
			moovhttp.Problem(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(microDeposits)
	}
}
