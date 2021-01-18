<p align="center">
    <a alt="Backend" href="https://biblebot.xyz">
        <img alt="Backend" width="400px" src="https://i.imgur.com/JVBY24z.png">
    </a>
    <br>
    <br>
    <a href="#">
        <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/BibleBot/backend?label=go">
    </a>
    <a alt="Go Report Card" href="https://goreportcard.com/report/github.com/BibleBot/backend">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/biblebot/backend">
    </a>
    <a alt="Discord" href="https://discord.gg/H7ZyHqE">
        <img alt="Discord" src="https://img.shields.io/discord/362503610006765568?label=discord">
    </a>
    <a href="#">
        <img alt="GitHub" src="https://img.shields.io/github/license/BibleBot/backend">
    </a>
</p>
<p align="center">
    The Backend API for BibleBot.
</p>

## Internal Organization

While the backend repository itself is a monolith, multiple packages will exist in `internal/` that all hook into the base API. These packages are internal, so cannot be imported standalone.

## Development Setup
```bash
git clone https://github.com/BibleBot/backend && cd backend

# just fill in random information for this self-signed cert
openssl req -x509 -newkey rsa:4096 -keyout https/ssl.key -out https/ssl.cert -days 365 -nodes -sha256
```

## Production Setup
```bash
git clone https://github.com/BibleBot/backend && cd backend
# place the production cert + key in https/ at this point
mkdir -p bin && go build -o bin
bin/backend.exe
```

## Special Thanks

To our financial supporters to help us keep this project's lights on.  
To our volunteer translators helping BibleBot be more accessible to everyone.  
To our licensing coordinators for helping us sift through all the darn permissions requests.  
To our outreach team for helping others use BibleBot.