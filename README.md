# crawl-links

[![Circle CI](https://img.shields.io/circleci/project/crackcomm/crawl-links.svg)](https://circleci.com/gh/crackcomm/crawl-links)

Links crawler.

## Usage

Example usage from command line:

```sh
# Install command line application for crawl scheduling
$ go install github.com/crackcomm/crawl/nsq/crawl-schedule
# Schedule crawl of google search results
$ crawl-schedule \
      --nsq-topic crawl_links \
      --callback github.com/crackcomm/crawl-links/spider.Links \
      "https://www.google.com/search?q=Github"
```

## License

                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

## Authors

* [Łukasz Kurowski](https://github.com/crackcomm)
