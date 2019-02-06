#!/usr/bin/env python3

class RecursiveLoopException(Exception):
    """Recursive loop"""
    pass

def isNumber(s):
    try:
        int(s)
        return True
    except ValueError:
        return False

def sum(path, previouses=[]):
    for prev in previouses:
        if prev is path:
            print(previouses)
            raise RecursiveLoopException
    subtotal = 0
    with open(path, 'r') as f:
        while True:
            line = f.readline().strip()
            if not line:
                break
            if isNumber(line):
                val = int(line)
                subtotal += val
            else:
                previouses.append(path)
                sum(line, previouses)
    print('Subtotal for {:s} is {:d} and Total is {:d}'.format(path, subtotal, total))
