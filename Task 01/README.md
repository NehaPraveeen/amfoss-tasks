## LEVEL 1 — AWAKENING AT LOGUETOWN REEF
 1. Basic text understanding: After reading the text I understood that there are 4 divisions named sectors A,B,C  where multiple duplicate copies of devil_fruit is present and the need was to find out the right fruit i.e the odd one out.

 2. Next i proceeded into learning terminal commands. The commands learnt are as follows:
    **pwd**: tells about the present location 
    **ls** : tells about what is that file
    **ls -l**:Gives detailed information about what's there in that file
    **cd filename**: goes into that particular file
    **cd..**: goes back

 3. With the help of the above commands I found the unique fruit, identification being all the files looked same after typing ls -l except for the devil_fruit_06 in sector_C which had -x which wasn't there in any other.
  
 4. After suspecting one fruit , the code ./eat.sh sector_C/devil_fruit_6.txt was run which gave the o/p as:
<img width="1768" height="787" alt="Screenshot From 2026-08-24 23-01-13" src="https://github.com/user-attachments/assets/002a61d3-c79b-4f5c-80d2-44238cd0e47f" />

## LEVEL 2 — THE TWO FACES OF WHISKEY PEAK
 1.After reading the text i understood that there is another version of the file that is hidden. from this it was understandable that Git commands were to be used.
 
 2.I went into the whiskey_peak directory and put ls -la command which showed feast_manifest.txt , tried cat feast_manifest.txt but dint give required result.
 
 3. hence i searched and found about git tool
   
 4. used git branch -a and got a list with whiskey_peak_investigation followed by it checkout whiskey_peak_investigation and ls -la
  
 5.then found .baroque_works_cache and procedure followed after is as follows
<img width="1920" height="1080" alt="Screenshot From 2026-08-26 23-13-30" src="https://github.com/user-attachments/assets/b6231fc5-69b8-436c-b2a2-cf18b32fe0d4" />

 6. Then i rechecked the awaken signature and finally got o/p:
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Loguetown_Reef$ cd ../Whiskey_Peak
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak$ git checkout whiskey_peak_investigation
Already on 'whiskey_peak_investigation'
Your branch is up to date with 'origin/whiskey_peak_investigation'.
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak$ cd .baroque_works_cache
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ export AWEKENING_SIGNATURE="ONE_PIECE{GITO_GITO_NO_AWAKENING}"
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ ./unlock_vault.sh
[ACCESS DENIED] Environmental Scan Failed. System user unauthorized.
Hint: Did you export the 'AWAKENING_SIGNATURE' variable inside this session?
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ export AWEKENING_AWAKENING_SIGNATURE="ONE_PIECE{GITO_GITO_NO_AWAKENING}" ./unlock_vault.sh
SIGNATURE="ONE_PIECE{GITO_GITO_NO_AWAKENING}"
bash: export: `./unlock_vault.sh': not a valid identifier
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ unset AWAKENING_SIGNATURE
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ AWAKENING_SIGNATURE="ONE_PIECE{GITO_GITO_NO_AWAKENING}" ./unlock_vault.sh
[SIGNATURE MATCH] Devil Fruit aura detected. Bypassing proxy firewall...
[SUCCESS] Decrypting Baroque transmission streams...
Files dropped: 'marine_intercept.log' and 'bounty_hunter_feed.log'. Run diff to compare.
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ diff marine_intercept.log bounty_hunter_feed.log
42c42
< LOG_STREAM_ENTRY_SECURE_NODE_042_VALID
neha-praveen@neha-praveen-LOQ-15ARP9:~/Terminal-Voyage-User-Edition/GrandLine/Whiskey_Peak/.baroque_works_cache$ 





