A minimalist example of calling SIMD from Go. Tests included to preserve sanity when copy and pasting blocks of ASM.

The go toolchain is sweet!

Adapted from this [stackoverflow question](http://stackoverflow.com/questions/25460967/go-isnt-linking-my-assembly-undefined-external-function)

Dev Note: There must be at least a single blank line following the body of an assembly file or compilation fails.