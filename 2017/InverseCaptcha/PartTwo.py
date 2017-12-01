#! /usr/bin/env python3
import sys

def pairwise(iterable, offset=1):
    for index, a in enumerate(iterable):
        b = iterable[(index + offset) % len(iterable)]
        yield a, b

captcha = sys.stdin.readline().strip();

assert len(captcha) % 2 == 0 # Is even

result=sum(int(char) for char, next in pairwise(captcha, len(captcha)//2) if char == next)
print(result)