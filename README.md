# ledgerquote
ledgerquote, or `lq`, is a small command line utility for downloading quotes for a Ledger CLI price file.

## Usage

Get quotes for current day
```
$ lq prices.dat --from 2021-01-01 --to 2020-01-10 --overwrite 
```

## Current support
Nothing at the moment. Tool is being developed.

## Roadmap
- [ ] parse prices.dat file
- [ ] date range
- [ ] single commodity name
- [ ] overwrite flag
- [ ] first source of fund price quotes
- [ ] first source of equity price quotes
- [ ] config file: date format
- [ ] config file: ignores
- [ ] config file: mapping
