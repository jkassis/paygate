/*
 * Paygate Admin API
 *
 * PayGate is a RESTful API enabling first-party Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transfers to be created without a deep understanding of a full NACHA file specification. First-party transfers initiate at an Originating Depository Financial Institution (ODFI) and are sent off to other Financial Institutions.  Refer to the [client endpoints](https://moov-io.github.io/paygate/) for customr facing operations.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin

// Error struct for Error
type Error struct {
	// An error message describing the problem intended for humans.
	Error string `json:"error"`
}
