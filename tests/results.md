
# First
```

    checks_total.......: 173875  1396.662774/s
    checks_succeeded...: 100.00% 173875 out of 173875
    checks_failed......: 0.00%   0 out of 173875

    ✓ create profile status 200
    ✓ login status 200
    ✓ stack status 200
    ✓ get profile status 200
    ✓ update preferences status 200
    ✓ swipe status 200
    ✓ refresh token status 200
    ✓ delete status 200

    HTTP
    http_req_duration..............: avg=74.69ms min=356.85µs med=11.6ms max=513.99ms p(90)=208.91ms p(95)=231.59ms
      { expected_response:true }...: avg=74.69ms min=356.85µs med=11.6ms max=513.99ms p(90)=208.91ms p(95)=231.59ms
    http_req_failed................: 0.00%  0 out of 173875
    http_reqs......................: 173875 1396.662774/s

    EXECUTION
    iteration_duration.............: avg=9.97s   min=280.2ms  med=7.59s  max=24.92s   p(90)=20.35s   p(95)=22.86s  
    iterations.....................: 1625   13.052923/s
    vus............................: 35     min=4           max=200
    vus_max........................: 200    min=200         max=200

    NETWORK
    data_received..................: 39 MB  314 kB/s
    data_sent......................: 141 MB 1.1 MB/s
```

# Second
scenarios: (100.00%) 1 scenario, 500 max VUs, 3m0s max duration (incl. graceful stop):
default: Up to 500 looping VUs for 2m30s over 2 stages(gracefulRampDown: 30s, gracefulStop: 30s)
```

    checks_total.......: 223416  1342.852885/s
    checks_succeeded...: 100.00% 223416 out of 223416
    checks_failed......: 0.00%   0 out of 223416

    ✓ create profile status 200
    ✓ login status 200
    ✓ stack status 200
    ✓ get profile status 200
    ✓ update preferences status 200
    ✓ swipe status 200
    ✓ refresh token status 200
    ✓ delete status 200

    HTTP
    http_req_duration..............: avg=205.51ms min=372.22µs med=12.88ms max=5.32s p(90)=280.5ms p(95)=1.8s  
      { expected_response:true }...: avg=205.51ms min=372.22µs med=12.88ms max=5.32s p(90)=280.5ms p(95)=1.8s  
    http_req_failed................: 0.00%  0 out of 223416
    http_reqs......................: 223416 1342.852885/s

    EXECUTION
    iteration_duration.............: avg=27.66s   min=299.18ms med=29.36s  max=44.7s p(90)=38.72s  p(95)=39.84s
    iterations.....................: 2088   12.550027/s
    vus............................: 95     min=10          max=499
    vus_max........................: 500    min=500         max=500

    NETWORK
    data_received..................: 50 MB  302 kB/s
    data_sent......................: 182 MB 1.1 MB/s

running (2m46.4s), 000/500 VUs, 2088 complete and 0 interrupted iterations
default ✓ [======================================] 000/500 VUs  2m30s
```

# Third
scenarios: (100.00%) 1 scenario, 700 max VUs, 3m0s max duration (incl. graceful stop

default: Up to 700 looping VUs for 2m30s over 2 stages (gracefulRampDown: 30s, gracefulStop: 30s)

```

    checks_total.......: 242332 1346.27215/s
    checks_succeeded...: 99.02% 239965 out of 242332
    checks_failed......: 0.97%  2367 out of 242332

    ✗ create profile status 200
      ↳  93% — ✓ 2309 / ✗ 164
    ✗ login status 200
      ↳  92% — ✓ 2276 / ✗ 197
    ✗ stack status 200
      ↳  93% — ✓ 2262 / ✗ 168
    ✗ get profile status 200
      ↳  99% — ✓ 2254 / ✗ 8
    ✗ update preferences status 200
      ↳  99% — ✓ 2256 / ✗ 6
    ✗ swipe status 200
      ↳  99% — ✓ 224275 / ✗ 1777
    ✗ refresh token status 200
      ↳  99% — ✓ 2181 / ✗ 15
    ✗ delete status 200
      ↳  98% — ✓ 2152 / ✗ 32

    HTTP
    http_req_duration..............: avg=267.18ms min=183.06µs med=15.84ms max=9.85s p(90)=631.05ms p(95)=1.98s 
      { expected_response:true }...: avg=265.01ms min=377.71µs med=16.1ms  max=9.85s p(90)=626.53ms p(95)=1.97s 
    http_req_failed................: 0.97%  2367 out of 242332
    http_reqs......................: 242332 1346.27215/s

    EXECUTION
    iteration_duration.............: avg=32.01s   min=1.28ms   med=35.82s  max=1m13s p(90)=51.5s    p(95)=55.62s
    iterations.....................: 2395   13.30539/s
    vus............................: 78     min=10             max=699
    vus_max........................: 700    min=700            max=700

    NETWORK
    data_received..................: 55 MB  305 kB/s
    data_sent......................: 206 MB 1.1 MB/s

running (3m00.0s), 000/700 VUs, 2395 complete and 78 interrupted iterations
default ✓ [======================================] 077/700 VUs  2m30s
```
