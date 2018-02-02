<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="description" content="">
        <meta name="author" content="">

        <title>Server</title>

        <!-- Bootstrap Core CSS -->
        <link href="/assets/startmin/css/bootstrap.min.css" rel="stylesheet">

        <!-- MetisMenu CSS -->
        <link href="/assets/startmin/css/metisMenu.min.css" rel="stylesheet">

        <!-- Timeline CSS -->
        <link href="/assets/startmin/css/timeline.css" rel="stylesheet">

        <!-- Custom CSS -->
        <link href="/assets/startmin/css/startmin.css" rel="stylesheet">

        <!-- Morris Charts CSS -->
        <link href="/assets/startmin/css/morris.css" rel="stylesheet">

        <!-- Custom Fonts -->
        <link href="/assets/startmin/css/font-awesome.min.css" rel="stylesheet" type="text/css">

        <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
        <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
        <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
        <![endif]-->
    </head>
    <body>

        <div id="wrapper">

            <!-- Navigation -->
            <nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
                <div class="navbar-header">
                    <a class="navbar-brand" href="/">Server</a>
                </div>

                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>

                <ul class="nav navbar-nav navbar-left navbar-top-links">
                    <li><a href="#"><i class="fa fa-home fa-fw"></i> Website</a></li>
                </ul>

                <ul class="nav navbar-right navbar-top-links">
                    <li class="dropdown">
                        <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                            <i class="fa fa-user fa-fw"></i> secondtruth <b class="caret"></b>
                        </a>
                        <ul class="dropdown-menu dropdown-user">
                            <li><a href="#"><i class="fa fa-user fa-fw"></i> User Profile</a>
                            </li>
                            <li><a href="#"><i class="fa fa-gear fa-fw"></i> Settings</a>
                            </li>
                            <li class="divider"></li>
                            <li><a href="login.html"><i class="fa fa-sign-out fa-fw"></i> Logout</a>
                            </li>
                        </ul>
                    </li>
                </ul>
                <!-- /.navbar-top-links -->

                <div class="navbar-default sidebar" role="navigation">
                    <div class="sidebar-nav navbar-collapse">
                        <ul class="nav" id="side-menu">
                            <li>
                                <a href="index.html" class="active"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <div id="page-wrapper">
                <div class="row">
                    <div class="col-lg-12">
                        <h1 class="page-header">Dashboard</h1>
                    </div>
                    <!-- /.col-lg-12 -->
                </div>

    <br>
    ホスト名:{{.Ip}}
    <br>

    <br>
    <div class="row">
      <div class="col-lg-6">

<canvas id="CPU"></canvas>

      </div>
      <div class="col-lg-6">
        <canvas id="Mem"></canvas>
          </div>
    </div>

            </div>
            <!-- /#page-wrapper -->


        </div>
        <!-- /#wrapper -->

        <!-- jQuery -->
        <script src="/assets/startmin/js/jquery.min.js"></script>

        <!-- Bootstrap Core JavaScript -->
        <script src="/assets/startmin/js/bootstrap.min.js"></script>

        <!-- Metis Menu Plugin JavaScript -->
        <script src="/assets/startmin/js/metisMenu.min.js"></script>

        <!-- Morris Charts JavaScript -->
        <script src="/assets/startmin/js/raphael.min.js"></script>
        <script src="/assets/startmin/js/morris.min.js"></script>
        <script src="/assets/startmin/js/morris-data.js"></script>

        <!-- Custom Theme JavaScript -->
        <script src="/assets/startmin/js/startmin.js"></script>


        <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.18.1/moment.js"></script>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.6.0/Chart.js"></script>
<script type="text/javascript" src="https://github.com/nagix/chartjs-plugin-streaming/releases/download/v1.2.0/chartjs-plugin-streaming.js"></script>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/pusher/4.1.0/pusher.js"></script><script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.18.1/moment.js"></script>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.6.0/Chart.js"></script>
<script type="text/javascript" src="https://github.com/nagix/chartjs-plugin-streaming/releases/download/v1.2.0/chartjs-plugin-streaming.js"></script>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/pusher/4.1.0/pusher.js"></script>
<script type="text/javascript">
  var buf = {};
  buf['CPU'] = [[],[]];
  buf['Mem'] = [[],[]];
      setInterval(function() {
      $.ajax({
        url: "/api/status?id={{.Id}}",
      })
      .done(function(data) {
        console.log(data);
        buf['CPU'][0].push({x:data.Time,y:data.CpuPer[0]})
        if(data.CpuPer.length >1){
            buf['CPU'][1].push({x:data.Time,y:data.CpuPer[1]})
        }
        buf['Mem'][0].push({x:data.Time,y:data.Vmem.used})
        buf['Mem'][1].push({x:data.Time,y:data.Vmem.free})
      })
      .fail(function(data) {
      });
    }  , 5000 );


