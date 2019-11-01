# Fast

A tool to compare a client-side performance for a feature branch against master.
Specifically, [Lighthouse](https://developers.google.com/web/tools/lighthouse/)
is used for the comparison.

Run the tool without any arguments to see available command line options.

## Build

This will build a binary `fast` which you can optionally move to a location on
your `$PATH`:

```
$ go build -o fast main.go
```
