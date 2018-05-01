# Tiny Redirect
Tiny https redirect application written in go for the sole purpose of redirecting web sites. This is designed to have an ultra small foot print. Takes an http request reads the host and path then returns them as an HTTPS redirect unless configured to do otherwise.

# Configuration
| Environment Variable | Description               |
| ---------------------| ------------------------- |
| REDIRECT_HOSTNAME    | Specifies a host that will replace the host in the request when the redirect url is built. |
| USE_HTTP             | Causes the redirect to use HTTP instead of HTTPS. Specifying anything Anything will activate this. (eg,. USE_HTTP=true) |
| WHITELISTED_SUFFIX   | Turns on host checking to prevent host header poisioning. If enabled it will respond with the URL from REDIRECT_URL in the host header doesn't end with the string. |
| REDIRECT_URL         | URL minus the http(s):// that will be used if a request with a host suffix that doesn't match. |
| REDIRECT_CODE        | Code to use for the redirect. Defaults to a temporary 302 redirect. If you don't supply a valid integer it will use the default.  |

