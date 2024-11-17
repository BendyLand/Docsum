# docsum

A simple CLI tool which recursively reads all the files in a given directory matching a certain file extension, then writes their names and contents into an single file called 'summary.txt'.

## Examples

The intended usage can be seen with the test directory in this project.

Running:
```bash
docsum . c
```
Produces the 'summary.txt' that is found in the root directory of the project, also shown below:
```txt
// 1.c
int main()
{
    
}

// 2.c
#include <stdio.h>

int main(void)
{
    return 0;
}

// 3.c
int main() {
    
}
```

Likewise, running:
```bash
docsum test py
```
Would write the following the a 'summary.txt':
```txt
# 1.py
print("Hello world!")
```

If you have a file type that doesn't support comments, like .txt, or you're using a language that isn't currently supported, the file names will be preceded with a hyphen:
```bash
docsum test txt
```
Would produce:
```txt
- 1.txt
Here is some sample
text for you
to read
in the text file.
```

## Future Additions

 - I want to add the option to simply print the result instead of writing it to a file.
 - I would also like to add a --maxdepth flag, in case you want to limit the amount of recursive searching.
