package main

import (
  tea "github.com/charmbracelet/bubbletea"
)

type model struct {
  text string
  next *model
}
func initialModel() *model {
  initModel := model{
    text: "...",
  }
  nextModel := model{
    text: "---",
    next: &initModel,
  }
  initModel.next = &nextModel
  return &initModel
}
func (m model) Init() tea.Cmd {
  return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "n":
      return m.next, nil
    case ".":
      m.text += "."
    case "-":
      m.text += "-"
    case "q":
      return m, tea.Quit
    }
  }
  return m, nil
}

func (m model) View() string {
  s := "press '-' or '.'. Switch scenes with 'n': " + m.text
  return s
}
func main () {
  p := tea.NewProgram(initialModel(), tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    panic(err)
  }
}
