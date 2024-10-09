package sup

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

var Client *supabase.Client

func InitSupabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading.env file")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	Client = supabase.CreateClient(supabaseURL, supabaseKey)
	if Client != nil {
		fmt.Println("Connected to Supabase")
	}
}
