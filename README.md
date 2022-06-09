# Base Project

[API Specs](https://gospecs.monstercode.net)


[Restful api MUSIC (galang Bootcamp)]

# Mater Data



#_Artist : 
    ID
    Name
    About 
```

# Song : 
    ID
    Name
    Gendre
    Year
    Description
    ArtistID 
```
#_Album : 
    ID
    AlbumName
    RealeseDate
    ArtistID
    GendreID
    
```
#_Gendre : 
    ID
    GenreName
    GenreDescr
    
```
#_PlaylistType : 
    ID
    PlayListTypeName
    
```
#_Playlist : 
    ID
    PlaylistName
    PlaylistDesc
    PlaylistPicture
    CreationDate
    PlaylistTypeID
    
```
<!-- dan akan terus di update jika memiliki waktu luang  :) -->

Running test:
```bash
docker-compose up test
# or
go test -v -coverprofile cover.txt ./...
```
#  or specific file
go test -v -coverprofile cover.txt ./app/controller
```

Get Coverage:
```bash
go tool cover -html cover.txt




Grade: A+ (100.0%)
Files: 63
Issues: 0
gofmt: 100%
go_vet: 100%
gocyclo: 100%
ineffassign: 100%
license: 100%
misspell: 100%
$ go mod download github.com/onsi/ginkgo
$ go get github.com/go-redis/redis/v8@v8.11.5
$ go get github.com/jackc/pgconn@v1.12.1
$ if [ "$FAIL_ON_CONTROLLER" = "1" ]; then \ # collapsed multi-line command
testing: warning: no tests to run
PASS
ok  	final-project/app/controller	0.055s [no tests to run]
Cleaning up file based variables
00:01
Job succeeded



//Unit Testing 

PASS
coverage: 100.0% of statements
ok  	final-project/app/controller/song	0.094s	coverage: 100.0% of statements
=== RUN   TestHTTPRequest
--- PASS: TestHTTPRequest (0.00s)
=== RUN   TestGetTest
--- PASS: TestGetTest (0.00s)
=== RUN   TestPostTest
--- PASS: TestPostTest (0.00s)
=== RUN   TestPutTest
--- PASS: TestPutTest (0.00s)
=== RUN   TestDeleteTest
--- PASS: TestDeleteTest (0.00s)
=== RUN   TestLoadEnvironment
--- PASS: TestLoadEnvironment (0.00s)
=== RUN   TestCipherEncryptDecrypt
--- PASS: TestCipherEncryptDecrypt (0.00s)
=== RUN   TestGeneratePassword
--- PASS: TestGeneratePassword (0.00s)
=== RUN   TestRandomChars
--- PASS: TestRandomChars (0.00s)
=== RUN   TestMerge
--- PASS: TestMerge (0.00s)
=== RUN   TestPasswordEncrypt
--- PASS: TestPasswordEncrypt (0.15s)
=== RUN   TestPasswordCompare
--- PASS: TestPasswordCompare (0.24s)
=== RUN   TestUUIDPtr
--- PASS: TestUUIDPtr (0.00s)
=== RUN   TestIntptr
--- PASS: TestIntptr (0.00s)
=== RUN   TestInt64ptr
--- PASS: TestInt64ptr (0.00s)
=== RUN   TestStrptr
--- PASS: TestStrptr (0.00s)
=== RUN   TestBoolptr
--- PASS: TestBoolptr (0.00s)
=== RUN   TestFloat64ptr
--- PASS: TestFloat64ptr (0.00s)
=== RUN   TestGetXUserID
--- PASS: TestGetXUserID (0.00s)
=== RUN   TestGetXAgentID
--- PASS: TestGetXAgentID (0.00s)
=== RUN   TestGetLanguage
--- PASS: TestGetLanguage (0.00s)
=== RUN   TestBodyParser
--- PASS: TestBodyParser (0.00s)
=== RUN   TestSend
--- PASS: TestSend (0.00s)
=== RUN   TestErrorBadRequest
--- PASS: TestErrorBadRequest (0.00s)
=== RUN   TestErrorNotFound
--- PASS: TestErrorNotFound (0.00s)
=== RUN   TestErrorConflict
--- PASS: TestErrorConflict (0.00s)
=== RUN   TestErrorInternal
--- PASS: TestErrorInternal (0.00s)
=== RUN   TestOK
--- PASS: TestOK (0.00s)
=== RUN   TestCurrentTime
--- PASS: TestCurrentTime (0.00s)
=== RUN   TestTimeNow
--- PASS: TestTimeNow (0.00s)
=== RUN   TestRangeDate
--- PASS: TestRangeDate (0.00s)
=== RUN   TestDateTimeAhead
--- PASS: TestDateTimeAhead (0.00s)
PASS
coverage: 100.0% of statements
ok  	final-project/app/lib	0.472s	coverage: 100.0% of statements
=== RUN   TestOauth2Authentication
--- PASS: TestOauth2Authentication (0.00s)
PASS
coverage: 100.0% of statements
ok  	final-project/app/middleware	0.027s	coverage: 100.0% of statements
?   	final-project/app/migrations	[no test files]
?   	final-project/app/model	[no test files]
?   	final-project/app/routes	[no test files]
?   	final-project/app/services	[no test files]
# or
go tool cover -func cover.txt
```
