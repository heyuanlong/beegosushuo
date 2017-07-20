<style type="text/css">
.box_right{
  background-color: white;
  float: left;
    width: 18%;
    margin-bottom: 20px;
    margin-left: 2%;
  padding-left: 1%;
    word-wrap:break-word;
}
.box_title{
  color: black;
  font-weight:bold;
  font-size: 16px;
}
.box_ul{
  padding-left:0%;
  list-style-type:none;
}
.box_ul li{
  padding-bottom: 2%;
}

.box_right_article_title{
  color: gray;
  font-weight:bold;
  text-decoration: none;
  font-size: 13px;
}
.box_right_article_title:hover{
  color: black;
  font-weight:bold;
  font-size: 14px;
  text-decoration: none;
}


</style>

    <div class="box_right">
      <div>
        <div class="box_title">最热博文</div>
        <ul class="box_ul">
          {{range $k,$v := .hottestviews}}
          <li>
            <a class="box_right_article_title" href="/article/{{$v.id}}">>{{$v.title}}</a>
            <span class="view_count">{{$v.views}}</span>
          </li>
          {{end}}
        </ul>
      </div>
    </div>
    