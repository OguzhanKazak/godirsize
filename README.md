# godirsize

`godirsize` lists the files and directories under a given directory and calculates the size of each file and directory in bytes. The results are displayed in a tree-like structure.

## Installation

```bash
go install github.com/OguzhanKazak/godirsize
```

## Usage

```bash
godirsize <directory name>
```

The `directory name` argument should be the full path of the directory whose sizes are to be calculated.

For example:

```bash
godirsize /Users/johndoe/Documents
```

This command lists the files and directories under `/Users/johndoe/Documents` and displays their sizes.