// buf['Bitfinex'] = [[], []];
// var ws = new WebSocket('wss://api.bitfinex.com/ws/');
// ws.onopen = function() {
//     ws.send(JSON.stringify({      // 購読リクエストを送信
//         "event": "subscribe",
//         "channel": "trades",
//         "pair": "BTCUSD"
//     }));
// };
// ws.onmessage = function(msg) {     // メッセージ更新時のコールバック
//     var response = JSON.parse(msg.data);
//     if (response[1] === 'te') {    // メッセージタイプ 'te' だけを見る
//     console.log(response[3] * 1000);
//         // buf['Bitfinex'][response[5] > 0 ? 0 : 1].push({
//         //     x: response[3] * 1000, // タイムスタンプ（ミリ秒）
//         //     y: response[4]         // 価格（米ドル）
//         // });
//     }
// }
var id = 'CPU';
var ctx = document.getElementById(id).getContext('2d');
var chart = new Chart(ctx, {
  type: 'line',
  data: {
      datasets: [{
          data: [],
          label: 'CPU1',                     // 買い取引データ
          borderColor: 'rgb(255, 99, 132)', // 線の色
          backgroundColor: 'rgba(255, 99, 132, 0.5)', // 塗りの色
          fill: false,                      // 塗りつぶさない
          lineTension: 0                    // 直線
      }, {
          data: [],
          label: 'CPU2',                    // 売り取引データ
          borderColor: 'rgb(54, 162, 235)', // 線の色
          backgroundColor: 'rgba(54, 162, 235, 0.5)', // 塗りの色
          fill: false,                      // 塗りつぶさない
          lineTension: 0                    // 直線
      }]
  },
  options: {
      title: {
          text: 'ステータス (' + id + ')', // チャートタイトル
          display: true
      },
      scales: {
          xAxes: [{
              type: 'realtime' // X軸に沿ってスクロール
          }],
          yAxes: [{
            ticks: {
                  beginAtZero: true,
                  min: 0,
                  max: 200
                }
          }]
      },
      plugins: {
          streaming: {
              duration: 300000, // 300000ミリ秒（5分）のデータを表示
              onRefresh: function(chart) { // データ更新用コールバック
                  Array.prototype.push.apply(
                      chart.data.datasets[0].data, buf[id][0]
                  );            // 買い取引データをチャートに追加
                  Array.prototype.push.apply(
                      chart.data.datasets[1].data, buf[id][1]
                  );            // 売り取引データをチャートに追加
                  buf[id] = [[],[]]; // バッファをクリア
              }
          }
      }
  }
});
var id2 = 'Mem';
var ctx2 = document.getElementById(id2).getContext('2d');
var chart2 = new Chart(ctx2, {
  type: 'line',
  data: {
      datasets: [{
          data: [],
          label: 'Used',                     // 買い取引データ
          borderColor: 'rgb(255, 99, 132)', // 線の色
          backgroundColor: 'rgba(255, 99, 132, 0.5)', // 塗りの色
          fill: false,                      // 塗りつぶさない
          lineTension: 0                    // 直線
      }, {
          data: [],
          label: 'Free',                    // 売り取引データ
          borderColor: 'rgb(54, 162, 235)', // 線の色
          backgroundColor: 'rgba(54, 162, 235, 0.5)', // 塗りの色
          fill: false,                      // 塗りつぶさない
          lineTension: 0                    // 直線
      }]
  },
  options: {
      title: {
          text: 'ステータス (' + id2 + ')', // チャートタイトル
          display: true
      },
      scales: {
          xAxes: [{
              type: 'realtime' // X軸に沿ってスクロール
          }],
          yAxes: [{
            ticks: {
                beginAtZero: true,
                min: 0,
                
              }

          }]
      },
      plugins: {
          streaming: {
              duration: 300000, // 300000ミリ秒（5分）のデータを表示
              onRefresh: function(chart) { // データ更新用コールバック
                  Array.prototype.push.apply(
                      chart.data.datasets[0].data, buf[id2][0]
                  );            // 買い取引データをチャートに追加
                  Array.prototype.push.apply(
                      chart.data.datasets[1].data, buf[id2][1]
                  );            // 売り取引データをチャートに追加
                  buf[id2] = [[],[]]; // バッファをクリア
              }
          }
      }
  }
});
</script>

    </body>
</html>
