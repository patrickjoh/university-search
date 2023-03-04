# Assignment 1 - PROG2005

## Description

A REST web appliction made in Golang that provides the client a service to retrieve information about universities that matches the search criteria, as explained below.
The first endpoint of the service provides contextual information about language and country that the university is situated in.

The second capability of this service is to find matching universities in countries other than the one you specified. This may for example help in finding universities to apply to, if you are looking to study abroad.

## Deployment

* Run the following command in the root directory of the project:
```go
go build ./cmd/main.go
```
## Dependencies

### External APIs
* http://universities.hipolabs.com/
* Documentation/Source under: https://github.com/Hipo/university-domains-list/
* https://restcountries.com/
* Documentation/Source under: https://gitlab.com/amatos/rest-countries

## Usage

### Endpoints

The web service have three resource root paths:
```http
/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
```

If the web service is running on localhost, port 8080,
the full paths to the resources would look something like this:
```http
http://localhost:8080/unisearcher/v1/uniinfo/
http://localhost:8080/unisearcher/v1/neighbourunis/
http://localhost:8080/unisearcher/v1/diag/
```

### Retrieve information for a given university
The initial endpoint focuses on returning information about
all universities that match a given name, complete or partial.

**Request**
```http
Method: GET
Path: uniinfo/{:partial_or_complete_university_name}/
```

Note: The name of the university can be partial or complete, and may return a single ("Cambridge") or multiple universities (e.g., "Middle").

*Example request:*
```
uniinfo/norwegian%20university%20of%20science%20and%20technology/
```

**Response**
* Content-Type: application/json
* Status code: 200 OK
```json
[
    {
        "name": "Norwegian University of Science and Technology",
        "country": "Norway",
        "iso_code": "NO",
        "webpages": [
            "http://www.ntnu.no/"
        ],
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "maps": {
            "googleMaps": "https://goo.gl/maps/htWRrphA7vNgQNdSA",
            "openStreetMaps": "https://www.openstreetmap.org/relation/2978650"
        }
    }
]
```

### Retrieve universities with same name components in neighbouring countries
The second endpoint provides an overview of universities in neighbouring countries to a given country that have the same name component (e.g., "Middle") in their institution name. This should not include universities from the country that was specified in the request.

**Request**
```http
Method: GET
Path: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}
```

* `{:country_name}` refers to the English name for the country that is the basis of the search. The service will find universities in the neighbouring countries of the one specified.
* `{:partial_or_complete_university_name}` is the partial or complete university name, for which universities with similar name are sought in neighbouring countries
* `{?limit={:number}}` is an optional parameter that limits the number of universities in bordering countries (number) that are reported.

*Example request:*
```
neighbourunis/norway/science?limit=5
```

**Response**
* Content-Type: application/json
* Status code: 200 OK
```json
[
    {
        "name": "Norwegian University of Science and Technology",
        "country": "Norway",
        "iso_code": "NO",
        "webpages": [
            "http://www.ntnu.no/"
        ],
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "maps": {
            "googleMaps": "https://goo.gl/maps/htWRrphA7vNgQNdSA",
            "openStreetMaps": "https://www.openstreetmap.org/relation/2978650"
        }
    }
]
```

### Diagnostics interface
The diagnostics interface indicates the availability of the individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.

**Request**

```http
Method: GET
Path: diag/
```

**Response**
* Content-Type: application/json
* Status code: 200 OK
```json
{
  "universities_api": "200 OK",
  "countries_api": "200 OK",
  "version": "v1",
  "uptime": "35m5.666422318s"
}
```

## Credits

### Contributors
* Patrick Johannessen (Owner)

### Aknowledgements
* Christopher Frantz (Lecturer)





