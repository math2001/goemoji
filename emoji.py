# -*- encoding: utf-8 -*-

import sys
import codecs
import json

def load_emojis():
    with open('emoji.json', encoding='utf-8') as fp:
        return json.load(fp)

def print_usage():
    print("Just pipe me!")
    print('\necho hello world :tada: | python emoji.py')

def main():
    return print(unicode('😀'))
    if not sys.stdin.isatty():
        data = sys.stdin.read()
    else:
        print_usage()
        exit()

    emojis = load_emojis()

    for emoji in emojis:
        if 'emoji' not in emoji.keys():
            continue
        for alias in emoji['aliases']:
            data = data.replace(':' + alias + ':', emoji['emoji'])

    with open(sys.stdout.fileno(), 'w', encoding='utf-8') as fp:
        print(data, file=fp)
    print('😀')

if __name__ == '__main__':
    main()
