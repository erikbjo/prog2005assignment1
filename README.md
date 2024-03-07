# Cloud Assignment 1

[TOC]

## Introduction

<p>
In this assignment, I developed a REST web application in Go that provides the client with information about books 
available in a given language or languages based on the Gutenberg library. The service further determines the number 
of potential readers (as a second endpoint) presumed to be able to read books in that language.
</p>

## Endpoints

### Syntax

`{:parameter}` is a required parameter. <br>
`{?parameter}` is an optional parameter. <br>
`{parameter+}` is a list of parameters, separated by a `,`. <br>
All parameters can be combined in any way, _if a list of parameters is required, there needs to be atleast one
parameter_.

---

### GET /librarystats/v1/bookcount

#### Description

<p>
Returns the number of books in the library for each requested language.
Also returns the number of authors that have written books in the language.
Finally, returns the fraction of books written in the language compared to the total number of books in the library.
The language codes are defined by the <a href="https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes">ISO 639-1 standard</a>.
</p>

#### Request

```
/librarystats/v1/bookcount/?language={:two_letter_language_code+}/
```

Example requests:

```
/librarystats/v1/bookcount/?language=no,sv
/librarystats/v1/bookcount/?language=no
```

<p>
Needs to be a list of two letter language codes, separated by a comma, with at least one language code.
The language codes are defined by the <a href="https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes"> ISO 639-1 standard</a>.
</p>

#### Response

* Content-Type: `application/json`
* Status: `200 OK` if successful, relevant error code otherwise.

```json
[
  {
    "language": "no",
    "books": 21,
    "authors": 16,
    "fraction": 0.00028
  },
  {
    "language": "sv",
    "books": 230,
    "authors": 139,
    "fraction": 0.00315
  },
  {
    "language": "ra",
    "books": 0,
    "authors": 0,
    "fraction": 0
  }
]
```

---

### GET /librarystats/v1/readership

#### Description

<p>
Returns the number of potential readers for a given language, based on the population of the countries where the language is spoken.
The language codes are defined by the <a href="https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes">ISO 639-1 standard</a>. 
Also returns the country name and the ISO 3166-1 alpha-2 code for the country.
</p>
<p>
The number of books and unique authors for the language is also returned. This is based on the books in the Gutenberg library.
</p>
<p>
An optional parameter, limit, can be used to limit the number of countries returned. If not specified, all countries are returned.
</p>

#### Request

```
/librarystats/v1/readership/{:two_letter_language_code}{?limit={:number}}
```

Example request:

```
/librarystats/v1/readership/no?limit=3
/librarystats/v1/readership/sv
```

<p>
The language code is defined by the <a href="https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes">ISO 639-1 standard</a>. The limit parameter is optional, and can be any positive
integer.
</p>

#### Response

* Content-Type: `application/json`
* Status: `200 OK` if successful, relevant error code otherwise.

```json
[
  {
    "country": "Iceland",
    "isocode": "IS",
    "books": 21,
    "authors": 16,
    "readership": 366425
  },
  {
    "country": "Norway",
    "isocode": "NO",
    "books": 21,
    "authors": 16,
    "readership": 5379475
  },
  {
    "country": "Svalbard and Jan Mayen Islands",
    "isocode": "SJ",
    "books": 21,
    "authors": 16,
    "readership": 2562
  }
]
```

---

### GET /librarystats/v1/status

#### Description

<p>
Returns the status of the used services, and total uptime.
</p>

#### Request

<p>
No request parameters.
</p>

#### Response

* Content-Type: `application/json`
* Status: `200 OK` if successful, relevant error code otherwise.

```json
{
  "gutendexapi": 200,
  "languageapi": 200,
  "countriesapi": 200,
  "version": "v1",
  "uptime": 1234
}
```

---

## Comments and notes

### Deployment

The application is deployed on Render, and can be accessed [here](https://prog2005assignment1.onrender.com/). 
The deployment has not been used a lot, and is now "spun down" due to inactivity. It can take up to 50 seconds to
spin up the first time you access it. This, in addition to changes taking time to deploy, made me use a local deployment
for most of the development.

### Deviations from the specification

#### Bookcount

When using the bookcount endpoint, you are "allowed" to use any language code, even if it is not a valid ISO 639-1 code.
This is dealt with gracefully, and the response will ignore the invalid language code. Example:

```http request
GET /librarystats/v1/bookcount/?language=norge,erik,no
```
```json
[
  {
    "language": "no",
    "books": 21,
    "authors": 16,
    "fraction": 0.00028
  }
]
```

This will however not happen if there's only invalid language codes. Example:

```http request
GET /librarystats/v1/bookcount/?language=erik
```
```
Invalid language code. Please specify one or more valid two letter language codes.
```

### Known issues

When using the readership endpoint, the country name is not always returned correctly. Example:

```json
{
    "country": "Ã…land Islands",
    "isocode": "AX",
    "books": 230,
    "authors": 139,
    "readership": 29458
}
```

### Future improvements

#### Testing

The testing is not as thorough as I wanted it to be. The problem is that the structure of the code is not very testable:
It relies heavily on external services, and the responses from these services are not always predictable. 
This makes it hard, if not impossible, to write good unit tests. One way to solve this is to use a mocking library,
but I did not have time to implement this. This is something I would like to do in the future. The code also 
relies heavily on constants, I tried to use different go:build tags to make it easier to test, but I did not have time to
implement this.

The tests that are implemented are mostly sanity checks, and not very thorough. This is something I would like to improve
in the future. I would also like to implement more integration tests, to make sure that the different parts of the
application work together as expected.

## Usage

### How to run

```bash
go run cmd/main.go
```

### How to test

```bash
go test ./...
```

### How to build

```bash
go build cmd/main.go
```

then run the binary.

## Contact

For more information, please contact me on [email](mailto:erbj@stud.ntnu.no).
