# ini2env

The goal of this project is to turn the content of an ini file into environment variable.

To do so, it outputs the content of the ini file into:
```
INI__<section>__<key>="<value>"
```
lines that can be eval'd into a shell and then referenced directly.

To mitigate any security issue with eval, it is recommended to parse the output to
weed out anything that doesn't follow the format above. For instance, you can use:
```
grep '^INI__.*=".*"$'
```

You can either build the binary directly via `go get && go build` and use it as such:
```
eval $(./ini2env | grep '^INI__.*=".*"$')
```
, or build the docker container via `docker build -t ini2env .` and then use it as:
```
eval $(docker run --rm  -v $PWD/settings.ini:/file.ini:ro ini2env | grep '^INI__.*=".*"$')
```

## Arguments

You can provide the following optional arguments to `ini2env`:

- `-file`: to set the name of the ini file you want to parse (default is `file.ini`)
- `-prefix`: to change the prefix from "INI" to another value
- `-booleans`: to transform boolean-like values to "1" (for truthy values) or "0"
