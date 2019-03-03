# 開発ガイドライン
## ディレクトリ構成とその役割

## ディレクトリ構成とその従属性と関連

1. 1 action has 1 parameter
コントローラーの各アクセスに対して、parameterファイルを一つ作成する
作成場所は /src/controller/param
パラメータはかぶることもあるけれど、影響範囲を小さくするためにほぼ同じでもコピーして作成する。
命名規則はそれで一意としてわかるように
{request_method}{url}_param.go
ex
GET /city/build_count
get_city_build_count_param.go

2. 1 controller has 1 interactor

3. 1 action has 1 interactors method

4. 1 interactor has many repositories

5. 1 repository has 1 or less data access object

6. 1 interactor has 1 response struct


ControllerでのFindなのかGetなのかの命名規則
1. GetはCitiesに対して一つのCityを返す場合
0県の場合はstatuscode 404
1. FindはCitiesに対して複数のCityを返す場合
0県の場合は[] 200


repositoryの分割
1. repositoryが肥大化してきた時には分割して対応する。

update
