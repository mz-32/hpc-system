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
                              <a href="/" ><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
                          </li>
                          <li>
                              <a href="/add" class="active">add node</a>
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
                <form role="form" method="post">
                    <div class="form-group">
                        <label>HOSTNAME</label>
                        <input name="IP" class="form-control" value="{{if .Ip}}{{.Ip}}{{end}}">
                    </div>

                    <button type="submit" class="btn btn-default">Submit Button</button>
                  </form>


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

    </body>
</html>
