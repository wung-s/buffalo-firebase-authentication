This example repository is part of a [blog post](https://medium.com/@wung_s/integrate-firebase-with-buffalo-for-api-authentication-d3d33ede8e0f) on implementing API authentication with [Buffalo](http://gobuffalo.io) and [Firebase](https://firebase.google.com/)

### Clone Repository

    $ git clone git@github.com:wung-s/buffalo-firebase-authentication.git firebase_authentication && cd firebase_authentication

### Create Your Databases

    $ buffalo db create -a

## Starting the Application

    $ buffalo dev

## Heroku

### Create app

    $ heroku create my-app-name

Note: Necessary only if you have **NOT** deployed to Heroku yet

### Set Environment Variables

    $ heroku config:set GO_ENV=production --app=my-app-name
    $ heroku config:set FB_SERVICE_AC_KEY=serviceAccountKey.json --app=my-app-name

### Deploy

    $ heroku container:login
    $ heroku container:push web

[Powered by Buffalo](http://gobuffalo.io)
