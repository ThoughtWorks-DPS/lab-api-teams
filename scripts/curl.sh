curl http://localhost:8080/teams --include --header     "Content-Type: application/json"     --request "POST" --data '{"TeamID": "team-sapphire", "TeamType": "normal", "TeamDescription": "Sapphire frontend team", "TeamRAM": 32, "TeamCPU": 12, "TeamRamLimit": 64, "TeamCpuLimit": 24}'
curl http://localhost:8080/teams

curl http://localhost:8080/teams --include --header     "Content-Type: application/json"     --request "POST" --data '{"NamespaceID": "default", "NamespaceType": "master", "NamespaceRam": 32, "NamespaceCpu": 12, "NamespaceInMesh": true, "NamespaceFromDefault": false}'
