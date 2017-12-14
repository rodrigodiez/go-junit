# go-junit
`go-junit` is a collection of types to marshall and unmarshall [JUnit](http://junit.org/junit5/) XML reports using [Go](https://golang.org/).

I created this project to help [Smocha](https://github.com/rodrigodiez/smocha) to generate compatible JUnit xml reports

# Example
```go
package main

import (
  "github.com/rodrigodiez/go-junit/types"
  "encoding/xml"
  "fmt"
)

func main(){
  testsuites := &types.Testsuites{
    Id: "The ID of the scan",
    Name: "The label of the scan",
    Tests: 2,
    Failures: 2,
    Time: 0.01,
    Testsuites: []*types.Testsuite{{
      Id: "The ID of the provider",
      Name: "The label of the provider",
      Tests: 2,
      Failures: 2,
      Time: 0.01,
      Testcases: []*types.Testcase{{
        Id: "The ID of the rule",
        Name: "The label of the rule",
        Time: 0.01,
        Failures: []*types.Failure{{
          Message: "The source code file, the line number, and the rule that is violated",
          Type: "The severity of the rule",
          Text: "The text of the rule and the severity"}}}}}}}

  output, _ := xml.MarshalIndent(testsuites, "  ", "    ")
  fmt.Println(string(output[:]))
}
```
## Output
```xml
<testsuites id="The ID of the scan" name="The label of the scan" tests="2" failures="2" time="0.01">
    <testsuite id="The ID of the provider" name="The label of the provider" tests="2" failures="2" time="0.01">
        <testcase id="The ID of the rule" name="The label of the rule" time="0.01">
            <failure message="The source code file, the line number, and the rule that is violated" type="The severity of the rule">The text of the rule and the severity</failure>
        </testcase>
    </testsuite>
</testsuites>
```
