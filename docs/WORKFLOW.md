# WORKFLOW.md
## プレフィックス

|  prefix  |  stands for  |  discription  |
| :---- | :---- | :---- |
|  docs  |  docmentations  |  ドキュメント  |
|  feat  |  features  |  新機能  |
|  refac  |  refactoring  |  リファクタリング  |
|  fix  |    |  バグの修正や間違いの訂正  |
|  chore  |    |  パッケージのアップデートなど小規模の変更  |
|  style  |  |  デザイン面の変更  |

## ブランチの命名規則
```
[prefix]/[issue_number]_[issue_title]
```

example
```
feat/21_add_authentications
```

## コミットメッセージの命名規則
```
[prefix]: [a discription of changes] [#issue_number]
```

example
```
fix: delete extra comments #15
```