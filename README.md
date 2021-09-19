# go-spotify-account-generator
Spotify account generator, just for fun purpose.

## Install
```
$ go get github.com/HamzaAnis/go-spotify-account-generator
```
# How to use

```properties
$ go-spotify-account-generator -h
Usage of go-spotify-account-generator:
  -routines int
    	The number of go routines to run (default 10)
  -size int
    	The total number of accounts to generate (default 100)

$ go-spotify-account-generator -routines 5 -size 10
Spotify Account: XVlBzgbaiCMRAjW@gmail.com, YzRyWJjPjzpfRFE
Spotify Account: whTHctcuAxhxKQF@gmail.com, gmotaFetHsbZRjx
Spotify Account: lgTeMaPEZQleQYh@gmail.com, TCoaNatyyiNKARe
Spotify Account: DaFpLSjFbcXoEFf@gmail.com, EkXBAkjQZLCtTMt
Spotify Account: KJyiXJrscctNswY@gmail.com, NsGRussVmaozFZB
Spotify Account: sbOJiFQGZsnwTKS@gmail.com, mVoiGLOpbUOpEdK
Spotify Account: RsWxPLDnJObCsNV@gmail.com, AwnwekrBEmfdzdc
Spotify Account: updOMeRVjaRzLNT@gmail.com, XYeUCWKsXbGyRAO
Spotify Account: mBTvKSJfjzaLbtZ@gmail.com, syMGeuDtRzQMDQi
Spotify Account: YCOhgHOvgSeycJP@gmail.com, JHYNufNjJhhjUVR
```
## Example: 

To use in another package
```go
package main

import "github.com/HamzaAnis/go-spotify-account-generator/pkg/spotify"

func main() {
	spotify.CreateAccount()
}

```

>Spotify Account: XVlBzgbaiCMRAjW@gmail.com, whTHctcuAxhxKQF