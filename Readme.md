# HTML2PDF

![Cover](images/htmlpdf10.png)


A Tool to get URLS from different sources and convert them into a PDF representation.
This tool uses chrome in its latest version. 

Use ```./install.sh''' to get the latest version.

### Basic usage

Get the binary from releases, unzip and install the only dependency

```bash
cd /temp/

curl https://github.com/epyphite/html2pdf/releases/download/0.1-alpha/linux-dist.tar.gz

tar zxvf linux-dist.tar.gz

cd linux

ls

drwxr-xr-x 2 0 0 4,0K jul  3 13:46 bin
-rw-r--r-- 1 0 0  740 jul  3 13:47 install.sh
-rw-r--r-- 1 0 0  195 jul  3 13:47 sample.cfg

bash install.sh

bin/html2pdf --url [URL]

```

## Building

A ```makefile``` has been added to ease the way feel free to contribute, among perks int he make file


To build the binary

```bash
make build/html2pdf-linux 
```

To Create a copy config files and packages 

```bash
make package
```

If you have token for github, the builds are done with ```release.sh``` but you can issue.

```bash
make distribute
```
