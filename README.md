# shell-mommy-promax
the super awesome, highly customizable, revolutionary terminal assistant that doesn't exist (WIP).

![Copilot_20250606_213843](https://github.com/user-attachments/assets/1a179924-9888-4c0a-85f6-8e0aa1d927de)


## Installation Instruction
just build and copy to some where in path (or just use absolute path when calling who cares)
```sh
go build ./cmd/mommy.go
```

## Usage
for bash add this to your ~/.bashrc
```bash
PROMPT_COMMAND="/path/to/mommy $PROMPT_COMMAND"
```

for zsh add this to your ~/.zshrc
```zsh
precmd() { /path/to/mommy }
```

or if you want to be fancy:
```zsh
precmd_functions+=(mommyFunction)
mommyFunction() { /path/to/mommy }
```

## To-Do:
- [x] this command can run after every command exits from shell
- [ ] detect executed commands and respond to them somehow
- [ ] fish support
