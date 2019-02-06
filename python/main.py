#!/usr/bin/env python3

import sys;

from tlk import sum, RecursiveLoopException;

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
    try:
        print('Total is {:d}'.format(sum(sys.argv[1])))
    except RecursiveLoopException as e:
        print(e)
    except FileNotFoundError as e:
        print(e)

if __name__ == '__main__':
    main()
