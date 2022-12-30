import * as linux from '@pulumi/linux';

const random = new linux.File('my-file', {
    path: 'my-file.txt',
});

export const output = random.result;
