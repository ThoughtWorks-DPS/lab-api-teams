curl http://localhost:8080/teams --include --header     "Content-Type: application/json"     --request "POST" --data '{"TeamID": "team-sapphire", "TeamType": "normal", "TeamDescription": "Sapphire frontend team", "TeamRAM": 32, "TeamCPU": 12, "TeamRamLimit": 64, "TeamCpuLimit": 24}'
curl http://localhost:8080/teams
