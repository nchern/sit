SIT - stands for the Simplest Issue Tracker

Prove of Concept!

## Main ideas

Main ideas behind this tool are:
* Be able to keep tickets / issues together with the code, not somewhere else
* Being distributed(as git)
* Full integration with git
* Open standards / data governance // needs explanation?
* Allow non-cli users as well
* SIMPLICITY

## HOWTOs

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

# show help
sit help
```

## Known (serious) questions

1. How to make it user friendly?
2. Short and easily usable / memorizable issue IDs
3. Authentication / authorisation(the current idea is to rely on git signed commits)
4. Workspaces / different projects 
   1. How this will work together?
   1. What should be the data model?
5. How different teams could be accounted here?
