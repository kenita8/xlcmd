# example for chart
..\csv2xlsx --input data --xlsx outputChart.xlsx
..\chart --config chart.yml --xlsx outputChart.xlsx
..\cellset --xlsx outputChart.xlsx --sheet cpu.tsv --range "A1:IJ1" --pattern "\\\\\\\\LAPTOP\\\\" --replacement ""

# example for ezchart
..\csv2xlsx --input data\cpu.tsv --xlsx outputEzChart.xlsx
..\ezchart --xlsx outputEzChart.xlsx
