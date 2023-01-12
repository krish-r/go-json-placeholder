# go-json-placeholder

A simple clone of [JSON Placeholder][json_placeholder], written in go with [chi][chi].

## Usage

-   Clone the Repository
-   Run

    ```sh
    # Runs on port 3000
    go run .
    ```

## Resources

|             |         |             |
| ----------- | ------- | ----------- |
| `/users`    | &check; | 3 users     |
| `/posts`    | &check; | 8 posts     |
| `/comments` | &check; | 12 comments |
| `/albums`   | &cross; |             |
| `/photos`   | &cross; |             |
| `/todos`    | &cross; |             |

## Routes

-   `/users`

    |          |                  |
    | -------- | ---------------- |
    | `GET`    | `/users`         |
    | `GET`    | `/users/1`       |
    | `GET`    | `/users/1/posts` |
    | `POST`   | `/users`         |
    | `PUT`    | `/users/1`       |
    | `PATCH`  | `/users/1`       |
    | `DELETE` | `/users/1`       |

-   `/posts`

    |          |                     |
    | -------- | ------------------- |
    | `GET`    | `/posts`            |
    | `GET`    | `/posts/1`          |
    | `GET`    | `/posts/1/comments` |
    | `POST`   | `/posts`            |
    | `PUT`    | `/posts/1`          |
    | `PATCH`  | `/posts/1`          |
    | `DELETE` | `/posts/1`          |

-   `/comments`

    |          |                       |
    | -------- | --------------------- |
    | `GET`    | `/comments`           |
    | `GET`    | `/comments/1`         |
    | `GET`    | `/comments/?postId=1` |
    | `POST`   | `/comments`           |
    | `PUT`    | `/comments/1`         |
    | `PATCH`  | `/comments/1`         |
    | `DELETE` | `/comments/1`         |

[json_placeholder]: https://jsonplaceholder.typicode.com/
[chi]: https://github.com/go-chi/chi
