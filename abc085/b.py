#! /usr/bin/env python3

n = int(input())
d = []
result = 1

for _ in range(0, n):
    d.append(int(input()))

d.sort()
for i in range(1, len(d)):
    if d[i-1] < d[i]:
        result += 1

print(result)
