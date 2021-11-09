package gocurl

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/rs/zerolog/log"
)

func Exec(filepath string, v interface{}) error {
	cmd := []string{"sh", filepath}
	log.Info().
		Interface("command", strings.Join(cmd, " ")).
		Msg("executing command")
	bs, err := exec.
		Command(cmd[0], cmd[1:]...).
		Output()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, v)
	if err != nil {
		return err
	}
	return nil
}
