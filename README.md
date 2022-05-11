# whatip

whatip is a simple ip resolver written in Go. Used in https://whatip.info


## Installation

```
VERSION=0.1.0 
OS=darwin
ARCH=amd64

wget -qO- https://github.com/oschvr/whatip/releases/download/$VERSION/whatip-$VERSION-$OS-$ARCH.tar.gz | tar -xzvf - -C /usr/local/bin

mv /usr/local/bin/whatip-$VERSION-$OS-$ARCH /usr/local/bin/whatip

whatip
             _               _     _         
 __      __ | |__     __ _  | |_  (_)  _ __  
 \ \ /\ / / | '_ \   / _` | | __| | | | '_ \ 
  \ V  V /  | | | | | (_| | | |_  | | | |_) |
   \_/\_/   |_| |_|  \__,_|  \__| |_| | .__/ 
                                      |_|    

Version: 0.1.0 
----------------
2022/05/10 20:34:48 🔵 [INFO] whatip running on port :8080
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[AGPLv3](https://choosealicense.com/licenses/agpl-3.0/)
