// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package depository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moov-io/base/admin"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/paygate/internal"

	"github.com/go-kit/kit/log"
)

func RegisterAdminRoutes(logger log.Logger, svc *admin.Server, depRepo internal.DepositoryRepository) {
	svc.AddHandler("/depositories/{depositoryId}", overrideDepositoryStatus(logger, depRepo))
}

type request struct {
	Status internal.DepositoryStatus `json:"status"`
}

func overrideDepositoryStatus(logger log.Logger, depRepo internal.DepositoryRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = internal.Wrap(logger, w, r)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if r.Method != "PUT" {
			moovhttp.Problem(w, fmt.Errorf("unsupported HTTP verb: %s", r.Method))
			return
		}

		id := internal.GetDepositoryID(r)
		requestID, userID := moovhttp.GetRequestID(r), moovhttp.GetUserID(r)

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		// read the depository so we know it exists
		dep, err := depRepo.GetDepository(id)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if err := depRepo.UpdateDepositoryStatus(id, req.Status); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		// re-read for marshaling
		dep, err = depRepo.GetUserDepository(id, dep.UserID())
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		logger.Log(
			"depositories", fmt.Sprintf("updated depository=%s to %s", id, req.Status),
			"requestID", requestID, "userID", userID)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&dep)
	}
}
