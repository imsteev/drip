[drip](https://drip.fly.dev)

Purge data on Fly.io:
```
# find the volume
fly volumes list

# delete it
fly volumes delete <volume name>

# create a 1gb volume in boston
fly volumes create data --region bos --size 1
```
