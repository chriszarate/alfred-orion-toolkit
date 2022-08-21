# Alfred Orion Toolkit

This is an [Alfred](https://www.alfredapp.com) workflow for [Kagi's Orion browser](https://browser.kagi.com). It is based on [`alfred-safari-toolkit`](https://github.com/addozhang/alfred-safari-toolkit).

## Keywords

- `ob` to search Orion bookmarks
- `oh` to search Orion history
- `or` to search Orion reading list

## Install

1. Download the workflow from [release page](https://github.com/chriszarate/alfred-orion-toolkit/releases) and double-click to add it to Alfred.
2. Because the produced binary is not signed, you will need to trust it manually as a one-time task:
   - Right-click the workflow in Alfred and "Show in Finder".
	 - Right-click `alfred-orion-toolkit` and "Open with... > Terminal.app". Accept the prompt.

## Build for development

Install [go-alfred](https://github.com/jason0x43/go-alfred) for workflow packaging:

1. `go install github.com/jason0x43/go-alfred/alfred@latest`
2. Run `CGO_ENABLED=1 alfred build` to produce an execution binary (`./workflow/alfred-orion-toolkit`)
3. Run `alfred pack` to produce a workflow package (`./alfred-orion-toolkit-VERSION.alfredworkflow`)
