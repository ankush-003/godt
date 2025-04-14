# Go Decision Trees 🌲

A high-performance implementation of tree-based machine learning algorithms in Go. This repository aims to provide efficient implementations of common decision tree algorithms and gradient boosting frameworks for machine learning applications.

![gopher with tree](https://github.com/user-attachments/assets/117bdabc-007d-4fd8-a4aa-264a310e29f9)


## Installation 🚀

To install the Go Decision Trees library, you can use the following command:

```bash
go get github.com/ankush-003/godt
```

## Usage 📚

### Decision Tree Classifier

```go
package main

import (
    "fmt"
    "github.com/ankush-003/godt/decision_tree_regressor"
    "github.com/ankush-003/godt/internal"
)

func main() {
    // Load data from CSV
    // Format: features columns with target as the last column
    data := internal.NewData(150, internal.FromCSV("path/to/data.csv"))

    // Initialize decision tree classifier
    // Parameters:
    // - data: training data
    // - maxDepth: maximum depth of the tree (12 in this example)
    // - minSamplesSplit: minimum samples required to split a node (10 in this example)
    // - targetColumn: index of the target column (4 in this example)
    dt := decision_tree_regressor.NewDecisionTreeClassifier(
        *data,
        12,  // maxDepth
        10,  // minSamplesSplit
        4,   // targetColumn
    )

    // Train the model
    dt.Fit()

    // Make predictions
    testData := internal.NewData(10, internal.FromCSV("path/to/test_data.csv"))
    predictions := dt.Predict(*testData)

    fmt.Println("Predictions:", predictions)
}
```

## Todo

- [x] Decision Tree Classifier
- [ ] Decision Tree Regressor
- [ ] Training Optimisation
- [ ] Random Forest
- [ ] Gradient Boost Regression Trees
