# Oauth2.0 with JWT and social login (FB, Google)
> Exchange user credentials or social login grant for JWT access token to IAfoosball services. 
## Flow
``` text
     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)--- JWT Access Token -----|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)--- JWT Access Token ---->|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|               |
     +--------+                               +---------------+
```
## Endpoints
```
Token access:
POST /oauth/login
POST /oauth/logout
POST /oauth/verify

Social login:
GET /oauth/facebook
GET /oauth/google
```
## Examples
* Login with Basic Auth (issue token)
```bash
curl -X POST \
  http://localhost:8001/oauth/login \
  -H 'Authorization: Basic dnlyd3U6dnlyd3U='
```
* Login with FB/Google
```bash
curl -X GET \
  http://localhost:8001/oauth/facebook
  
  or
  
curl -X GET \
  http://localhost:8001/oauth/google
```
* Verify JWT
```bash
curl -X POST \
  http://localhost:8001/oauth/verify \
  -H 'Authorization: JWT eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c3IiOiJ2eXJ3dSIsImV4cCI6MTU0MTk2OTA4OCwianRpIjoiUlpHTEg2QTVPVyJ9.FWUhvRnszVHG3wcTq97i8RhezyZgmf3w3NYk50iYfmrnBoPSD0QMJxDl60gButJvENYdvp9dmAGti1F8S7rVHTFhGriPrEtrncBtpz1TGbvw0wNW1nmf6umC7F9DfcB71bDlXhH-sIRkHA5P0P9zPnsQCF1C9rAOXvQxsCp0FTk'
```
* Logout aka. revoke token
```bash
curl -X POST \
  http://localhost:8001/oauth/logout \
  -H 'Authorization: JWT eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c3IiOiJ2eXJ3dSIsImV4cCI6MTU0MTk2OTA4OCwianRpIjoiUlpHTEg2QTVPVyJ9.FWUhvRnszVHG3wcTq97i8RhezyZgmf3w3NYk50iYfmrnBoPSD0QMJxDl60gButJvENYdvp9dmAGti1F8S7rVHTFhGriPrEtrncBtpz1TGbvw0wNW1nmf6umC7F9DfcB71bDlXhH-sIRkHA5P0P9zPnsQCF1C9rAOXvQxsCp0FTk'
```


