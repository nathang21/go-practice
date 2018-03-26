Example echo server for simple hosting a container on GCE (Google Compute Engine). 

* Used this a baseline to setup a global load balancer ingress
* Load balancer contained 1 instance pool for each GCP region.
* Each region contained 1 Instance group.
* Each instance group was built off the same template with a single node instace running this container.
