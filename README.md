# goth-template
A template for the goth stack

Go + Htmx + Tailwind + DaisyUi + Templ

You have to install goose and sqlc respective cli to use goose and sqlc.

```console
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Of course you'll need the templ cli.
```console
go install github.com/a-h/templ/cmd/templ@latest
```

```console
git clone https://github.com/MKamelll/goth-template.git
cd goth-template
npm install
npm run dev
```

To create a migration
```console
npm run goose:create <Your migration name>
```

To migrate
```console
npm run goose:migrate
```

To generate the go code from your sql files
```console
npm run sqlc:generate
```
