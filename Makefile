gin-install:
	go get github.com/codegangsta/gin

gin-run:
	gin -p 2000 -a 1323 run  main.go
