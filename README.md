# Go + HTMX Starter

I didn't like any of the current GO/HTMX starter out there, so I decided to build my own.
I hadn't found any templates out there that had proper hot reloading, which IMO is very important for UI work.
So, my goal was to create the simplest base I could start working with. The HR definitively jank, but it works most of the time.

## Development

This project uses Bun to handle client side stuff. [Pay them a visit to install it first!](https://bun.sh/)
You'll probably need to install [ Air ](https://github.com/air-verse/air) too. And you know, Go stuff...
Use the following command to start your web page at http://localhost:3000

```sh
bun install # for the first time building it
go mod install
bun dev
```

## Deployment

Dockerfile and docker-compose are included, just run `docker compose up --build` to run locally

## Formatting

Besides using `gofmt` for most of the formatting, this also uses `prettier` to handle Tailwind classes sorting. So make sure your editor is ready to add `prettier` to `.templ` extensions, i.e. for conform.nvim: 

```lua
      formatters_by_ft = {
        templ = { 'gofmt', 'prettier' },
      },
```
