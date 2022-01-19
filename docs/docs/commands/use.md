This command generates your project from any bare-template on github(for now).<br/>
You can select the variant of template you want, the variable in the template, so you dont have to do any new configurations to start your project.

### Usage

```shell
bare use <user>/<repo> <output path>
```

- **User** : User or organization name of repo on github.
- **Repo** : Repo name
- **output path** : The path where you want to create the project.

For example, [**`bare-cli`**](https://github.com/bare-cli) has an example repository.

```
bare use bare-cli/vanilla-js-template my-new-project/
```

This will prompt you to select the variant of the template and all the variables required by the template.
Boom! your project is ready !!!
