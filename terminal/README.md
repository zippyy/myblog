# Tech Relay Terminal

An SSH-accessible terminal version of [Tech Relay](https://techrelay.xyz). It reads the Hugo-generated JSON feed and presents posts in a responsive terminal UI.

## Features

- Anonymous SSH access with no server shell
- Latest-post list with keyboard navigation
- Search across titles, summaries, categories, and tags
- Scrollable article reader
- Automatic feed refresh every 15 minutes
- Manual refresh with `r`
- Read-only container with a persistent SSH host key
- HTTP health endpoint at `/healthz`

## Start it

```bash
cd terminal
mkdir -p data
docker compose up -d --build
```

Connect locally:

```bash
ssh -p 2222 guest@localhost
```

Wish accepts any SSH username. The username is only part of the SSH client syntax; it does not create a shell account or grant operating-system access.

## Make `ssh terminal.techrelay.xyz` work

Point `terminal.techrelay.xyz` at the Docker host. If port 22 is available exclusively for this service, change the Compose mapping to:

```yaml
ports:
  - "22:23234"
```

Then connect with:

```bash
ssh guest@terminal.techrelay.xyz
```

To use the exact command `ssh terminal.techrelay.xyz` while keeping the service on port 2222, add this client configuration to `~/.ssh/config`:

```sshconfig
Host terminal.techrelay.xyz
    HostName terminal.techrelay.xyz
    Port 2222
    User guest
```

SSH does not provide hostname-based SNI routing, so sharing one public port 22 between a normal SSH daemon and this app requires a separate public IP or an SSH-aware front end.

## Configuration

| Variable | Default | Purpose |
|---|---|---|
| `TECHRELAY_FEED_URL` | `https://techrelay.xyz/index.json` | Hugo JSON feed |
| `TECHRELAY_SITE_URL` | `https://techrelay.xyz` | Canonical website |
| `TECHRELAY_LISTEN_ADDR` | `0.0.0.0:23234` | SSH listen address |
| `TECHRELAY_HEALTH_ADDR` | `0.0.0.0:8080` | Health endpoint address |
| `TECHRELAY_HOST_KEY` | `/data/ssh_host_ed25519` | Persistent SSH host key |
| `TECHRELAY_REFRESH_INTERVAL` | `15m` | Background feed refresh interval |

## Keyboard controls

### Post list

- `j` / `k` or arrow keys: move
- `Enter`: read
- `/`: search
- `r`: refresh
- `a`: about
- `?`: help
- `q`: disconnect

### Article reader

- `j` / `k` or arrow keys: scroll
- `Space`, `PgDn`, `PgUp`: page
- `g` / `G`: top / bottom
- `b`, `Esc`, or `q`: back
- `Ctrl+C`: disconnect

## Hugo integration

The site already enables JSON output for the home page. `layouts/index.json.json` defines the feed consumed by this application. After deployment, verify it with:

```bash
curl -fsSL https://techrelay.xyz/index.json | jq '.[0] | {title, date, url}'
```

The feed includes the plain-text article body so the SSH application never has to scrape the rendered HTML website.
