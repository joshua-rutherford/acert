// Copyright 2019 Decipher Technology Studios
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

package authorities

import (
	"github.com/greymatter-io/acert/cmd/authorities/create"
	"github.com/greymatter-io/acert/cmd/authorities/delete"
	"github.com/greymatter-io/acert/cmd/authorities/export"
	"github.com/greymatter-io/acert/cmd/authorities/issue"
	"github.com/greymatter-io/acert/cmd/authorities/list"
	"github.com/spf13/cobra"
)

// Command returns a command that lists and manages authority certificates.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "authorities",
		Short: "Manage authorities",
	}

	command.AddCommand(create.Command())
	command.AddCommand(delete.Command())
	command.AddCommand(export.Command())
	command.AddCommand(issue.Command())
	command.AddCommand(list.Command())

	return command
}
