# BARE

A tool to manage all your boilerplate from cli and generate files for you !
## Installation
Currently available for debian based system will release for other systems later
```
npm i -g barego
```
## Using Recipe
### Adding a new bare
```
bare init <bare name> 
```
This makes a `recipe.json` file and a folder at `~/.bare/<bare-name>` directory

Edit the `recipe.json`. 

- `BareName`  : the name of the boilerplate
- `Version` : adds version of the boilerplate `in development`
- `Include` : list all the files in array For eg. `[src/index.js, Dockerfile]` and bare will save this file as your boilerplate
- `Barepath` : DO NOT CHANGE this unless you know what you are doing. This field points to the path where the bare in present in your home directory.

After editing `recipe.json` according to your preference use 
```
bare add
```
Reads your json file and create the template in `~/.bare/<bare-name>`
And you can use this files next time you want to create project!!

### More feature are in development 🏗️ 
### Contribute to make it better
- Give your idea and in report bugs in the `issues` tab
