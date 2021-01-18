# ledgerquote
ledgerquote, or `lq`, is a small command line utility for downloading quotes for a Ledger CLI price file.

## Usage

Get quotes for current day
```
Usage: lq <pricefile>
  -config string
        Config file to use
  -from string
        Date range start (inclusive)
  -nooverwrite
        Don't overwrite existing prices (default false)
  -to string
        Date range end (inclusive)
```

By default `lq` looks for configuration file `~/.lqrc` for configuration.

### Examples

Updated all prices in file `prices.dat` only for todays date:
```
> lq prices.dat
```

Updated all prices in file `prices.dat` for today's date, but do not overwrite existing prices:
```
> lq prices.dat --nooverwrite
```

Updated all prices in file `prices.dat` from given start date (inclusive) `2021-01-01` until today (inclusive):
```
> lq prices.dat --from 2021-01-01
```

Updated all prices in file `prices.dat` from given start date (inclusive) `2021-01-01` until given end date (inclusive):
```
> lq prices.dat --from 2021-01-01 --to 2021-01-10
```

Updated all prices in file `prices.dat`, given a specific configuration file as input
```
> lq prices.dat --config lq.conf
```

## Configuration file
By default `lq` looks for, and uses, configuration file
`~/.lqrc`. In this file you can specify  

Here is an example file:
```
file=~/prices.dat
nooverwrite
```

Which specifies that `lq` should always use the `--nooverwrite flag` and it should always use the input file `~/prices.dat`.
With such a config file, one could simply run: 

```
> lq
``` 

which, without the config file, would correspond to running:

```
> lq ~/prices.dat --nooverwrite
```

### Mappings
TBD

## Current support
TBD

## Roadmap
- [ ] config file: default arguments
- [ ] parse prices.dat file
- [ ] date range
- [ ] single commodity name
- [ ] overwrite flag
- [ ] first source of fund price quotes
- [ ] first source of equity price quotes
- [ ] config file: date format
- [ ] config file: ignores
- [ ] config file: mapping
