<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="description" content="">
        <meta name="author" content="">

        <title>Server</title>

          <!-- 以下をGoogleににホストされたものに変更 -->
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


                <!-- /.navbar-top-links -->

                <div class="navbar-default sidebar" role="navigation">
                    <div class="sidebar-nav navbar-collapse">
                        <ul class="nav" id="side-menu">
                          <li>
                              <a href="/" class="active"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
                          </li>
                          <li>
                              <a href="/add">add node</a>
                          </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <div id="page-wrapper">
                <div class="row">
                    <div class="col-lg-12">
                        <h1 class="page-header">status</h1>
                    </div>
                    <!-- /.col-lg-12 -->
                </div>

    <br>
    ホスト名:{{.Hostname}}
    <br>

    <br>
    <div class="row">
      {{range $i, $v := .CpuCount}}
      <div class="col-lg-3 col-xs-6">
      <canvas id="CPU{{$i}}" class="CPU"></canvas>
    </div>
      {{end}}
    </div>
    <div class="row">
        <canvas id="Mem"></canvas>
    </div>

            </div>
            <!-- /#page-wrapper -->


        </div>
        <!-- /#wrapper -->

          <!-- 以下をGoogleににホストされたものに変更 -->
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

  var cpucount =$('.CPU').length;
  var buf = [];
  buf['CPU'] = [];
  for (var i = 0; i < cpucount; i++) {
    buf['CPU'].push([])
  }
  console.log(buf);
  buf['Mem'] = [[],[]];
      setInterval(function() {
      $.ajax({
        url: "/api/status?id={{.Id}}",
      })
      .done(function(data) {
        for (var i = 0; i < cpucount; i++) {

          buf['CPU'][i].push({x:data.Time,y:data.CpuPer[i]});
        }
        buf['Mem'][0].push({x:data.Time,y:data.Vmem.used})
        buf['Mem'][1].push({x:data.Time,y:data.Vmem.free})
        console.log(buf);
      })
      .fail(function(data) {
      });
    }  , 5000 );

$('.CPU').each(function(i, elem){
  var id = elem.id;
  var ctx = document.getElementById(id).getContext('2d');
  var chart = new Chart(ctx, {
    type: 'line',
    data: {
        datasets: [{
            data: [],
            label: id,                     // 買い取引データ
            borderColor: 'rgb(255, 99, 132)', // 線の色
            backgroundColor: 'rgba(255, 99, 132, 0.5)', // 塗りの色
            fill: false,                      // 塗りつぶさない
            lineTension: 0                    // 直線
        }]
    },
    options: {
        title: {
            text: id, // チャートタイトル
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
                    max: 100
                  }
            }]
        },
        legend: {
            display: false
         },
        plugins: {
            streaming: {
                duration: 600000, // 300000ミリ秒（5分）のデータを表示
                onRefresh: function(chart) { // データ更新用コールバック
                    Array.prototype.push.apply(
                        chart.data.datasets[0].data, buf["CPU"][i]
                    );            // 買い取引データをチャートに追加
                    buf["CPU"][i] = []; // バッファをクリア
                }
            }
        }
    }
  });
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
              duration: 600000, // 300000ミリ秒（5分）のデータを表示
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
