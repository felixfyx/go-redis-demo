# go-redis-demo
An attempt at working with go and redis

In this demo, the server will be the publisher and the client will be the subscriber. Whena server push out a single message, every client should receive said message.

Following tutorial from here: https://dev.to/franciscomendes10866/using-redis-pub-sub-with-golang-mf9

## Notes: Running redis in a VM / another machine
Assuming we installed redis using `sudo apt install redis-server`, we need to modify the `redis.conf` file that may be sitting in the `/etc/redis` folder.

Find the setting that goes along the line of 
```
bind 127.0.0.1 ::-1
```
Change it to the ip to connect with

Also turn off the protected mode
```
protected-mode off
```