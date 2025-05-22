As a privacy advocate myself I do my best to respect your privacy while simultaneously keeping my server safe. I also think it's
important to be transparent about this kind of stuff so I wrote this document to tell you how exactly I process or store your
information.

## DNS & SSL
Currently I use cloudflare's name servers, however cloudflare doesn't own my domain (I didn't purchase the domain from cloudflare)
nor it proxies any of the traffic. All of my records use the *DNS only*" mode, meaning they are just DNS records and they point
to my server, not cloudflare. This also means cloudflare doesn't control any of the SSL certificates. All the certificates are
stored on my server and they are created with Let's Encrypt. So I own the certificates and it's not possible for cloudlfare to
suddenly switch DNS records (it would break SSL).

## Usage metrics
I don't have any kind of request or traffic monitoring on my server. So no, your HTTP(S) requests or other network
connections are not processed to generate colorful graphs, pie charts and stuff.

You may have realized that in the bottom of this page there is total visitor number tracker. It is the only
usage/visitor metric tracking I have, which is implemented by website's API, which is free (as in freedom) so you can
go audit it's code yourself.

I want to make it clear that this metric tracker does not store any information about you or your HTTP(S) requests to
a database. It temporarily stores your IP address' SHA1 hash, in memory, so it doesn't recount the same visitor again
when they refresh the website or visit it multiple times in a short period of time. After a certain amount of requests,
your IP address' SHA1 hash will be removed from the memory and it will be replaced with a new visitor's SHA1 hash instead.

## Logs
All the HTTP(S) services are proxied with nginx, and nginx logs all of them to a file on the disk. This file (`access.log`)
is only readable by the root user, and it is contents are deleted every 4 hours (with `shred` to make sure it doesn't leave
anything on the disk). The logs include *only* the following information:

- Request time
- Requested host
- Requested path
- HTTP request method
- HTTP response code

This is the minimal information I need to trace any issues if something goes wrong, which is the main reason why I use
logging at all, to make it easier to find any issues.

## Data removal
If you want to remove any of your data from my server, [you can send me an email](mailto:ngn@ngn.tf). And yes this includes
removing usage metrics and logs.
