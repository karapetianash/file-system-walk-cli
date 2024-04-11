## Description
This cross-platform CLI tool crawls into file system directories looking for specific files. When the tool finds the files it’s looking for (filtered files), it can list, archive, or delete them.
This is also a useful tool that helps you back up and clean up file
systems.

## Supported flags
This version accepts next command-line parameters:
- `-root root/directory/path`\
The root of the directory tree to start the search. The default is the
current directory.
- `-list`\
List files found by the tool. When specified, no other actions will
be executed.
- `-ext ".ext1 .ext2..."`\
File extensions (one or more) to search. When specified, the tool will only match
files with thees extensions.
- `-size int`\
Minimum file size in bytes. When specified, the tool will only
match files whose size is larger than this value.
- `-del`\
Deletes specified file or all filtered ones.
- `-log file/path`\
If specified, provides feedback about deleted files in log file. Otherwise, prints to STDOUT.
- `-archive file/path`\
Archives compressed files before deleting.
- `-since time_in_RFC822Z_format`\
File last modification time to search. When specified, the tool will only 
match files which last time modified after this value.

## Usage
To build the app, run the following command in the root folder:

```
> go build .
```
Above command will generate `walk` file. This name is defined in the `go.mod` file, and it will be the initialized module name.

After that you can run the program using the cmd.\
Example of listing filtered files of root directory:

```
> .\walk.exe -root root/directory/ -ext ".txt .rar"
```
