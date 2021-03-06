// Copyright 2015-2017 HenryLee. All Rights Reserved.
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

package tp

// Packet types
const (
	TypeUndefined int32 = 0
	TypePull      int32 = 1
	TypeReply     int32 = 2 // reply to pull
	TypePush      int32 = 3
	// TypeAuth      int32 = 4
	// TypeHeartbeat int32 = 5
)

// TypeText returns the packet type text.
// If the type is undefined returns 'Undefined'.
func TypeText(typ int32) string {
	switch typ {
	case TypePull:
		return "PULL"
	case TypeReply:
		return "REPLY"
	case TypePush:
		return "PUSH"
	default:
		return "Undefined"
	}
}

// Common response header status as registered with IANA.
// Note: Recommended custom status code is greater than 1000.
const (
	StatusUndefined = 0

	StatusWriteFailed = 100
	StatusConnClosed  = 101

	StatusOK = 200

	StatusBadUri               = 400
	StatusUnauthorized         = 401
	StatusNotFound             = 404
	StatusConflict             = 409
	StatusUnsupportedTx        = 410
	StatusUnsupportedCodecType = 415
	// StatusFailedPlugin         = 424

	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusVariantAlsoNegotiates         = 506
	StatusInsufficientStorage           = 507
	StatusLoopDetected                  = 508
	StatusNotExtended                   = 510
	StatusNetworkAuthenticationRequired = 511
)

var statusText = map[int]string{
	StatusUndefined:            "Undefined Status",
	StatusWriteFailed:          "Write Failed",
	StatusConnClosed:           "Connection Closed",
	StatusOK:                   "OK",
	StatusBadUri:               "Bad URI",
	StatusUnauthorized:         "Unauthorized",
	StatusNotFound:             "Not Found",
	StatusConflict:             "Conflict",
	StatusUnsupportedTx:        "Unsupported transaction type",
	StatusUnsupportedCodecType: "Unsupported codec type",
	// StatusFailedPlugin:         "Failed Plugin",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

// StatusText returns a text for the Response Header status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
