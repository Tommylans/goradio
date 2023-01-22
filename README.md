<p align="center">
  <img src="https://link.storjshare.io/s/jw72n4eodygmsee5qaavopbdbvkq/devv/github/GoRadio/logo.png?wrap=0" height="365px" alt="GoRadio Logo"/>
</p>
<h1 align="center">GoRadio</h1>
<p align="center">
 <a href="https://github.com/tommylans/GoRadio/actions/workflows/ci.yml?query=branch%3Amaster"><img src="https://github.com/tommylans/GoRadio/actions/workflows/go.yml/badge.svg" alt="Build Status"></a>
 <a href="https://goreportcard.com/report/github.com/tommylans/GoRadio"><img src="https://goreportcard.com/badge/github.com/tommylans/GoRadio" alt="Go Report Card"></a>
 <a href="https://github.com/tommylans/GoRadio/releases"><img src="https://img.shields.io/github/v/release/tommylans/GoRadio?display_name=tag&sort=semver" alt="GitHub Releases"></a>
</p>

---
**Warning: Some things may not work correctly yet and the code is currently in spaghetti phase**

This is a cli tool to play some internet radio's made with Go.

## Integrations
This Radio client also has Discord RPC support.

## Screenshot

<details>
  <summary>Click me</summary>
  
  <img src="https://link.storjshare.io/s/jwfffvjuj2ofruocy2iv5nywxwja/devv/github/GoRadio/Tui-Screenshot.png?wrap=0" alt="Screenshot of the terminal with the tui open" />
</details>

## Usage

Compile the project and then just run the binary in the terminal.

## Installation
### Go
```bash
$ go install github.com/tommylans/goradio@latest

$ goradio
```

## Keybindings

| **Key**                | **Action**      |
|------------------------|-----------------|
| <center>m</center>     | Mute            |
| <center>s</center>     | Stop            |
| <center>+ / =</center> | Increase volume |
| <center>-</center>     | Decrease volume |
| <center>0</center>     | Reset volume    |
| <center>q</center>     | Quit            |

## Todo

* Create a proper Readme
