SIT - stands for the Simplest Issue Tracker

Prove of Concept!

### Install

To edit tickets the app uses `$EDITOR` if you have it set. If it is empty `vim` is used by default.

```sh
make install
```

### Usage

```sh
# initialize
sit init

# create issue
sit create

# list issues in this repo
sit list

# edit issue, accepts parts of the id
sit edit <partial-issue-id>

# delete issue
sit delete <partial-issue-id>
```
