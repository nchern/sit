Important design questions and ideas


### Update repo dir structure.

Introduce new subdirs to store issues. The basic idea here is to keep new issues in a separate folder.

Example 1:

`.sit/issues/new`

`.sit/issues/all`

Example 2:

`.sit/issues/new`

`.sit/issues/in-progress`

`.sit/issues/completed`

`.sit/issues/undefined` - special folder to put issues that do not fall into any of the previous categories



### Should an issue be stored in a single md file or split into user editable file and system fields file?

One `main.md` file is a very simple solution

Separating out system fields(modifiable by the utility) and user editable part protects from bad edits of system fields.

Example: `Id` field can be easily edited now in editor.
