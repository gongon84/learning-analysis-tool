<br>

[![Maintainability](https://api.codeclimate.com/v1/badges/6deb19b3248771740587/maintainability)](https://codeclimate.com/github/gongon84/learning-golangci/maintainability)[![Test Coverage](https://api.codeclimate.com/v1/badges/6deb19b3248771740587/test_coverage)](https://codeclimate.com/github/gongon84/learning-golangci/test_coverage)[![lint status](https://github.com/goark/koyomi/workflows/lint/badge.svg)](https://github.com/gongon84/learning-golangci/actions)

# 概要
* [code climate](https://codeclimate.com/)

    * 保守性
        * コードの保守性を評価。適宜コメントで指摘をしてくれる
    
    * テストカバレッジ
        * テストカバレッジを計測し表示。

* [golangci-lint]()
    
    * linterが実行される

* [lighthouse-ci](https://github.com/GoogleChrome/lighthouse-ci)

    * プルリクエストへのコメント（/run lighthouse）をトリガーに以下を実行
        * lighthouseでサイトパフォーマンスを計測（ngrokで一時公開したものを計測）
        * 計測結果のURLをプルリクでコメントしてくれる

# キャプチャ

* CI連携

    <img src="https://github.com/gongon84/learning-golangci/assets/57177320/d74f43d2-8717-40b5-a0de-26af0be98a90" width="600">

* lighthouse 計測
    * コメント

       <img src="https://github.com/gongon84/learning-analysis-tool/assets/57177320/0d965d01-4a76-494f-a9e5-6e4d048c5d9d" width="500"> 

    * 計測結果

       <img src="https://github.com/gongon84/learning-analysis-tool/assets/57177320/3f298c16-05be-409f-a374-a399638940f1" width="500">
  