## The engineering issues identified:
1. The main issue I found in the east-blue directory was a warning about __Mixed Protocol Versions__. This happens when a project reads a list of data entries, but the entries are written in different version languages that don't align together.
2. The first station (station-alpha) was using an old language called v1.The second station (station-beta) was using a newer language called v2.

## The approach used to investigate them:
1. First I went into east-blue. After which i ran the command _cat logs/restore.log_ where I noticed a warning sign saying _"legacy station snapshot contained mixed protocol versions"_ which made me realize that the files are mixed up.
2. Since it mentioned leagacy-stations and I remembered a similar file name in the directory where archives was there, i inspected that using command _cat legacy-stations.yml_ and found out that one had proto "v1" and other had proto "v2".

