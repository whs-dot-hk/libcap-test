```sh
go build test.go
sudo ./test
sudo tar --xattrs --xattrs-include='*' -xvf test.tar -C test2
getcap test2/web
fakeroot bash -c "setcap cap_net_bind_service=+ep web&& tar --xattrs --xattrs-include='*' -cvf test2.tar web"
```
