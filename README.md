# Emerge things in an interactive way

With this tool, you'll be able to run complex emerge commands without having the head in the manual.

## Todo

- [ ] Find a smart way to generate suggestions from `emerge` manual.
- [ ] Auto-complete package name when installing / removing a package.
- [ ] Add this tool to `go⁻overlay` when the two previous points will be done.
- [ ] Provide binary

## Installing from sources

Currently, you need to have `golang` installed and `GO111MODULE=on` to get dependencies. Later, a binary will be available on release page.

```shell
$ git clone https://github.com/tormath1/emerge-interactive
$ cd emerge-interactive
$ make install
```

Make sure that your `${GOPATH}/bin` is in your system ${PATH}.

## Profit

```shell
$ emerge-interactive 
emerge-interactive v0.0.1
Run `Ctrl+d` or `exit` to quit
> 
```

## Uninstalling

Just run `$ make uninstall`

