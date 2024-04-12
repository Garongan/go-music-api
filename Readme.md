## Simple Crud with gin framework

Crud albums music with gin framework without connection to database

## list of feature:

1. ### get all albums

   ------

   

   - exampe request:

     ```bash
     curl --location 'localhost:8080/albums'
     ```
     
     
     
   - example response:

     ```javascript
     {
         "data": [
             {
                 "id": "1",
                 "title": "Dumes",
                 "artist": "Denny Chaknan",
                 "price": 100000
             },
             {
                 "id": "2",
                 "title": "Lamunan",
                 "artist": "Shinta Arshinta",
                 "price": 200000
             },
             {
                 "id": "3",
                 "title": "Kelangan",
                 "artist": "Guyon Waton",
                 "price": 250000
             }
         ],
         "message": "success retrieve albums"
     }
     ```
     
     

2. ### get album by id

    ------

    

    - example request:

      ```bash
      curl --location 'localhost:8080/albums/1'
      ```

      

    - example response:

      ```javascript
      {
          "data": {
              "id": "1",
              "title": "Dumes",
              "artist": "Denny Chaknan",
              "price": 100000
          },
          "message": "success retrieve album"
      }
      ```

      

3. ### create new album

    ------

    > [!NOTE]
    >
    > don't insert the id field, because its generated

    

    - example request:

      ```bash
      curl --location 'localhost:8080/albums' \
      --header 'Content-Type: application/json' \
      --data '{
          "title": "Perlahan",
          "artist": "Guyon Waton",
          "price": 150000
      }'
      ```

      

    - example response:

      ```javascript
      {
          "data": {
              "id": "3",
              "title": "Perlahan",
              "artist": "Guyon Waton",
              "price": 150000
          },
          "message": "success created album"
      }
      ```

      

      

4. ### update album by id

      ---

      > [!NOTE]

       can't update if has different id in request param and body

      - example request:

        ```bash
        curl --location --request PUT 'localhost:8080/albums/1' \
        --header 'Content-Type: application/json' \
        --data '{
            "id": "1",
            "title": "Perlahan",
            "artist": "Guyon Waton",
            "price": 0
        }'
        ```

        

      - example response:

        ```javascript
        {
            "data: ": {
                "id": "1",
                "title": "Perlahan",
                "artist": "Guyon Waton",
                "price": 0
            },
            "message": "success updated album"
        }
        ```

        

5. ### delete album by id

      ---

      - example request:

        ```bash
        curl --location --request DELETE 'localhost:8080/albums/4' \
        --data ''
        ```

        

      - example response:

        ```javascript
        {
            "data": {
                "id": "4",
                "title": "Perlahan",
                "artist": "Guyon Waton",
                "price": 150000
            },
            "message": "album has been deleted"
        }
        ```

        


## Status Code In This Api

| Status Code | Description   |
| ----------- | ------------- |
| 200         | `OK`        |
| 201         | `CREATED`   |
| 404         | `NOT FOUND` |
