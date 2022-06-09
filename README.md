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
       final-project   [no test files]
?       final-project/app/config        [no test files]
=== RUN   TestGetAPIIndex
--- PASS: TestGetAPIIndex (0.01s)
=== RUN   TestGetAPIInfo
--- PASS: TestGetAPIInfo (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller    0.037s  coverage: 100.0% of statements
=== RUN   TestAlbumDelete
--- PASS: TestAlbumDelete (0.00s)
=== RUN   TestAlbumGetById
--- PASS: TestAlbumGetById (0.01s)
=== RUN   TestAlbumGetAll
--- PASS: TestAlbumGetAll (0.00s)
=== RUN   TestAlbumPost
--- PASS: TestAlbumPost (0.00s)
=== RUN   TestAlbumPut
error : unexpected end of JSON input
--- PASS: TestAlbumPut (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller/album      0.063s  coverage: 100.0% of statements
=== RUN   TestArtistDelete
--- PASS: TestArtistDelete (0.01s)
=== RUN   TestArtistGetById
--- PASS: TestArtistGetById (0.01s)
=== RUN   TestArtistGet
--- PASS: TestArtistGet (0.01s)
=== RUN   TestArtistPost
--- PASS: TestArtistPost (0.00s)
=== RUN   TestArtistPut
error : unexpected end of JSON input
--- PASS: TestArtistPut (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller/artist     0.053s  coverage: 100.0% of statements
=== RUN   TestGendreDelete
--- PASS: TestGendreDelete (0.00s)
=== RUN   TestGendreGetById
--- PASS: TestGendreGetById (0.01s)
=== RUN   TestGendreGetAll
--- PASS: TestGendreGetAll (0.00s)
=== RUN   TestGendrePost
--- PASS: TestGendrePost (0.00s)
=== RUN   TestGendrePut
error : unexpected end of JSON input
--- PASS: TestGendrePut (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller/gendre     0.034s  coverage: 100.0% of statements
=== RUN   TestPlaylistDelete
--- PASS: TestPlaylistDelete (0.01s)
=== RUN   TestPlaylistGetById
--- PASS: TestPlaylistGetById (0.01s)
=== RUN   TestPlaylistGet
--- PASS: TestPlaylistGet (0.00s)
=== RUN   TestPlaylistPost
--- PASS: TestPlaylistPost (0.00s)
=== RUN   TestPlaylistPut
error : unexpected end of JSON input
--- PASS: TestPlaylistPut (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller/playlist   0.033s  coverage: 100.0% of statements
=== RUN   TestPlaylistTypeDelete
--- PASS: TestPlaylistTypeDelete (0.00s)
=== RUN   TestPlaylistTypeGetById
--- PASS: TestPlaylistTypeGetById (0.00s)
=== RUN   TestPlaylistTypeGet
--- PASS: TestPlaylistTypeGet (0.00s)
=== RUN   TestPlaylistTypePost
--- PASS: TestPlaylistTypePost (0.00s)
=== RUN   TestPlaylistTypePut
error : unexpected end of JSON input
--- PASS: TestPlaylistTypePut (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller/playlist_type      0.032s  coverage: 100.0% of statements
=== RUN   TestSongDeleted
--- PASS: TestSongDeleted (0.00s)
=== RUN   TestSongGet
--- PASS: TestSongGet (0.00s)
=== RUN   TestSongGetById
--- PASS: TestSongGetById (0.00s)
=== RUN   TestPostSong
--- PASS: TestPostSong (0.00s)
=== RUN   TestPutSong
error : unexpected end of JSON input
--- PASS: TestPutSong (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/controller/song       0.032s  coverage: 100.0% of statements
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
--- PASS: TestPasswordEncrypt (0.08s)
=== RUN   TestPasswordCompare
--- PASS: TestPasswordCompare (0.25s)
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
ok      final-project/app/lib   0.357s  coverage: 100.0% of statements
=== RUN   TestOauth2Authentication
--- PASS: TestOauth2Authentication (0.00s)
PASS
coverage: 100.0% of statements
ok      final-project/app/middleware    0.005s  coverage: 100.0% of statements
?       final-project/app/migrations    [no test files]
?       final-project/app/model [no test files]
?       final-project/app/routes        [no test files]
?       final-project/app/services      [no test files]
?       final-project/docs      [no test files]
go tool cover -func cover.txt
```
