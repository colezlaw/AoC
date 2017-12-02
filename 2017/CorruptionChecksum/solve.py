#! /usr/bin/env python3
import sys
from itertools import permutations


def read(problem):
    with open(problem + '.tsv') as f:
        return f.read().strip()


def diff_checksum(table):
    return sum(max(row) - min(row) for row in table)


def div_checksum(table):
    return sum(next(x // y for x, y in permutations(row, 2) if x % y == 0) for row in table)


if __name__ == '__main__':
    spreadsheet = [[int(s) for s in row.split('\t')] for row in read('spreadsheet').split('\n')]

    print(f'Part One: {diff_checksum(spreadsheet)}')
    print(f'Part Two: {div_checksum(spreadsheet)}')


