// Copyright 2024 kenita8
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package excelize

import (
	"github.com/xuri/excelize/v2"
)

type ChartOptions struct {
	ChartOptions []ChartOption `yaml:"Charts"`
}

type ChartOption struct {
	Sheet string   `yaml:"Sheet"`
	Cell  string   `yaml:"Cell"`
	Chart *Chart   `yaml:"Chart"`
	Combo []*Chart `yaml:"Combo"`
}

func (co *ChartOption) ConvertExcelizeOption() (*ExcelizeChartOption, error) {
	chart, err := co.Chart.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	combos := []*excelize.Chart{}
	for _, combo := range co.Combo {
		excelizeComb, err := combo.ConvertExcelizeOption()
		if err != nil {
			return nil, err
		}
		combos = append(combos, excelizeComb)
	}
	return &ExcelizeChartOption{
		Sheet: co.Sheet,
		Cell:  co.Cell,
		Chart: chart,
		Combo: combos,
	}, nil
}

type Chart struct {
	Type         ChartType      `yaml:"Type"`
	Series       []ChartSeries  `yaml:"Series"`
	Format       GraphicOptions `yaml:"Format"`
	Dimension    ChartDimension `yaml:"Dimension"`
	Legend       ChartLegend    `yaml:"Legend"`
	Title        []RichTextRun  `yaml:"Title"`
	VaryColors   *bool          `yaml:"VaryColors"`
	XAxis        ChartAxis      `yaml:"XAxis"`
	YAxis        ChartAxis      `yaml:"YAxis"`
	PlotArea     ChartPlotArea  `yaml:"PlotArea"`
	Fill         Fill           `yaml:"Fill"`
	Border       ChartLine      `yaml:"Border"`
	ShowBlanksAs string         `yaml:"ShowBlanksAs"`
	BubbleSize   int            `yaml:"BubbleSize"`
	HoleSize     int            `yaml:"HoleSize"`
}

func (c *Chart) ConvertExcelizeOption() (*excelize.Chart, error) {
	chart := &excelize.Chart{}
	var err error
	chart.Type, err = c.Type.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	for _, s := range c.Series {
		es, err := s.ConvertExcelizeOption()
		if err != nil {
			return nil, err
		}
		chart.Series = append(chart.Series, es)
	}
	chart.Format, err = c.Format.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.Dimension, err = c.Dimension.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.Legend, err = c.Legend.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	for _, t := range c.Title {
		et, err := t.ConvertExcelizeOption()
		if err != nil {
			return nil, err
		}
		chart.Title = append(chart.Title, et)
	}
	chart.VaryColors = c.VaryColors
	chart.XAxis, err = c.XAxis.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.YAxis, err = c.YAxis.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.PlotArea, err = c.PlotArea.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.Fill, err = c.Fill.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.Border, err = c.Border.ConvertExcelizeOption()
	if err != nil {
		return nil, err
	}
	chart.ShowBlanksAs = c.ShowBlanksAs
	chart.BubbleSize = c.BubbleSize
	chart.HoleSize = c.HoleSize
	return chart, nil
}

type ChartType string

