import sys

import ruamel.yaml as yaml

GET_QUOTE_RESPONSE = ' '.join(sys.argv[1:])

file = "../kubernetes-manifests/shippingservice.yaml"
with open(file, "r") as stream:
    d = list(yaml.safe_load_all(stream))

d[0]['spec']['template']['spec']['containers'][0]['env'][1]['value'] = GET_QUOTE_RESPONSE

with open(file, "w") as stream:
    yaml.dump_all(
        d,
        stream,
        default_flow_style=False
    )

file = "../kubernetes-manifests/testservice.yaml"
with open(file, "r") as stream:
    d = list(yaml.safe_load_all(stream))

d[0]['spec']['template']['spec']['containers'][0]['env'][1]['value'] = GET_QUOTE_RESPONSE

with open(file, "w") as stream:
    yaml.dump_all(
        d,
        stream,
        default_flow_style=False
    )
