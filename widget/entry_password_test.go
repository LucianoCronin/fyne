package widget_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestNewPasswordEntry(t *testing.T) {
	p := widget.NewPasswordEntry()
	p.Text = "visible"
	r := test.WidgetRenderer(p)

	cont := r.Objects()[2].(*container.Scroll).Content.(fyne.Widget)
	r = test.WidgetRenderer(cont)
	rich := r.Objects()[1].(*widget.RichText)
	r = test.WidgetRenderer(rich)

	assert.Equal(t, "•••••••", r.Objects()[0].(*canvas.Text).Text)
}
