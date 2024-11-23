# Todo CLI

A basic todo list app that keeps track of tasks. Tasks are tracked on files and will remain between restarts. This system should be portable between all operating systems, but it is only tested and maintained for MacOS.

## List Info

Lists are stored in `[install_directory]/todoLists/`. They are stored in the .csv format and are not encrypted. 

Lists are stored with 0664 permissions, such that the owner and group of the owner can read and write, while all other users can read each list. If you wish to change this behavior, you will need to alter the file permission constant in `writeList`. Here are some suggested alternatives:

| Permission | Owner | Group | Other | Description |
| ---------- | ----- | ----- | ----- | ----------- |
| 0666       | RW    | RW    | RW    | Fully public lists, readable and writeable by anyone. |
| 0664       | RW    | RW    | R     | <ins>Default.</ins> Public lists, only writeable by the owner and their group. |
| 0644       | RW    | R     | R     | Public lists, only writeable by the owner. |
| 0644       | RW    | RW    |       | Readble and writeable by the owner and their group. |
| 0640       | RW    | R     |       | Readable by the owner and their group, only writeable by owner. |
| 0600       | RW    |       |       | Private lists, only readable and writeable by the owner. |

## Compilation

```go build main.go```

## Run

```./main```

## Exit

Whenever prompted for a command, you may run the `exit` command to exit safely.

Data processing and file writing is only performed when full operations are ready to complete, so you may also exit at any time a prompt is visible with `Ctrl+C` and not risk losing your data.

## Create a new task list

The first thing you must do is create a list to track your tasks.

1. Start the cli with `./main` or any alias you have set up to start the program.
2. Call the `init` command inside of the cli
3. Provide the name of the new list when prompted
4. You may now exit. Your lists are saved, so you don't have to worry about losing them when you exit.

Lists must be uniquely named, so you will fail to init a list if a name with that list already exists.

## Select a list

Interacting with a list requires having a list selected for commands to run on. You can select a list with the `select` command

1. Start the cli with `./main` or any alias you have set up to start the program.
2. Call the `select` command inside of the cli.
3. Provide the name of the list you want to select. It must be the name of an existing list you have created with the `init` command before.
4. That list will remain selected until you select a new list or exit the program.

## Operations on a selected list

You may perform any of the following operations on your selected list whenever you have one selected:

1. `ls`: Prints all items in that list
2. `push`: Adds a new item to the end of the list.
3. `pop`: Removes the latest item from the list and prints it.
4. `enqueue`: Adds a new item to the start of the list.
5. `dequeue`: Removes an item from the end of the list and prints it.

# Todo items

## Expected

The following items are expected features that are still in progress:

1. ~~Queue functionality: support for `enqueue` and `dequeue` commands.~~
2. Editing items
3. Viewing individual items
4. Use structs for data rather than raw string slices
5. Standardize use of error responses and error handling.

## Desired

The following items are features that may be implemented in the future, but are not guaranteed:

1. Transaction-based editing that requires saving rather than immediately saving all actions
2. Deleting lists completely
3. Dedicated priority values, rather than a free string, as well as sorting by priority
4. Dedicated due date format, rather than a free string, as well as sorting by due date
5. Inserting at specific indices
6. Compiling command inputs as args to each command, rather than separate inputs











