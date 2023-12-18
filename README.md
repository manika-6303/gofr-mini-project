# Music Album CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API for managing music albums. The API is built in Go using the gofr framework and utilizes a SQLite database for storage.
Table of Contents

    Endpoints
        1. Get Album by ID
        2. Get Album by Title
        3. Get Album by Artist
        4. Delete Album by ID
        5. Add Album
        6. Update Album by ID
        7. Get All Albums
    Usage
    Dependencies

Endpoints
1. Get Album by ID

    Endpoint: /album/getid/{id}
    Method: GET
    Description: Retrieves a specific album by its ID.
    Example: /album/getid/1

2. Get Album by Title

    Endpoint: /album/gettitle/{title}
    Method: GET
    Description: Retrieves a specific album by its title.
    Example: /album/gettitle/ExampleTitle

3. Get Album by Artist

    Endpoint: /album/getartist/{artist}
    Method: GET
    Description: Retrieves all albums by a specific artist.
    Example: /album/getartist/ExampleArtist

4. Delete Album by ID

    Endpoint: /album/delete/{id}
    Method: POST
    Description: Deletes a specific album by its ID.
    Example: /album/delete/1

5. Add Album

    Endpoint: /album/{data}
    Method: POST
    Description: Adds a new album.
    Example: /album/1,ExampleTitle,ExampleArtist,100

6. Update Album by ID

    Endpoint: /album/updateid/{id}/{data}
    Method: POST
    Description: Updates a specific album by its ID.
    Example: /album/updateid/1/title=NewTitle,artist=NewArtist,price=200

7. Get All Albums

    Endpoint: /album
    Method: GET
    Description: Retrieves all albums.

Usage

    Ensure you have Go installed on your machine.
    Install the required dependencies using go get.
    Run the main application file: go run main.go.
    Access the API endpoints as described above.

Dependencies

    gofr: A lightweight Go web framework
