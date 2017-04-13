# revel-stream

A rudimentary multi-user DASH streaming service with user authentication and a simple web interface,
using Revel and nginx.

### Configuration:

This app is designed to be run on top of nginx, using [nginx-rtmp-module](https://github.com/sergey-dryabzhinsky/nginx-rtmp-module).
An example nginx configuration is included in the root directory of this repository.

The main app configuration can be found in `conf/app.conf`. Here you can change
the location of your SQLite database file, your DASH path, and your output stream's
base URL, along with the standard `app.conf` settings (more information [here](https://revel.github.io/manual/appconf.html)).

Routes can be changed to fit the needs of your service in `conf/routes`. See [here](https://revel.github.io/manual/routing.html) for more information.

You will also need to set up a SQLite database to store your user information in. Currently,
there is no user sign-up component, so users must be manually added to this database.
The database schema should be as follows:

```
CREATE TABLE `users` (
	`user`	TEXT NOT NULL UNIQUE,
	`key`	TEXT NOT NULL UNIQUE
);
```

(Where `user` is the username of each user, and `key` is their unique, private stream key)

### Building the frontend:

The frontend utilizes Sass and JavaScript with Node.js, and utilizes npm for package management.
Information about installing Node.js and npm can be found [here](https://www.npmjs.com/get-npm?utm_source=house&utm_medium=homepage&utm_campaign=free%20orgs&utm_term=Install%20npm).

To install all Node.js packages, run:

```
npm install
```

The source files can then be compiled by running:

```
npm run build:webpack
npm run build:gulp
```

Alternatively, you can have webpack and gulp watch your source files for changes:

```
npm run watch:webpack
```

and

```
npm run watch:gulp
```

### Starting the web server:

The web server can be started by running:

```
npm run serve
```

In production, it should be run with:

```
revel run github.com/allie/revel-stream prod
```

## Repository structure:

The directory structure of this repository is as follows:

```
conf/             Configuration directory
    app.conf      Main app configuration file
    routes        Routes definition file

app/              App sources
    init.go       Interceptor registration
    controllers/  App controllers go here
    views/        Templates directory

public/           Public static assets
    css/          Compiled CSS files
    js/           Compiled Javascript files
    assets/       Image files

frontend/         Frontend source files
    sass/         Sass source files
	js/           Javascript source files
	assets/       Image files
```

## Help with Revel:

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).
