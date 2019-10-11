window.downloadVideo = function (url,filename) {
  var timeout = Math.random() * 59000;
  this.setTimeout(function () {
    var xhr = new XMLHttpRequest();
    xhr.open("get", url, true);
    xhr.responseType = "blob";
    xhr.onload = function () {
      if (this.status == 200) {
        var blob = this.response;
        var downloadElement = document.createElement('a');
        var href = window.URL.createObjectURL(blob); //创建下载的链接
        downloadElement.href = href;
        downloadElement.download = filename; //下载后文件名
        document.body.appendChild(downloadElement);
        downloadElement.click(); //点击下载
        document.body.removeChild(downloadElement); //下载完成移除元素
        window.URL.revokeObjectURL(href); //释放掉blob对象
      }
    }
    xhr.send()
  }, timeout);
}

window.getAllTs=function() {
  var url = "https://vod.se0271.com//201904/7598deb0/";
  // "https://vod.se0271.com//201904/7598deb0/m3u8.m3u8"
  for(let i=0;i <=59;i++){
    let down_URL;
    let filename;
    if(i < 10){
      filename = "000"+i+".ts";
    } else if( i < 100){
      filename = "00"+i+".ts";
    } else{
      filename = "0"+i+".ts";
    }
    down_URL = url + filename;
    console.log("正在下载:"+ down_URL);
    downloadVideo(down_URL,filename);
  }
}