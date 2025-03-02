#!/bin/bash

# Example migration script
echo "Running database migrations..."
# Add migration commands here 
/Users/minjungsung/go/bin/migrate -database "postgres://postgres.btmccsouvoknkwmrxgwe:helloworld_3090!@db.btmccsouvoknkwmrxgwe.supabase.co:5432/postgres" -path migrations up
