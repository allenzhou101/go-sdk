package auth

import "net/http"

type IAuth interface {
	// SignInOTP - use to login a user based on the given identifier either email or a phone
	// and choose the selected delivery method for verification.
	// returns an error upon failure.
	SignInOTP(method DeliveryMethod, identifier string) error
	// SignUpOTP - use to create a new user based on the given identifier either email or a phone
	// and choose the selected delivery method for verification.
	// optional to add user metadata for farther user details such as name and more.
	// returns an error upon failure.
	SignUpOTP(method DeliveryMethod, identifier string, user *User) error

	// VerifyCode - Use to verify a SignIn/SignUp based on the given identifier either an email or a phone
	// followed by the code used to verify and authenticate the user.
	// returns a list of set-cookie data or an error upon failure.
	VerifyCode(method DeliveryMethod, identifier string, code string) ([]*http.Cookie, error)
	// VerifyCodeEmail - Use to verify a SignIn/SignUp based on the email identifier
	// followed by the code used to verify and authenticate the user.
	// returns a list of set-cookie data or an error upon failure.
	VerifyCodeEmail(identifier string, code string) ([]*http.Cookie, error)
	// VerifyCodeSMS - Use to verify a SignIn/SignUp based on the phone identifier
	// followed by the code used to verify and authenticate the user.
	// returns a list of set-cookie data or an error upon failure.
	VerifyCodeSMS(identifier string, code string) ([]*http.Cookie, error)
	// VerifyCodeWhatsApp - Use to verify a SignIn/SignUp based on the phone identifier
	// followed by the code used to verify and authenticate the user.
	// returns a list of set-cookie data or an error upon failure.
	VerifyCodeWhatsApp(identifier string, code string) ([]*http.Cookie, error)

	// ValidateSessionRequest - Use to validate a session of a given request.
	// Should be called before any private API call that requires authorization.
	// returns true upon success or false and/or error upon failure.
	ValidateSessionRequest(request *http.Request) (bool, error)

	// AuthenticationMiddleWare - middleware used to validate session and invoke if provided a failure and
	// success callbacks after calling ValidateSessionRequest().
	// onFailure will be called when the authentication failed, if empty, will write unauthorized (401) on the response writer.
	// onSuccess will be called when the authentication succeded, if empty, the request will be forwarded to the next handle.
	AuthenticationMiddleWare(onFailure func(http.ResponseWriter, *http.Request, error), onSuccess func(http.ResponseWriter, *http.Request)) func(next http.Handler) http.Handler
}
