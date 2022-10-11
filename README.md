![build](https://github.com/gobi-tools/bee-cli-tool/actions/workflows/go.yml/badge.svg)

# Bee
This is a command line utility to help out with managing your aliases, functions and scripts. 

<img src="/assets/main-screen.jpg?raw=true" alt="Bee Screenshot" width="100%"/>

# Installation

To install, execute this command in your terminal:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gobi-tools/bee-cli-tool/main/install.sh)"
```

This will download [this script](https://github.com/gobi-tools/bee-cli-tool/blob/main/install.sh). 

In turn, this will download a binary into a new folder on your system: `$HOME/.local/bin/bee/`. 

The binary depends on your platform. You can find all options [here](https://github.com/gobi-tools/bee-cli-tool/tree/main/dist).

Finally, in your `.bashrc`, `.zshrc`, `.profile` file, add the following line, to create a shorthand for the script:

```
alias bee='$HOME/.local/bin/bee/bee'
```

# Details

More details [here](https://bee-cli.com/tutorial#title).
