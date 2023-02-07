## Movie Service Endpoint List

1. /Movies

Method: POST

```
    "data": [
        {
            "id": 2,
            "title": "Pengabdi Setan 3",
            "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
            "rating": 7.3,
            "image": "",
            "created_at": "2023-02-07T19:09:24.678902Z",
            "updated_at": "2023-02-07T19:09:24.682421Z"
        }
    ],
```

2. /Movies/:ID

Method: GET

```
"data": {
    "id": 2,
    "title": "Pengabdi Setan 3",
    "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
    "rating": 7.3,
    "image": "",
    "created_at": "2023-02-07T19:09:24.678902Z",
    "updated_at": "2023-02-07T19:09:24.682421Z"
},
```

3. /Movies

Method: POST

```
{
    "id": 2,
    "title": "Pengabdi Setan 3",
    "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
    "rating": 7.3,
    "image": "",
    "created_at": "2022-08-01 10:56:31",
    "updated_at": "2022-08-13 09:30:23"
}
```

4. /Movies/:ID

Method: PATCH

```
{
    "id": 2,
    "title": "Pengabdi Setan 3",
    "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
    "rating": 7.3,
    "image": "",
    "created_at": "2022-08-01 10:56:31",
    "updated_at": "2022-08-13 09:30:23"
}
```

5. /Movies/:ID

Method: DELETE
