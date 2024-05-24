# Pocket-Litefs

An example repo for deploying pocketbase as a framework to multiple regions on Fly.io with data replicated by LiteFS.

#### Install Golang project dependencies
```bash
go mod download
```

#### Install Air to watch and rebuild in dev
```bash
go install github.com/cosmtrek/air@latest
```

#### Important Paths in Docker/Fly Instance
- `/app` build directory for Go Pocketbase application codebase
- `/app/bin/pocketbase` Pocketbase binary that LiteFS will run and manage sqlite file locking
- `/litefs` place where sqlite files should be placed by Pocketbase for replication by LiteFS

#### Deploying to Fly
- Copy this project or its config files (Dockerfile, litefs.yml, fly.toml)
- Launch but don't deploy. You can skip this if you already have an existing Fly app.
- When Fly has created an app in your Dashboard, run `fly attach consul`
- Then you can run `fly deploy`. Volumes and machines will be created for your primary region.
- At this point your service should be reachable.
- Scale to two other regions with `fly scale count 2 --region sea,dfw`.
- This will create two new 1Gb volumes to store replicated sqlite databases and two new Machines running Pocketase + LiteFS.
- Have fun.
