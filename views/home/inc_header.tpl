  
<style type="text/css">

.header_go{
  background: none repeat scroll 0% 0% #ffd52b;
  width: 100%;
  padding-top: 1%;

}
.header_go a{
  text-decoration: none;
}
.header_go a:hover {
  text-decoration: none;
  color: #000;
}
.header_img{
  margin-left: 19%;
}
.header_reglogin{
  float: right;
  margin-right: 33%;
  margin-top: 2%;
}

.header_nav{
  height: 52px;
    background-color: #fff;
    opacity: 0.9;
    clear: both;
}
.header_nav ul{
  padding-left: 0px;
    width: 61.8%;
    margin: 0px auto;
}
.header_nav li{
    position: relative;
    float: left;
    margin-right: 1px;
    display: block;
    min-width: 130px;
    width: auto;
    line-height: 52px;
    font-size: 20px;
    color: #fff;
    text-align: left;
}
.header_nav li:hover{
  background-color: #ffd52b;
}
.header_nav li a{
  text-decoration: none;
  color: #000;
  font-family: Arial, Helvetica,'Microsoft Yahei', 'Hiragino Sans GB', sans-serif;
}

</style>

  <div>
    <div class="header_go">
      <span class="header_img">
        <a href="/">
          <img src="/static/img/header.png">
        </a>
      </span>
      <span class="header_reglogin">
      {{if .account }}
      <a href="/"><span>{{.account}}</span></a>|
      <a href="/home/logout"><span class="header_logout">退出</span></a>
      {{else}}
      <a href="/home/reg"><span class="header_reg">注册</span></a>|
      <a href="/home/login"><span class="header_login">登录</span></a>
      {{end}}
      </span>
    </div>
    <div class="header_nav">
      <ul>
        <li><a href="/">最新</a></li>
        <li><a href="/">最热</a></li>
      </ul>
    </div>
  </div>