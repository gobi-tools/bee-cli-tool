# Bee
This is a command line utility to help out with managing your aliases, functions and scripts. 

# Installation

To install, execute this command in your terminal:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gobi-tools/bee-cli-tool/main/install.sh)"
```

This will download [this script](https://github.com/gobi-tools/bee-cli-tool/blob/main/install.sh). 

In turn, this will download a binary into a new folder on your system: `$HOME/.local/bin/scripthub/`. 

The binary depends on your platform. You can find all options [here](https://github.com/gobi-tools/bee-cli-tool/tree/main/dist).

Finally, in your `.bashrc`, `.zshrc`, `.profile` file, add the following line, to create a shorthand for the script:

```
alias bee='$HOME/.local/bin/scripthub/bee'
```

# Details

More details [here](https://google.com).