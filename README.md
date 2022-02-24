# gobuild-tu
A per-package time usage analytics tool for go build.

### Example
```bash
$ go build -o YOUR_BIN_PATH/gobuildtu cmd/main.go
$ cd ~/gobuild-tu
$ gobuildtu cmd/main.go
Compile Time	Package Name
851.752ms       runtime
344.623ms       reflect
230.72ms        syscall
152.464ms       time
135.262ms       strconv
125.929ms       strings
125.13ms        math
120.448ms       encoding/json
115.443ms       fmt
113.393ms       bytes
109.972ms       unicode
103.612ms       os
103.247ms       command-line-arguments
83.994ms        internal/reflectlite
83.103ms        internal/cpu
74.937ms        sync/atomic
72.967ms        os/exec
69.473ms        runtime/internal/atomic
69.05ms         encoding/binary
67.998ms        internal/poll
61.983ms        internal/bytealg
55.334ms        math/bits
55.245ms        io
50.98ms         io/fs
43.294ms        path/filepath
41.702ms        sync
41.071ms        unicode/utf8
40.447ms        context
40.372ms        runtime/internal/sys
40.097ms        internal/race
36.856ms        path
36.668ms        sort
33.44ms         runtime/internal/math
32.95ms         io/ioutil
30.711ms        internal/fmtsort
28.51ms         encoding
28.218ms        internal/unsafeheader
27.564ms        unicode/utf16
27.051ms        encoding/base64
20.893ms        internal/syscall/unix
19.911ms        errors
18.691ms        internal/syscall/execenv
16.92ms         command-line-arguments
15.907ms        internal/testlog
15.497ms        github.com/gobuild-tu/analyze
14.641ms        internal/oserror
811µs           command-line-arguments
0s              unsafe
```
