# SSHPKM

![](img/diagram.png)

Documentation

### Managing User

Creating user
```
create user test
```

### Managing Server Access

Create host
```
create host adminVM
```

## Configure SSH

Use the command `AuthorizedKeysCommand` to hook the SSH public key.

`vim /etc/ssh/sshd_config`:
```
AuthorizedKeysCommand /bin/getkey
```

`vim /bin/getkey`:
```
#!/bin/bash
curl https://<ip>:<port>/ curl -H "SSH-Host: adminVM" -H "SSH-User: teste"
```