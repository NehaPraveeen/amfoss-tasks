## The engineering issues identified:
1. The issue found was __Mixed Protocol__ . First i went into east-blue. after that i ran the command _cat logs/restore.log_ where i noticed a warning sign saying _legacy station snapshot contained mixed protocol versions_ which made me realise that the files are mixed up.
2. since it mentioned leagacy atations and I remembered a similar file name in the directory where archives was there. i inspected that using command _cat legacy-stations.yml_ and found that one had proto "v1" and other had proto "v2".

