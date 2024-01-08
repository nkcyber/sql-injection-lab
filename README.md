# sql-injection-lab
This is a bare bones document viewer used to practice SQL injections

This application features a simple document viewer, that takes a user's 'security code'
and returns all documents tagged with that code

> [!WARNING]  
> This application contains intentional SQL injection vulnerabilities.

> [!NOTE]  
> This will be reused across CTFs on **February 18th, 2024** and **March 23rd, 2024** (and maybe more).

## Screenshot

![image](https://github.com/nkcyber/sql-injection-lab/assets/46602241/444ad589-380a-456c-ab53-2292aa8a1287)


## Run locally

1. Install [templ](https://templ.guide/quick-start/installation) and [Go](https://go.dev/doc/install).
2. ```bash
   templ generate 
   go run main.go 
   ```

## Project Overview

This website contains a single webpage that takes a secret "security code", and returns all documents with that security code.

It is intended to illustrate a simplistic example of how SQL injections can be used to extract more information from databases. 
Pedagogically, the "security code" is an example of passwords, usernames, or any other text input that may be intended to constrain a query.

> [!NOTE]  
> ```
> Usage of ./sql-injection-lab:
>  -ip string
>    	The ip address to listen and serve HTTP on (default "localhost")
>  -port int
>    	The port to listen and serve HTTP on (default 8080)
>  -seedPath string
>    	The path to the SQL script with seed data;
>    	The script will be executed on server initalization (default "./example_seed.sql")
> ```
> For example,
> ```bash
> go run main.go -seedPath='/your/path/here'
> ```


### Tech Stack

| Technology                                        | Used for                |
|---------------------------------------------------|-------------------------|
| [Go](https://go.dev/)                             | Programming language    |
| [templ](https://github.com/a-h/templ)             | HTML Templating         |
| [Chroma](https://github.com/alecthomas/chroma)    | SQL syntax highlighting |
| [slog](https://golang.org/x/exp/slog)             | Structured Logging      |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | Database Driver         |
| [SQLite3](https://www.sqlite.org/index.html)      | Minimal Database        |
| [Tailwind CSS](https://tailwindcss.com/)          | CSS Framework           |

This project intentionally does not include user sessions or account management to limit scope. All URL paths are treated equally. The page is intended to interact well with [Burp Suite](https://en.wikipedia.org/wiki/Burp_Suite).
