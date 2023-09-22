# ステータス

[![Maintainability](https://api.codeclimate.com/v1/badges/6deb19b3248771740587/maintainability)](https://codeclimate.com/github/gongon84/learning-golangci/maintainability)[![Test Coverage](https://api.codeclimate.com/v1/badges/6deb19b3248771740587/test_coverage)](https://codeclimate.com/github/gongon84/learning-golangci/test_coverage)[![lint status](https://github.com/goark/koyomi/workflows/lint/badge.svg)](https://github.com/gongon84/learning-golangci/actions)
  
<br>

# 概要
* [lighthouse-ci](https://github.com/GoogleChrome/lighthouse-ci)

    * プルリクエストへのコメント（/run lighthouse）をトリガーに以下が実行される
        * lighthouseでサイトパフォーマンスを計測（ngrokで一時公開したものを計測）
        * 計測結果のURLをプルリクでコメントしてくれる

* [code climate](https://codeclimate.com/)

    * コードの保守性を評価。適宜コメントで指摘をしてくれる
    
    * テストカバレッジを計測し表示。

* [golangci-lint]()
    
    * push時にlinterが実行


# キャプチャ

* CI連携

    <img src="https://github.com/gongon84/learning-golangci/assets/57177320/d74f43d2-8717-40b5-a0de-26af0be98a90" width="600">

* lighthouse 計測
    * コメント

       <img src="https://github.com/gongon84/learning-analysis-tool/assets/57177320/be6b2706-1d33-4f15-bf35-3632e31ee211" width="500"> 

    * 計測結果

       <img src="https://github.com/gongon84/learning-analysis-tool/assets/57177320/3f298c16-05be-409f-a374-a399638940f1" width="500">
  