<p align="center">
    <a alt="Backend" href="https://biblebot.xyz">
        <img alt="Backend" width="400px" src="https://i.imgur.com/JVBY24z.png">
    </a>
    <br>
    <br>
    <a href="https://github.com/BibleBot/backend/actions?query=workflow%3A%22docker+%28dev%29%22">
        <img alt="GitHub Workflow Status" src="https://github.com/BibleBot/backend/workflows/docker%20(dev)/badge.svg">
    </a>
    <a href="https://github.com/BibleBot/backend/actions?query=workflow%3A%22docker+%28prod%29%22">
        <img alt="GitHub Workflow Status" src="https://github.com/BibleBot/backend/workflows/docker%20(prod)/badge.svg">
    </a>
    <br>
    <a href="https://github.com/BibleBot/backend/blob/master/go.mod">
        <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/BibleBot/backend?label=go">
    </a>
    <a alt="Go Report Card" href="https://goreportcard.com/report/github.com/BibleBot/backend">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/biblebot/backend">
    </a>
    <a alt="Discord" href="https://discord.gg/H7ZyHqE">
        <img alt="Discord" src="https://img.shields.io/discord/362503610006765568?label=discord">
    </a>
    <a href="https://github.com/BibleBot/backend/blob/master/LICENSE.txt">
        <img alt="MPL-2.0" src="https://img.shields.io/github/license/BibleBot/backend">
    </a>

</p>
<p align="center">
    The Backend API for BibleBot.
</p>

## Internal Organization

While the backend repository itself is a monolith, multiple packages will exist in `internal/` that all hook into the base API. These packages are internal, since they'd be pretty useless on their own.

## Prerequisites

- Go v1.15
- Docker

## Self-Host Setup
```bash
git clone https://github.com/BibleBot/backend && cd backend
cp config.example.yml && $EDITOR config.yml

# build production container
# the build-arg is optional if you're wanting localhost *without* HTTPS
docker build --build-arg DOMAIN=<domain> -t backend .
docker run -dp 443:443 backend
```

## Special Thanks

To our financial supporters to help us keep this project's lights on.  
To our volunteer translators helping BibleBot be more accessible to everyone.  
To our licensing coordinators for helping us sift through all the darn permissions requests.  
To our outreach team for helping others use BibleBot.
