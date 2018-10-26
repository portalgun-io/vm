

### TTY vs PTY

**TTY**
[TTY Programming Guide](http://www.linusakesson.net/programming/tty/)


### PTMX 

```
ptmx, pts - pseudoterminal master and slave
Description
The file /dev/ptmx is a character file with major number 5 and minor number 2, usually of mode 0666 and owner.group of root.root. It is used to create a pseudoterminal master and slave pair.

When a process opens /dev/ptmx, it gets a file descriptor for a pseudoterminal master (PTM), and a pseudoterminal slave (PTS) device is created in the /dev/pts directory. Each file descriptor obtained by opening /dev/ptmx is an independent PTM with its own associated PTS, whose path can be found by passing the descriptor to ptsname(3).

Before opening the pseudoterminal slave, you must pass the master's file descriptor to grantpt(3) and unlockpt(3).

Once both the pseudoterminal master and slave are open, the slave provides processes with an interface that is identical to that of a real terminal.

Data written to the slave is presented on the master descriptor as input. Data written to the master is presented to the slave as input.

In practice, pseudoterminals are used for implementing terminal emulators such as xterm(1), in which data read from the pseudoterminal master is interpreted by the application in the same way a real terminal would interpret the data, and for implementing remote-login programs such as sshd(8), in which data read from the pseudoterminal master is sent across the network to a client program that is connected to a terminal or terminal emulator.

Pseudoterminals can also be used to send input to programs that normally refuse to read input from pipes (such as su(1), and passwd(1)). 
```

### Linux Console Device
```
console - console terminal and virtual consoles
Description
A Linux system has up to 63 virtual consoles (character devices with major number 4 and minor number 1 to 63), usually called /dev/ttyn with 1 <= n <= 63. The current console is also addressed by /dev/console or /dev/tty0, the character device with major number 4 and minor number 0. The device files /dev/* are usually created using the script MAKEDEV, or using mknod(1), usually with mode 0622 and owner root.tty.

Before kernel version 1.1.54 the number of virtual consoles was compiled into the kernel (in tty.h: #define NR_CONSOLES 8) and could be changed by editing and recompiling. Since version 1.1.54 virtual consoles are created on the fly, as soon as they are needed.

Common ways to start a process on a console are: (a) tell init(8) (in inittab(5)) to start a mingetty(8) (or agetty(8)) on the console; (b) ask openvt(1) to start a process on the console; (c) start X--it will find the first unused console, and display its output there. (There is also the ancient doshell(8).)

Common ways to switch consoles are: (a) use Alt+Fn or Ctrl+Alt+Fn to switch to console n; AltGr+Fn might bring you to console n+12 [here Alt and AltGr refer to the left and right Alt keys, respectively]; (b) use Alt+RightArrow or Alt+LeftArrow to cycle through the presently allocated consoles; (c) use the program chvt(1). (The key mapping is user settable, see loadkeys(1); the above mentioned key combinations are according to the default settings.)

The command deallocvt(1) (formerly disalloc) will free the memory taken by the screen buffers for consoles that no longer have any associated process.
Properties
Consoles carry a lot of state. I hope to document that some other time. The most important fact is that the consoles simulate vt100 terminals. In particular, a console is reset to the initial state by printing the two characters ESC c. All escape sequences can be found in console_codes(4).
Files
/dev/console
/dev/tty*
```
