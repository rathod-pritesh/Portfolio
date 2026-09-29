# Agent Authentication Walkthrough for Pritesh Rathod Portfolio

This document specifies how AI agents and automated clients authenticate and interact with Pritesh Rathod's developer portfolio API and services, following the WorkOS `auth.md` specification.

## 1. Discover

Agents can discover authorization endpoints and resource constraints via RFC 9728 protected resource metadata:
- Protected Resource Metadata: `/.well-known/oauth-protected-resource`
- When accessing protected endpoints without credentials, the API responds with a `WWW-Authenticate` header:
```http
HTTP/1.1 401 Unauthorized
WWW-Authenticate: Bearer resource_metadata="https://priteshrathod.vercel.app/.well-known/oauth-protected-resource", error="invalid_token"
```

The discovery document advertises the `agent_auth` block and supported registration mechanisms.

## 2. Pick a method

Supported agent authentication methods:
- **anonymous**: Open read access for browsing public portfolio data (projects, skills, about, education, certifications). No registration required.
- **identity_assertion**: For automated agents submitting contact inquiries or integrating via A2A/MCP, agents may provide an `id-jag` or verified identity assertion token.

Decision criteria:
1. If reading public portfolio resources: Use **anonymous** method with no auth header.
2. If submitting contact inquiries programmatically: Use standard JSON POST with sender details or provide a signed identity assertion bearer token.

## 3. Register

For authenticated agent sessions requiring dedicated rate limits or telemetry:
- **register_uri**: `https://priteshrathod.vercel.app/api/agent/register`

Registration request example:
```http
POST /api/agent/register HTTP/1.1
Host: priteshrathod.vercel.app
Content-Type: application/json

{
  "client_name": "CodingAssistantBot",
  "identity_type": "identity_assertion",
  "redirect_uris": ["https://agent.example.com/callback"]
}
```

## 4. Claim

To complete the authentication ceremony and verify agent identity:
- **claim_uri**: `https://priteshrathod.vercel.app/api/agent/claim`

The agent presents the challenge response or user verification code obtained during registration.

## 5. Use the credential

Once authenticated, agents include the bearer token in the `Authorization` HTTP header:
```http
GET /api/projects HTTP/1.1
Host: priteshrathod.vercel.app
Authorization: Bearer <access_token>
Accept: application/json
```

For public queries, the `Authorization` header is optional.

## 6. Errors

The API returns standard RFC 6750 error responses with JSON bodies:
- `400 Bad Request`: `{"error": "invalid_request", "error_description": "Missing required field"}`
- `401 Unauthorized`: Returned with `WWW-Authenticate` header when a token is invalid or expired.
- `403 Forbidden`: `{"error": "insufficient_scope", "error_description": "Operation requires write permissions"}`
- `429 Too Many Requests`: `{"error": "rate_limit_exceeded", "error_description": "Too many requests. Please slow down."}`

## 7. Revocation

To revoke an active agent token or session:
- **revocation_uri**: `https://priteshrathod.vercel.app/api/agent/revoke`

Revocation request:
```http
POST /api/agent/revoke HTTP/1.1
Host: priteshrathod.vercel.app
Content-Type: application/x-www-form-urlencoded

token=<access_token>&token_type_hint=access_token
```
