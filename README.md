# Go PAM WebSSO module
Go version of the SRAM websso module

# Install
```
apt install libpam0g-dev

make module

make install
```

/etc/pam.d/sshd:

```auth required pam_websso.so```
