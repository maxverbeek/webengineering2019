#GET
#songs
curl -s -H "Accept: application/json" "localhost:8080/api/v1/songs?limit=10&sort=hotttnesss" > doc/songs200.json
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/songs?limit=10&sort=hotttnesss" > doc/songs200.csv
#song
curl -s -H "Accept: application/json" "localhost:8080/api/v1/songs/SOCWJDB12A58A776AF" > doc/song200.json
curl -s -H "Accept: application/json" "localhost:8080/api/v1/songs/IDONTEXIST" > doc/song404.json
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/songs/SOCWJDB12A58A776AF" > doc/song200.csv
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/songs/IDONTEXIST" > doc/song404.csv
#artists
curl -s -H "Accept: application/json" "localhost:8080/api/v1/artists?limit=10&sort=hotttnesss" > doc/artists200.json
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/artists?limit=10&sort=hotttnesss" > doc/artists200.csv
#artist
curl -s -H "Accept: application/json" "localhost:8080/api/v1/artists/ARWPYQI1187FB4D55A" > doc/artist200.json
curl -s -H "Accept: application/json" "localhost:8080/api/v1/artists/IDONTEXIST" > doc/artist404.json
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/artists/ARWPYQI1187FB4D55A" > doc/artist200.csv
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/artists/IDONTEXIST" > doc/artist404.csv
#stats
curl -s -H "Accept: application/json" "localhost:8080/api/v1/artists/ARWPYQI1187FB4D55A/stats" > doc/stats200.json
curl -s -H "Accept: application/json" "localhost:8080/api/v1/artists/IDONTEXIST/stats" > doc/stats404.json
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/artists/ARWPYQI1187FB4D55A/stats" > doc/stats200.csv
curl -s -H "Accept: text/csv" "localhost:8080/api/v1/artists/IDONTEXIST/stats" > doc/stats404.csv

# PUT

curl -s -X PATCH --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs/IDONTEXIST" > doc/putSong404.json
curl -s -H "Accept: text/csv" -X PATCH --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs/IDONTEXIST" > doc/putSong404.csv

curl -sL -w "%{http_code}\n" -X PATCH --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs/SOCWJDB12A58A776AF"

#DELETE

curl -sL -w "%{http_code}\n" -X DELETE "localhost:8080/api/v1/songs/SOCWJDB12A58A776AF"
# POST for json
curl -s -X POST --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs" > doc/postSong201.json
curl -s -X POST --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs" > doc/postSong409.json
curl -s -X POST --data-binary "@doc/whoami.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs" > doc/postSong400.json

#DELETE
curl -sL -w "%{http_code}\n" -X DELETE "localhost:8080/api/v1/songs/SOCWJDB12A58A776AF"
# POST for csv
curl -s -H "Accept: text/csv" -X POST --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs" > doc/postSong201.csv
curl -s -H "Accept: text/csv" -X POST --data-binary "@doc/rickroll.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs" > doc/postSong409.csv
curl -s -H "Accept: text/csv" -X POST --data-binary "@doc/whoami.json" -H "Content-Type: application/json" "localhost:8080/api/v1/songs" > doc/postSong400.csv

sed -i 's%$%\n%' doc/*.csv
