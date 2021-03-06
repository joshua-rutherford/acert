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

package list

import (
	"fmt"

	"github.com/greymatter-io/acert/certificates"
	"github.com/greymatter-io/acert/config"
	"github.com/spf13/cobra"
)

// Command returns a command that lists the authorities.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "list",
		Short: "List the authorities",
		RunE: func(command *cobra.Command, args []string) error {

			authorities, err := config.Authorities()
			if err != nil {
				return err
			}

			identities, err := authorities.List()
			if err != nil {
				return err
			}

			for _, identity := range identities {

				expiration := certificates.Expiration(identity.Certificate)
				fingerprint := certificates.Fingerprint(identity.Certificate)
				name := certificates.CommonName(identity.Certificate)

				fmt.Printf("%s\t%s\t%v\n", fingerprint, name, expiration)
			}

			return nil
		},
	}

	return command
}
