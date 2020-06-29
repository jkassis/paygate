/*
 * Paygate API
 *
 * PayGate is a RESTful API enabling first-party Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transfers to be created without a deep understanding of a full NACHA file specification. First-party transfers initiate at an Originating Depository Financial Institution (ODFI) and are sent off to other Financial Institutions.  Tenants are the largest grouping in PayGate and are typically a vendor who is reselling ACH services or a company making ACH payments themselves. A legal entity is linked off a Tenant as the primary Customer used to KYC and in transfers with the Tenant itself.  An Organization is a grouping within a Tenant which typically represents an entity making ACH transfers. These include clients of an ACH reseller or business accepting payments over ACH. A legal entity is linked off an Organization as the primary Customer used to KYC and in transfers with the Organization itself.  ![](https://raw.githubusercontent.com/moov-io/paygate/master/docs/images/tenant-in-paygate.png)  There are also [admin endpoints](https://moov-io.github.io/paygate/admin/) for back-office operations.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// ReturnCode struct for ReturnCode
type ReturnCode struct {
	// Optional NACHA return code for this Transfer
	Code string `json:"code"`
	// Short NACHA description of return code
	Reason string `json:"reason"`
	// Long form explanation of return code
	Description string `json:"description"`
}
