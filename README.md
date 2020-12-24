# iconic-drive
[![Go Report Card](https://goreportcard.com/badge/github.com/isaec/iconic-drive)](https://goreportcard.com/report/github.com/isaec/iconic-drive)
![badge](https://img.shields.io/badge/badge-badge-orange)
![useful badge](https://img.shields.io/badge/useful-maybe-yellow)
![contributions](https://img.shields.io/badge/contributions-yes%20please-FF00FF)

This tool sets drive icons that can be seen on windows and mac os, with a gui and codebase that should work on linux, windows, and mac. I have tested it with png and jpeg, but it may work with other formats.
## method
<details>
<summary>After getting the target drive and image, this tool...</summary>
  
- deletes the following files to prevent weird issues overwriting
  - ``.autorun.ico``
  - ``autorun.inf``
  - ``.VolumeIcon.icns``
  - ``._``
  - ``._.VolumeIcon.icns``
- to apply windows icon it
  - writes the icon as an ico file
  - points to it with the ``autorun.inf`` file
- to apply mac icon it
  - writes the icon as an icns file
  - writes the ``._`` and ``._.VolumeIcon.icns`` files
    - these files seem to be needed for the icon to appear - I don't know why

</details>

## known current limitations
on mac os and unix systems, the autorun.inf file will be visible, as windows will not use it if it starts with a period

on windows systems, all files will be visible, as I haven't figured out a way to set the windows hidden property on non windows systems

## screenshots
![select target drive](https://user-images.githubusercontent.com/72410860/103106849-d9e28980-45ed-11eb-8f95-e32acee54af9.png)
![loading bar](https://user-images.githubusercontent.com/72410860/103106897-39409980-45ee-11eb-973e-d2da282ec0a8.png)
![finished screen](https://user-images.githubusercontent.com/72410860/103106872-0b5b5500-45ee-11eb-868b-b978a33935f0.png)

## goals

- [ ] bring code up to golang standerds
- [ ] set windows hidden property for all written files - read fatattr's code?
- [ ] add option to remove backgrounds from filetypes without transparency
- [ ] add options for how to crop non square files, instead of leaving it up to the os
- [ ] catch and log errors

## licence
MIT - Go wild!

## credits
error icon from [tahsintahil](https://www.iconfinder.com/tahsintahil)

<div>App Icon made by <a href="https://www.flaticon.com/authors/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a></div>
