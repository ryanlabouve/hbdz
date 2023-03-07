Installs:
* ansible
* vagrant
* vmware player
* vmware vagrant utility (https://developer.hashicorp.com/vagrant/downloads/vmware)
  * `$ vagrant plugin install vagrant-vmware-desktop`


## Installation

```bash
# ~/.ssh/config
Host vagrant*
  Hostname 127.0.0.1
  User vagrant
  UserKnownHostsFile /dev/null
  StrictHostKeyChecking no
  PasswordAuthentication no
  IdentityFile ~/.vagrant.d/insecure_private_key 
  IdentitiesOnly yes
  LogLevel FATAL
```

## Handy Tricks

* vagrant ssh-config

```bash
ansible-inventory --graph
ansible testserver -i inventory/vagrant.ini -m ping
ansible testserver -m command -a uptime
ansible testserver -b -a "tail /var/log/syslog"
ansible all -a "uname -a"

```


```bash
vagrant destroy -f
vagrant provision
vagrant reload
```

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -subj /CN=localhost -keyout files/nginx.key -out files/nginx.crt
```


prometheus.service.js inspired from https://github.com/prometheus-community/ansible/blob/main/roles/prometheus/templates/prometheus.service.j2