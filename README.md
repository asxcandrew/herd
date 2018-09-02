# HERD - Federated publishing platform

> Federated publishing platform based on ActivityPub connected to the world of free social networking

Demo instance [herd.anonbecon.com](http://herd.anonbecon.com/)

It's not ready for production yet, check for updates our mastodon account [`@herd@mastodon.technology`](https://mastodon.technology/@herd) or matrix room: [`#herd:matrix.org`](https://riot.im/app/#/room/#herd:matrix.org/)

---

## Tech stack

This project combines several modern technologies that I like:

- vue.js on frontend
- Golang (go-gin) for backend
- Protocol Buffers and gRPC
- Docker

> Anyone who's slightly familiar or want to try one of this technologies is welcomed to join.

*([Submit a pull request](https://github.com/asxcandrew/herd/pulls))*

## Development
With [docker-compose](https://docs.docker.com/compose/) installed, run

```
$ make setup-dev
```

You can now execute the `make run-dev` command to start up the development.
Project server will be running at *http://localhost:90*

## Deployment

You need to create running postgres with created database before deploy.

```
> psql -c "CREATE DATABASE herd_production_db"
CREATE DATABASE
```
 - deployment section here

## What's next

Checklist to gauge how Herd`s development is going:

- [x] Setup application with development environment
- [ ] Full-featured text editor on client (in progress)
- [ ] Integrate basic ActivityPub features
- [ ] Use external service for image processing
- [ ] HTTPS by default
- [ ] IPFS for pictures storage as option
- [ ] Integration, unit testing

## Why?

Software that implements ActivityPub can be wildly different in how it looks and what it does! But the social graph–what we call the people and their connections–is the same.

- For rich juicy blogs, for users used to modern functionality Herd is the choice.
- Mastodon is the software built around 500-character text posts.
- You want a video platform? That’s PeerTube.
- You want something centered on photos and pictures? PixelFed is here.
- You want to write long, rich blog posts? Plume is in development.

## See Also

[`Why ActivityPub is the future`](https://blog.joinmastodon.org/2018/06/why-activitypub-is-the-future/)!

## License

