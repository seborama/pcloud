package sdk

import (
	"fmt"
	"net/url"
	"time"
)

// ClientOption is a Go functional parameter signature.
// This is used by most SDK methods to pass global parameters such as username,
// getauth,id, authexpire, etc.
type ClientOption func(q *url.Values)

// WithGlobalOptionID if set to anything, you will get it back in the reply (no matter
// successful or not). This might be useful if you pipeline requests from many places over
// single connection.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionID(id string) ClientOption {
	return func(q *url.Values) {
		q.Add("id", id)
	}
}

// WithGlobalOptionTimeFormatAsUnixUTCTimestamp DO NOT USE THIS OPTION
// It is here only to remind me not to implement it :)
// The reason for not implementing WithGlobalOptionTimeFormatAsUnixUTCTimestamp is that the time
// format contract is between pCloud's API and this SDK via `sdk.APITime`!!
// This SDK uses Go's standard `time.Time`. Use `time.Format()`, etc to reformat the time
// as desired.
//
// The default datetime format is Thu, 21 Mar 2013 18:31:45 +0000 (rfc 2822), exactly 31 bytes
// long.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionTimeFormatAsUnixUTCTimestamp() ClientOption {
	return func(q *url.Values) {
		panic("do not use this option. see comment for `WithGlobalOptionTimeFormatAsUnixUTCTimestamp`")
	}
}

// WithGlobalOptionGetAuth if set, upon successful authentication an auth token will be returned.
// Auth tokens are at most 64 bytes long and can be passed back instead of username/password
// credentials by auth parameter.
// This token is especially good for setting the auth cookie to keep the user logged in.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionGetAuth() ClientOption {
	return func(q *url.Values) {
		q.Add("getauth", "1")
	}
}

// WithGlobalOptionUsername sets the username in plain text.
// Should only be used over SSL connections.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionUsername(username string) ClientOption {
	return func(q *url.Values) {
		q.Add("username", username)
	}
}

// WithGlobalOptionPassword sets the password in plain text.
// Should only be used over SSL connections.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionPassword(password string) ClientOption {
	return func(q *url.Values) {
		q.Add("password", password)
	}
}

// WithGlobalOptionAuthExpire defines the expire value of authentication token, when it is
// requested. This field is in seconds and the expire will the moment after these seconds
// since the current moment.
// Defaults to 31536000 and its maximum is 63072000.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionAuthExpire(authExpire time.Duration) ClientOption {
	const (
		defaultSeconds = 31536000
		maxSeconds     = 63072000
	)

	return func(q *url.Values) {
		e := int64(authExpire.Seconds())
		if e < 0 {
			e = defaultSeconds
		}
		if e > maxSeconds {
			e = maxSeconds
		}
		q.Add("authexpire", fmt.Sprintf("%d", e))
	}
}

// WithGlobalOptionAuthInactiveExpire defines the expire_inactive value of authentication token,
// when it is requested. This field is in seconds and the expire_incative will the moment
// after these seconds since the current moment. Defaults to 2678400 and its maximum is 5356800.
// https://docs.pcloud.com/methods/intro/global_parameters.html
func WithGlobalOptionAuthInactiveExpire(authInactiveExpire time.Duration) ClientOption {
	const (
		defaultSeconds = 2678400
		maxSeconds     = 5356800
	)

	return func(q *url.Values) {
		e := int64(authInactiveExpire.Seconds())
		if e < 0 {
			e = defaultSeconds
		}
		if e > maxSeconds {
			e = maxSeconds
		}
		q.Add("authinactiveexpire", fmt.Sprintf("%d", e))
	}
}
