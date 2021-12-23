# scraperCookie

This repo is a generalized library. Each specific implementation should include 2-3 components:

    > Proxy (optional)
    > Data scraper
        Implemented with a builder pattern designed to not expose client code
            to a scraper that is only partially constructed
            https://golangbyexample.com/builder-pattern-golang/?__cf_chl_managed_tk__=Yel6VzlV22y4b1iWKlNVx7STpGlu2tQHo52ZSr.RWV0-1639966741-0-gaNycGzNChE
    > Data storer
