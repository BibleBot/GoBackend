<div align="center"><p>
<a alt="Backend" href="https://biblebot.xyz"><img alt="Backend" width="400px" src="https://i.imgur.com/JVBY24z.png"></a>
</p><p>
The Backend API for BibleBot.
</p></div>

## Prerequisites

- Go v1.16
- Docker

## Notes

- Tests that end with `_U` in the function name are unimplemented, to be made later.

## Self-Host Setup
```bash
git clone https://internal.kerygma.digital/kerygma-digital/biblebot/backend && cd backend
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
