fssh
===
fsshは指定したリソース量の空きのあるサーバーへ接続するsshコマンドです。

# 使い方


## 書式　　
  ``
  fssh [オプション]
  ``
## オプション    
 + -c 数値  
 使用CPU数を指定する。デフォルトは２
 ---
 以下は未実装
 + -m 数値
 使用メモリーサイズ (MB)を指定する。
 + -g
 gpu使用のする場合
 + -s　hostname1,hostname2,...
 接続ホストの制限


## 設定  
 ssh接続の設定はsshコマンドを内部で利用しているため、~/.ssh/configに以下のように書き加えると反映される。 　　

```
Host file-server
HostName 192.218.174.32
User [自分のアカウント]

Host ss0[1~4 のいずれか]
HostName ss0[1~4 のいずれか]
User [自分のアカウント]
LocalForward [学籍番号下 4 桁] localhost:[学籍番号下 4 桁]
ProxyCommand  ssh -CW %h:%p file-server
```
