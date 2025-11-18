# Hello, Edgerunner
---

To run this piece of code, you should install go on your computer

Than clone this repo via git
```
git clone https://github.com:CatKap/sixPsyh.git 
```

Install (sqlc)[ https://docs.sqlc.dev/en/latest/overview/install.html]



And after that run
```
go run main.go
```

Check avalable on http://localhost:8080/health/


## Docs 
--- 
Alaviable enpoints is 

Path /cathegorys/ returns all cathegorys in 
```json
[
    {
    "id":<id>,
    "name":"<name>"
    },
    ...
] 
```

Path /events/new/
Accepts POST request with json payload.
`ctg` is a cathegory. Accepts only string name, adding via id is prohibited.

```bash
curl -X POST http://localhost:8080/events/new/ \
  -H "Content-Type: application/json" \   # Content type
    # Payload
  -d '{\
    "name":"Meeting of psycopaths",\
    "desc":"Everyone will be preasured!",\ # Description of event
    "ctg":"BestCath",\ # Category of event, only name are accepted!
    "time":50230482304\ # Time shoud be int, UTC format
}'
```

Path /events/, GET request - will be return list of all events in json with fiels `id`, `name`, `desc`, `ctg`, `time`. Time in UTC. 



Path /entity/ accepts only DELETE http request with such payload:
```json
    {
    "table": "<cathegory || event",
    "ids": [<list of ids to delete, such 1, 2, 3>]
}
```
In return Ok or error.







