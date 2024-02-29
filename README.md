# Cloud Assignment 1

## Introduction

Bla bla

## Endpoints

### Syntax

`{:parameter}` is a required parameter. <br>
`{?parameter}` is an optional parameter. <br>
`{parameter+}` is a list of parameters, separated by a `,`. <br>
All parameters can be combined in any way, _if a list of parameters is required, there needs to be atleast one
parameter_.

### GET /librarystats/v1/bookcount

#### Description

Returns the number of books in the library for each requested language.
Also returns the number of authors that have written books in the language.
Finally, returns the fraction of books written in the language compared to the total number of books in the library.

#### Request

```
/librarystats/v1/bookcount/?language={:two_letter_language_code+}/
```

Example request:

```
/librarystats/v1/bookcount/?language=no,sv
```

Needs to be a list of two letter language codes, separated by a `,`, with at least one language code.
The language codes are defined by the ISO 639-1 standard.

#### Response

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
  }
]
```

---

### GET /librarystats/v1/readership

#### Description

<p>
Returns the number of potential readers for a given language, based on the population of the countries where the language is spoken.
The language codes are defined by the ISO 639-1 standard. Also returns the country name and the ISO 3166-1 alpha-2 code for the country.
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
```

The language code is defined by the ISO 639-1 standard. The limit parameter is optional, and can be any positive
integer.

#### Response

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

Returns the status of the used services, and total uptime.

#### Request

No request parameters.

#### Response

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

### Testing

### Deployment

### Deviations from the specification

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

## Usage

### How to run

```bash
go run cmd/main.go
```

### How to test

```bash

```

### How to build

```bash
go build cmd/main.go
```

then run the binary.

## Contact

For more information, please contact the author at [email](mailto:erbj@stud.ntnu.no).
