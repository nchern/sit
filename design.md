Important design questions and ideas


### Update repo dir structure.

Introduce new subdirs to store issues



### Should an issue be stored in a single md file or split into user editable file and system fields file?

1 main.md file is a very simple solution

Separating out system fields(modifiable by the utility) and user editable part protects from bad edits of system fields.
Example: `Id` field can be easily edited now in editor.
