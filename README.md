# Deprecation
This module is deprecated due to the fact that Go routines and invocation of PAM from SSH are not compatible.

The issue is caused by SSH forking when starting the PAM subprocess, while Go wants to start a thread on goroutine. These don't go together well and result in an indefinate mutex lock wait in the PAM module.

# Go PAM WebSSO module
Go version of the SRAM websso module

# Developing

In order to develop this module, a working environment is started if you wish. To make use this feature, you need to have **Docker** installed on your host environment.
In this work environment you can build, change, test and git commmi, push your changes.
In parallel you can continue to use your host editor to change files as well.

Enter the working enviroment, by:
```
./work.sh
```

# Install
```
apt install libpam0g-dev

make module

make install
```

/etc/pam.d/sshd:

```auth required pam_websso.so```
