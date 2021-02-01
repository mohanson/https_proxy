# HTTPS Proxy

Serves HTTPs requests and proxy it to local HTTP server. It is used to maintain my personal blog [http://accu.cc](http://accu.cc), and it works well.

Note that HTTP service is provided by [https://github.com/mohanson/http_server](https://github.com/mohanson/http_server), and HTTPs certificate is managed by [certbot](https://certbot.eff.org/).

```text
Usage of https_proxy:
  -cert string
        path to tls certificate file
  -priv string
        path to tls private key file
```
