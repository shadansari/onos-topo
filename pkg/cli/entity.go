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
		Use:   "get-entity <id>",
		Args:  cobra.MinimumNArgs(1),
		Short: "Get a topo entity",
		RunE:  runGetEntityCommand,
	}
	/*
		cmd.Flags().StringP("id", "i", "", "the id of the entity")
		cmd.Flags().BoolP("verbose", "v", false, "whether to print the entity with verbose output")

		_ = cmd.MarkFlagRequired("id")
	*/

	return cmd
}

func runGetEntityCommand(cmd *cobra.Command, args []string) error {

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
		reference := &topo.Reference{
			ID: topo.ID(args[0]),
		}
		refs := []*topo.Reference{reference}
		response, err := client.Read(ctx, &topo.ReadRequest{Refs: refs})
		if err != nil {
			cli.Output("get error")
			return err
		}

		if len(response.Objects) != 0 {
			switch obj := response.Objects[0].Obj.(type) {
			case *topo.Object_Entity:
				_, _ = fmt.Fprintf(writer, "ID\t%s\n", response.Objects[0].Ref.GetID())
				_, _ = fmt.Fprintf(writer, "Type\t%s\n", obj.Entity.GetType())
			case nil:
				cli.Output("No object is set")
				// No object is set
			default:
				cli.Output("get error")
				// return ERROR
			}
		}
	}
	return writer.Flush()
}

func getAddEntityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-entity <type> <id> [args]",
		Args:  cobra.MinimumNArgs(2),
		Short: "Add an entity",
		RunE:  runAddEntityCommand,
	}
	return cmd
}

func runAddEntityCommand(cmd *cobra.Command, args []string) error {
	entityType := args[0]
	id := args[1]

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
