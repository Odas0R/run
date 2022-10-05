package editor

import (
	"fmt"
	"strings"

	"github.com/odas0r/cmd/pkg/fs"
	"github.com/odas0r/cmd/pkg/shell"
)

func Edit(path string) error {
	if fs.NotExists(path) {
		return fmt.Errorf("file does not exist on given path")
	}

	shell.Exec(fmt.Sprintf(`
    if [[ ! -S "$NVIM_SOCKET" ]]; then
      nvim "+${2}" --listen "$NVIM_SOCKET" "%s"
    else
      nvim --server "$NVIM_SOCKET" --remote-send "<C-\><C-N>:wincmd p | edit %s<CR>"
    fi
  `, path, path))

	return nil
}

func Fzf(content []string, prompt string) string {
	contentStr := strings.Join(content, "\n")

	echo := fmt.Sprintf("echo \"%s\"", contentStr)
	fzf := "fzf-tmux -p 50%"
	fzfPrompt := fmt.Sprintf("--prompt=\"%s\"", prompt)

	fzfCommand := fmt.Sprintf(`
    %s | %s %s
  `, echo, fzf, fzfPrompt)

	// execute the bash command and return the output as a string
	output := shell.ExecOutput(fzfCommand)

	// replace new lines with empty string
	output = strings.ReplaceAll(output, "\n", "")

	return strings.TrimSpace(output)
}

func FzfPrintQuery(content []string, prompt string) string {
	contentStr := strings.Join(content, "\n")

	echo := fmt.Sprintf("echo \"%s\"", contentStr)
	fzf := "fzf-tmux -p 50% --print-query"
	fzfPrompt := fmt.Sprintf("--prompt=\"%s\"", prompt)

	fzfCommand := fmt.Sprintf(`
    readarray -t lines < <(%s | %s %s)
    query="${lines[0]}"
    selected="${lines[1]}"
    if [[ -n "$selected" ]]; then
      echo "$selected"
      exit 0
    fi
    if [[ -n "$query" ]]; then
      echo "$query"
      exit 0
    fi
  `, echo, fzf, fzfPrompt)

	// execute the bash command and return the output as a string
	output := shell.ExecOutput(fzfCommand)

	// replace new lines with empty string
	output = strings.ReplaceAll(output, "\n", "")

	return strings.TrimSpace(output)
}

func Notify(text string) {
	shell.Exec(fmt.Sprintf(`
    if [[ -S "$NVIM_SOCKET" ]]; then
      nvim --server "$NVIM_SOCKET" --remote-send '<ESC>:lua vim.notify("%s", nil, {title="Pomodoro"})<CR>'
    fi
  `, text))
}

func NotifyByType(text string, ntype string) {
	shell.Exec(fmt.Sprintf(`
    if [[ -S "$NVIM_SOCKET" ]]; then
      nvim --server "$NVIM_SOCKET" --remote-send '<ESC>:lua vim.notify("%s", "%s", {title="Pomodoro"})<CR>'
    fi
  `, text, ntype))
}
