# Abbreviation - The de-abbreviating CLI!

**The idea is simple;** you enter an abbreviation/acronym/initialism, and it outputs the meaning(s) of it.

## Get Started

Follow [here](#alternate-installation-methods) if you want to install the CLI another way.

1. Install the CLI globally to your system with Go.
```sh
go install github.com/devitzer/abbreviation@latest
```
2. Search for the abbreviation you want!
```
abbreviation search idk
Meaning: I don't know
```

That's all you have to do to start using it!

## Alternate Installation Methods

There is more than one way to install the abbreviation CLI.

### Method 1:

There is binaries for the CLI attached to every GitHub release. We support the following platforms:

- Windows/amd64
- Windows/arm64
- Linux/amd64
- Linux/arm64
- macOS/amd64
- macOS/arm64

You can download the binary and then add it to your path or just use the raw executable whenever you'd like.

### Method 2:

You can build the source code from scratch. You still need Go for this, but it provides a little flexibility if you want to adjust the CLI. <br>
Here is the script we use to build the project. Replace the version with a proper version, such as v1.0.0 for example:

```sh
go build -ldflags="-X 'github.com/devitzer/abbreviation/internal/helpers.version=(**INSERT VERSION HERE**)'"
```

This will provide you a binary in your platform's architecture and OS, which should function well. (cannot guarantee proper functionality if modified.)

## Contributing

There are many abbreviations that exist in the world, as a single developer I can't keep up with every single one that exists. If you want to add your own abbreviations, check out the abbreviation.json file and go ahead!