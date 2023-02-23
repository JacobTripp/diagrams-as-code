/*
Copyright © 2023 Jacob Tripp <jake@jaketripp.dev>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JacobTripp/diagrams-as-code/internal/newDiagram"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	DiagramConfigNotFound int = iota + 1
	GetYamlFlagError
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "Define a new diagram type.",
	Long: `Given a yaml file containing a diagram definition this will
generate the code for a new diagram.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("yaml")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(GetYamlFlagError)
		}
		d := loadDiagramConfig(path)
		err = newDiagram.Render(os.Stdout, viper.GetString("templatePath"), d)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(GetYamlFlagError)
		}
	},
}

func loadDiagramConfig(path string) *viper.Viper {
	diagramYaml := viper.New()
	fileDir, fileName := filepath.Split(path)
	diagramYaml.SetConfigName(fileName)
	fileDir, err := filepath.Abs(fileDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(DiagramConfigNotFound)
	}
	diagramYaml.AddConfigPath(fileDir)
	diagramYaml.SetConfigType("yaml")
	err = diagramYaml.ReadInConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(DiagramConfigNotFound)
	}
	return diagramYaml
}

func init() {
	rootCmd.AddCommand(defineCmd)

	defineCmd.Flags().StringP("yaml", "y", "", "path to diagram definition")
	defineCmd.MarkFlagRequired("yaml")

}
