#!/usr/bin/env python3

import unittest

from tlk import sum, RecursiveLoopException

class TestSumFunction(unittest.TestCase):

    def test_normal(self):
        self.assertEqual(sum('tests/test_set/a.txt'), 221)
        self.assertEqual(sum('tests/test_set_simple/a.txt'), 10)

    def test_simple_loop(self):
        self.assertRaises(RecursiveLoopException, sum, 'tests/test_set_loop_simple/a.txt')

    def test_complex_loop(self):
        self.assertRaises(RecursiveLoopException, sum, 'tests/test_set_loop_complex/a.txt')

if __name__ == '__main__':
    unittest.main()
