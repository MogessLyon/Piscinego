echo $(curl -s https://learn.zone01dakar.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"mousdieng\"}}){id}}"}' | grep -Eo '[0-9]{1,4}')


