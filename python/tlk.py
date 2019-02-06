class RecursiveLoopException(Exception):
    def __init__(self, message='Recursive loop', file=None):
        self.message = message
        self.file = file
    def __str__(self):
        return str(self.message)
    pass

def isNumber(s):
    try:
        int(s)
        return True
    except ValueError:
        return False

def sum(path, previouses=[]):
    total = 0
    for prev in previouses:
        if prev == path:
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
                total += sum(line, previouses + [path])
    print('Subtotal for {:s} is {:d}'.format(path, subtotal))
    return total + subtotal
