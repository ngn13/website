# My Website | [ngn13.fun](https://ngn13.fun)
This repo contains the source code of my personal website.
It's written NuxtJS and supports full SSR. As database,
it uses mongodb. It's just a personal project that
I am working on.

## Setup
For some reason if you want to setup my website localy
install `nodejs` and `npm`, then run the following:
```bash
git clone https://github.com/ngn13/ngn13.fun.git && cd ngn13.fun &&
npm i
echo "PASS=password" > .env
echo "DATABASE=mongodb://127.0.0.1" > .env
npm run build
npm run start
```
