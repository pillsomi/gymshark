# Gym Shark
## Task

Build a server that serves REST/API in order to calculate the most optimal number of packages to be delivered
for a given number of items.
It also provides a simple UI, that allows to check the current package sizes, update the list of package sizes,
and calculate number of packages to be delivered.

## Features

- Check current package sizes
- Update package sizes
- Calculate optimal number of packages for a given number of items.

## Tech

Technologies used for the task:

- [Golang] - Used for API server!
- [HTLM] - Used for the simple UI!
- [Docker] - Used for easy testing and Containerization.

## Installation

Gym Shark golang to be installed locally in order to test it, or docker.

Make sure to have either in order to run the web server.

Clone the repository locally.

Runing locally with go:

```sh
cd gymshark
go run main.go
```

The server will be working on port 8080.

Runing locally with docker:

```sh
cd gymshark
docker build --tag gymshark .
docker run -p 8080:8080 gymshark
```

## API and UI
Once the server is runing, you can access the UI at http://localhost:8080. It will show the index page, where you
can find the current list of package sizes, and the calculator.
You also have the possibility to change the current list of package sizes, by clicking on the update
button under the list of paccakges.

The API is also accessible directly by the endpoints:
POST: /calculate - return the packages and number for the given number of items
Request body example:
```json
{
    "number_of_items": 5000
}
```

Response body example:
```json
{
    "number_of_boxes": [
        {
            "size": 31,
            "number": 4
        },
        {
            "size": 53,
            "number": 92
        }
    ]
}
```

GET: /package/sizes - return the current list of package sizes
Response body example:
```json
{
    "package_sizes": [
        250,
        500,
        1000,
        2000,
        5000
    ]
}
```

POST: /package/sizes/update - update the list of package sizes in memory
Request body example:
```json
{
    "packages": [
        200, 
        500,
        1000'
    ]
}
```

Response body example:
```json
{
    "package_sizes": [
        250,
        500,
        1000
    ]
}
```

Also in case if errors, beside the status code, there is also an error structure returned in the response
Example error response:
```json
{
    "error_description": "please provide a non empty list of positive integers"
}
```
Usually this are user input errors, and as such the status code is 400, Bad Request.