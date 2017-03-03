### SETUP
##### install
    go get github.com/golang/mock/gomock
    go get github.com/golang/mock/mockgen

##### creating mock file
###### user
    mockgen -source=repository/user.go -destination=repository/mock_user.go -package=repository -imports=.=simple-api/models