import sys

import ruamel.yaml as yaml

kubernetes_manifests = ['adservice.yaml', 'cartservice.yaml', 'checkoutservice.yaml',
						'currencyservice.yaml', 'emailservice.yaml', 'emailservice.yaml',
						'paymentservice.yaml', 'productcatalogservice.yaml', 
						'recommendationservice.yaml', 'shippingservice.yaml']

extraLatency = sys.argv[1]

for file in kubernetes_manifests:
	file = "../kubernetes-manifests/" + file
	with open(file, "r") as stream:
	    d = list(yaml.safe_load_all(stream))

	if d[0]['spec']['template']['spec']['containers'][0]['env'][0]['value'].endswith('s'):
		d[0]['spec']['template']['spec']['containers'][0]['env'][0]['value'] = extraLatency + 's'
	else:
		d[0]['spec']['template']['spec']['containers'][0]['env'][0]['value'] = extraLatency

	with open(file, "w") as stream:
	    yaml.dump_all(
	        d,
	        stream,
	        default_flow_style=False
	    )
