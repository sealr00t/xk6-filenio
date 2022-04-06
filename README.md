# xk6-file
[k6](https://github.com/grafana/k6) extension for writing files, implemented using the
[xk6](https://github.com/grafana/xk6) system.

## Build
```shell
xk6 build v0.37.0 --with github.com/sealr00t/xk6-file@latest
```

## Example
```javascript
import file from 'k6/x/filenio';

const filepath = 'sample-output.txt';

export default function () {
    file.writeString(filepath, 'New file. First line.\n');
    file.appendString(filepath, `Second line. VU: ${__VU}  -  ITER: ${__ITER}`);
    file.removeLines(filepath, 1, 1);
}
```

## Run sample script
```shell
./k6 run sample-script.js
```
