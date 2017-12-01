#! /usr/bin/env python3
import sys

def pairwise(iterable, offset=1):
    for index, a in enumerate(iterable):
        b = iterable[(index + offset) % len(iterable)]
        yield a, b

captcha = sys.stdin.readline().strip();

result=sum(int(char) for char, next in pairwise(captcha) if char == next)
print(result)