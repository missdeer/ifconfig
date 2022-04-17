<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang ="en">
  <head>
    <meta http-equiv="content-Type" content="text/html; charset=utf-8" />
    <meta http-equiv="content-style-type" content="text/css" />
    <meta http-equiv="content-script-type" content="text/javascript" />
    <meta http-equiv="content-language" content="en" />
    <meta http-equiv="pragma" content="no-cache" />
    <meta http-equiv="cache-control" content="no-cache" />
    <meta name="description" content="Get my IP Address" />
    <meta name="keywords" content="ip address ifconfig" />
    <meta name="author" content="{{.Author}}" />
    <link rel="shortcut icon" href="/static/img/favicon.ico" />
    <link rev="made" href="mailto: {{.Email}}" />
    <title>What Is My IP Address?</title>
    <style type="text/css">
    * {margin:0; padding:0;font-style:normal;font-weight:normal;font-family:Arial,sans-serif;font-size:13px;color:#333;}
      body{}
      #container{background:white;width:750px;margin:10px auto;margin-bottom:10px;border:solid 1px #888;}
      #header{height:50px;padding:15px 0 10px 10px;}
      #header table{width:100%;}
      h1 a{font-size:20px;margin:10px;letter-spacing:1px;text-decoration:none;color:#555;}
      #ads{height:20px;margin-bottom:15px;text-align:right;background:#777;padding-top:4px;border-left:solid 1px #888;border-right:solid 1px #888;}
      #info_area{float:left;margin-left:11px;}
      #info_table{border-collapse:collapse;margin:0 auto;line-height:20px;width:728px;table-layout:fixed;}
      #info_table tr{height:30px;}
      #info_table td{border: solid 1px #888;padding: 0px 10px;}
      .info_table_label{width:100px;}
      #table_ads{width:260px;text-align:center;}
      #middle{height:110px;clear:both;text-align:center;padding-top:20px;}
      #cli_wrap{margin:0 10px;}
      h2{height:30px;line-height:30px;font-size:18px;margin:0 5px;padding-left:5px;}
      #cli_table{border:solid 1px #888;border-collapse:collapse;width:728px;margin:0 auto;line-height:20px;table-layout:fixed;}
      #cli_table tr{height:30px;}
      #cli_table td{border-bottom: solid 1px #888;padding-top:3px;padding-bottom:3px;overflow:hidden;}
      .cli_command{width:190px;padding-left:10px;}
      .cli_arrow{width:40px;text-align:center;}
      #footer{background:#777;height:20px;line-height:20px;color:#ddd;text-align:center;margin-top:15px;border-top:none;}
      #ip_address{font-size:1.9em;}
      #ip_address_cell{padding:5px;height:50px;}
      
    #plungins{height:20px;float:right;}
    .plungin{float:right;}
    #button_twitter{width:100px;}
    #button_plusone{width:70px;}
    #button_facebook{width:100px;}
    </style>
    <script async src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script>
  </head>
  <body>
    <div id="container" class="clearfix">
      <div id="header">
        <table>
          <tr>
            <td>
              <h1><a href="https://{{.BaseUrl}}">What Is My IP Address? - Yet Another ifconfig</a></h1>
            </td>
            <td></td>
          </tr>
          <tr>
            <td></td>
            <td>
            </td>
          </tr>
        </table>
      </div>
      <div id="ads">
<!-- Ads begin -->
<ins class="adsbygoogle"
    style="display:inline-block;width:728px;height:15px"
    data-ad-client="ca-pub-3435940217424489"
    data-ad-slot="9166587836"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>
<!-- Ads end -->
      </div>
      <div id="info_area">
        <h2>Your Connection</h2>
        <table id="info_table" summary="info">
          <tr>
            <td class="info_table_label">IP Address</td>
            <td id="ip_address_cell"><strong id="ip_address">{{.IP}}</strong></td>
            <td rowspan="12" id="table_ads">
<!-- Ads begin -->
<ins class="adsbygoogle"
    style="display:inline-block;width:250px;height:360px"
    data-ad-client="ca-pub-3435940217424489"
    data-ad-slot="9166587836"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>
<!-- Ads end -->

            </td>
          </tr>
          <tr><td class="info_table_label">Remote Host</td><td>{{.Host}}</td></tr>
          <tr><td class="info_table_label">QQWry</td><td>{{.QQWry}}</td></tr>
          <tr><td class="info_table_label">Geoip2</td><td>{{.Geoip2}}</td></tr>
          <tr><td class="info_table_label">IPIP.net</td><td>{{.IPIP}}</td></tr>
          <tr><td class="info_table_label">IP2Region</td><td>{{.IP2Region}}</td></tr>
          <tr><td class="info_table_label">User Agent</td><td>{{.UserAgent}}</td></tr>
          <tr><td class="info_table_label">Port</td><td>{{.Port}}</td></tr>
          <tr><td class="info_table_label">Language</td><td>{{.Lang}}</td></tr>
          <tr><td class="info_table_label">Referer</td><td>{{.Referer}}</td></tr>
          <tr><td class="info_table_label">Connection</td><td>{{.Connection}}</td></tr>
          <tr><td class="info_table_label">KeepAlive</td><td>{{.Keepalive}}</td></tr>
          <tr><td class="info_table_label">Method</td><td colspan="2">{{.Method}}</td></tr>
          <tr><td class="info_table_label">Encoding</td><td colspan="2">{{.Encoding}}</td></tr>
          <tr><td class="info_table_label">MIME Type</td><td colspan="2">{{.Mime}}</td></tr>
          <tr><td class="info_table_label">Charset</td><td colspan="2">{{.Charset}}</td></tr>
          <tr><td class="info_table_label">Via</td><td colspan="2">{{.Via}}</td></tr>
          <tr><td class="info_table_label">X-Forwarded-For</td><td colspan="2">{{.Forwarded}}</td></tr>
        </table>
      </div>
      <div id="middle">
<!-- Ads begin -->
<ins class="adsbygoogle"
    style="display:inline-block;width:728px;height:90px"
    data-ad-client="ca-pub-3435940217424489"
    data-ad-slot="9166587836"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>
<!-- Ads end -->
      </div>
      <div id="cli_wrap">
        <h2>Command Line Interface</h2>
        <table id="cli_table" summary="cli">
          <tr><td class="cli_command">$ curl {{.BaseUrl}}</td><td class="cli_arrow">&rArr;</td><td>{{.IP}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/ip</td><td class="cli_arrow">&rArr;</td><td>{{.IP}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/host</td><td class="cli_arrow">&rArr;</td><td>{{.Host}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/ua</td><td class="cli_arrow">&rArr;</td><td>{{.UserAgent}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/port</td><td class="cli_arrow">&rArr;</td><td>{{.Port}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/lang</td><td class="cli_arrow">&rArr;</td><td>{{.Lang}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/keepalive</td><td class="cli_arrow">&rArr;</td><td>{{.Keepalive}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/connection</td><td class="cli_arrow">&rArr;</td><td>{{.Connection}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/encoding</td><td class="cli_arrow">&rArr;</td><td>{{.Encoding}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/mime</td><td class="cli_arrow">&rArr;</td><td>{{.Mime}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/charset</td><td class="cli_arrow">&rArr;</td><td>{{.Charset}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/via</td><td class="cli_arrow">&rArr;</td><td>{{.Via}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/forwarded</td><td class="cli_arrow">&rArr;</td><td>{{.Forwarded}}</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/all</td><td class="cli_arrow">&rArr;</td><td>ip_addr: {{.IP}}<br />remote_host: {{.Host}} <br />user_agent: {{.UserAgent}}<br />port: {{.Port}}<br />lang: {{.Lang}}<br />connection: {{.Connection}}<br />keep_alive: {{.Keepalive}}<br />encoding: {{.Encoding}}<br />mime: {{.Mime}}<br />charset: {{.Charset}}<br />via: {{.Via}}<br />forwarded: {{.Forwarded}}<br /></td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/all.xml</td><td class="cli_arrow">&rArr;</td><td>&lt;info&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;charset&gt;{{.Charset}}&lt;/charset&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;connection&gt;{{.Connection}}&lt;/connection&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;encoding&gt;{{.Encoding}}&lt;/encoding&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;forwarded&gt;{{.Forwarded}}&lt;/forwarded&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;ip_addr&gt;{{.IP}}&lt;/ip_addr&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;keep_alive&gt;{{.Keepalive}}&lt;/keep_alive&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;lang&gt;{{.Lang}}&lt;/lang&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;mime&gt;{{.Mime}}&lt;/mime&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;port&gt;{{.Port}}&lt;/port&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;remote_host&gt;{{.Host}}&lt;/remote_host&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;user_agent&gt;{{.UserAgent}}&lt;/user_agent&gt;<br />&nbsp;&nbsp;&nbsp;&nbsp;&lt;via&gt;{{.Via}}&lt;/via&gt;<br />
&lt;/info&gt;<br />
</td></tr>
          <tr><td class="cli_command">$ curl {{.BaseUrl}}/all.json</td><td class="cli_arrow">&rArr;</td><td>{"connection":"{{.Connection}}","ip_addr":"{{.IP}}","lang":"{{.Lang}}","remote_host":"{{.Host}}","user_agent":"{{.UserAgent}}","charset":"{{.Charset}}","port":"{{.Port}}","via":"{{.Via}}","forwarded":"{{.Forwarded}}","mime":"{{.Mime}}","keep_alive":"{{.Keepalive}}","encoding":"{{.Encoding}}"}</td></tr>
        </table>
      </div>
      <div id="bottom">
<!-- Ads begin -->
<ins class="adsbygoogle"
    style="display:inline-block;width:728px;height:90px"
    data-ad-client="ca-pub-3435940217424489"
    data-ad-slot="9166587836"></ins>
<script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script>
<!-- Ads end -->
      </div>
      <div id="footer">&copy; 2014 - 2022 <a href="https://{{.BaseUrl}}">https://{{.BaseUrl}}</a></div>
    </div>
  </body>	
</html>
