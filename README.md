# zzz

zzz is a highly opinionated fork of nap

To move to zzz just rename all your previously defined `NAP_` env vars to `ZZZ_`

<img width="1200" alt="zzz" src="https://user-images.githubusercontent.com/42545625/202545409-eb53f92a-233a-4f78-b598-a59c65248ad3.png">

<sub><sub>z</sub></sub><sub>z</sub>z

Zzz is a code snippet manager for your terminal. Create and access new snippets
quickly with the command-line interface or browse, manage, and organize them with the
text-user interface. Keep your code snippets safe, sound, and well-rested in your terminal.

<br />

<p align="center">
<img width="1000" src="https://user-images.githubusercontent.com/42545625/202577549-f2e0887a-b740-41f4-9408-c2f53673503f.gif" />
</p>

<br />

## Text-based User Interface

Launch the interactive interface:

```bash
zzz
```

<img width="1000" src="https://user-images.githubusercontent.com/42545625/202768989-caf2ab62-b69d-4e2d-ac93-1517eab7f2ad.gif" />

<details>

<summary>Key Bindings</summary>

<br />

| Action                               | Key                            |
| :----------------------------------- | :----------------------------- |
| Create a new snippet                 | <kbd>n</kbd>                   |
| Edit selected snippet (in `$EDITOR`) | <kbd>e</kbd>                   |
| Copy selected snippet to clipboard   | <kbd>c</kbd>                   |
| Paste clipboard to selected snippet  | <kbd>p</kbd>                   |
| Delete selected snippet              | <kbd>x</kbd>                   |
| Move selected snippet up             | <kbd>K</kbd>                   |
| Move selected snippet down           | <kbd>J</kbd>                   |
| Rename selected snippet              | <kbd>r</kbd>                   |
| Rename selected folder               | <kbd>R</kbd>                   |
| Move to next pane                    | <kbd>l</kbd> <kbd>→</kbd>      |
| Move to previous pane                | <kbd>h</kbd> <kbd>←</kbd>      |
| Search for snippets                  | <kbd>/</kbd>                   |
| Toggle help                          | <kbd>?</kbd>                   |
| Quit application                     | <kbd>q</kbd> <kbd>ctrl+c</kbd> |

</details>

## Command Line Interface

Create new snippets:

```bash
# Quick save an untitled snippet.
zzz < main.go

# From a file, specify Notes/ folder and Go language.
zzz Notes/FizzBuzz.go < main.go

# Save some code from the internet for later.
curl https://example.com/main.go | zzz Notes/FizzBuzz.go

# Works great with GitHub gists
gh gist view 4ff8a6472247e6dd2315fd4038926522 | zzz
```

<img width="600" src="https://user-images.githubusercontent.com/42545625/202767159-134d679f-490f-4ad2-8875-cda604aa7b13.gif" />

Output saved snippets:

```bash
# Fuzzy find snippet.
zzz fuzzy

# Write snippet to a file.
zzz go/boilerplate > main.go

# Copy snippet to clipboard.
zzz foobar | pbcopy
zzz foobar | xclip
zzz foobar | wl-copy
```

<img width="600" src="https://user-images.githubusercontent.com/42545625/202240249-d724fd73-2f90-4036-b9fc-6d2ccef982b3.gif" />

List snippets:

```bash
zzz list
```

<img width="600" src="https://user-images.githubusercontent.com/42545625/202242653-1696dda6-2527-4c38-b673-74d67ad1517f.gif" />

Fuzzy find a snippet (with [Gum](https://github.com/charmbracelet/gum)).

```bash
zzz $(zzz list | gum filter)
```

<img width="600" src="https://user-images.githubusercontent.com/42545625/202240268-3a71fde6-73c3-4b0a-b129-f87ec1bb1b88.gif" />

## Installation

###### Install with nix (using flake + home manager):

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    home-manager = {
      url = "github:nix-community/home-manager";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    zzz.url = "github:isabelroses/zzz";
  };

  outputs = { self, nixpkgs, home-manager, zzz }: {
    homeConfigurations."user@hostname" = home-manager.lib.homeManagerConfiguration {
      modules = [
        home-manager.homeManagerModules.default
        zzz.homeManagerModules.default
        {
          programs.zzz = {
            settings = {
              home = "~/.zzz";
            };
            colors = {
              background = "#252B2E";
              foreground = "#D9E4DC";
              primary_color = "#B2C98F";
              primary_color_subdued = "#6E8585";
              green = "#B2C98F";
              bright_green = "#83C092";
              red = "#E67E80";
              bright_red = "#E69875";
              textinvert = "#46545B";
              gray = "#343E44";
            };
          };
        }
      ];
    };
  }
}
```

###### Install with Go:

```sh
go install github.com/isabelroses/zzz@main
```

Or download a binary from the [releases](https://github.com/isabelroses/zzz/releases).

## Customization

Zzz is customized through a configuration file located at `ZZZ_CONFIG` (`$XDG_CONFIG_HOME/zzz/config.yaml`).

```yaml
# Configuration
home: ~/.zzz
default_language: go
theme: nord

# Colors
background: "0"
foreground: "7"
primary_color: "#AFBEE1"
primary_color_subdued: "#64708D"
green: "#527251"
bright_green: "#BCE1AF"
bright_red: "#E49393"
red: "#A46060"
gray: "240"
textinvert: "#373B41"
```

The configuration file can be overridden through environment variables:

```bash
# Configuration
export ZZZ_CONFIG="~/.zzz/config.yaml"
export ZZZ_HOME="~/.zzz"
export ZZZ_DEFAULT_LANGUAGE="go"
export ZZZ_THEME="nord"

# Colors
export ZZZ_PRIMARY_COLOR="#AFBEE1"
export ZZZ_RED="#A46060"
export ZZZ_GREEN="#527251"
export ZZZ_FOREGROUND="7"
export ZZZ_BACKGROUND="0"
export ZZZ_GRAY="240"
export ZZZ_TEXTINVERT="#373B41"
```

<br />

<p align="center">
  <img
    width="1000"
    alt="image"
    src="https://user-images.githubusercontent.com/42545625/202867429-5bcf8fae-5dd7-478c-b958-638aa5765d97.png"
  />
</p>

## License

[MIT](https://github.com/isabelroses/zzz/blob/master/LICENSE)

## Feedback

I'd love to hear your feedback on improving `zzz`.

Feel free to reach out via:

- [Contact](https://isabel.contact)
- [Email](mailto:isabel@isabelroses.com)
- [GitHub issues](https://github.com/isabelroses/zzz/issues/new)

---

<sub><sub>z</sub></sub><sub>z</sub>z
