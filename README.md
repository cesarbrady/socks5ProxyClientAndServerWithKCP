# socks5ProxyClientAndServerWithKCP

Usage

On server side, use “s” to run kcp service

```
$ ./run s  
12-09 15:37:10   1 [TRAC] (main.go:12) listen kcp on port  64321
```

On client side, use “c” to connect to kcp service

```
$ ./run c 127.0.0.1   
12-09 15:38:22   1 [TRAC] (main.go:29) listen tcp on port 9999
```

Now use a socks5 client connect to tcp port 9999, it should work.