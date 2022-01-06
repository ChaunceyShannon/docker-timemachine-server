I am using an external disk as a time machine backup destination before, but when I am moving to another place I have to unplug the disk first and I always forget to do so, and if the disk is full, it always backup failed even it says it will delete the oldest backup when the disk is full, so I am thinking about deploying a time machine server on the LAN in my another laptop that runs a single-node Kubernetes cluster.

If you are using docker, simply use this command 

```
docker run -d \
	-e AFP_VOL_NAME=backup \ # The shared volume name of the afp server
	-e AFP_USER=username \ # Username for authorize  
	-e AFP_PASSWORD=password \ # Password for authorize 
	-e AFP_USER_UID=1000 \ # UID for the user, so does the files 
	-e AFP_USER_GID=1000 \ # GID for the user, so does the files 
	-e AFP_SIZE_LIMIT=2048000 \ # The shared volume size, in MB
	-v /path/to/data/on/the/host:/timemachine \ # Mount the data directory on the host to `/timemachine` inside the container 
	chaunceyshannon/timemachine
```

Or you want to build the image for yourself 

```
docker build . -t timemachine 
```

For kubernetes Yaml files, visit: https://github.com/ChaunceyShannon/flux-sample/tree/main/timemachine

To use the timemachine you have to connect to the AFP server manually first

![image](https://user-images.githubusercontent.com/87258078/148397539-10cb4b75-0d00-4bcf-add9-7981bfca10f1.png)

![image](https://user-images.githubusercontent.com/87258078/148397781-42a163a0-a02b-46e8-a6a5-fcd6eacb82d5.png)

And then use it inside the timemachine 

![image](https://user-images.githubusercontent.com/87258078/148397935-bc01ed86-1afe-48d5-a001-338898daf20b.png)

To find the IP address of the service, using this command 

```
$ kubectl get svc                                                                           
NAME                 TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)         AGE
kubernetes           ClusterIP      10.96.0.1       <none>         443/TCP         20d
timemachine-svc      LoadBalancer   10.98.200.154   192.168.1.14   548:40790/TCP   26h
```


