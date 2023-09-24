```sh
go build test.go
sudo ./test
sudo tar --xattrs --xattrs-include='*' -xvf test.tar -C test2
getcap test2/web
```
