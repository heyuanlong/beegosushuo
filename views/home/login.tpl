<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="Pragma" content="no-cache">
<meta http-equiv="Cache-Control" content="no-cache">
<meta http-equiv="Expires" content="0">
<title>resume</title>



<!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

<!-- 可选的 Bootstrap 主题文件（一般不用引入） -->
<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">

<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>




<link href="/static/css/resume.css" type="text/css" rel="stylesheet" charset="utf-8" />
<script type="text/javascript">

</script>

</head>
<body>
  
  {{template "home/inc_header.tpl" .}}
  <div class="container_w">
    <div class="container">

      		<form action="/home/login" method="post" class="center">
				<div class="input-group input-group-lg">
				  <span class="input-group-addon" id="sizing-addon1">账号</span>
				  <input type="text"  name="account" class="form-control" placeholder="账号" aria-describedby="sizing-addon1">
				</div>
				<div class="input-group input-group-lg">
				  <span class="input-group-addon" id="sizing-addon1">密码</span>
				  <input type="password" name="password" class="form-control" placeholder="密码" aria-describedby="sizing-addon1">
				</div>

				<button type="submit" class="btn btn-primary">登录</button>
			</form>
    </div>
  </div>

  {{template "home/inc_footer.tpl" .}}

</body>

<script type="text/javascript">

</script>

</html>
