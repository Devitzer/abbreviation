# v0.6.0

Custom abbreviations are now supported! <br>
The new commands save the custom abbreviations in these directories:

- Windows: C:/Users/(your user)/AppData/abbreviation
- macOS/Linux: ~/.config/abbreviation

## Added

- `add` command: Adds custom abbreviations to the database.
- `remove` command: Removes custom abbreviations from the database.
- `list-user` command: Lists all custom abbreviations from the database.
- New **LICENSE** file to the repository.
- Added `sjson` dependency for editing JSON data.

## Changed

- Version command now lists details about the OS and architecture it's running on.
- Production binaries in GitHub releases are now much more optimized.
- Help command now includes the new commands and has been *slightly* reformatted.
- Search command now also searches for user made abbreviations.
- Modified error messages.

# v0.5.1

## Bug Fixes

- Fixed unfinished line in v0.5.0 changelog notes.
- Fixed issues regarding the version command

# v0.5.0

You may be asking, wait a minute, why is this project starting on v0.5.0? <br>
This project used to be a project built in Rust, but a language conversion was decided after v0.4.1. We will continue as normal from here!

## Added

- Added a couple new abbreviations.

## Changed

- Changed source code language from Rust to Go.
- Init command is gone now! The abbreviation.json file is now embedded into the CLI.