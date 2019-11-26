package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewLXCManager() *cobra.Command {
	cmd := &cobra.Command{
		Use:	"woyectl"
}
