// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"context"
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/onosproject/onos-lib-go/pkg/cli"
	"github.com/onosproject/onos-topo/api/topo"
	"github.com/spf13/cobra"
)

func getGetEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "entity <id>",
		Aliases: []string{"entities"},
		Args:    cobra.MaximumNArgs(1),
		Short:   "Get a topo entity",
		RunE:    runGetEntityCommand,
	}
	cmd.Flags().BoolP("verbose", "v", false, "whether to print the entity with verbose output")
	cmd.Flags().Bool("no-headers", false, "disables output headers")
	return cmd
}

func runGetEntityCommand(cmd *cobra.Command, args []string) error {
	//verbose, _ := cmd.Flags().GetBool("verbose")
	//noHeaders, _ := cmd.Flags().GetBool("no-headers")

	conn, err := cli.GetConnection(cmd)
	if err != nil {
		return err
	}
	defer conn.Close()
	outputWriter := cli.GetOutput()
	writer := new(tabwriter.Writer)
	writer.Init(outputWriter, 0, 0, 3, ' ', tabwriter.FilterHTML)

	client := topo.CreateTopoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if len(args) == 0 {
		// TODO - implement List function
	} else {
		refs := make([]*topo.Reference, 1)
		refs[0].ID = topo.ID(args[0])
		response, err := client.Read(ctx, &topo.ReadRequest{Refs: refs})
		if err != nil {
			cli.Output("get error")
			return err
		}

		switch obj := response.Objects[0].Obj.(type) {
		case *topo.Object_Entity:
			_, _ = fmt.Fprintf(writer, "ID\t%s\n", response.Objects[0].Ref.GetID())
			_, _ = fmt.Fprintf(writer, "Type\t%s\n", obj.Entity.GetType())
		case nil:
			// No object is set
		default:
			// return ERROR
		}
	}
	return writer.Flush()
}

func getAddEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "entity <id> [args]",
		Aliases: []string{"entities"},
		Args:    cobra.ExactArgs(1),
		Short:   "Add an entity",
		RunE:    runAddEntityCommand,
	}
	cmd.Flags().StringP("type", "t", "", "the type of the entity")
	cmd.Flags().StringToString("attributes", map[string]string{}, "an arbitrary mapping of attributes")

	_ = cmd.MarkFlagRequired("type")
	return cmd
}

func runAddEntityCommand(cmd *cobra.Command, args []string) error {
	id := args[0]
	entityType, _ := cmd.Flags().GetString("type")
	//attributes, _ := cmd.Flags().GetStringToString("attributes")

	conn, err := cli.GetConnection(cmd)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := topo.CreateTopoClient(conn)

	updates := make([]*topo.Update, 1)

	updates[0] = &topo.Update{
		Type: topo.Update_INSERT,
		Object: &topo.Object{
			Ref: &topo.Reference{
				ID: topo.ID(id)},
			Type: topo.Object_ENTITY,
			Obj: &topo.Object_Entity{
				Entity: &topo.Entity{
					Type: entityType,
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = client.Write(ctx, &topo.WriteRequest{Updates: updates})
	if err != nil {
		return err
	}
	cli.Output("Added entity %s \n", id)
	return nil
}
