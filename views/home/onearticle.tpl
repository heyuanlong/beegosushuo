<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta http-equiv="Pragma" content="no-cache">
<meta http-equiv="Cache-Control" content="no-cache">
<meta http-equiv="Expires" content="0">
<title>resume</title>
<link href="/static/css/resume.css" type="text/css" rel="stylesheet" charset="utf-8" />

<script type="text/javascript">

</script>

</head>
<body>
  
  {{template "home/inc_header.tpl" .}}
  <div class="container">

    <div class="itemcontent">
      <div class="content_list">
          <div class="item">
            <div class="title"><a class="article_title" href="/article/{{.id}}">{{.title}}</a></div>
            <div class="content_attr">
              time:<span class="time">{{.time}}</span>
              view:<span class="view_count">{{.views}}</span>
              classify:<span class="classify">{{.comments}}</span>
            </div>
            <div id="content" class="content_n">
                  <textarea style="display:none;">{{.content}}</textarea>          
            </div>
          </div>
      </div>
    </div>

  </div>

  {{template "home/inc_footer.tpl" .}}

</body>

<script type="text/javascript">

</script>

</html>

