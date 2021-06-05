SIT - stands for the Simplest Issue Tracker

Proof of Concept!

## Main ideas

* Be able to keep tickets / issues together with the code, not somewhere else
* Being distributed(as git)
* Full integration with git
* Open standards / data governance (needs explanation)
* Allow non-cli users as well
* SIMPLICITY

## HOWTOs

### Install

You need Golang to install it from source.

```sh
make install
```

### Usage

To edit tickets the app uses `$EDITOR` if you have it set, otherwise `vim` is used by default.

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

### Collaborate

Sit repo is just a folder with plain text files. It can be a part of any code repo
including standalone repos using any scm, e.g. git. The current view is that git / git workflow
can cover the majority of collaboration use cases, like sharing, having a central repo for all issues, etc.
In future tighter integrations with git could be added along with more tooling.

## Issues

[issues.md](issues.md)

## Known (serious) questions

1. How to make it user friendly?
2. Short and easily usable / memorizable issue IDs
3. Authentication / authorisation(the current idea is to rely on git signed commits)
4. Workspaces / different projects 
   1. How this will work together?
   1. What should be the data model?
5. How different teams could be accounted here?
