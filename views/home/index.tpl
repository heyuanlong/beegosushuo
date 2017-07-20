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
  <div class="container_w">
    <div class="container">

      <div class="itemcontent">
        <div class="content_list">
          {{range $k,$v := .list}}
            <div class="item">
              <div class="title"><a class="article_title" href="/article/{{$v.id}}">{{$v.title}}</a></div>
              <div class="content_attr">{{$v.time}}|发布人：xxx|评论：{{$v.comments}}人|浏览:{{$v.views}}</div>
              <div id="content_{{$v.id}}" class="content_n">
                    {{$v.content}}      
              </div>
            </div>
          {{end}}
        </div>

        <div class="page_forword">
          {{if .prev_page_flag}}
          <a class="prev_page" href="{{.prev_page}}" class="page-nav">上一页</a>
          {{end}}
          {{if .next_page_flag}}
          <a class="next_page" href="{{.next_page}}" class="page-nav">下一页</a>
          {{end}}
        </div>
      </div>

    </div>
  </div>

  {{template "home/inc_footer.tpl" .}}

</body>

<script type="text/javascript">

</script>

</html>
