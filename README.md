# scraperCookie

This repo is a generalized library. Each specific implementation should include 2-3 components:

    > Proxy (optional)
        TODO

    > Scraper
        Implemented with a builder pattern designed to not expose client code
        to a scraper that is only partially constructed
        https://golangbyexample.com/builder-pattern-golang/?__cf_chl_managed_tk__=Yel6VzlV22y4b1iWKlNVx7STpGlu2tQHo52ZSr.RWV0-1639966741-0-gaNycGzNChE

    > Store
        Automated store as an abstract interface
            type IStore interface {
                Init()
                Store(l Locator, data io.Reader) error
                Read(l Locator) []byte
            }