func (t *ChartType) ConvertExcelizeOption() (excelize.ChartType, error) {
	if *t == "Area" {
		return excelize.Area, nil
	} else if *t == "AreaStacked" {
		return excelize.AreaStacked, nil
	} else if *t == "AreaPercentStacked" {
		return excelize.AreaPercentStacked, nil
	} else if *t == "Area3D" {
		return excelize.Area3D, nil
	} else if *t == "Area3DStacked" {
		return excelize.Area3DStacked, nil
	} else if *t == "Area3DPercentStacked" {
		return excelize.Area3DPercentStacked, nil
	} else if *t == "Bar" {
		return excelize.Bar, nil
	} else if *t == "BarStacked" {
		return excelize.BarStacked, nil
	} else if *t == "BarPercentStacked" {
		return excelize.BarPercentStacked, nil
	} else if *t == "Bar3DClustered" {
		return excelize.Bar3DClustered, nil
	} else if *t == "Bar3DStacked" {
		return excelize.Bar3DStacked, nil
	} else if *t == "Bar3DPercentStacked" {
		return excelize.Bar3DPercentStacked, nil
	} else if *t == "Bar3DConeClustered" {
		return excelize.Bar3DConeClustered, nil
	} else if *t == "Bar3DConeStacked" {
		return excelize.Bar3DConeStacked, nil
	} else if *t == "Bar3DConePercentStacked" {
		return excelize.Bar3DConePercentStacked, nil
	} else if *t == "Bar3DPyramidClustered" {
		return excelize.Bar3DPyramidClustered, nil
	} else if *t == "Bar3DPyramidStacked" {
		return excelize.Bar3DPyramidStacked, nil
	} else if *t == "Bar3DPyramidPercentStacked" {
		return excelize.Bar3DPyramidPercentStacked, nil
	} else if *t == "Bar3DCylinderClustered" {
		return excelize.Bar3DCylinderClustered, nil
	} else if *t == "Bar3DCylinderStacked" {
		return excelize.Bar3DCylinderStacked, nil
	} else if *t == "Bar3DCylinderPercentStacked" {
		return excelize.Bar3DCylinderPercentStacked, nil
	} else if *t == "Col" {
		return excelize.Col, nil
	} else if *t == "ColStacked" {
		return excelize.ColStacked, nil
	} else if *t == "ColPercentStacked" {
		return excelize.ColPercentStacked, nil
	} else if *t == "Col3D" {
		return excelize.Col3D, nil
	} else if *t == "Col3DClustered" {
		return excelize.Col3DClustered, nil
	} else if *t == "Col3DStacked" {
		return excelize.Col3DStacked, nil
	} else if *t == "Col3DPercentStacked" {
		return excelize.Col3DPercentStacked, nil
	} else if *t == "Col3DCone" {
		return excelize.Col3DCone, nil
	} else if *t == "Col3DConeClustered" {
		return excelize.Col3DConeClustered, nil
	} else if *t == "Col3DConeStacked" {
		return excelize.Col3DConeStacked, nil
	} else if *t == "Col3DConePercentStacked" {
		return excelize.Col3DConePercentStacked, nil
	} else if *t == "Col3DPyramid" {
		return excelize.Col3DPyramid, nil
	} else if *t == "Col3DPyramidClustered" {
		return excelize.Col3DPyramidClustered, nil
	} else if *t == "Col3DPyramidStacked" {
		return excelize.Col3DPyramidStacked, nil
	} else if *t == "Col3DPyramidPercentStacked" {
		return excelize.Col3DPyramidPercentStacked, nil
	} else if *t == "Col3DCylinder" {
		return excelize.Col3DCylinder, nil
	} else if *t == "Col3DCylinderClustered" {
		return excelize.Col3DCylinderClustered, nil
	} else if *t == "Col3DCylinderStacked" {
		return excelize.Col3DCylinderStacked, nil
	} else if *t == "Col3DCylinderPercentStacked" {
		return excelize.Col3DCylinderPercentStacked, nil
	} else if *t == "Doughnut" {
		return excelize.Doughnut, nil
	} else if *t == "Line" {
		return excelize.Line, nil
	} else if *t == "Line3D" {
		return excelize.Line3D, nil
	} else if *t == "Pie" {
		return excelize.Pie, nil
	} else if *t == "Pie3D" {
		return excelize.Pie3D, nil
	} else if *t == "PieOfPie" {
		return excelize.PieOfPie, nil
	} else if *t == "BarOfPie" {
		return excelize.BarOfPie, nil
	} else if *t == "Radar" {
		return excelize.Radar, nil
	} else if *t == "Scatter" {
		return excelize.Scatter, nil
	} else if *t == "Surface3D" {
		return excelize.Surface3D, nil
	} else if *t == "WireframeSurface3D" {
		return excelize.WireframeSurface3D, nil
	} else if *t == "Contour" {
		return excelize.Contour, nil
	} else if *t == "WireframeContour" {
		return excelize.WireframeContour, nil
	} else if *t == "Bubble" {
		return excelize.Bubble, nil
	} else if *t == "Bubble3D" {
		return excelize.Bubble3D, nil
	}
	return 0, ErrChartType.Details("CharType", *t)
}

