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

```bash
# install templ and go
templ generate 
go run main.go 
```

## Project Overview

This website contains a single webpage that takes a secret "security code", and returns all documents with that security code.

It is intended to illustrate a simplistic example of how SQL injections can be used to extract more information from databases.

> [!NOTE]  
> To configure seed data for use in a CTF, use the `-seedPath` script to your `.sql` file.
> For example,
> ```bash
> go run main.go -seedPath='/your/path/here'
> ```
> See [`/example_seed.sql`](./example_seed.sql) as an example.

### Tech Stack

| Technology                                        | Used for                |
|---------------------------------------------------|-------------------------|
| [Go](https://go.dev/)                             | Programming language    |
| [templ](https://github.com/a-h/templ)             | HTML Templating         |
| [Chroma](https://github.com/alecthomas/chroma)    | SQL syntax highlighting |
| [Slog](https://golang.org/x/exp/slog)             | Structured Logging      |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | Database Driver         |
| [SQLite3](https://www.sqlite.org/index.html)      | Minimal Database        |
| [Tailwind CSS](https://tailwindcss.com/)          | CSS Framework           |
