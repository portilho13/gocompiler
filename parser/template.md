Program
  ├── IncludeDirective: <stdio.h>
  └── FunctionDeclaration: main
      └── CompoundStatement
          ├── VariableDeclaration: x (int)
          │   └── AssignmentExpression: x = 10
          ├── VariableDeclaration: y (int)
          │   └── AssignmentExpression: y = 20
          ├── VariableDeclaration: sum (int)
          │   └── AssignmentExpression: sum = BinaryExpression (+)
          │       ├── Identifier: x
          │       └── Identifier: y
          ├── FunctionCallExpression: printf("Sum: %d\n", sum)
          │   ├── StringLiteral: "Sum: %d\n"
          │   └── Identifier: sum
          └── ReturnStatement: 0