const (
	Area                        ChartType = "Area"
	AreaStacked                 ChartType = "AreaStacked"
	AreaPercentStacked          ChartType = "AreaPercentStacked"
	Area3D                      ChartType = "Area3D"
	Area3DStacked               ChartType = "Area3DStacked"
	Area3DPercentStacked        ChartType = "Area3DPercentStacked"
	Bar                         ChartType = "Bar"
	BarStacked                  ChartType = "BarStacked"
	BarPercentStacked           ChartType = "BarPercentStacked"
	Bar3DClustered              ChartType = "Bar3DClustered"
	Bar3DStacked                ChartType = "Bar3DStacked"
	Bar3DPercentStacked         ChartType = "Bar3DPercentStacked"
	Bar3DConeClustered          ChartType = "Bar3DConeClustered"
	Bar3DConeStacked            ChartType = "Bar3DConeStacked"
	Bar3DConePercentStacked     ChartType = "Bar3DConePercentStacked"
	Bar3DPyramidClustered       ChartType = "Bar3DPyramidClustered"
	Bar3DPyramidStacked         ChartType = "Bar3DPyramidStacked"
	Bar3DPyramidPercentStacked  ChartType = "Bar3DPyramidPercentStacked"
	Bar3DCylinderClustered      ChartType = "Bar3DCylinderClustered"
	Bar3DCylinderStacked        ChartType = "Bar3DCylinderStacked"
	Bar3DCylinderPercentStacked ChartType = "Bar3DCylinderPercentStacked"
	Col                         ChartType = "Col"
	ColStacked                  ChartType = "ColStacked"
	ColPercentStacked           ChartType = "ColPercentStacked"
	Col3D                       ChartType = "Col3D"
	Col3DClustered              ChartType = "Col3DClustered"
	Col3DStacked                ChartType = "Col3DStacked"
	Col3DPercentStacked         ChartType = "Col3DPercentStacked"
	Col3DCone                   ChartType = "Col3DCone"
	Col3DConeClustered          ChartType = "Col3DConeClustered"
	Col3DConeStacked            ChartType = "Col3DConeStacked"
	Col3DConePercentStacked     ChartType = "Col3DConePercentStacked"
	Col3DPyramid                ChartType = "Col3DPyramid"
	Col3DPyramidClustered       ChartType = "Col3DPyramidClustered"
	Col3DPyramidStacked         ChartType = "Col3DPyramidStacked"
	Col3DPyramidPercentStacked  ChartType = "Col3DPyramidPercentStacked"
	Col3DCylinder               ChartType = "Col3DCylinder"
	Col3DCylinderClustered      ChartType = "Col3DCylinderClustered"
	Col3DCylinderStacked        ChartType = "Col3DCylinderStacked"
	Col3DCylinderPercentStacked ChartType = "Col3DCylinderPercentStacked"
	Doughnut                    ChartType = "Doughnut"
	Line                        ChartType = "Line"
	Line3D                      ChartType = "Line3D"
	Pie                         ChartType = "Pie"
	Pie3D                       ChartType = "Pie3D"
	PieOfPie                    ChartType = "PieOfPie"
	BarOfPie                    ChartType = "BarOfPie"
	Radar                       ChartType = "Radar"
	Scatter                     ChartType = "Scatter"
	Surface3D                   ChartType = "Surface3D"
	WireframeSurface3D          ChartType = "WireframeSurface3D"
	Contour                     ChartType = "Contour"
	WireframeContour            ChartType = "WireframeContour"
	Bubble                      ChartType = "Bubble"
	Bubble3D                    ChartType = "Bubble3D"
)

type ChartSeries struct {
	Name              string                     `yaml:"Name"`
	Categories        string                     `yaml:"Categories"`
	Values            string                     `yaml:"Values"`
	Sizes             string                     `yaml:"Sizes"`
	Fill              Fill                       `yaml:"Fill"`
	Line              ChartLine                  `yaml:"Line"`
	Marker            ChartMarker                `yaml:"Marker"`
	DataLabelPosition ChartDataLabelPositionType `yaml:"DataLabelPosition"`
}

