# Tiny Redirect
Tiny https redirect application written in go for the sole purpose of redirecting web sites. This is designed to have an ultra small foot print. Takes an http request reads the host and path then returns them as an HTTPS redirect unless configured to do otherwise.

# Configuration
| Environment Variable | Description               |
| ---------------------| ------------------------- |
| REDIRECT_HOSTNAME    | Specifies a host that will replace the host in the request when the redirect url is built. |
| USE_HTTP             | Causes the redirect to use HTTP instead of HTTPS. Specifying anything Anything will activate this. (eg,. USE_HTTP=true) |

