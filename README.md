# train-server
A simple wrapper server for the MTA realtime train data API. More info about the API can be found here: [Realtime Data Feeds](https://api.mta.info/#/landing).

### Usage
Server runs on :8080. Requests can be made to `/api/arrivalTimes` with the following json body:
```
{
    "count": 5,
    "line":"L",
    "station":"L08N"
}
```

Response will be a list like this:
```
[
    "6 mins (8:52PM)",
    "13 mins (8:59PM)",
    "20 mins (9:06PM)",
    "28 mins (9:14PM)",
    "36 mins (9:22PM)"
]
```

### Todo
- Caching
- refactor mta_client.go