func (s *ChartSeries) ConvertExcelizeOption() (excelize.ChartSeries, error) {
	ret := excelize.ChartSeries{
		Name:       s.Name,
		Categories: s.Categories,
		Values:     s.Values,
		Sizes:      s.Sizes,
	}
	var err error
	ret.Fill, err = s.Fill.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	ret.Line, err = s.Line.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	ret.Marker, err = s.Marker.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	ret.DataLabelPosition, err = s.DataLabelPosition.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

type Fill struct {
	Type    string   `yaml:"Type"`
	Pattern int      `yaml:"Pattern"`
	Color   []string `yaml:"Color"`
	Shading int      `yaml:"Shading"`
}

func (f *Fill) ConvertExcelizeOption() (excelize.Fill, error) {
	return excelize.Fill{
		Type:    f.Type,
		Pattern: f.Pattern,
		Color:   f.Color,
		Shading: f.Shading,
	}, nil
}

type ChartLine struct {
	Type   ChartLineType `yaml:"Type"`
	Smooth bool          `yaml:"Smooth"`
	Width  float64       `yaml:"Width"`
}

func (l *ChartLine) ConvertExcelizeOption() (excelize.ChartLine, error) {
	ret := excelize.ChartLine{
		Smooth: l.Smooth,
		Width:  l.Width,
	}
	var err error
	ret.Type, err = l.Type.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

type ChartLineType string

func (t *ChartLineType) ConvertExcelizeOption() (excelize.ChartLineType, error) {
	if *t == "ChartLineSolid" {
		return excelize.ChartLineSolid, nil
	} else if *t == "ChartLineNone" {
		return excelize.ChartLineNone, nil
	} else if *t == "ChartLineAutomatic" {
		return excelize.ChartLineAutomatic, nil
	}
	return 0, nil
}

const (
	ChartLineSolid     ChartLineType = "ChartLineSolid"
	ChartLineNone      ChartLineType = "ChartLineNone"
	ChartLineAutomatic ChartLineType = "ChartLineAutomatic"
)

type ChartMarker struct {
	Fill   Fill   `yaml:"Fill"`
	Symbol string `yaml:"Symbol"`
	Size   int    `yaml:"Size"`
}

func (m *ChartMarker) ConvertExcelizeOption() (excelize.ChartMarker, error) {
	ret := excelize.ChartMarker{
		Symbol: m.Symbol,
		Size:   m.Size,
	}
	var err error
	ret.Fill, err = m.Fill.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

type ChartDataLabelPositionType string

func (t *ChartDataLabelPositionType) ConvertExcelizeOption() (excelize.ChartDataLabelPositionType, error) {
	if *t == "ChartDataLabelsPositionUnset" {
		return excelize.ChartDataLabelsPositionUnset, nil
	} else if *t == "ChartDataLabelsPositionBestFit" {
		return excelize.ChartDataLabelsPositionBestFit, nil
	} else if *t == "ChartDataLabelsPositionBelow" {
		return excelize.ChartDataLabelsPositionBelow, nil
	} else if *t == "ChartDataLabelsPositionCenter" {
		return excelize.ChartDataLabelsPositionCenter, nil
	} else if *t == "ChartDataLabelsPositionInsideBase" {
		return excelize.ChartDataLabelsPositionInsideBase, nil
	} else if *t == "ChartDataLabelsPositionInsideEnd" {
		return excelize.ChartDataLabelsPositionInsideEnd, nil
	} else if *t == "ChartDataLabelsPositionLeft" {
		return excelize.ChartDataLabelsPositionLeft, nil
	} else if *t == "ChartDataLabelsPositionOutsideEnd" {
		return excelize.ChartDataLabelsPositionOutsideEnd, nil
	} else if *t == "ChartDataLabelsPositionRight" {
		return excelize.ChartDataLabelsPositionRight, nil
	} else if *t == "ChartDataLabelsPositionAbove" {
		return excelize.ChartDataLabelsPositionAbove, nil
	}
	return 0, nil
}

const (
	ChartDataLabelsPositionUnset      ChartDataLabelPositionType = "ChartDataLabelsPositionUnset"
	ChartDataLabelsPositionBestFit    ChartDataLabelPositionType = "ChartDataLabelsPositionBestFit"
	ChartDataLabelsPositionBelow      ChartDataLabelPositionType = "ChartDataLabelsPositionBelow"
	ChartDataLabelsPositionCenter     ChartDataLabelPositionType = "ChartDataLabelsPositionCenter"
	ChartDataLabelsPositionInsideBase ChartDataLabelPositionType = "ChartDataLabelsPositionInsideBase"
	ChartDataLabelsPositionInsideEnd  ChartDataLabelPositionType = "ChartDataLabelsPositionInsideEnd"
	ChartDataLabelsPositionLeft       ChartDataLabelPositionType = "ChartDataLabelsPositionLeft"
	ChartDataLabelsPositionOutsideEnd ChartDataLabelPositionType = "ChartDataLabelsPositionOutsideEnd"
	ChartDataLabelsPositionRight      ChartDataLabelPositionType = "ChartDataLabelsPositionRight"
	ChartDataLabelsPositionAbove      ChartDataLabelPositionType = "ChartDataLabelsPositionAbove"
)

type GraphicOptions struct {
	AltText         string  `yaml:"AltText"`
	PrintObject     *bool   `yaml:"PrintObject"`
	Locked          *bool   `yaml:"Locked"`
	LockAspectRatio bool    `yaml:"LockAspectRatio"`
	AutoFit         bool    `yaml:"AutoFit"`
	OffsetX         int     `yaml:"OffsetX"`
	OffsetY         int     `yaml:"OffsetY"`
	ScaleX          float64 `yaml:"ScaleX"`
	ScaleY          float64 `yaml:"ScaleY"`
	Hyperlink       string  `yaml:"Hyperlink"`
	HyperlinkType   string  `yaml:"HyperlinkType"`
	Positioning     string  `yaml:"Positioning"`
}

func (t *GraphicOptions) ConvertExcelizeOption() (excelize.GraphicOptions, error) {
	return excelize.GraphicOptions{
		AltText:         t.AltText,
		PrintObject:     t.PrintObject,
		Locked:          t.Locked,
		LockAspectRatio: t.LockAspectRatio,
		AutoFit:         t.AutoFit,
		OffsetX:         t.OffsetX,
		OffsetY:         t.OffsetY,
		ScaleX:          t.ScaleX,
		ScaleY:          t.ScaleY,
		Hyperlink:       t.Hyperlink,
		HyperlinkType:   t.HyperlinkType,
		Positioning:     t.Positioning,
	}, nil
}

type ChartDimension struct {
	Width  uint `yaml:"Width"`
	Height uint `yaml:"Height"`
}

func (d *ChartDimension) ConvertExcelizeOption() (excelize.ChartDimension, error) {
	return excelize.ChartDimension{
		Width:  d.Width,
		Height: d.Height,
	}, nil
}

type ChartLegend struct {
	Position      string `yaml:"Position"`
	ShowLegendKey bool   `yaml:"ShowLegendKey"`
}

func (l *ChartLegend) ConvertExcelizeOption() (excelize.ChartLegend, error) {
	return excelize.ChartLegend{
		Position:      l.Position,
		ShowLegendKey: l.ShowLegendKey,
	}, nil
}

type RichTextRun struct {
	Font *Font  `yaml:"Font"`
	Text string `yaml:"Text"`
}

func (r *RichTextRun) ConvertExcelizeOption() (excelize.RichTextRun, error) {
	ret := excelize.RichTextRun{
		Text: r.Text,
	}
	if r.Font == nil {
		ret.Font = nil
	} else {
		font, err := r.Font.ConvertExcelizeOption()
		if err != nil {
			return ret, nil
		}
		ret.Font = &font
	}
	return ret, nil
}

type Font struct {
	Bold         bool    `yaml:"Bold"`
	Italic       bool    `yaml:"Italic"`
	Underline    string  `yaml:"Underline"`
	Family       string  `yaml:"Family"`
	Size         float64 `yaml:"Size"`
	Strike       bool    `yaml:"Strike"`
	Color        string  `yaml:"Color"`
	ColorIndexed int     `yaml:"ColorIndexed"`
	ColorTheme   *int    `yaml:"ColorTheme"`
	ColorTint    float64 `yaml:"ColorTint"`
	VertAlign    string  `yaml:"VertAlign"`
}

func (f *Font) ConvertExcelizeOption() (excelize.Font, error) {
	return excelize.Font{
		Bold:         f.Bold,
		Italic:       f.Italic,
		Underline:    f.Underline,
		Family:       f.Family,
		Size:         f.Size,
		Strike:       f.Strike,
		Color:        f.Color,
		ColorIndexed: f.ColorIndexed,
		ColorTheme:   f.ColorTheme,
		ColorTint:    f.ColorTint,
		VertAlign:    f.VertAlign,
	}, nil
}

type ChartAxis struct {
	None           bool          `yaml:"None"`
	MajorGridLines bool          `yaml:"MajorGridLines"`
	MinorGridLines bool          `yaml:"MinorGridLines"`
	MajorUnit      float64       `yaml:"MajorUnit"`
	TickLabelSkip  int           `yaml:"TickLabelSkip"`
	ReverseOrder   bool          `yaml:"ReverseOrder"`
	Secondary      bool          `yaml:"Secondary"`
	Maximum        *float64      `yaml:"Maximum"`
	Minimum        *float64      `yaml:"Minimum"`
	Font           Font          `yaml:"Font"`
	LogBase        float64       `yaml:"LogBase"`
	NumFmt         ChartNumFmt   `yaml:"NumFmt"`
	Title          []RichTextRun `yaml:"Title"`
}

func (a *ChartAxis) ConvertExcelizeOption() (excelize.ChartAxis, error) {
	ret := excelize.ChartAxis{
		None:           a.None,
		MajorGridLines: a.MajorGridLines,
		MinorGridLines: a.MinorGridLines,
		MajorUnit:      a.MajorUnit,
		TickLabelSkip:  a.TickLabelSkip,
		ReverseOrder:   a.ReverseOrder,
		Secondary:      a.Secondary,
		Maximum:        a.Maximum,
		Minimum:        a.Minimum,
		LogBase:        a.LogBase,
	}
	var err error
	ret.Font, err = a.Font.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	ret.NumFmt, err = a.NumFmt.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	runs := []excelize.RichTextRun{}
	for _, title := range a.Title {
		t, err := title.ConvertExcelizeOption()
		if err != nil {
			return ret, err
		}
		runs = append(runs, t)
	}
	ret.Title = runs
	return ret, nil
}

type ChartPlotArea struct {
	SecondPlotValues int         `yaml:"SecondPlotValues"`
	ShowBubbleSize   bool        `yaml:"ShowBubbleSize"`
	ShowCatName      bool        `yaml:"ShowCatName"`
	ShowLeaderLines  bool        `yaml:"ShowLeaderLines"`
	ShowPercent      bool        `yaml:"ShowPercent"`
	ShowSerName      bool        `yaml:"ShowSerName"`
	ShowVal          bool        `yaml:"ShowVal"`
	Fill             Fill        `yaml:"Fill"`
	NumFmt           ChartNumFmt `yaml:"NumFmt"`
}

func (a *ChartPlotArea) ConvertExcelizeOption() (excelize.ChartPlotArea, error) {
	ret := excelize.ChartPlotArea{
		SecondPlotValues: a.SecondPlotValues,
		ShowBubbleSize:   a.ShowBubbleSize,
		ShowCatName:      a.ShowCatName,
		ShowLeaderLines:  a.ShowLeaderLines,
		ShowPercent:      a.ShowPercent,
		ShowSerName:      a.ShowSerName,
		ShowVal:          a.ShowVal,
	}
	var err error
	ret.Fill, err = a.Fill.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	ret.NumFmt, err = a.NumFmt.ConvertExcelizeOption()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

type ChartNumFmt struct {
	CustomNumFmt string `yaml:"CustomNumFmt"`
	SourceLinked bool   `yaml:"SourceLinked"`
}

func (f *ChartNumFmt) ConvertExcelizeOption() (excelize.ChartNumFmt, error) {
	return excelize.ChartNumFmt{
		CustomNumFmt: f.CustomNumFmt,
		SourceLinked: f.SourceLinked,
	}, nil
}

type ExcelizeChartOption struct {
	Sheet string
	Cell  string
	Chart *excelize.Chart
	Combo []*excelize.Chart
}

func ConvertExelizeChartOption(chartOpts []ChartOption) ([]*ExcelizeChartOption, error) {
	result := []*ExcelizeChartOption{}
	for _, chartOpt := range chartOpts {
		excelizeChartOpt, err := chartOpt.ConvertExcelizeOption()
		if err != nil {
			return nil, err
		}
		result = append(result, excelizeChartOpt)
	}
	return result, nil
}

func GenerateAutoGraphChartOption(valueRange string, title string) *excelize.Chart {
	return &excelize.Chart{
		Type: excelize.Line,
		Series: []excelize.ChartSeries{
			{
				Values: valueRange,
			},
		},
		Title: []excelize.RichTextRun{
			{
				Text: title,
			},
		},
	}
}
