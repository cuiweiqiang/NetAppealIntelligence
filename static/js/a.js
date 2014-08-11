<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/2000/REC-xhtml1-20000126/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=gb2312" />
<meta name="keywords" content="站长,网页特效,js特效,广告代码,zzjs,zzjs.net,sky,www.zzjs.net,站长特效 网" />
<meta name="description" content="www.zzjs.net,站长特效网，站长必备js特效及广告代码。大量高质量js特效，提供高质量广告代码下载,尽在站长特效网" />
<script type="text/javascript">
function encode(s){
  return s.replace(/&/g,"&").replace(/</g,"<").replace(/>/g,">").replace(/([\\\.\*\[\]\(\)\$\^])/g,"\\$1");
}//欢迎来到站长特效网，我们的网址是www.zzjs.net，很好记，zz站长，js就是js特效，本站收集大量高质量js代码，还有许多广告代码下载。
function decode(s){
  return s.replace(/\\([\\\.\*\[\]\(\)\$\^])/g,"$1").replace(/>/g,">").replace(/</g,"<").replace(/&/g,"&");
}
function highlight(s){
  if (s.length==0){
    alert('搜索关键词未填写！');
    return false;
  }
  s=encode(s);
  var obj=document.getElementsByTagName("body")[0];
  var t=obj.innerHTML.replace(/<span\s+class=.?highlight.?>([^<>]*)<\/span>/gi,"$1");
  obj.innerHTML=t;
  var cnt=loopSearch(s,obj);
  t=obj.innerHTML
  var r=/{searchHL}(({(?!\/searchHL})|[^{])*){\/searchHL}/g
  t=t.replace(r,"<span class='highlight'>$1</span>");
  obj.innerHTML=t;
  alert("搜索到关键词"+cnt+"处")
}//欢迎来到站长特效网，我们的网址是www.zzjs.net，很好记，zz站长，js就是js特效，本站收集大量高质量js代码，还有许多广告代码下载。
function loopSearch(s,obj){
  var cnt=0;
  if (obj.nodeType==3){
    cnt=replace(s,obj);
    return cnt;
  }
  for (var i=0,c;c=obj.childNodes[i];i++){
    if (!c.className||c.className!="highlight")
      cnt+=loopSearch(s,c);
  }
  return cnt;
}
function replace(s,dest){
  var r=new RegExp(s,"g");
  var tm=null;
  var t=dest.nodeValue;
  var cnt=0;
  if (tm=t.match(r)){
    cnt=tm.length;
    t=t.replace(r,"{searchHL}"+decode(s)+"{/searchHL}")
    dest.nodeValue=t;
  }
  return cnt;
}//欢迎来到站长特效网，我们的网址是www.zzjs.net，很好记，zz站长，js就是js特效，本站收集大量高质量js代码，还有许多广告代码下载。
</script>
<style type="text/css">
.highlight{background:green;font-weight:bold;color:white;}
</style>
</head>
<body>
<a href="http://www.zzjs.net/">站长特效网</a>,站长必备的高质量网页特效和广告代码。zzjs.net，站长js特效。<hr>
<!--欢迎来到站长特效网，我们网站收集大量高质量js特效，提供许多广告代码下载，网址：www.zzjs.net，zzjs@msn.com,用.net打造靓站-->
<script type="text/javascript" src="http://www.zzjs.net/ad/tc.js"></script>
<form onsubmit="highlight(this.s.value);return false;">
<p><input name="s" id="s" title="输入站长特效看看"/><input type="submit" value="搜索"/></p>
</form>
<div id="content">
<font face=Verdana>站长特效网,www.zzjs.net<br />
专注于网页特效及广告代码,供站长和js爱好者参考.<br />
<br />
打造国内最帅的网页特效站!
<br />
欢迎收藏本站,您的选择是不会错的.
<br />
站长特效,永远是最帅的!</font>
</div>
</body>
</html>