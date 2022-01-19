# Contributing

Thank you for interest in our project. We appreciate the time you took for contributing. Follow the steps to get started.

If you have some issue that you want to bring to our focus, Please head [here](https://github.com/bare-cli/bare/issues) to file an issue and let us know what's troubling you.

If you want to contribute by code:

## Steps for setup

- Fork the repo from [here](https://github.com/bare-cli/bare)
- Clone the forked repo by git:

  ```
  git clone https://github.com/<username>/bare <folder-name>
  ```

  or by [gh](https://github.com/cli/cli) :

  ```
  gh repo clone <username>/bare <folder-name>`
  ```

- Change directory to your local copy:
  `cd <folder-name>`
- Setup the project with:
  ```
  make setup
  ```
- Run development envrionment

  > We use [`air`](https://github.com/cosmtrek/air) to live reload while in development envrionment

      make dev

- Build the tool using make:
  ```
  make build
  ```
  Now you will have your tool in `tmp` folder, you can try it out following `Readme.md`.
  If you want to make it available globally you can add `alias` in `.bashrc` or `.zshrc`
