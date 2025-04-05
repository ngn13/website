My website's API, [api.ngn.tf](https://api.ngn.tf), stores information about my
self-hosted services, it also allows me to publish news and updates about these
services using an Atom feed and it keeps track of visitor metrics. The API itself is
written in Go and uses SQLite for storage.

This documentation contains information about all the available API endpoints.

## Version 1 Endpoints
Each version 1 endpoint, can be accessed using the `/v1` route.

All the endpoints return JSON formatted data.

### Errors
If any error occurs, you will get a non-200 response. And the JSON data will have an
`error` key, which will contain information about the error that occured, in the
string format. This is the only JSON key that will be set in non-200 responses.

### Results
If no error occurs, `error` key will be set to an emtpy string (""). If the endpoint
returns any data, this will be stored using the `result` key. The `result` have a
different expected type and a format for each endpoint.

### Multilang
Some `result` formats may use a structure called "Multilang". This is a simple JSON
structure that includes one key for each supported language. The key is named after
the language it represents. Currently only supported languages are:
- English (`en`)
- Turkish (`tr`)

So each multilang structure, will have **at least** one of these keys.

Here is an example multilang structure:
```
{
  "en": "Hello, world!",
  "tr": "Merhaba, dünya!"
}
```
If a `result` field is using a multilang structure, it will be specified as "Multilang"
in the rest of the documentation.

### Administrator routes
The endpoints under the `/v1/admin` route, are administrator-only routes. To access
these routes you'll need to specfiy a password using the `Authorization` header. If
the password you specify, matches with the password specified using the `API_PASSWORD`
environment variable, you will be able to access the route.

### GET /v1/services
Returns a list of available services. Each service has the following JSON format:
```
{
  "name": "Test Service",
  "desc": {
    "en": "Service used for testing the API",
    "tr": "API'ı test etmek için kullanılan servis"
  },
  "check_time": 1735861944,
  "check_res": 1,
  "check_url": "http://localhost:7001",
  "clear": "http://localhost:7001",
  "onion": "",
  "i2p": ""
}
```
Where:
- `name`: Service name (string)
- `desc`: Service description (Multilang)
- `check_time`: Last time status check time for the service, set 0 if status checking is
not supported for this service/status checking is disabled (integer, UNIX timestamp)
- `check_res`: Last service status check result (integer)
  * 0 if the service is down
  * 1 if the service is up
  * 2 if the service is up, but slow
  * 3 if the service doesn't support status checking/status checking is disabled
- `check_url`: URL used for service's status check (string, empty if none)
- `clear`: Clearnet URL for the service (string, empty string if none)
- `onion`: Onion (TOR) URL for the service (string, empty string if none)
- `i2p`: I2P URL for the service (string, empty string if none)

You can also get information about a specific service by specifying it's name using
a URL query named `name`.

### GET /v1/news/:language
Returns a Atom feed of news for the given language. Supports languages that are supported
by Multilang.

### GET /v1/metrics
Returns metrics about the API usage. The metric data has the following format:
```
{
  "since":1736294400,
  "total":8
}
```
Where:
- `since`: Metric collection start date (integer, UNIX timestamp)
- `total`: Total number of visitors (integer)

### GET /v1/admin/logs
Returns a list of administrator logs. Each log has the following JSON format:
```
{
  "action": "Added service \"Test Service\"",
  "time": 1735861794
}
```
Where:
- `action`: Action that the administrator performed (string)
- `time`: Time when the administrator action was performed (integer, UNIX timestamp)

### PUT /v1/admin/service/add
Creates a new service. The request body needs to contain JSON data, and it needs to
have the JSON format used to represent a service. See `/v1/services/all` route to
see this format.

Returns no data on success.

### DELETE /v1/admin/service/del
Deletes a service. The client needs to specify the name of the service to delete, by
setting the URL query `name`.

Returns no data on success.

### GET /v1/admin/service/check
Forces a status check for all the services.

Returns no data on success.

### PUT /v1/admin/news/add
Creates a news post. The request body needs to contain JSON data, and it needs
to use the following JSON format:
```
{
  "id": "test_news",
  "title": {
    "en": "Very important news",
    "tr": "Çok önemli haber"
  },
  "author": "ngn",
  "content": {
    "en": "Just letting you know that I'm testing the API",
    "tr": "Sadece API'ı test ettiğimi bilmenizi istedim"
  }
}
```
Where:
- `id`: Unique ID for the news post (string)
- `title`: Title for the news post (Multilang)
- `author`: Author of the news post (string)
- `content`: Contents of the news post (Multilang)

Returns no data on success.

### DELETE /v1/admin/news/del
Deletes a news post. The client needs to specify the ID of the news post to delete,
by setting the URL query `id`.

Returns no data on success.
