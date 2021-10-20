# Go Spotify Ad Blocker [![release](https://github.com/kooroshh/SpotifyADBlocker/actions/workflows/release.yml/badge.svg?branch=1.0)](https://github.com/kooroshh/SpotifyADBlocker/actions/workflows/release.yml)
simple Socks5 Server to block spotify ads  
run the spotify ad blocker cli using 
```
./spotifyadblocker --port 51081
```
and set the socks settings in spotify application  
```
SYNOPSIS:
    spotifyadblocker.exe --port|-p <int> [--address|-a <string>] [--blacklist|-b <string>]
               [--help|-h|-?] [--upstream|-u <string>] [<args>]

REQUIRED PARAMETERS:
    --port|-p <int>            Listening Port

OPTIONS:
    --address|-a <string>      Listening Address (default: "127.0.0.1")

    --blacklist|-b <string>    Blacklist File path (default: "black_list")

    --help|-h|-?               (default: false)

    --upstream|-u <string>     UpStream Socks5 Address (ex: 127.0.0.1:1080) (default: "")
```
