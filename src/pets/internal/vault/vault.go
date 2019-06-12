// vault reads ansible-vault data
// requires: ansible-vault in $PATH and ansible.cfg in cwd with vault_password_file set
package vault

import (
  "os/exec"

  "gopkg.in/yaml.v2"
)

func View(cwd, vaultPath string, v interface{}) error {
  Cmd := exec.Command("ansible-vault", "view", vaultPath)
  Cmd.Dir = cwd
	res, err := Cmd.CombinedOutput()
  if err != nil {
    return err
  }
  return yaml.Unmarshal(res, v)
}
