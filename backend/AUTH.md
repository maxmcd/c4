# API Authentication

Notes on API authentication methods. Currently the only supported auth method is JWT.

The HTTP api endpoint for authentication is `v0/authenticate`. See `pkg/api/auth_service.proto` for service/message definitions.

## Via gRPC
See `examples/client_jwt_auth.go`.

## Via REST

- HTTP endpoints:
    - Supply JWT token via the `Authorization` header
    ```
    Authorization: JWTTOKEN
    ```
- Websocket endpoints:
    - Supply JWT token via the `Sec-Websocket-Protocol` header.
    ```
    example:
      Sec-Websocket-Protocol: Bearer, foobar
    is converted to:
      Authorization: Bearer foobar
    ```

NOTE: gRPC-go does not want `Bearer` in the `Authorization` header but many
      JWT client/server implementations will try to place `Bearer` before the token.
