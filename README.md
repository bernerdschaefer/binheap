# binheap

The binheap package provides a min-heap implementation for storing arbitrary
values with priority values. Unlike other heap packages, the priority is
defined as separate from the value stored, where the priority can be any
integer, and the smaller the number the higher the priority.

See `example_test.go` for an imagined implementation of a work queue
with two heaps -- priority and delay -- that store jobs.
