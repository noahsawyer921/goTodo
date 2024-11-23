# Todo CLI

A basic todo list app that keeps track of tasks. Tasks are tracked on files and will remain between restarts.

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
2. `push`: Prompts you to create a new item by letting you fill out all of the necessary fields. This item is then appended to the end of the list
3. `pop`: Removes the latest item from the list and prints it.

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











