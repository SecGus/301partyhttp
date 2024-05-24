git clone https://github.com/SecGus/301partyhttp.git
cd 301partyhttp/
go mod init 301party
go get -u github.com/MadAppGang/httplog
go get -u github.com/gorilla/mux
go mod tidy
go build 301party.go
