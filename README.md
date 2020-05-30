# grand
[![godoc](https://godoc.org/github.com/hashamali/grand?status.svg)](http://godoc.org/github.com/hashamali/grand)
[![sec](https://img.shields.io/github/workflow/status/hashamali/grand/security?label=security&style=flat-square)](https://github.com/hashamali/grand/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/grand)](https://goreportcard.com/report/github.com/hashamali/grand)
[![license](https://badgen.net/github/license/hashamali/grand)](https://opensource.org/licenses/MIT)

Generates random strings of a specified length with optional seed types.

#### Methods

* `RandomStringWithSeed`: Will seed rand with the given seed, and generate a random string of length N.
* `RandomString`: Will seed rand with current Unix timestamp (in nanoseconds), and generate a random string of length N.
* `TimeSeed`: Returns a timestamp based seed value, the current Unix timestamp (in nanoseconds).

#### Testing

None. ‾\\\_(ツ)\_/‾
