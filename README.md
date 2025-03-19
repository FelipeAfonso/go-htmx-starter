# Go + HTMX Starter

I didn't like any of the current GO/HTMX starter out there, so I decided to build my own.
I hadn't found any templates out there that had proper hot reloading, which IMO is very important for UI work.
So, my goal was to create the simplest base I could start working with. 

This is a Go Fiber boilerplate, that combines Vite/Air to handle Hot Reloading, improving the dev experience. 
It uses Tailwindcss@v4 (with the Prettier plugin to sort classes); Vite to handle bundling of JS dependencies; and Bun to handle the installation of new dependencies.

## Development

This project uses Bun to handle client side stuff. [Pay them a visit to install it first!](https://bun.sh/)
You'll probably need to install [ Air ](https://github.com/air-verse/air) too, and also [ Templ ](https://templ.guide/quick-start/installation). And you know, [ Go ](https://go.dev/doc/install) stuff...

Run these the first time to install deps:
```sh
go mod download
bun install # for the first time building it
```


And the following commands to start your web page at ** http://localhost:3000 ** (Do not confuse with the vite's localhost:5173 that is used for bundling and tooling)
```sh
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

## Notes

I'm personally using this in a small but active production application, and I couldn't be happier.
I've been updating the starter as I made changes to the current active project. Its not a finished
starter by any means, but I'm aiming to keep using and updating it.
