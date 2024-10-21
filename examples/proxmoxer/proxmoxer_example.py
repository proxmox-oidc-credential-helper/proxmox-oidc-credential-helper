#!/usr/bin/env python3

from proxmoxer import ProxmoxAPI
import json
import subprocess

proxmox_host = 'proxmox.example.com'
proxmox_port = 8006
realm = "realm1"

creds_helper_output = subprocess.run(f"proxmox-oidc-credential-helper -proxmox-url https://{proxmox_host}:{proxmox_port} -realm {realm} -output=json", shell=True, capture_output=True)
creds_helper_json = json.loads(creds_helper_output.stdout)

pve = ProxmoxAPI(proxmox_host, user = creds_helper_json['data']['username'], password = creds_helper_json['data']['ticket'], verify_ssl=True)
print(pve.nodes.get())
