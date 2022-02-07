# Notes on Implementation

- Tailscale utilizes a unix socket (defaults to `/var/run/tailscale/tailscaled.sock`).
- On every poll the socket is update with a `FilesWaiting` section that represents the local machine's inbox

## Receiving files 

- To receive a file from your inbox, you basically do a `curl --unix-socket /var/run/tailscale/tailscaled.sock http://local-tailscaled.sock/localapi/v0/files/{FILENAME}`.
- Once received, you can delete this file from the inbox by sending `curl --unix-socket /var/run/tailscale/tailscaled.sock -X DELETE http://local-tailscaled.sock/localapi/v0/files/{FILENAME}`

## Sending files

- When sending a file to a target, you basically do `curl --unix-socket /var/run/tailscale/tailscaled.sock -X PUT http://local-tailscaled.sock/localapi/v0/file-put/{TARGET}/{FILENAME}`
