// In this program we use CLI like viper,cobra \
// CLI :- A command-line interface (CLI) is a text-based user interface (UI) used to run programs,
// manage computer files and interact with the computer.
//CLI are also called "command-line user interfaces", "console user interfaces" and "character user interfaces".
// CLIs accept as input commands that are entered by keyboard;
//the commands invoked at the command prompt are then run by the computer.

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The Davinci003 engine is a Go library for creating and running distributed applications.
//To use the GPT-3 Davinci API in Go lang,
//you can make HTTP requests to the OpenAI API endpoints using the net/http package.

func GetResponse(client gpt3.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},

		MaxTokens: gpt3.IntPtr(4000),

		// we create a  pointer to an integer variable called maxTokens and initialize it with the value gpt3.IntPtr(4000).
		//The gpt3.IntPtr function returns a pointer to an integer variable with the value passed as its argument

		// here tempreture is variable ,of type float32 and assign it a value of 0.

		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)

	})

	if err != nil {
		fmt.Println(err)
		os.Exit(13)

	}
	fmt.Printf("\n")

}

// The NullWriter type is a type that implements the io.Writer interface in Go lang.
//It is a type that discards any data written to it,
// and always returns a successful write with a byte count of 0 and no error.

type NullWriter int

func (NullWriter) Write([]byte) (int, error) {
	return 0, nil
}

func main() {

	// viper is a popular configuration management library in Go that
	// provides a convenient way to manage application configurations.
	// It supports various configuration file formats like JSON, YAML, etc

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	apikey := viper.GetString("API_KEY")

	if apikey == "" {

		panic("API_KEY is missing")

	}

	ctx := context.Background()

	client := gpt3.NewClient(apikey)

	// cobra is a powerful and easy-to-use CLI framework for Go that provides
	//  features such as sub-commands, flags, and command-line completions.

	//we create a cobra.Command named rootCmd with the name "chatgpt",
	// Run function that will be executed when the command is run.

	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "chat with chatgpt in terminal",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false
			for !quit {
				fmt.Println("Ask anything ('quit' to end ): ")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
				default:
					GetResponse(client, ctx, question)
				}

			}

		},
	}
	rootCmd.Execute()
}
