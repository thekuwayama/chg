# chg

[![Build Status](https://travis-ci.org/thekuwayama/chg.svg?branch=master)](https://travis-ci.org/thekuwayama/chg)
[![MIT licensed](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://raw.githubusercontent.com/thekuwayama/chg/master/LICENSE.txt)

`chg` is Concurrent Http Get command.


## Usage

```bash
$ chg -help
Usage of chg:
  -n int
        number of workers (default 4)
```

```bash
$ cat << EOS | chg
> https://YOUR_URL_1
> https://YOUR_URL_2
> EOS
```


## License

The CLI is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
