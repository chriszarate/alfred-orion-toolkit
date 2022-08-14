# Alfred Orion Toolkit

This is an [Alfred](https://www.alfredapp.com) workflow for [Kagi's Orion browser](https://browser.kagi.com). It is based on [`alfred-safari-toolkit`](https://github.com/addozhang/alfred-safari-toolkit).

## How to install

1. Download the workflow from [release page](https://github.com/chriszarate/alfred-orion-toolkit/releases) and double-click to add it to Alfred.
2. Because the produced binary is not signed, you will need to trust it manually as a one-time task:
   - Right-click the workflow in Alfred and "Show in Finder".
	 - Right-click `alfred-orion-toolkit` and "Open with... > Terminal.app". Accept the prompt.

## How to build

It's recommended to [go-alfred](https://github.com/jason0x43/go-alfred) for workflow packaging.

1. First, install is by executing `go install github.com/jason0x43/go-alfred/alfred@latest`.
2. After running `CGO_ENABLED=1 alfred build` to build project, you will get the execution binary under `workflow` folder. 
3. At last, run `alfred pack` and the workflow package will present in root folder. 
