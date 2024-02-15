# csvconvert

Simple pipe friendly tool for converting CSV as JSON.

### Installation
`go install github.com/gaigals/csvconvert@latest`

### Usage samples:
```sh
// Can be used to pipe in csv content.
cat some_file.csv | csvconvert | jq .
```

```sh
// Directly provide CSV content as input.
csvconvert -i `csv_content` 
```

```sh
// Provide CSV file path. File has to have .csv extenstion.
csvconvert -i `some_file.csv`
```

More info:
```
csvconvert -h
```
