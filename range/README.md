Installs:
* ansible
* vagrant
* vmware player
* vmware vagrant utility (https://developer.hashicorp.com/vagrant/downloads/vmware)
  * `$ vagrant plugin install vagrant-vmware-desktop`



## Handy Tricks

* vagrant ssh-config

```bash
ansible-inventory --graph
ansible testserver -i inventory/vagrant.ini -m ping
ansible testserver -m command -a uptime
ansible testserver -b -a "tail /var/log/syslog"
```


```bash
vagrant destroy -f
vagrant provision
vagrant reload
```

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -subj /CN=localhost -keyout files/nginx.key -out files/nginx.crt
```