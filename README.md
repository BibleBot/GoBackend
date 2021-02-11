<div align="center"><p>
<a alt="Backend" href="https://biblebot.xyz"><img alt="Backend" width="400px" src="https://i.imgur.com/JVBY24z.png"></a>
</p><p>
The Backend API for BibleBot.
</p></div>

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
To our i18n team for helping BibleBot be more accessible to everyone.  
To our licensing team for helping us sift through all the darn permissions requests.  
To our support team for helping others use BibleBot.
