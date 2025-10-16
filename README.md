# Neuron

## How to start?

```
make start
```

```
docker ps
```

## API Calls

#### Admin Login
```
curl -X POST http://localhost:8090/api/v1/admin/login \
     -H "Content-Type: application/json" \
     -d '{"login": "admin", "password": "password"}'
```

#### Create Device
```
curl -X POST http://localhost:8090/api/v1/device/ \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3OTE5ODMzNTMsInVzZXJfaWQiOjF9.yRjm0QEclWuAFr0ZXg-9AOjP7Cbs4v_bT3zohmeEIHI" \
     -d '{"name": "device1"}' 
```

#### Execute Commands
```
curl -X POST http://localhost:8090/api/v1/execute/ \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VJRCI6ImQ1MWYzM2JjLTI1YzItNDhmNy1hYTYwLTM2MDM0OTVkYzQ3NSIsImV4cCI6MTc5MTk4MzYzNH0.0fCrL2gFdO9zQhKFn8vv-zKLgaVpQUPKscSZVVMo8Mk" \
     -d '{
       "timestamp": "2025-10-14T08:20:00Z",
       "location": "Downtown",
       "roadSegments": [
         {"id": "R1", "vehicleCount": 150, "avgSpeed": 12.5, "capacity": 300, "type": "arterial"},
         {"id": "R2", "vehicleCount": 80, "avgSpeed": 35.0, "capacity": 200, "type": "highway"},
         {"id": "R3", "vehicleCount": 60, "avgSpeed": 18.0, "capacity": 150, "type": "local"}
       ],
       "incidents": [
         {"segmentID": "R1", "type": "accident", "severity": 2},
         {"segmentID": "R2", "type": "construction", "severity": 1}
       ],
       "weather": {"condition": "rain", "visibility": 0.7},
       "policy": {"congestionCharge": true, "publicTransportPriority": true}
     }'

```
