import file from 'k6/x/filenio';

const filepath = 'sample-output.txt';

export default function () {
    file.writeString(filepath, 'New file. First line.\n');
    file.appendString(filepath, `Second line. VU: ${__VU}  -  ITER: ${__ITER}`);
    file.removeLines(filepath, 1, 1);
}