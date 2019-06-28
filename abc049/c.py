#!/usr/bin/env python3
# https://atcoder.jp/contests/abc049/tasks/arc065_a

s = input()[::-1]
collector = ''
t = ''
words = ['remaerd', 'maerd', 'resare', 'esare']

for c in s:
    collector += c
    if collector in words:
        t += collector
        collector = ''

if s == t:
    print('YES')
else:
    print('NO')
