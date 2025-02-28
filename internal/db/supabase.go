package db

import (
	"log"

	"encoding/json"

	"github.com/supabase-community/supabase-go"
)

func InsertDataToSupabase(client *supabase.Client, table string, data interface{}) error {
	// Assuming the correct method chain for inserting data
	_, _, err := client.From(table).Insert(data, false, "", "", "").Execute()
	if err != nil {
		return err
	}

	log.Println("✅ Data inserted into Supabase")
	return nil
}

func SelectDataFromSupabase(client *supabase.Client, table string, filter map[string]string) ([]map[string]interface{}, error) {
	response, _, err := client.From(table).Select("*", "exact", false).Match(filter).Execute()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Data retrieved from Supabase")
	return result, nil
}

func UpdateDataInSupabase(client *supabase.Client, table string, filter map[string]string, update map[string]interface{}) error {
	_, _, err := client.From(table).Update(update, "", "").Match(filter).Execute()
	if err != nil {
		return err
	}

	log.Println("✅ Data updated in Supabase")
	return nil
}

func DeleteDataFromSupabase(client *supabase.Client, table string, filter map[string]string) error {
	_, _, err := client.From(table).Delete("", "").Match(filter).Execute()
	if err != nil {
		return err
	}

	log.Println("✅ Data deleted from Supabase")
	return nil
}
