import sys

import ruamel.yaml as yaml

TRAFFIC = sys.argv[1]

file = "../kubernetes-manifests/loadgenerator.yaml"
with open(file, "r") as stream:
    d = list(yaml.safe_load_all(stream))

d[0]['spec']['template']['spec']['containers'][0]['env'][0]['value'] = TRAFFIC

with open(file, "w") as stream:
    yaml.dump_all(
        d,
        stream,
        default_flow_style=False
    )