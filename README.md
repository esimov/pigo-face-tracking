# pigo-face-tracking

[![CI](https://github.com/esimov/pigo-face-tracking/workflows/CI/badge.svg)](https://github.com/esimov/pigo-face-tracking/actions)
[![License](https://img.shields.io/github/license/esimov/pigo-face-tracking)](https://github.com/esimov/pigo-face-tracking/blob/master/LICENSE)

This is a real time face tracking application using the [Pigo](https://github.com/esimov/pigo) face detection library to convert face movements into keyboard actions like <kbd>UP</kbd>, <kbd>DOWN</kbd>, <kbd>LEFT</kbd>, <kbd>RIGHT</kbd>. This means you can play games without being physically attached to a keyboard device or you can navigate through your web browser like you would navigate with the arrow keys. 

This is achieved due to the library high performance facial landmark points detection capabilities.

<p align="center"><img src="https://github.com/esimov/pigo-face-tracking/raw/master/capture.gif" alt="Screen capture"/></p>

## How does it work?

The **Pigo** library is capable for high accuracy facial landmark points detection, but out of the existing **15** facial landmark points, it proves that the nose (**LP93**) landmark point provides the best accuracy for face tracking, so this has been used to track the head movement. Once a head movement is detected a keyboard press event is triggered through the OS system events as you would press the key physically.

## Install
**Notice: at least Go 1.13 is required!**

```bash
$ go get -u -v github.com/esimov/pigo-face-tracking

```

## Run
To run it is as simple as to type a single `make` command from the project root directory:

1. This will open `localhost:5000` in a new tab, will activate the webcam and it will start tracking your head.
2. Find some Atari like online games and start playing. 
3. Start moving your head <kbd>&uparrow;</kbd> and <kbd>&downarrow;</kbd>, <kbd>&leftarrow;</kbd> and <kbd>&rightarrow;</kbd> as you would press the arrow keys.

## OS Support
**This program has been tested on Linux and MacOS, but normally it should also run on Windows.**

Because of the OS imposed security constrains there are some important steps you need to take:

#### MacOS:
In MacOS you must set the accessibility authorization for the terminal you are running from.

<img src="https://user-images.githubusercontent.com/705503/80077645-11c09b00-854e-11ea-8b52-ad130b42028b.png" width=300/>

### Linux:
On Linux the library used for triggering the keyboard events uses ***uinput***, which on the major distributions requires root permissions.
The easy solution is executing with `sudo`. A worse way is by changing the executable's permissions by using `chmod`. For this reason you can run the accompanied `permission.sh` shell file.

## Author

* Endre Simo ([@simo_endre](https://twitter.com/simo_endre))

## License

Copyright © 2020 Endre Simo

This software is distributed under the MIT license. See the [LICENSE](https://github.com/esimov/pigo-face-tracking/blob/master/LICENSE) file for the full license text.
