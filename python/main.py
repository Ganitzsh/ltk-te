#!/usr/bin/env python3

import sys;

from tlk import sum;

def printUsage():
    print("""
Usage: sums file.txt

Description

This program takes a file and counts the total of each numbers at each line
including the subtotal of other subsequent files specified instead of a
number
    """)

def main():
    if len(sys.argv) is 1 or len(sys.argv) > 2:
        printUsage()
        return
    total = 0
    try:
        sum(total, sys.argv[1])
    except Exception as e:
        print('Something went wrong')
    print('Total is {:d}'.format(total))

if __name__ == '__main__':
    main()
