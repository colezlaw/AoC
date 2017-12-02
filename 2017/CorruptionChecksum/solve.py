#! /usr/bin/env python3
import sys
from itertools import permutations


def read(problem):
    with open(problem + '.txt') as f:
        return f.read().strip()


def diff_checksum(table):
    return sum(max(row) - min(row) for row in table)


def div_checksum(table):
    return sum(next(x // y for x, y in permutations(row, 2) if x % y == 0) for row in table)


if __name__ == '__main__':
    table1 = [[int(s) for s in row.split('\t')]
             for row in read('partone').split('\n')]

    print(f'Part One: {diff_checksum(table1)}')

    table2 = [[int(s) for s in row.split('\t')]
              for row in read('parttwo').split('\n')]

    print(f'Part Two: {diff_checksum(table2)}')


