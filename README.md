# Server Monitor
## Description
Monitor GPU usages across multi servers.

## Usage
### Server
```shell script
# build the server or download it from the release page
go build
# start server with a daemon (here we use pm2)
pm2 start ./server_monitor --name server-monitor -- 80
```

### Client
1. Edit the server url in `client.sh`.
2. `chmod u+x ./client.sh`.
3. `crontab -e`.
4. Add a new schedule task: `0 * * * * /path/to/client.sh` (run every 1 hour).

## TODO
- [ ] Notify specified people when specified card is available.