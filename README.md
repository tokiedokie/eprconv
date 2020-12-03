# eprconv

eprconv converts Bruker's binary file into text one

[EasySpin](https://easyspin.org/) has similar features, but it needs to use MATLAB which is not free/open source, so I create this software

This is still unstable

Please feel free to open issue/PR

## Build

```sh
git clone https://github.com/tokiedokie/eprconv
cd eprconv
go build
```

## Usage

```sh
./eprconv -c <path/to/config(*.DSC)> -d <path/to/data(*.DTA)> -o <path/to/output>
```
