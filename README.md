# ChiConverter

ChiConverter is a command line tool to convert chinese from Traditional to Simplified or Simplified to Traditional.

## Command

| Command | Description |
| ---- | ---- |
| download | download and remove current text config files |

## SubCommand

| Command | Description |
| ---- | ---- |
| s2t | convert Simplified Chinese to Traditional Chinese | 
| t2s | convert Traditional Chinese to Simplified Chinese | 
| s2tw | convert Simplified Chinese to Traditional Chinese (Taiwan Standard) | 
| tw2s | convert Traditional Chinese (Taiwan Standard) to Simplified Chinese | 
| s2hk | convert Simplified Chinese to Traditional Chinese (Hong Kong Standard) | 
| hk2s | convert Traditional Chinese (Hong Kong Standard) to Simplified Chinese | 
| s2twp | convert Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom | 
| tw2sp | convert Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom | 
| t2tw | convert Traditional Chinese (OpenCC Standard) to Taiwan Standard | 
| t2hk | convert Traditional Chinese (OpenCC Standard) to Hong Kong Standard | 

## SubCommand Flag

| Flag | Description |
| ---- | ---- |
| -s | source file name |
| -d | output file location |
| -h | show all `chiconverter` command and flag description |

## How To Use

```
# download config files, first use `chiconverter` should download config files
$ ./chiconverter download

# convert chinese
$ ./chiconverter <subcommnad> -s <source file> -d <output location>
```

## Support

* [x] linux
* [ ] mac
* [ ] windows
