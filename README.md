# wikintersection
> Little Go application for finding the intersection of two wiki articles. 

![](https://img.shields.io/badge/golang-1.11.5-orange.svg)

This is essentially a learning application for the Go language. It takes two Wikipedia articles from the main English language version and calculates the link data intersection between the two.

The reasoning here is that such an application is a good way to learn the fundamentals of a language. Using this, a developer can explore string handling, network activity, recursion, modularity, mapping, data structures and much more.

## Installation

Windows, OS X & Linux with Golang >1.11:

```sh
go get github.com/dotloadmovie/wikintersection
go build projectmain.go -o wikintersection
```


## Usage example

This isn't a stupendously useful application but it does potentially return better clarifications than a random search-engine search. It's handy for winning pub arguments but ideally needs to be harnessed to the Web UI interface at https://www.github.com/dotloadmovie/wikinteraction-web to be properly user friendly

to run it at command line:

```sh
wikintersection compare ARTICLE_1 ARTICLE_2
```
to return a list of the articles shared by both

```sh
wikintersection search ARTICLE_NAME
```
will return a list of fuzzy matched wikipedia articles to feed the comparison engine


## Release History

* 0.0.1
    * Initial version and first playground

## Meta

Dave Tickle – dave@shiftinteraction.com

Distributed under the unlicense license. See ``UNLICENSE.txt`` for more information.

https://github.com/dotloadmovie/wikintersection

## Contributing

Fork it if you like but I really really wouldn't. This is essentially a learning process so is unlikely to contain perfectly idiomatic Go or the most performant code. It changes at my whim too.


